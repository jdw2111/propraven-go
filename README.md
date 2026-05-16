# propraven-go

The Go SDK for [PropRaven](https://propraven.com) — national property intelligence over 230M+ U.S. parcels.

```bash
go get github.com/jdw2111/propraven-go@latest
```

## Quickstart

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    propraven "github.com/jdw2111/propraven-go"
)

func main() {
    client, err := propraven.NewClient(
        propraven.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    parcel, err := client.Parcels.Get(ctx, "06037:1234-567-890")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", parcel)

    // List all parcels in a county — Iter ranges over every page.
    for p, err := range client.Parcels.Iter(ctx, propraven.ListOptions{
        StateFIPS:  "06",
        CountyFIPS: "037",
    }) {
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(p.ParcelID, p.OwnerName)
    }
}
```

## Webhook verification

Every PropRaven webhook delivery carries an `X-PropRaven-Signature` header. Verify it server-side:

```go
import propraven "github.com/jdw2111/propraven-go"

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    sig := r.Header.Get("X-PropRaven-Signature")

    if err := propraven.VerifyWebhook(sig, body, os.Getenv("PROPRAVEN_WEBHOOK_SECRET")); err != nil {
        http.Error(w, "invalid signature", http.StatusUnauthorized)
        return
    }

    // Process the event...
    w.WriteHeader(http.StatusOK)
}
```

Verification is constant-time, enforces a 5-minute replay window by default, and protects against tampered bodies.

## Error handling

Every non-2xx response unwraps to a typed `*propraven.Error`:

```go
parcel, err := client.Parcels.Get(ctx, "missing")
if err != nil {
    var pe *propraven.Error
    if errors.As(err, &pe) {
        fmt.Printf("Status: %d, Code: %s, RequestID: %s\n", pe.Status, pe.Code, pe.RequestID)
    }
    if propraven.IsNotFound(err) { /* ... */ }
    if propraven.IsRateLimited(err) { /* back off and retry */ }
}
```

Always quote `RequestID` in support tickets — it's the join key for our logs.

## Resources covered in v0.1

- `client.Parcels.Get(ctx, parcelID)`
- `client.Parcels.List(ctx, ListOptions{...})` + `.Iter(...)` for ranging
- `client.Owners.Get(ctx, ownerEntityID)`
- `client.Owners.Search(ctx, OwnerSearchOptions{...})`
- `client.Deeds.Search(ctx, DeedSearchOptions{...})`
- `client.Permits.Search(ctx, PermitSearchOptions{...})`
- `client.Webhooks.Create/List/Delete(...)`
- `client.Health.Check(ctx)`
- `propraven.VerifyWebhook(sig, body, secret)`

Additional endpoints (coverage, deals, market, audit) ship over v0.x as customer demand surfaces. Full REST is always usable directly through `WithHTTPClient(...)` if you need an endpoint we haven't wrapped yet.

## Options

```go
client, _ := propraven.NewClient(
    propraven.WithAPIKey(os.Getenv("PROPRAVEN_API_KEY")),
    propraven.WithBaseURL("https://api.propraven.com"),       // override for self-host
    propraven.WithTimeout(60 * time.Second),                  // default 30s
    propraven.WithHTTPClient(myInstrumentedClient),           // inject OTel transport
    propraven.WithUserAgent("acme-pricing/1.0"),              // identifies your app
)
```

## Source of truth

This SDK is hand-crafted against the [OpenAPI 3.1 spec](https://api.propraven.com/openapi.json). Both files are co-versioned; SDK releases tag in sync with breaking spec changes.

## Status

v0.1.x — under active development; minor versions add resources without breaking existing signatures. v1.0 freezes the surface.

## License

MIT — see [LICENSE](./LICENSE).
