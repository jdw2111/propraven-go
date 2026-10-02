# Changelog

## v0.3.0 (unreleased)

v0.3.0 rewrites the SDK. The Stainless-generated client is gone. The SDK now has a small hand-written core (transport, retries, errors, pagination, webhooks) and a layer generated in-repo from the vendored `openapi.json` by `go run ./internal/cmd/generate`. The new layer covers all 69 operations of the v1 API, including the 32 that v0.2.0 lacked.

### Breaking changes from v0.2.0

- **Client and namespaces.** Namespaces are fields directly on the client and follow the spec's `x-sdk-group`. `client.V1.Parcels.Get(ctx, id)` becomes `client.Parcels.Get(ctx, id, params)`. `client.V1.GetCoverage(ctx, params)` becomes `client.Coverage.Get(ctx, params)`. Method names follow `x-sdk-method`: `GetOwner` becomes `Owner`, `GetPermits` becomes `Permits`, `NewEndpoint` becomes `Create`, and so on.
- **Signatures.** Every method takes `(ctx, pathParams..., params *XParams, opts ...RequestOption)`. The single params struct holds query, header and JSON-body fields, and it may be nil. Optional fields are plain pointers (`*string`, `*int64`, `*float64`, `*bool`); set them with `propraven.String/Int/Float/Bool/Ptr`. This replaces `param.Field`/`param.Opt`.
- **Options.** The `option` package is gone. Options now live in the root package: `propraven.WithAPIKey`, `WithBaseURL`, `WithHTTPClient`, `WithMaxRetries`, `WithTimeout` (replaces `option.WithRequestTimeout`), `WithHeader`, `WithLogger` and `WithRawResponse`. All of them work per client and per request.
- **Numbers are numbers.** Numeric fields are `float64`/`int64`, not strings. Identifiers (`parcel_id`, `apn`, `county_fips`, `state_fips`, `zip`) stay strings. During the API's serializer rollout the decoder also accepts quoted decimals.
- **Errors.** `*propraven.Error` is parsed from RFC 7807 problem+json. It has `StatusCode`, `Type`, `Title`, `Detail`, `Code`, `Errors`, `RequestID`, `RetryAfter`, `Accepts` (x402), `Header` and `Body`. New helpers: `IsNotFound`, `IsRateLimited`, `IsPaymentRequired`, `IsAuthentication`, `IsBadRequest`, `IsPermissionDenied`, `IsConflict`, `IsServerError`, `IsConnectionError`, `IsTimeout`, `AsError`. When no response arrives, the call returns `*propraven.ConnectionError`.
- **Webhooks.** The signature is now `VerifyWebhook(payload []byte, signature, secret string, opts ...WebhookOption) (json.RawMessage, error)`. Payload comes first, the call returns the verified event, every failure wraps `ErrWebhookSignature`, the default tolerance is 300 s (`WithWebhookTolerance`), and multiple `v1=` entries are accepted. `VerifyWebhookOptions` is removed.
- **Removed packages.** `option`, `packages/param`, `packages/respjson`, `shared/constant`, the `aliases.go` re-exports and all `internal/api*` encoders. The module now has no third-party dependencies (`tidwall/gjson` and `tidwall/sjson` are dropped).
- **Response types.** Every operation has a named `<Namespace><Method>Response` type. Responses that are a component schema are aliases of it (`ParcelsGetResponse = Parcel`). Bare-array-or-envelope responses (`Parcels.Permits`, `Parcels.Deeds`) are unions with `Array` / `Object` fields. Other unions and untyped schemas are `json.RawMessage`.

### Added

- Retries on network errors, timeouts, 429, 503 and 504 (and other 5xx for idempotent methods). They honour `Retry-After` and `X-RateLimit-Reset`, and never wait longer than 60 s.
- An `XIter` method on every paginated operation, with `IterOptions{PageSize, MaxItems}`. Covers offset (query or body) and cursor (`Search.FullIter`) pagination.
- `client.LastRateLimit()`, `WithRawResponse`, `SignWebhook` (for testing handlers), and a `pz_` key-prefix warning.
- `internal/cmd/generate` (generator), `internal/cmd/livesmoke` (manual production smoke test), a CI check that generated code is current, and a daily `regenerate` workflow that opens a `spec-sync` PR.
