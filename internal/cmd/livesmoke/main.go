// Command livesmoke exercises the SDK against the live API with a real key.
// It is NOT run in CI.
//
//	PROPRAVEN_API_KEY=pz_... go run ./internal/cmd/livesmoke
//
// Guard rails, enforced by the transport below rather than by convention:
//   - at most 12 HTTP requests, at most 2 per second, retries disabled;
//   - only the read-only operations listed in allowed (no paid endpoints,
//     no webhook/watch/cohort creation, the only POST is /api/v1/search);
//   - a request carrying X-PAYMENT is refused before it leaves the process;
//   - the API key is never printed.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jdw2111/propraven-go"
)

const (
	maxCalls   = 12
	minSpacing = 600 * time.Millisecond // keeps us under 2 requests/second
)

var allowed = []struct {
	method string
	path   *regexp.Regexp
}{
	{"POST", regexp.MustCompile(`^/api/v1/search$`)},
	{"GET", regexp.MustCompile(`^/api/v1/parcels/[^/]+$`)},
	{"GET", regexp.MustCompile(`^/api/v1/parcels/[^/]+/permits$`)},
	{"GET", regexp.MustCompile(`^/api/v1/search/full$`)},
	{"GET", regexp.MustCompile(`^/api/v1/deals/absentee$`)},
	{"GET", regexp.MustCompile(`^/api/v1/market/counties$`)},
	{"GET", regexp.MustCompile(`^/api/v1/owners/[^/]+$`)},
	{"GET", regexp.MustCompile(`^/api/v1/coverage$`)},
	{"GET", regexp.MustCompile(`^/api/v1/freshness$`)},
	{"GET", regexp.MustCompile(`^/api/v1/account/usage$`)},
}

type guard struct {
	mu    sync.Mutex
	calls int
	last  time.Time
	next  http.RoundTripper
}

func (g *guard) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Header.Get("X-PAYMENT") != "" {
		return nil, errors.New("livesmoke guard: refusing a request with X-PAYMENT")
	}
	ok := false
	for _, a := range allowed {
		if r.Method == a.method && a.path.MatchString(r.URL.Path) {
			ok = true
		}
	}
	if !ok {
		return nil, fmt.Errorf("livesmoke guard: %s %s is not on the read-only allowlist", r.Method, r.URL.Path)
	}
	g.mu.Lock()
	if g.calls >= maxCalls {
		g.mu.Unlock()
		return nil, fmt.Errorf("livesmoke guard: call budget of %d exhausted", maxCalls)
	}
	g.calls++
	if wait := minSpacing - time.Since(g.last); !g.last.IsZero() && wait > 0 {
		time.Sleep(wait)
	}
	g.last = time.Now()
	g.mu.Unlock()
	return g.next.RoundTrip(r)
}

type result struct {
	name, outcome, detail string
	status                int
	ok                    bool
}

func main() {
	if strings.TrimSpace(os.Getenv("PROPRAVEN_API_KEY")) == "" {
		fmt.Fprintln(os.Stderr, "livesmoke: PROPRAVEN_API_KEY is not set")
		os.Exit(2)
	}
	g := &guard{next: http.DefaultTransport}
	client := propraven.NewClient(
		propraven.WithHTTPClient(&http.Client{Transport: g}),
		propraven.WithMaxRetries(0),
		propraven.WithTimeout(45*time.Second),
	)
	ctx := context.Background()
	var results []result
	record := func(name string, raw *propraven.RawResponse, err error, detail string, wantErr func(error) bool) {
		r := result{name: name, status: raw.StatusCode, detail: detail}
		switch {
		case wantErr != nil && err != nil && wantErr(err):
			r.ok, r.outcome = true, "OK (expected error)"
			if e := propraven.AsError(err); e != nil {
				r.status = e.StatusCode
				r.detail = fmt.Sprintf("%s: %s", e.Class(), e.Error())
			}
		case wantErr != nil:
			r.outcome = "FAIL (wanted an error)"
			if err != nil {
				r.detail = err.Error()
			}
		case err != nil:
			r.outcome = "FAIL"
			r.detail = err.Error()
			if e := propraven.AsError(err); e != nil {
				r.status = e.StatusCode
				r.detail = fmt.Sprintf("%s: %s", e.Class(), e.Error())
			}
		default:
			r.ok, r.outcome = true, "OK"
		}
		results = append(results, r)
	}
	str := func(p *string) string {
		if p == nil {
			return "<nil>"
		}
		return *p
	}

	// 1. search.parcels (POST /api/v1/search) — discovers a real parcel id.
	parcelID, ownerName := "37:119:12104406", "MARKEY ENTERPRISES INC"
	{
		var raw propraven.RawResponse
		res, err := client.Search.Parcels(ctx, &propraven.SearchParcelsParams{
			Bounds: &propraven.SearchParcelsParamsBounds{North: 35.215, South: 35.205, East: -80.855, West: -80.865},
			Limit:  propraven.Int(2),
		}, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			detail = fmt.Sprintf("rows=%d total=%d has_more=%t", len(res.Data), res.Total, res.HasMore)
			if len(res.Data) > 0 {
				d := res.Data[0]
				if len(d.CountyFIPS) == 5 && d.ParcelID != "" {
					parcelID = d.StateFIPS + ":" + d.CountyFIPS[2:] + ":" + d.ParcelID
				}
				if d.OwnerName != nil && *d.OwnerName != "" {
					ownerName = *d.OwnerName
				}
				detail += fmt.Sprintf(" first=%s value=%v", parcelID, deref(d.TotalAssessedValue))
			}
		}
		record("search.parcels", &raw, err, detail, nil)
	}

	// 2. parcels.get
	{
		var raw propraven.RawResponse
		p, err := client.Parcels.Get(ctx, parcelID, nil, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			detail = fmt.Sprintf("parcel_id=%s county_fips=%s address=%s total_assessed_value=%v year_built=%v",
				p.ParcelID, p.CountyFIPS, str(p.Address), deref(p.TotalAssessedValue), derefInt(p.YearBuilt))
			if p.OwnerName != nil && *p.OwnerName != "" {
				ownerName = *p.OwnerName
			}
		}
		record("parcels.get", &raw, err, detail, nil)
	}

	// 3. parcels.permits (envelope shape)
	{
		var raw propraven.RawResponse
		res, err := client.Parcels.Permits(ctx, parcelID, &propraven.ParcelsPermitsParams{
			Shape: propraven.String(propraven.ParcelsPermitsParamsShapeEnvelope),
		}, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			switch {
			case res.Object != nil:
				detail = fmt.Sprintf("envelope rows=%d permit_count=%d basis=%s", len(res.Object.Data), res.Object.PermitCount, res.Object.PermitCountBasis)
			case res.Array != nil:
				detail = fmt.Sprintf("bare array rows=%d", len(res.Array))
			default:
				detail = "unrecognised shape"
			}
		}
		record("parcels.permits", &raw, err, detail, nil)
	}

	// 4-5. search.full: first page + one cursor page through the iterator.
	{
		var raw propraven.RawResponse
		it := client.Search.FullIter(ctx, &propraven.SearchFullParams{Q: propraven.String("main st"), State: propraven.String("NC")},
			propraven.IterOptions{PageSize: 2, MaxItems: 3}, propraven.WithRawResponse(&raw))
		var ids []string
		for it.Next() {
			ids = append(ids, it.Current().ParcelID)
		}
		detail := fmt.Sprintf("pages=%d items=%d ids=%v", it.Pages(), len(ids), ids)
		if it.Err() == nil && it.Pages() < 2 {
			detail += " (no second page: cursor absent)"
		}
		record("search.full (+cursor page via FullIter)", &raw, it.Err(), detail, nil)
	}

	// 6-7. deals.absentee: 2 pages of 2 via the offset iterator.
	{
		var raw propraven.RawResponse
		it := client.Deals.AbsenteeIter(ctx, &propraven.DealsAbsenteeParams{CountyFIPS: propraven.String("37119")},
			propraven.IterOptions{PageSize: 2, MaxItems: 4}, propraven.WithRawResponse(&raw))
		n := 0
		var first string
		for it.Next() {
			if n == 0 {
				first = it.Current().ParcelID
			}
			n++
		}
		record("deals.absentee (AbsenteeIter, 2x2)", &raw, it.Err(), fmt.Sprintf("pages=%d items=%d first=%s", it.Pages(), n, first), nil)
	}

	// 8. owners.get
	{
		var raw propraven.RawResponse
		o, err := client.Owners.Get(ctx, ownerName, nil, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			detail = fmt.Sprintf("owner=%q property_count=%v total_assessed_value=%v", ownerName, derefInt(o.PropertyCount), deref(o.TotalAssessedValue))
		} else {
			detail = fmt.Sprintf("owner=%q", ownerName)
		}
		record("owners.get", &raw, err, detail, nil)
	}

	// 9. coverage.get
	{
		var raw propraven.RawResponse
		c, err := client.Coverage.Get(ctx, &propraven.CoverageGetParams{State: propraven.String("NC")}, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			detail = fmt.Sprintf("rows=%d state=%s total_counties=%v", len(c.Data), str(c.State), derefInt(c.TotalCounties))
		}
		record("coverage.get", &raw, err, detail, nil)
	}

	// 10. account.usage
	{
		var raw propraven.RawResponse
		u, err := client.Account.Usage(ctx, nil, propraven.WithRawResponse(&raw))
		detail := ""
		if err == nil {
			detail = fmt.Sprintf("tier=%s calls_used=%d calls_remaining=%d", u.Tier, u.CallsUsed, u.CallsRemaining)
		}
		record("account.usage", &raw, err, detail, nil)
	}

	// 11. error path: unknown parcel -> 404
	{
		var raw propraven.RawResponse
		_, err := client.Parcels.Get(ctx, "37:119:sdk-livesmoke-does-not-exist", nil, propraven.WithRawResponse(&raw))
		record("parcels.get (missing id -> NotFoundError)", &raw, err, "", propraven.IsNotFound)
	}

	// 12. error path: invalid limit -> 400
	{
		var raw propraven.RawResponse
		_, err := client.Market.Counties(ctx, &propraven.MarketCountiesParams{Limit: propraven.Int(-1)}, propraven.WithRawResponse(&raw))
		record("market.counties (limit=-1 -> BadRequestError)", &raw, err, "", propraven.IsBadRequest)
	}

	fmt.Printf("propraven-go %s live smoke against %s (%s)\n", propraven.Version, client.BaseURL(), time.Now().UTC().Format(time.RFC3339))
	failed := 0
	for i, r := range results {
		if !r.ok {
			failed++
		}
		fmt.Printf("%2d. %-45s %-22s http=%d  %s\n", i+1, r.name, r.outcome, r.status, r.detail)
	}
	if rl := client.LastRateLimit(); rl != nil {
		fmt.Printf("last rate limit: limit=%d remaining=%d reset=%d\n", rl.Limit, rl.Remaining, rl.Reset)
	}
	fmt.Printf("HTTP requests made: %d (budget %d)\n", g.calls, maxCalls)
	fmt.Printf("result: %d/%d checks passed\n", len(results)-failed, len(results))
	if failed > 0 {
		os.Exit(1)
	}
}

func deref(p *float64) any {
	if p == nil {
		return "<nil>"
	}
	return strconv.FormatFloat(*p, 'f', -1, 64)
}

func derefInt(p *int64) any {
	if p == nil {
		return "<nil>"
	}
	return *p
}
