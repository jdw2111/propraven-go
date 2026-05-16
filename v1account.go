// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/jdw2111/propraven-go/internal/apijson"
	"github.com/jdw2111/propraven-go/internal/requestconfig"
	"github.com/jdw2111/propraven-go/option"
	"github.com/jdw2111/propraven-go/packages/respjson"
)

// Account-scoped usage, quota, and key-level reporting.
//
// V1AccountService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountService] method instead.
type V1AccountService struct {
	options []option.RequestOption
}

// NewV1AccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1AccountService(opts ...option.RequestOption) (r V1AccountService) {
	r = V1AccountService{}
	r.options = opts
	return
}

// Returns the calling key's current-period API usage, included allotment,
// remaining calls, configured per-minute and per-day rate limits, and the hard-cap
// status. Per-user (all of a user's API keys roll up to the same monthly counter,
// since they share a Stripe subscription).
func (r *V1AccountService) GetUsage(ctx context.Context, opts ...option.RequestOption) (res *V1AccountGetUsageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/account/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V1AccountGetUsageResponse struct {
	// Plan allotment for the current period.
	CallsIncluded  int64 `json:"calls_included"`
	CallsRemaining int64 `json:"calls_remaining"`
	CallsUsed      int64 `json:"calls_used"`
	// Whether further calls will be hard-rejected vs allowed-and-billed.
	HardCapped bool                               `json:"hard_capped"`
	Period     V1AccountGetUsageResponsePeriod    `json:"period"`
	RateLimit  V1AccountGetUsageResponseRateLimit `json:"rate_limit"`
	// Any of "free", "starter", "pro", "scale", "api_100k".
	Tier V1AccountGetUsageResponseTier `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallsIncluded  respjson.Field
		CallsRemaining respjson.Field
		CallsUsed      respjson.Field
		HardCapped     respjson.Field
		Period         respjson.Field
		RateLimit      respjson.Field
		Tier           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetUsageResponsePeriod struct {
	End   time.Time `json:"end" format:"date-time"`
	Label string    `json:"label"`
	Start time.Time `json:"start" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Label       respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetUsageResponsePeriod) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetUsageResponsePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetUsageResponseRateLimit struct {
	PerDay    int64 `json:"per_day"`
	PerMinute int64 `json:"per_minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PerDay      respjson.Field
		PerMinute   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetUsageResponseRateLimit) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetUsageResponseRateLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetUsageResponseTier string

const (
	V1AccountGetUsageResponseTierFree    V1AccountGetUsageResponseTier = "free"
	V1AccountGetUsageResponseTierStarter V1AccountGetUsageResponseTier = "starter"
	V1AccountGetUsageResponseTierPro     V1AccountGetUsageResponseTier = "pro"
	V1AccountGetUsageResponseTierScale   V1AccountGetUsageResponseTier = "scale"
	V1AccountGetUsageResponseTierAPI100k V1AccountGetUsageResponseTier = "api_100k"
)
