# propraven-go

The Go SDK for the [PropRaven](https://propraven.com) REST API: national property intelligence over U.S. parcels, owners, deeds, permits and markets.

- Standard library only, no third-party dependencies.
- Typed methods for every operation in the [OpenAPI spec](https://propraven.com/openapi.json), generated in-repo from the vendored `openapi.json`.
- Retries with backoff, RFC 7807 error parsing, auto-pagination iterators, webhook signature verification.

Requires Go 1.22 or newer.

```bash
go get github.com/jdw2111/propraven-go@v0.3.0
```

## Documentation

- Developer hub: https://propraven.com/developers
- REST API v1 reference: https://propraven.com/docs/v1
- Hosted MCP server: https://propraven.com/docs/mcp
- Go reference: https://pkg.go.dev/github.com/jdw2111/propraven-go

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jdw2111/propraven-go"
)

func main() {
	client := propraven.NewClient() // reads PROPRAVEN_API_KEY

	ctx := context.Background()
	parcel, err := client.Parcels.Get(ctx, "37:119:12104406", nil)
	if err != nil {
		log.Fatal(err)
	}
	if parcel.Address != nil {
		fmt.Println(parcel.ParcelID, *parcel.Address)
	}

	permits, err := client.Parcels.Permits(ctx, parcel.ID, &propraven.ParcelsPermitsParams{
		Shape: propraven.String(propraven.ParcelsPermitsParamsShapeEnvelope),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(permits.Object.PermitCount)
}
```

Every method has the shape `client.<Namespace>.<Method>(ctx, pathParams..., params, opts...)`:

- Path parameters are positional, in path order.
- `params` is a pointer to `<Namespace><Method>Params` holding the query parameters, header parameters and JSON-body fields. It may be nil.
- `opts` are optional per-request options (see [Options](#options)).
- The result is a pointer to `<Namespace><Method>Response`. Endpoints that return CSV return a `string`.

Optional fields are pointers. Use the helpers `propraven.String`, `propraven.Int`, `propraven.Float`, `propraven.Bool` and `propraven.Ptr`. Enumerations are string types with constants, for example `propraven.ParcelsPermitsParamsShapeEnvelope`. Numbers are `float64` (JSON number) or `int64` (JSON integer). Identifiers such as `parcel_id`, `county_fips`, `state_fips` and `zip` are strings.

The decoder ignores unknown response fields. It accepts a numeric field sent either as a JSON number or as a quoted decimal string (`"19388200.00"`, which older API deployments send). If a field's JSON type does not match the spec, the decoder leaves that field unset and still returns the rest of the response.

## Authentication

Pass the key with `propraven.WithAPIKey("pz_...")`, or set `PROPRAVEN_API_KEY`. The client sends it as `Authorization: Bearer <key>`. Keys start with `pz_`. If a key does not, the client logs a warning and uses it anyway. You can create a client without a key, because several endpoints don't need one. Endpoints that do need a key answer 401.

**Use the client on servers only.** API keys are secrets, and the REST API sends no CORS headers, so browser code cannot call it anyway.

## Options

```go
client := propraven.NewClient(
	propraven.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
	propraven.WithBaseURL("https://propraven.com"), // or PROPRAVEN_BASE_URL
	propraven.WithMaxRetries(3),                     // default 2
	propraven.WithTimeout(30*time.Second),           // per attempt, default 60s
	propraven.WithHeader("X-Trace-Id", "abc"),
	propraven.WithHTTPClient(&http.Client{Transport: myTransport}),
	propraven.WithLogger(nil),                       // silence warnings (default: stderr)
)
```

You can also pass any option to a single call, where it applies to that call only:

```go
var raw propraven.RawResponse
usage, err := client.Account.Usage(ctx, nil,
	propraven.WithTimeout(5*time.Second),
	propraven.WithMaxRetries(0),
	propraven.WithRawResponse(&raw), // raw status, headers and body
)
```

The base URL defaults to `https://propraven.com`; every path already includes `/api/v1`. `https://api.propraven.com` serves the same API.

## Errors

A non-2xx response returns a `*propraven.Error`. The client parses the body as RFC 7807 `application/problem+json`, as an x402 payment-required envelope, or as a legacy `{"error": "..."}` body:

```go
parcel, err := client.Parcels.Get(ctx, "37:119:missing", nil)
var apiErr *propraven.Error
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Code, apiErr.Detail, apiErr.RequestID)
	for _, fe := range apiErr.Errors { // per-parameter validation failures
		fmt.Println(fe.Param, fe.Message)
	}
}
switch {
case propraven.IsNotFound(err):
case propraven.IsRateLimited(err):     // apiErr.RetryAfter (seconds) when the server sent one
case propraven.IsPaymentRequired(err): // apiErr.Accepts holds x402 payment requirements
case propraven.IsAuthentication(err):
case propraven.IsTimeout(err), propraven.IsConnectionError(err): // no response received
}
```

`Error()` reads `"<status> <code>: <detail>"`. `apiErr.Class()` gives the error class name that the TypeScript and Python SDKs use (`NotFoundError`, `RateLimitError`, `PaymentRequiredError`, ...). Other predicates: `IsBadRequest`, `IsPermissionDenied`, `IsConflict` and `IsServerError`.

The SDK never signs or sends x402 payments. A 402 response comes back as a `PaymentRequired` error.

## Retries

By default the client retries a failed request up to 2 times:

- It retries network errors, timeouts, 429, 503 and 504 for any method. It retries other 5xx responses only for GET, HEAD, DELETE and OPTIONS.
- It never retries other 4xx responses, and never 402. The monthly-cap `Retry-After` is measured in days.
- To pick the wait, it uses `Retry-After` (seconds or an HTTP date) first. If that header is absent and `X-RateLimit-Remaining` is 0, it waits until `X-RateLimit-Reset`. Otherwise it backs off exponentially: 0.5 s × 2^attempt, ±25% jitter.
- It waits at most 60 s. If the server asks for a longer wait, the client returns the error without retrying.

## Pagination

Every paginated operation `X` also has a method `XIter`. It returns an iterator that fetches pages lazily:

```go
it := client.Deals.AbsenteeIter(ctx, &propraven.DealsAbsenteeParams{
	CountyFIPS: propraven.String("37119"),
}, propraven.IterOptions{PageSize: 100, MaxItems: 1000})
for it.Next() {
	row := it.Current()
	fmt.Println(row.ParcelID)
}
if err := it.Err(); err != nil {
	log.Fatal(err)
}
```

For offset-paginated operations, the iterator advances `offset` by the page size, whether the page is in the query or in the JSON body (`POST /api/v1/search`). It stops when a page comes back short, when `offset` reaches `total`, or when `has_more` is false. If the server clamps the page size, the iterator follows the `limit` the server echoes back.

For cursor pagination (`client.Search.FullIter`), the iterator passes each page's `nextCursor` as `after` to get the next page. It stops when the cursor is null or absent, or when `hasMore` is false.

The iterator uses `Next`, `Current` and `Err` (like `bufio.Scanner`), so the module still supports Go 1.22. It does not use Go 1.23 range-over-func.

## Webhooks

```go
func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	event, err := propraven.VerifyWebhook(body, r.Header.Get("X-PropRaven-Signature"), os.Getenv("PROPRAVEN_WEBHOOK_SECRET"))
	if err != nil { // errors.Is(err, propraven.ErrWebhookSignature)
		http.Error(w, "invalid signature", http.StatusUnauthorized)
		return
	}
	var sold propraven.ParcelSoldEvent
	_ = json.Unmarshal(event, &sold)
	w.WriteHeader(http.StatusOK)
}
```

The signature header is `t=<unix_ms>,v1=<hex>`, where `v1` is the HMAC-SHA256 of `"<t>.<raw body>"`. The key is the secret exactly as issued (`whsec_...`). The default tolerance is 300 s; change it with `propraven.WithWebhookTolerance`. The verifier accepts more than one `v1=` entry, to allow secret rotation, and compares in constant time. To test your handler, generate a header with `propraven.SignWebhook(body, secret, time.Now())`.

## Rate-limit info

```go
rl := client.LastRateLimit() // nil until a response has carried rate-limit headers
if rl != nil {
	fmt.Println(rl.Limit, rl.Remaining, rl.ResetTime())
}
```

## Methods
<!-- BEGIN GENERATED METHODS (go run ./internal/cmd/generate) -->

### client.Parcels

| Method | HTTP | Summary |
|---|---|---|
| `AssessmentHistory(ctx, id, params)` | `GET /api/v1/parcels/{id}/assessment-history` | Get recorded annual assessment history |
| `Get(ctx, id, params)` | `GET /api/v1/parcels/{id}` | Get parcel by ID |
| `Owner(ctx, id, params)` | `GET /api/v1/parcels/{id}/owner` | Get parcel owner details and portfolio |
| `Permits(ctx, id, params)` | `GET /api/v1/parcels/{id}/permits` | Get parcel permits |
| `Deeds(ctx, id, params)` | `GET /api/v1/parcels/{id}/deeds` | Get parcel deed history |
| `Risks(ctx, id, params)` | `GET /api/v1/parcels/{id}/risks` | Get parcel risk assessment |
| `TaxStatus(ctx, id, params)` | `GET /api/v1/parcels/{id}/tax-status` | Property-tax delinquency status of a parcel |
| `Geojson(ctx, params)` | `GET /api/v1/parcels/geojson` | Parcel polygons as GeoJSON for a bounding box |
| `Report(ctx, id, params)` | `GET /api/v1/parcels/{id}/report` | Parcel dossier (paid, provenance-first) |
| `CompPack(ctx, id, params)` | `GET /api/v1/parcels/{id}/comp-pack` | Comp pack (paid, priced per pack) — with a FREE preview |
| `RiskScore(ctx, id, params)` | `GET /api/v1/parcels/{id}/risk-score` | Risk score (paid, priced per assessment) — with a FREE preview |
| `TrafficHistory(ctx, id, params)` | `GET /api/v1/parcels/{id}/traffic-history` | Nearest traffic station + AADT history |
| `Batch(ctx, params)` | `POST /api/v1/parcels/batch` | Fetch up to 100 parcels by (state, county, parcel) tuple |
| `Comps(ctx, id, params)` | `GET /api/v1/parcels/{id}/comps` | Comparable sales for a parcel |
| `Occupants(ctx, id, params)` | `GET /api/v1/parcels/{id}/occupants` | Business occupants of a parcel |
| `Violations(ctx, id, params)` | `GET /api/v1/parcels/{id}/violations` | Code violations on a parcel |
| `Pois(ctx, params)` | `GET /api/v1/parcels/poi` | Business parcels in a small bounding box |

### client.Search

| Method | HTTP | Summary |
|---|---|---|
| `Parcels(ctx, params)` | `POST /api/v1/search` | Search parcels |
| `ParcelsIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Parcels` |
| `Autocomplete(ctx, params)` | `GET /api/v1/search/autocomplete` | Address / place / parcel autocomplete |
| `Export(ctx, params)` | `GET /api/v1/search/export` | Export search results as CSV |
| `Full(ctx, params)` | `GET /api/v1/search/full` | Full paginated text + attribute search |
| `FullIter(ctx, params, iterOpts)` | cursor pages of `results` | iterator over `Full` |

### client.Coverage

| Method | HTTP | Summary |
|---|---|---|
| `Get(ctx, params)` | `GET /api/v1/coverage` | Get coverage statistics |
| `Map(ctx, params)` | `GET /api/v1/coverage/map` | County coverage map data |

### client.Deals

| Method | HTTP | Summary |
|---|---|---|
| `Absentee(ctx, params)` | `GET /api/v1/deals/absentee` | Find absentee owners |
| `AbsenteeIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Absentee` |
| `Flips(ctx, params)` | `GET /api/v1/deals/flips` | Find property flips |
| `FlipsIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Flips` |
| `Contractors(ctx, params)` | `GET /api/v1/deals/contractors` | Search contractors by permit activity |
| `ContractorsIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Contractors` |
| `Entities(ctx, params)` | `GET /api/v1/deals/entities` | Find entity-owned parcels (LLC, Corp, Trust, LP) |
| `EntitiesIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Entities` |
| `HighLandRatio(ctx, params)` | `GET /api/v1/deals/high-land-ratio` | Find parcels with high land-to-improvement ratio |
| `HighLandRatioIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `HighLandRatio` |
| `Lenders(ctx, params)` | `GET /api/v1/deals/lenders` | Search lender profiles |
| `LendersIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Lenders` |
| `LongHold(ctx, params)` | `GET /api/v1/deals/long-hold` | Find long-held parcels (10+ years) |
| `LongHoldIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `LongHold` |
| `Market(ctx, params)` | `GET /api/v1/deals/market` | County-quarter transaction summary or affordability index |
| `MarketIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Market` |
| `PortfolioOwners(ctx, params)` | `GET /api/v1/deals/portfolio-owners` | Find portfolio investors (owners of 2+ properties) |
| `PortfolioOwnersIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `PortfolioOwners` |

### client.Market

| Method | HTTP | Summary |
|---|---|---|
| `Counties(ctx, params)` | `GET /api/v1/market/counties` | Get county market statistics |
| `CountiesIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Counties` |
| `Trends(ctx, params)` | `GET /api/v1/market/trends` | Get market trends |
| `County(ctx, fips, params)` | `GET /api/v1/market/counties/{fips}` | Detailed view for a single county |
| `Flips(ctx, params)` | `GET /api/v1/market/flips` | Flip-activity summary grouped by county |
| `FlipsIter(ctx, params, iterOpts)` | offset pages of `data` | iterator over `Flips` |
| `Snapshot(ctx, params)` | `GET /api/v1/market/snapshot` | Market snapshot for a geography |
| `ZillowContext(ctx, params)` | `GET /api/v1/market/zillow/context` | Get qualified regional Zillow context for a property |
| `ZillowTimeseries(ctx, params)` | `GET /api/v1/market/zillow/timeseries` | Get one provider region monthly series |
| `CompareZillowMarkets(ctx, params)` | `GET /api/v1/market/zillow/compare` | Compare explicit provider regions at one common period |

### client.Owners

| Method | HTTP | Summary |
|---|---|---|
| `Search(ctx, params)` | `GET /api/v1/owners/search` | Search property owners |
| `Get(ctx, name, params)` | `GET /api/v1/owners/{name}` | Get owner profile |
| `Properties(ctx, name, params)` | `GET /api/v1/owners/{name}/properties` | Get owner's properties |
| `PropertiesIter(ctx, name, params, iterOpts)` | offset pages of `data` | iterator over `Properties` |
| `Portfolio(ctx, name, params)` | `GET /api/v1/owners/{name}/portfolio` | Get owner portfolio summary |
| `Report(ctx, name, params)` | `GET /api/v1/owners/{name}/report` | Owner intelligence report (paid, priced per resolution; account required) — with a free preview |
| `Transactions(ctx, name, params)` | `GET /api/v1/owners/{name}/transactions` | Recorded deed transactions for an owner |
| `Card(ctx, params)` | `GET /api/v1/owners/card` | Owner card -- the owner of record and their mailing contact (account required) |

### client.Webhooks

| Method | HTTP | Summary |
|---|---|---|
| `List(ctx, params)` | `GET /api/v1/webhooks` | List webhook endpoints |
| `Create(ctx, params)` | `POST /api/v1/webhooks` | Create a webhook endpoint |
| `Get(ctx, id, params)` | `GET /api/v1/webhooks/{id}` | Get a single webhook endpoint |
| `Delete(ctx, id, params)` | `DELETE /api/v1/webhooks/{id}` | Soft-disable a webhook endpoint |
| `Deliveries(ctx, id, params)` | `GET /api/v1/webhooks/{id}/deliveries` | Recent delivery attempts for a webhook |
| `RetryDelivery(ctx, id, deliveryID, params)` | `POST /api/v1/webhooks/{id}/deliveries/{deliveryId}/retry` | Re-queue a failed webhook delivery |

### client.Account

| Method | HTTP | Summary |
|---|---|---|
| `Usage(ctx, params)` | `GET /api/v1/account/usage` | Current-period usage and quota |

### client.Storefront

| Method | HTTP | Summary |
|---|---|---|
| `Catalog(ctx, params)` | `GET /api/v1/storefront/catalog` | Machine Storefront — sealed field catalog |
| `Availability(ctx, params)` | `GET /api/v1/storefront/availability` | Machine Storefront -- try-before-buy (jurisdiction coverage or per-parcel quote) |

### client.Leads

| Method | HTTP | Summary |
|---|---|---|
| `Find(ctx, params)` | `GET /api/v1/leads/find` | Lead feed (paid, priced per lead) — with a FREE preview |

### client.Credits

| Method | HTTP | Summary |
|---|---|---|
| `Topup(ctx, params)` | `GET /api/v1/storefront/credits/topup` | Fund a prepaid credit balance over x402 |
| `Balance(ctx, params)` | `GET /api/v1/storefront/credits/balance` | Read a prepaid credit balance + ledger |

### client.Watch

| Method | HTTP | Summary |
|---|---|---|
| `List(ctx, params)` | `GET /api/v1/watch` | List your watches |
| `Create(ctx, params)` | `POST /api/v1/watch` | Create a watch (free) |
| `Poll(ctx, id, params)` | `GET /api/v1/watch/{id}` | Poll a watch for new changes (priced per delta) |
| `Delete(ctx, id, params)` | `DELETE /api/v1/watch/{id}` | Delete a watch |

### client.Verify

| Method | HTTP | Summary |
|---|---|---|
| `Get(ctx, params)` | `GET /api/v1/verify` | Verify facts for one parcel |
| `Batch(ctx, params)` | `POST /api/v1/verify` | Verify facts (batch, paid per lookup) - FREE preview |

### client.Cohorts

| Method | HTTP | Summary |
|---|---|---|
| `Export(ctx, id, params)` | `GET /api/v1/cohorts/{id}/export` | Mail-merge export of one of your lists (account required; included for subscribers, per row otherwise) |
| `List(ctx, params)` | `GET /api/v1/cohorts` | List your saved parcel lists (cohorts) |

### client.Lookup

| Method | HTTP | Summary |
|---|---|---|
| `Get(ctx, params)` | `GET /api/v1/lookup` | Exact parcel lookup (UUID or APN) |
| `Batch(ctx, params)` | `POST /api/v1/lookup/batch` | Resolve up to 500 parcel queries in one call |

### client.CMBS

| Method | HTTP | Summary |
|---|---|---|
| `Exposure(ctx, params)` | `GET /api/v1/cmbs/exposure` | CMBS loan exposure for a parcel or an owner |

### client.Freshness

| Method | HTTP | Summary |
|---|---|---|
| `Get(ctx, params)` | `GET /api/v1/freshness` | How fresh the served parcel snapshot is |
| `Datasets(ctx, params)` | `GET /api/v1/freshness/datasets` | Per-dataset availability and freshness |

### client.Crime

| Method | HTTP | Summary |
|---|---|---|
| `Lookup(ctx, params)` | `GET /api/v1/crime/lookup` | Crime score near a point |

### client.Traffic

| Method | HTTP | Summary |
|---|---|---|
| `Stations(ctx, params)` | `GET /api/v1/traffic/stations` | Traffic count stations in a bounding box |

### client.Licensees

| Method | HTTP | Summary |
|---|---|---|
| `Firms(ctx, params)` | `GET /api/v1/licensees/firms` | Search licensed firms in a place |

### client.Intelligence

| Method | HTTP | Summary |
|---|---|---|
| `Signals(ctx, id, params)` | `GET /api/v1/parcels/{id}/signals` | Get evidence-backed property signals |
| `Run(ctx, runID, params)` | `GET /api/v1/intelligence/runs/{runId}` | Read an owned retained run and evidence |
| `CreateScenario(ctx, params)` | `POST /api/v1/intelligence/scenarios` | Save an explicit named residual scenario |
| `Handoff(ctx, runID, params)` | `GET /api/v1/intelligence/runs/{runId}/handoff` | Prepare an owned structured investigation handoff |
<!-- END GENERATED METHODS -->

## Regenerating from the spec

The generated layer (`gen_*.go` plus the method table above) comes from `openapi.json`:

```bash
curl -fsSL https://propraven.com/openapi.json -o openapi.json   # or copy a spec file
go run ./internal/cmd/generate
go vet ./... && go test -race ./...
```

The generator uses only the standard library and is deterministic. Running it twice produces no diff, and CI fails if the committed files are stale. The `regenerate` workflow runs these steps every day and opens a PR on the `spec-sync` branch when the published spec changes. The hand-written core (`client.go`, `transport.go`, `errors.go`, `pagination.go`, `lenient.go`, `webhook.go`, `ptr.go`) is never generated.

`internal/cmd/livesmoke` makes 12 or fewer real read-only calls against production (`PROPRAVEN_API_KEY=pz_... go run ./internal/cmd/livesmoke`). CI never runs it.

## Versioning

Versions are tagged `vX.Y.Z`, and `internal.PackageVersion` must match the tag. Until v1.0.0, a minor version may break the API surface (see [CHANGELOG.md](./CHANGELOG.md)).

## License

MIT. See [LICENSE](./LICENSE).
