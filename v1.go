// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/jdw2111/propraven-go/internal/apijson"
	"github.com/jdw2111/propraven-go/internal/apiquery"
	"github.com/jdw2111/propraven-go/internal/requestconfig"
	"github.com/jdw2111/propraven-go/option"
	"github.com/jdw2111/propraven-go/packages/param"
	"github.com/jdw2111/propraven-go/packages/respjson"
)

// Data coverage statistics.
//
// V1Service contains methods and other services that help with interacting with
// the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	options []option.RequestOption
	// Parcel lookup, owner details, permits, deeds, and risk data.
	Parcels V1ParcelService
	// Geographic and filtered parcel search.
	Search V1SearchService
	// Deal sourcing: absentee owners, property flips.
	Deals V1DealService
	// County-level market statistics and trends.
	Market V1MarketService
	// Owner search, profiles, and portfolios.
	Owners V1OwnerService
	// Webhook subscriptions and delivery history. Manage which events PropRaven pushes
	// to your endpoints.
	Webhooks V1WebhookService
	// Account-scoped usage, quota, and key-level reporting.
	Account V1AccountService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.options = opts
	r.Parcels = NewV1ParcelService(opts...)
	r.Search = NewV1SearchService(opts...)
	r.Deals = NewV1DealService(opts...)
	r.Market = NewV1MarketService(opts...)
	r.Owners = NewV1OwnerService(opts...)
	r.Webhooks = NewV1WebhookService(opts...)
	r.Account = NewV1AccountService(opts...)
	return
}

// Retrieve parcel coverage statistics at the state or county level.
func (r *V1Service) GetCoverage(ctx context.Context, query V1GetCoverageParams, opts ...option.RequestOption) (res *V1GetCoverageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1GetCoverageResponse struct {
	Data          []V1GetCoverageResponseData `json:"data"`
	StatesCovered int64                       `json:"states_covered"`
	TotalParcels  int64                       `json:"total_parcels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		StatesCovered respjson.Field
		TotalParcels  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GetCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1GetCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GetCoverageResponseData struct {
	CountyFips  string  `json:"county_fips" api:"nullable"`
	CountyName  string  `json:"county_name" api:"nullable"`
	GeocodedPct float64 `json:"geocoded_pct"`
	OwnerPct    float64 `json:"owner_pct"`
	ParcelCount int64   `json:"parcel_count"`
	StateFips   string  `json:"state_fips"`
	StateName   string  `json:"state_name"`
	ValuePct    float64 `json:"value_pct"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountyFips  respjson.Field
		CountyName  respjson.Field
		GeocodedPct respjson.Field
		OwnerPct    respjson.Field
		ParcelCount respjson.Field
		StateFips   respjson.Field
		StateName   respjson.Field
		ValuePct    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GetCoverageResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1GetCoverageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GetCoverageParams struct {
	// State FIPS code or abbreviation to filter coverage to a specific state and
	// return county-level breakdown.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1GetCoverageParams]'s query parameters as `url.Values`.
func (r V1GetCoverageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
