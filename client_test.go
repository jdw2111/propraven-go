package propraven

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// recorder captures requests and replays scripted responses.
type recorder struct {
	mu       sync.Mutex
	requests []*recorded
}

type recorded struct {
	Method string
	Path   string // escaped path
	Query  string
	Header http.Header
	Body   []byte
}

func (r *recorder) count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.requests)
}

func (r *recorder) last() *recorded {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.requests[len(r.requests)-1]
}

type testEnv struct {
	client *Client
	rec    *recorder
	waits  []time.Duration
	logs   *bytes.Buffer
	srv    *httptest.Server
}

// newEnv starts a server whose handler is called with the attempt index.
func newEnv(t *testing.T, handler func(w http.ResponseWriter, r *http.Request, attempt int), opts ...Option) *testEnv {
	t.Helper()
	t.Setenv("PROPRAVEN_API_KEY", "")
	t.Setenv("PROPRAVEN_BASE_URL", "")
	env := &testEnv{rec: &recorder{}, logs: &bytes.Buffer{}}
	env.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		env.rec.mu.Lock()
		n := len(env.rec.requests)
		env.rec.requests = append(env.rec.requests, &recorded{
			Method: r.Method, Path: r.URL.EscapedPath(), Query: r.URL.RawQuery, Header: r.Header.Clone(), Body: body,
		})
		env.rec.mu.Unlock()
		handler(w, r, n)
	}))
	t.Cleanup(env.srv.Close)
	all := append([]Option{WithBaseURL(env.srv.URL), WithLogger(log.New(env.logs, "", 0))}, opts...)
	env.client = NewClient(all...)
	env.client.core.sleep = func(_ <-chan struct{}, d time.Duration) bool {
		env.waits = append(env.waits, d)
		return true
	}
	env.client.core.jitter = func() float64 { return 0 }
	return env
}

func jsonReply(status int, body string, headers ...string) func(w http.ResponseWriter, r *http.Request, attempt int) {
	return func(w http.ResponseWriter, r *http.Request, attempt int) {
		for i := 0; i+1 < len(headers); i += 2 {
			w.Header().Set(headers[i], headers[i+1])
		}
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}
}

const parcelJSON = `{"id":"p1","county_fips":"37119","state_fips":"37","parcel_id":"12104406","address":"600 E 4TH ST","total_assessed_value":19388200,"year_built":1990}`

func TestAuthHeaderAndDefaults(t *testing.T) {
	env := newEnv(t, jsonReply(200, parcelJSON), WithAPIKey("pz_live_abc"))
	p, err := env.client.Parcels.Get(context.Background(), "37:119:12104406", nil)
	if err != nil {
		t.Fatal(err)
	}
	if p.ParcelID != "12104406" || p.Address == nil || *p.Address != "600 E 4TH ST" {
		t.Fatalf("decoded %+v", p)
	}
	if p.TotalAssessedValue == nil || *p.TotalAssessedValue != 19388200 {
		t.Fatalf("total_assessed_value = %v", p.TotalAssessedValue)
	}
	r := env.rec.last()
	if got := r.Header.Get("Authorization"); got != "Bearer pz_live_abc" {
		t.Errorf("Authorization = %q", got)
	}
	if got := r.Header.Get("User-Agent"); got != "propraven-go/"+Version {
		t.Errorf("User-Agent = %q", got)
	}
	if got := r.Header.Get("Accept"); got != "application/json" {
		t.Errorf("Accept = %q", got)
	}
	if got := r.Header.Get("Content-Type"); got != "" {
		t.Errorf("Content-Type on GET = %q", got)
	}
	if r.Method != "GET" || r.Path != "/api/v1/parcels/37:119:12104406" {
		t.Errorf("request %s %s", r.Method, r.Path)
	}
	if env.logs.Len() != 0 {
		t.Errorf("unexpected warning: %s", env.logs)
	}
}

func TestEnvironmentVariables(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"auth":"`+r.Header.Get("Authorization")+`"}`)
	}))
	defer srv.Close()
	t.Setenv("PROPRAVEN_API_KEY", "pz_from_env")
	t.Setenv("PROPRAVEN_BASE_URL", srv.URL+"/")
	c := NewClient()
	if c.BaseURL() != srv.URL {
		t.Fatalf("BaseURL = %q", c.BaseURL())
	}
	var raw RawResponse
	if _, err := c.Freshness.Get(context.Background(), nil, WithRawResponse(&raw)); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw.Body), "Bearer pz_from_env") {
		t.Fatalf("env key not sent: %s", raw.Body)
	}
	// An explicit option beats the environment.
	c2 := NewClient(WithAPIKey("pz_explicit"))
	if _, err := c2.Freshness.Get(context.Background(), nil, WithRawResponse(&raw)); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw.Body), "Bearer pz_explicit") {
		t.Fatalf("explicit key not sent: %s", raw.Body)
	}
}

func TestMissingKeySendsNoAuthorization(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`))
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Header.Get("Authorization"); got != "" {
		t.Fatalf("Authorization = %q, want none", got)
	}
	if env.logs.Len() != 0 {
		t.Fatalf("missing key should not warn: %s", env.logs)
	}
}

func TestKeyPrefixWarning(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`), WithAPIKey("sk_wrong"))
	if !strings.Contains(env.logs.String(), `does not start with "pz_"`) {
		t.Fatalf("no warning logged: %q", env.logs)
	}
	// It still sends the key.
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Header.Get("Authorization"); got != "Bearer sk_wrong" {
		t.Fatalf("Authorization = %q", got)
	}
}

func TestQueryAndPathEncoding(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{"data":[],"total":0,"limit":5,"offset":10}`))
	_, err := env.client.Deals.Absentee(context.Background(), &DealsAbsenteeParams{
		CountyFIPS: String("37119"),
		MinValue:   Float(250000.5),
		OutOfState: Bool(true),
		Limit:      Int(5),
		Offset:     Int(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	r := env.rec.last()
	want := "county_fips=37119&limit=5&min_value=250000.5&offset=10&out_of_state=true"
	if r.Query != want {
		t.Errorf("query = %q, want %q", r.Query, want)
	}
	// Unset params are omitted; false booleans are sent as "false".
	_, err = env.client.Deals.Absentee(context.Background(), &DealsAbsenteeParams{OutOfState: Bool(false)})
	if err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Query; got != "out_of_state=false" {
		t.Errorf("query = %q", got)
	}
	// Path params are escaped per segment.
	if _, err := env.client.Owners.Get(context.Background(), "SMITH JOHN/A", nil); err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Path; got != "/api/v1/owners/SMITH%20JOHN%2FA" {
		t.Errorf("path = %q", got)
	}
	if _, err := env.client.Webhooks.RetryDelivery(context.Background(), "wh 1", "d?2", nil); err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Path; got != "/api/v1/webhooks/wh%201/deliveries/d%3F2/retry" {
		t.Errorf("path = %q", got)
	}
}

func TestHeaderParams(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`))
	_, err := env.client.Credits.Balance(context.Background(), &CreditsBalanceParams{CreditToken: String("ct_123")})
	if err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Header.Get("X-Credit-Token"); got != "ct_123" {
		t.Fatalf("X-CREDIT-TOKEN = %q", got)
	}
	if got := env.rec.last().Header.Get("X-Payment"); got != "" {
		t.Fatalf("X-PAYMENT must not be sent: %q", got)
	}
}

func TestJSONBody(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{"data":[],"total":0,"total_is_estimate":false,"total_is_lower_bound":false,"has_more":false,"limit":50,"offset":0,"sort_applied":null}`))
	_, err := env.client.Search.Parcels(context.Background(), &SearchParcelsParams{
		Bounds: &SearchParcelsParamsBounds{North: 35.215, South: 35.205, East: -80.855, West: -80.865},
		Filters: &SearchParcelsParamsFilters{
			ValueRange: &SearchParcelsParamsFiltersValueRange{Min: Float(100000)},
			OwnerTypes: []SearchParcelsParamsFiltersOwnerTypes{SearchParcelsParamsFiltersOwnerTypesLLC},
		},
		Limit: Int(50),
	})
	if err != nil {
		t.Fatal(err)
	}
	r := env.rec.last()
	if r.Method != "POST" || r.Path != "/api/v1/search" {
		t.Fatalf("%s %s", r.Method, r.Path)
	}
	if got := r.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q", got)
	}
	var body map[string]any
	if err := json.Unmarshal(r.Body, &body); err != nil {
		t.Fatalf("body %s: %v", r.Body, err)
	}
	want := `{"bounds":{"east":-80.855,"north":35.215,"south":35.205,"west":-80.865},"filters":{"ownerTypes":["llc"],"valueRange":{"min":100000}},"limit":50}`
	got, _ := json.Marshal(body)
	if string(got) != want {
		t.Errorf("body = %s\nwant   %s", got, want)
	}
	// A required body with nil params still sends a JSON object.
	if _, err := env.client.Lookup.Batch(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if got := string(env.rec.last().Body); got != "{}" {
		t.Errorf("nil params body = %q", got)
	}
}

func TestCSVEndpointReturnsString(t *testing.T) {
	env := newEnv(t, jsonReply(200, "parcel_id,address\n1,MAIN ST\n", "Content-Type", "text/csv; charset=utf-8"))
	csv, err := env.client.Search.Export(context.Background(), &SearchExportParams{North: Float(35.2), South: Float(35.1), East: Float(-80.8), West: Float(-80.9)})
	if err != nil {
		t.Fatal(err)
	}
	if csv != "parcel_id,address\n1,MAIN ST\n" {
		t.Fatalf("csv = %q", csv)
	}
	if got := env.rec.last().Header.Get("Accept"); !strings.HasPrefix(got, "text/csv") {
		t.Errorf("Accept = %q", got)
	}
}

func TestMixedCSVOrJSON(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if r.URL.Query().Get("preview") == "true" {
			jsonReply(200, `{"preview":true,"parcels_listed":3}`)(w, r, attempt)
			return
		}
		jsonReply(200, "a,b\n", "Content-Type", "text/csv")(w, r, attempt)
	})
	out, err := env.client.Cohorts.Export(context.Background(), "c1", &CohortsExportParams{Preview: Bool(true)})
	if err != nil {
		t.Fatal(err)
	}
	if out.JSON == nil || out.Text != "" {
		t.Fatalf("preview: %+v", out)
	}
	out, err = env.client.Cohorts.Export(context.Background(), "c1", nil)
	if err != nil {
		t.Fatal(err)
	}
	if out.Text != "a,b\n" || !strings.HasPrefix(out.ContentType, "text/csv") {
		t.Fatalf("csv: %+v", out)
	}
}

func TestUnionResponse(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if r.URL.Query().Get("shape") == "envelope" {
			jsonReply(200, `{"data":[{"permit_number":"B1"}],"permit_count":"7","permit_count_basis":"exact","truncated":false,"row_cap":100}`)(w, r, attempt)
			return
		}
		jsonReply(200, `[{"permit_number":"B1"},{"permit_number":"B2"}]`)(w, r, attempt)
	})
	bare, err := env.client.Parcels.Permits(context.Background(), "p1", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(bare.Array) != 2 || bare.Object != nil {
		t.Fatalf("bare: %+v", bare)
	}
	env2, err := env.client.Parcels.Permits(context.Background(), "p1", &ParcelsPermitsParams{Shape: String(ParcelsPermitsParamsShapeEnvelope)})
	if err != nil {
		t.Fatal(err)
	}
	if env2.Object == nil || env2.Object.PermitCount != 7 || len(env2.Object.Data) != 1 || env2.Array != nil {
		t.Fatalf("envelope: %+v", env2.Object)
	}
	round, _ := json.Marshal(bare)
	if !strings.HasPrefix(string(round), "[") {
		t.Fatalf("union marshal = %s", round)
	}
}

func TestLenientNumbersAndTypeDrift(t *testing.T) {
	// Quoted decimals (older servers), plain numbers (newer), and a value
	// whose type drifted from the spec must all decode without error.
	body := `{"id":"p1","county_fips":"37119","state_fips":"37","parcel_id":"1",
		"total_assessed_value":"19388200.00","year_built":"1990","latitude":35.2,"state":"NC",
		"bedrooms":null,"address":12345}`
	env := newEnv(t, jsonReply(200, body))
	p, err := env.client.Parcels.Get(context.Background(), "p1", nil)
	if err != nil {
		t.Fatal(err)
	}
	if p.TotalAssessedValue == nil || *p.TotalAssessedValue != 19388200 {
		t.Errorf("total_assessed_value = %v", p.TotalAssessedValue)
	}
	if p.YearBuilt == nil || *p.YearBuilt != 1990 {
		t.Errorf("year_built = %v", p.YearBuilt)
	}
	if p.Latitude == nil || *p.Latitude != 35.2 {
		t.Errorf("latitude = %v", p.Latitude)
	}
	if p.Bedrooms != nil || p.State != nil {
		t.Errorf("bedrooms = %v state = %v; want nil", p.Bedrooms, p.State)
	}
	if p.Address != nil && *p.Address != "" {
		t.Errorf("address = %q; want unset", *p.Address)
	}
	if p.CountyFIPS != "37119" {
		t.Errorf("county_fips = %q", p.CountyFIPS)
	}

	// A drifted field in a struct without numeric fields: the rest still
	// decodes, and the mismatch is logged.
	env = newEnv(t, jsonReply(200, `{"owner_name":5,"mailing_address":"1 MAIN ST","owner":{"owner_city":"RALEIGH"}}`))
	o, err := env.client.Parcels.Owner(context.Background(), "p1", nil)
	if err != nil {
		t.Fatal(err)
	}
	if o.MailingAddress == nil || *o.MailingAddress != "1 MAIN ST" || o.Owner.OwnerCity == nil || *o.Owner.OwnerCity != "RALEIGH" {
		t.Errorf("owner = %+v", o)
	}
	if !strings.Contains(env.logs.String(), "owner_name") {
		t.Errorf("type drift not logged: %q", env.logs)
	}
}

func TestPerRequestOptions(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`), WithAPIKey("pz_a"), WithHeader("X-Trace", "client"))
	_, err := env.client.Freshness.Get(context.Background(), nil, WithAPIKey("pz_b"), WithHeader("X-Trace", "call"))
	if err != nil {
		t.Fatal(err)
	}
	r := env.rec.last()
	if r.Header.Get("Authorization") != "Bearer pz_b" || r.Header.Get("X-Trace") != "call" {
		t.Fatalf("headers %v", r.Header)
	}
	// The client's own settings are untouched afterwards.
	if _, err := env.client.Freshness.Get(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	r = env.rec.last()
	if r.Header.Get("Authorization") != "Bearer pz_a" || r.Header.Get("X-Trace") != "client" {
		t.Fatalf("headers %v", r.Header)
	}
	// WithHeader cannot override Authorization.
	if _, err := env.client.Freshness.Get(context.Background(), nil, WithHeader("Authorization", "Bearer nope")); err != nil {
		t.Fatal(err)
	}
	if got := env.rec.last().Header.Get("Authorization"); got != "Bearer pz_a" {
		t.Fatalf("Authorization = %q", got)
	}
}

func TestRateLimitInfo(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`, "X-RateLimit-Limit", "1000", "X-RateLimit-Remaining", "999", "X-RateLimit-Reset", "1700000000"))
	if env.client.LastRateLimit() != nil {
		t.Fatal("LastRateLimit before any call should be nil")
	}
	if _, err := env.client.Account.Usage(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	rl := env.client.LastRateLimit()
	if rl == nil || rl.Limit != 1000 || rl.Remaining != 999 || rl.Reset != 1700000000 {
		t.Fatalf("rate limit = %+v", rl)
	}
	if !rl.ResetTime().Equal(time.Unix(1700000000, 0)) {
		t.Fatalf("ResetTime = %v", rl.ResetTime())
	}
}

func TestTimeout(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		select {
		case <-r.Context().Done():
		case <-time.After(2 * time.Second):
		}
	}, WithTimeout(20*time.Millisecond), WithMaxRetries(1))
	_, err := env.client.Freshness.Get(context.Background(), nil)
	if !IsTimeout(err) || !IsConnectionError(err) {
		t.Fatalf("err = %v, want timeout", err)
	}
	if env.rec.count() != 2 {
		t.Fatalf("attempts = %d, want 2 (timeouts are retried)", env.rec.count())
	}
}

func TestCancelledContextIsNotRetried(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{}`))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := env.client.Freshness.Get(ctx, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v", err)
	}
	if env.rec.count() != 0 {
		t.Fatalf("requests = %d", env.rec.count())
	}
}

func TestConnectionErrorRetried(t *testing.T) {
	t.Setenv("PROPRAVEN_API_KEY", "")
	c := NewClient(WithBaseURL("http://127.0.0.1:1"), WithLogger(nil), WithMaxRetries(2))
	var waits int
	c.core.sleep = func(_ <-chan struct{}, d time.Duration) bool { waits++; return true }
	_, err := c.Freshness.Get(context.Background(), nil)
	if !IsConnectionError(err) || IsTimeout(err) {
		t.Fatalf("err = %v", err)
	}
	if waits != 2 {
		t.Fatalf("waits = %d, want 2", waits)
	}
}
