# propraven-go

The Go SDK for [PropRaven](https://propraven.com) — national property intelligence over 230M+ U.S. parcels.

Generated from the [OpenAPI 3.1 spec](https://api.propraven.com/openapi.json) via [Stainless](https://www.stainless.com). Webhook signature verification is hand-maintained alongside the generated code.

```bash
go get github.com/jdw2111/propraven-go@latest
```

## Documentation

- Developer hub: https://propraven.com/developers
- Hosted MCP server (31 read-only tools): https://propraven.com/docs/mcp
- REST API v1 reference: https://propraven.com/docs/v1

## Quickstart

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jdw2111/propraven-go"
    "github.com/jdw2111/propraven-go/option"
)

func main() {
    client := propraven.NewClient(
        option.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
    )

    ctx := context.Background()

    parcel, err := client.V1.Parcels.Get(ctx, "06037:1234-567-890")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", parcel.ParcelID)
}
```

`PROPRAVEN_API_KEY` and `PROPRAVEN_BASE_URL` are also read from the environment automatically by `NewClient`.

## API surface

All endpoints sit under `client.V1.*`:

- `client.V1.Parcels.Get(ctx, id)` — single parcel
- `client.V1.Parcels.GetOwner(ctx, id)` — current owner
- `client.V1.Parcels.GetPermits(ctx, id, params)` — permits on a parcel
- `client.V1.Parcels.GetDeeds(ctx, id, params)` — recorded deeds
- `client.V1.Search.Parcels(ctx, params)` — geo + filter search
- `client.V1.Owners.Get(ctx, entityID)` — consolidated owner entity
- `client.V1.Owners.GetPortfolio(ctx, entityID, params)` — owner's parcels
- `client.V1.Deals.Absentee(ctx, params)` / `.Flips(ctx, params)` — deal sourcing
- `client.V1.Market.GetCounty(ctx, params)` — county-level stats
- `client.V1.Webhooks.NewEndpoint(ctx, body)` / `ListEndpoints(ctx)` / `DisableEndpoint(ctx, id)` / `GetDeliveries(ctx, id)`
- `client.V1.Account.Usage(ctx)` — your API key's usage + quota
- `client.V1.GetCoverage(ctx, params)` — state/county data coverage

Full reference: [pkg.go.dev/github.com/jdw2111/propraven-go](https://pkg.go.dev/github.com/jdw2111/propraven-go).

## Webhook verification

`propraven.VerifyWebhook` does constant-time HMAC-SHA256 verification with a 5-minute replay window. Use it in your inbound handler:

```go
import "github.com/jdw2111/propraven-go"

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    sig := r.Header.Get("X-PropRaven-Signature")

    if err := propraven.VerifyWebhook(sig, body, os.Getenv("PROPRAVEN_WEBHOOK_SECRET")); err != nil {
        http.Error(w, "invalid signature", http.StatusUnauthorized)
        return
    }
    // Process the event…
    w.WriteHeader(http.StatusOK)
}
```

Tune the replay window with `propraven.VerifyWebhookOptions{MaxClockSkew: ...}`.

## Error handling

Non-2xx responses unwrap to `*propraven.Error`:

```go
parcel, err := client.V1.Parcels.Get(ctx, "missing")
if err != nil {
    var pe *propraven.Error
    if errors.As(err, &pe) {
        log.Printf("status=%d request_id=%s", pe.StatusCode, pe.Header.Get("X-Request-Id"))
    }
}
```

The client retries on 429 and idempotent 5xx by default. Disable via `option.WithMaxRetries(0)`.

## Options

```go
client := propraven.NewClient(
    option.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
    option.WithBaseURL("https://api.propraven.com"),
    option.WithMaxRetries(3),
    option.WithRequestTimeout(60 * time.Second),
    option.WithHeader("X-Trace-Id", "..."),
)
```

## Versioning

This SDK follows the version of the underlying OpenAPI spec; minor versions add resources without breaking existing signatures, and v1.0 freezes the surface. Track the [changelog](https://propraven.com/changelog) for spec changes.

## License

MIT — see [LICENSE](./LICENSE).
