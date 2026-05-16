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

// County-level market statistics and trends.
//
// V1MarketService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1MarketService] method instead.
type V1MarketService struct {
	options []option.RequestOption
	// County-level market statistics and trends.
	Counties V1MarketCountyService
}

// NewV1MarketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1MarketService(opts ...option.RequestOption) (r V1MarketService) {
	r = V1MarketService{}
	r.options = opts
	r.Counties = NewV1MarketCountyService(opts...)
	return
}

// Aggregated flip activity per county: count, average ROI, average hold days,
// total profit. Use for surfacing the hottest flip markets. Differs from
// /api/v1/deals/flips which returns the underlying transactions.
func (r *V1MarketService) GetFlipActivity(ctx context.Context, query V1MarketGetFlipActivityParams, opts ...option.RequestOption) (res *V1MarketGetFlipActivityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/market/flips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve quarterly time series of market metrics for one or more counties or a
// state.
func (r *V1MarketService) GetTrends(ctx context.Context, query V1MarketGetTrendsParams, opts ...option.RequestOption) (res *V1MarketGetTrendsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/market/trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1MarketGetFlipActivityResponse struct {
	Data   []V1MarketGetFlipActivityResponseData `json:"data"`
	Limit  int64                                 `json:"limit"`
	Offset int64                                 `json:"offset"`
	Total  int64                                 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketGetFlipActivityResponse) RawJSON() string { return r.JSON.raw }
func (r *V1MarketGetFlipActivityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketGetFlipActivityResponseData struct {
	AvgHoldDays float64 `json:"avg_hold_days" api:"nullable"`
	// Average profit percentage (e.g. 0.18 = 18%).
	AvgRoi      float64 `json:"avg_roi" api:"nullable"`
	CountyFips  string  `json:"county_fips"`
	FlipCount   int64   `json:"flip_count"`
	TotalProfit float64 `json:"total_profit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgHoldDays respjson.Field
		AvgRoi      respjson.Field
		CountyFips  respjson.Field
		FlipCount   respjson.Field
		TotalProfit respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketGetFlipActivityResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1MarketGetFlipActivityResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketGetTrendsResponse struct {
	Data []V1MarketGetTrendsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketGetTrendsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1MarketGetTrendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketGetTrendsResponseData struct {
	CountyFips string                                 `json:"county_fips"`
	CountyName string                                 `json:"county_name"`
	Quarters   []V1MarketGetTrendsResponseDataQuarter `json:"quarters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountyFips  respjson.Field
		CountyName  respjson.Field
		Quarters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketGetTrendsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1MarketGetTrendsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketGetTrendsResponseDataQuarter struct {
	AvgPrice    float64 `json:"avg_price"`
	MedianPrice float64 `json:"median_price"`
	Quarter     string  `json:"quarter"`
	SaleCount   int64   `json:"sale_count"`
	YoyChange   float64 `json:"yoy_change"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgPrice    respjson.Field
		MedianPrice respjson.Field
		Quarter     respjson.Field
		SaleCount   respjson.Field
		YoyChange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketGetTrendsResponseDataQuarter) RawJSON() string { return r.JSON.raw }
func (r *V1MarketGetTrendsResponseDataQuarter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketGetFlipActivityParams struct {
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 2-digit state FIPS filter.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1MarketGetFlipActivityParams]'s query parameters as
// `url.Values`.
func (r V1MarketGetFlipActivityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1MarketGetTrendsParams struct {
	// Comma-separated list of county FIPS codes.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	// State FIPS code. Used if county_fips is not provided.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1MarketGetTrendsParams]'s query parameters as
// `url.Values`.
func (r V1MarketGetTrendsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
