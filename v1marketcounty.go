// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"errors"
	"fmt"
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
// V1MarketCountyService contains methods and other services that help with
// interacting with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1MarketCountyService] method instead.
type V1MarketCountyService struct {
	options []option.RequestOption
}

// NewV1MarketCountyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1MarketCountyService(opts ...option.RequestOption) (r V1MarketCountyService) {
	r = V1MarketCountyService{}
	r.options = opts
	return
}

// Returns the full county profile: quarterly market stats (sale count, median
// price, YoY change, days on market), affordability index by year, parcel summary
// (count, avg assessed value), and flip activity. Use for county-detail
// dashboards.
func (r *V1MarketCountyService) GetDetail(ctx context.Context, fips string, opts ...option.RequestOption) (res *V1MarketCountyGetDetailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if fips == "" {
		err = errors.New("missing required fips parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/market/counties/%s", url.PathEscape(fips))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve real estate market statistics aggregated at the county level, including
// sale counts, median prices, and year-over-year changes.
func (r *V1MarketCountyService) GetStatistics(ctx context.Context, query V1MarketCountyGetStatisticsParams, opts ...option.RequestOption) (res *V1MarketCountyGetStatisticsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/market/counties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1MarketCountyGetDetailResponse struct {
	Affordability []AffordabilityRow                           `json:"affordability"`
	FlipSummary   V1MarketCountyGetDetailResponseFlipSummary   `json:"flip_summary"`
	MarketStats   []V1MarketCountyGetDetailResponseMarketStat  `json:"market_stats"`
	ParcelSummary V1MarketCountyGetDetailResponseParcelSummary `json:"parcel_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Affordability respjson.Field
		FlipSummary   respjson.Field
		MarketStats   respjson.Field
		ParcelSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketCountyGetDetailResponse) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetDetailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetDetailResponseFlipSummary struct {
	AvgHoldDays float64 `json:"avg_hold_days" api:"nullable"`
	AvgRoi      float64 `json:"avg_roi" api:"nullable"`
	FlipCount   int64   `json:"flip_count"`
	TotalProfit float64 `json:"total_profit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgHoldDays respjson.Field
		AvgRoi      respjson.Field
		FlipCount   respjson.Field
		TotalProfit respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketCountyGetDetailResponseFlipSummary) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetDetailResponseFlipSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetDetailResponseMarketStat struct {
	// Average days on market.
	AvgDom          float64 `json:"avg_dom" api:"nullable"`
	AvgSalePrice    float64 `json:"avg_sale_price" api:"nullable"`
	CountyFips      string  `json:"county_fips"`
	MedianSalePrice float64 `json:"median_sale_price" api:"nullable"`
	PriceYoyPct     float64 `json:"price_yoy_pct" api:"nullable"`
	Quarter         string  `json:"quarter"`
	SaleCount       int64   `json:"sale_count"`
	StateFips       string  `json:"state_fips"`
	TotalVolume     float64 `json:"total_volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgDom          respjson.Field
		AvgSalePrice    respjson.Field
		CountyFips      respjson.Field
		MedianSalePrice respjson.Field
		PriceYoyPct     respjson.Field
		Quarter         respjson.Field
		SaleCount       respjson.Field
		StateFips       respjson.Field
		TotalVolume     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketCountyGetDetailResponseMarketStat) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetDetailResponseMarketStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetDetailResponseParcelSummary struct {
	AvgAssessedValue float64 `json:"avg_assessed_value" api:"nullable"`
	ParcelCount      int64   `json:"parcel_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgAssessedValue respjson.Field
		ParcelCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketCountyGetDetailResponseParcelSummary) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetDetailResponseParcelSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetStatisticsResponse struct {
	Data   []V1MarketCountyGetStatisticsResponseData `json:"data"`
	Limit  int64                                     `json:"limit"`
	Offset int64                                     `json:"offset"`
	Total  int64                                     `json:"total"`
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
func (r V1MarketCountyGetStatisticsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetStatisticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetStatisticsResponseData struct {
	AvgDaysOnMarket int64   `json:"avg_days_on_market"`
	AvgPrice        float64 `json:"avg_price"`
	CountyFips      string  `json:"county_fips"`
	CountyName      string  `json:"county_name"`
	MedianPrice     float64 `json:"median_price"`
	Quarter         string  `json:"quarter"`
	SaleCount       int64   `json:"sale_count"`
	StateAbbr       string  `json:"state_abbr"`
	StateFips       string  `json:"state_fips"`
	// Year-over-year median price change as a decimal (e.g., 0.05 = 5%).
	YoyChange float64 `json:"yoy_change"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgDaysOnMarket respjson.Field
		AvgPrice        respjson.Field
		CountyFips      respjson.Field
		CountyName      respjson.Field
		MedianPrice     respjson.Field
		Quarter         respjson.Field
		SaleCount       respjson.Field
		StateAbbr       respjson.Field
		StateFips       respjson.Field
		YoyChange       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1MarketCountyGetStatisticsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1MarketCountyGetStatisticsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1MarketCountyGetStatisticsParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum number of sales in the period to include a county.
	MinSales param.Opt[int64] `query:"min_sales,omitzero" json:"-"`
	Offset   param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Specific quarter to retrieve (e.g., 2025Q4). Defaults to latest available.
	Quarter param.Opt[string] `query:"quarter,omitzero" json:"-"`
	// Filter by state FIPS code.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// Any of "asc", "desc".
	Order V1MarketCountyGetStatisticsParamsOrder `query:"order,omitzero" json:"-"`
	// Sort field.
	//
	// Any of "sale_count", "median_price", "yoy_change", "county_name".
	Sort V1MarketCountyGetStatisticsParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1MarketCountyGetStatisticsParams]'s query parameters as
// `url.Values`.
func (r V1MarketCountyGetStatisticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1MarketCountyGetStatisticsParamsOrder string

const (
	V1MarketCountyGetStatisticsParamsOrderAsc  V1MarketCountyGetStatisticsParamsOrder = "asc"
	V1MarketCountyGetStatisticsParamsOrderDesc V1MarketCountyGetStatisticsParamsOrder = "desc"
)

// Sort field.
type V1MarketCountyGetStatisticsParamsSort string

const (
	V1MarketCountyGetStatisticsParamsSortSaleCount   V1MarketCountyGetStatisticsParamsSort = "sale_count"
	V1MarketCountyGetStatisticsParamsSortMedianPrice V1MarketCountyGetStatisticsParamsSort = "median_price"
	V1MarketCountyGetStatisticsParamsSortYoyChange   V1MarketCountyGetStatisticsParamsSort = "yoy_change"
	V1MarketCountyGetStatisticsParamsSortCountyName  V1MarketCountyGetStatisticsParamsSort = "county_name"
)
