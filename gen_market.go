// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// MarketService groups the market operations. Use it as client.Market.
type MarketService struct {
	client *Client
}

// Counties: Get county market statistics
//
// Retrieve real estate market statistics aggregated at the county level, including sale counts,
// median prices, and year-over-year changes. API key optional (county aggregates, no person-level
// fields): anonymous callers are rate-limited per IP at the free tier; a present but invalid key
// is a 401; keyed calls are metered.
//
// HTTP: GET /api/v1/market/counties
func (s *MarketService) Counties(ctx context.Context, params *MarketCountiesParams, opts ...RequestOption) (*MarketCountiesResponse, error) {
	var out MarketCountiesResponse
	if err := s.client.do(ctx, buildMarketCountiesRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// CountiesIter iterates every item of [MarketService.Counties] across pages (offset pagination
// over "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *MarketService) CountiesIter(ctx context.Context, params *MarketCountiesParams, iter IterOptions, opts ...RequestOption) *Iter[MarketCountiesResponseData] {
	var p MarketCountiesParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[MarketCountiesResponseData](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildMarketCountiesRequest(&q)
	}, opts)
}

func buildMarketCountiesRequest(params *MarketCountiesParams) *apiRequest {
	req := newRequest("GET", "/api/v1/market/counties")
	if params != nil {
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "min_sales", params.MinSales)
		addQuery(req.query, "quarter", params.Quarter)
		addQuery(req.query, "sort", params.Sort)
		addQuery(req.query, "order", params.Order)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// MarketCountiesParams holds the query, header and JSON-body parameters of
// [MarketService.Counties]. Pass nil when you need none.
type MarketCountiesParams struct {
	// Filter by state FIPS code.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Minimum number of sales in the period to include a county.
	MinSales *int64 `query:"min_sales" json:"-"`

	// Exact quarter, e.g. 2025Q4; any other format is a 400.
	Quarter *string `query:"quarter" json:"-"`

	// Sort column. The former spellings median_price and yoy_change are accepted as deprecated
	// aliases; any other value is a 400.
	Sort *MarketCountiesParamsSort `query:"sort" json:"-"`

	// Sort direction (case-insensitive); any other value is a 400.
	Order  *MarketCountiesParamsOrder `query:"order" json:"-"`
	Limit  *int64                     `query:"limit" json:"-"`
	Offset *int64                     `query:"offset" json:"-"`
}

// MarketCountiesParamsSort is generated from the OpenAPI spec. It is a string; the
// MarketCountiesParamsSort* constants list the documented values.
type MarketCountiesParamsSort = string

// Documented values of MarketCountiesParamsSort.
const (
	MarketCountiesParamsSortMedianSalePrice MarketCountiesParamsSort = "median_sale_price"
	MarketCountiesParamsSortSaleCount       MarketCountiesParamsSort = "sale_count"
	MarketCountiesParamsSortTotalVolume     MarketCountiesParamsSort = "total_volume"
	MarketCountiesParamsSortPriceYoyPct     MarketCountiesParamsSort = "price_yoy_pct"
)

// MarketCountiesParamsOrder is generated from the OpenAPI spec. It is a string; the
// MarketCountiesParamsOrder* constants list the documented values.
type MarketCountiesParamsOrder = string

// Documented values of MarketCountiesParamsOrder.
const (
	MarketCountiesParamsOrderAsc  MarketCountiesParamsOrder = "asc"
	MarketCountiesParamsOrderDesc MarketCountiesParamsOrder = "desc"
)

// MarketCountiesResponse: Get county market statistics
type MarketCountiesResponse struct {
	Data    []MarketCountiesResponseData  `json:"data"`
	Summary MarketCountiesResponseSummary `json:"summary"`
	Total   int64                         `json:"total"`
	Limit   int64                         `json:"limit"`
	Offset  int64                         `json:"offset"`
}

// UnmarshalJSON decodes MarketCountiesResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketCountiesResponse) UnmarshalJSON(data []byte) error {
	type plain MarketCountiesResponse
	aux := struct {
		*plain
		Total  lenientNumber[int64] `json:"total"`
		Limit  lenientNumber[int64] `json:"limit"`
		Offset lenientNumber[int64] `json:"offset"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assign(&r.Total)
	aux.Limit.assign(&r.Limit)
	aux.Offset.assign(&r.Offset)
	return softTypeError(err)
}

// MarketCountiesResponseData is generated from the OpenAPI spec.
type MarketCountiesResponseData struct {
	CountyFIPS      string    `json:"county_fips"`
	StateFIPS       string    `json:"state_fips"`
	CountyName      *string   `json:"county_name,omitempty"`
	State           *string   `json:"state,omitempty"`
	RefreshedAt     *string   `json:"refreshed_at,omitempty"`
	Quarter         *string   `json:"quarter,omitempty"`
	SaleCount       *int64    `json:"sale_count,omitempty"`
	MedianSalePrice *float64  `json:"median_sale_price,omitempty"`
	AvgSalePrice    *float64  `json:"avg_sale_price,omitempty"`
	TotalVolume     *int64    `json:"total_volume,omitempty"`
	PriceYoyPct     *float64  `json:"price_yoy_pct,omitempty"`
	AvgDom          *float64  `json:"avg_dom,omitempty"`
	UnderReview     []*string `json:"under_review,omitempty"`
	StaleQuarter    *bool     `json:"stale_quarter,omitempty"`
	StateAbbr       *string   `json:"state_abbr,omitempty"`
	MedianPrice     *float64  `json:"median_price,omitempty"`
	AvgPrice        *float64  `json:"avg_price,omitempty"`

	// Year-over-year median price change as a decimal (e.g., 0.05 = 5%).
	YoyChange       *float64 `json:"yoy_change,omitempty"`
	AvgDaysOnMarket *int64   `json:"avg_days_on_market,omitempty"`
}

// UnmarshalJSON decodes MarketCountiesResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketCountiesResponseData) UnmarshalJSON(data []byte) error {
	type plain MarketCountiesResponseData
	aux := struct {
		*plain
		SaleCount       lenientNumber[int64]   `json:"sale_count"`
		MedianSalePrice lenientNumber[float64] `json:"median_sale_price"`
		AvgSalePrice    lenientNumber[float64] `json:"avg_sale_price"`
		TotalVolume     lenientNumber[int64]   `json:"total_volume"`
		PriceYoyPct     lenientNumber[float64] `json:"price_yoy_pct"`
		AvgDom          lenientNumber[float64] `json:"avg_dom"`
		MedianPrice     lenientNumber[float64] `json:"median_price"`
		AvgPrice        lenientNumber[float64] `json:"avg_price"`
		YoyChange       lenientNumber[float64] `json:"yoy_change"`
		AvgDaysOnMarket lenientNumber[int64]   `json:"avg_days_on_market"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SaleCount.assignPtr(&r.SaleCount)
	aux.MedianSalePrice.assignPtr(&r.MedianSalePrice)
	aux.AvgSalePrice.assignPtr(&r.AvgSalePrice)
	aux.TotalVolume.assignPtr(&r.TotalVolume)
	aux.PriceYoyPct.assignPtr(&r.PriceYoyPct)
	aux.AvgDom.assignPtr(&r.AvgDom)
	aux.MedianPrice.assignPtr(&r.MedianPrice)
	aux.AvgPrice.assignPtr(&r.AvgPrice)
	aux.YoyChange.assignPtr(&r.YoyChange)
	aux.AvgDaysOnMarket.assignPtr(&r.AvgDaysOnMarket)
	return softTypeError(err)
}

// MarketCountiesResponseSummary is generated from the OpenAPI spec.
type MarketCountiesResponseSummary struct {
	TotalCounties      int64     `json:"total_counties"`
	TotalSales         *int64    `json:"total_sales,omitempty"`
	OverallMedianPrice float64   `json:"overall_median_price"`
	TotalVolume        *int64    `json:"total_volume,omitempty"`
	AvgYoyPct          float64   `json:"avg_yoy_pct"`
	UnderReview        []*string `json:"under_review,omitempty"`
}

// UnmarshalJSON decodes MarketCountiesResponseSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketCountiesResponseSummary) UnmarshalJSON(data []byte) error {
	type plain MarketCountiesResponseSummary
	aux := struct {
		*plain
		TotalCounties      lenientNumber[int64]   `json:"total_counties"`
		TotalSales         lenientNumber[int64]   `json:"total_sales"`
		OverallMedianPrice lenientNumber[float64] `json:"overall_median_price"`
		TotalVolume        lenientNumber[int64]   `json:"total_volume"`
		AvgYoyPct          lenientNumber[float64] `json:"avg_yoy_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalCounties.assign(&r.TotalCounties)
	aux.TotalSales.assignPtr(&r.TotalSales)
	aux.OverallMedianPrice.assign(&r.OverallMedianPrice)
	aux.TotalVolume.assignPtr(&r.TotalVolume)
	aux.AvgYoyPct.assign(&r.AvgYoyPct)
	return softTypeError(err)
}

// Trends: Get market trends
//
// Retrieve quarterly time series of market metrics for one or more counties or a state. Requires
// an API key (or a signed-in session); metered.
//
// HTTP: GET /api/v1/market/trends
func (s *MarketService) Trends(ctx context.Context, params *MarketTrendsParams, opts ...RequestOption) (*MarketTrendsResponse, error) {
	var out MarketTrendsResponse
	if err := s.client.do(ctx, buildMarketTrendsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildMarketTrendsRequest(params *MarketTrendsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/market/trends")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
	}
	return req
}

// MarketTrendsParams holds the query, header and JSON-body parameters of [MarketService.Trends].
// Pass nil when you need none.
type MarketTrendsParams struct {
	// Comma-separated list of county FIPS codes.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// State FIPS code. Used if county_fips is not provided.
	StateFIPS *string `query:"state_fips" json:"-"`
}

// MarketTrendsResponse: Get market trends
type MarketTrendsResponse struct {
	Data []MarketTrendsResponseData `json:"data"`
	Mode string                     `json:"mode"`
}

// MarketTrendsResponseData is generated from the OpenAPI spec.
type MarketTrendsResponseData struct {
	Quarter     *string                            `json:"quarter,omitempty"`
	TotalSales  *int64                             `json:"total_sales,omitempty"`
	MedianPrice *float64                           `json:"median_price,omitempty"`
	TotalVolume *int64                             `json:"total_volume,omitempty"`
	AvgDom      *float64                           `json:"avg_dom,omitempty"`
	CountyFIPS  *string                            `json:"county_fips,omitempty"`
	CountyName  *string                            `json:"county_name,omitempty"`
	Quarters    []MarketTrendsResponseDataQuarters `json:"quarters,omitempty"`
}

// UnmarshalJSON decodes MarketTrendsResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketTrendsResponseData) UnmarshalJSON(data []byte) error {
	type plain MarketTrendsResponseData
	aux := struct {
		*plain
		TotalSales  lenientNumber[int64]   `json:"total_sales"`
		MedianPrice lenientNumber[float64] `json:"median_price"`
		TotalVolume lenientNumber[int64]   `json:"total_volume"`
		AvgDom      lenientNumber[float64] `json:"avg_dom"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalSales.assignPtr(&r.TotalSales)
	aux.MedianPrice.assignPtr(&r.MedianPrice)
	aux.TotalVolume.assignPtr(&r.TotalVolume)
	aux.AvgDom.assignPtr(&r.AvgDom)
	return softTypeError(err)
}

// MarketTrendsResponseDataQuarters is generated from the OpenAPI spec.
type MarketTrendsResponseDataQuarters struct {
	Quarter     *string  `json:"quarter,omitempty"`
	SaleCount   *int64   `json:"sale_count,omitempty"`
	MedianPrice *float64 `json:"median_price,omitempty"`
	AvgPrice    *float64 `json:"avg_price,omitempty"`
	YoyChange   *float64 `json:"yoy_change,omitempty"`
}

// UnmarshalJSON decodes MarketTrendsResponseDataQuarters, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketTrendsResponseDataQuarters) UnmarshalJSON(data []byte) error {
	type plain MarketTrendsResponseDataQuarters
	aux := struct {
		*plain
		SaleCount   lenientNumber[int64]   `json:"sale_count"`
		MedianPrice lenientNumber[float64] `json:"median_price"`
		AvgPrice    lenientNumber[float64] `json:"avg_price"`
		YoyChange   lenientNumber[float64] `json:"yoy_change"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SaleCount.assignPtr(&r.SaleCount)
	aux.MedianPrice.assignPtr(&r.MedianPrice)
	aux.AvgPrice.assignPtr(&r.AvgPrice)
	aux.YoyChange.assignPtr(&r.YoyChange)
	return softTypeError(err)
}

// County: Detailed view for a single county
//
// Returns the full county profile: quarterly market stats (sale count, median price, YoY change,
// days on market), affordability index by year, parcel summary (count, avg assessed value), and
// flip activity. Use for county-detail dashboards.
//
// HTTP: GET /api/v1/market/counties/{fips}
func (s *MarketService) County(ctx context.Context, fips string, params *MarketCountyParams, opts ...RequestOption) (*MarketCountyResponse, error) {
	var out MarketCountyResponse
	if err := s.client.do(ctx, buildMarketCountyRequest(fips, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildMarketCountyRequest(fips string, params *MarketCountyParams) *apiRequest {
	req := newRequest("GET", "/api/v1/market/counties/"+pathParam(fips))
	return req
}

// MarketCountyParams holds the query, header and JSON-body parameters of [MarketService.County].
// Pass nil when you need none.
type MarketCountyParams struct {
}

// MarketCountyResponse: Detailed view for a single county
type MarketCountyResponse = CountyDetail

// Flips: Flip-activity summary grouped by county
//
// Aggregated flip activity per county: count, average ROI, average hold days, total profit. Use
// for surfacing the hottest flip markets. Differs from /api/v1/deals/flips which returns the
// underlying transactions. Requires an API key (or a signed-in session); metered.
//
// HTTP: GET /api/v1/market/flips
func (s *MarketService) Flips(ctx context.Context, params *MarketFlipsParams, opts ...RequestOption) (*MarketFlipsResponse, error) {
	var out MarketFlipsResponse
	if err := s.client.do(ctx, buildMarketFlipsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// FlipsIter iterates every item of [MarketService.Flips] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *MarketService) FlipsIter(ctx context.Context, params *MarketFlipsParams, iter IterOptions, opts ...RequestOption) *Iter[MarketFlipsRow] {
	var p MarketFlipsParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[MarketFlipsRow](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildMarketFlipsRequest(&q)
	}, opts)
}

func buildMarketFlipsRequest(params *MarketFlipsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/market/flips")
	if params != nil {
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// MarketFlipsParams holds the query, header and JSON-body parameters of [MarketService.Flips].
// Pass nil when you need none.
type MarketFlipsParams struct {
	// 2-digit state FIPS filter.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// MarketFlipsResponse: Flip-activity summary grouped by county
type MarketFlipsResponse struct {
	Data   []MarketFlipsRow `json:"data"`
	Total  int64            `json:"total"`
	Limit  int64            `json:"limit"`
	Offset int64            `json:"offset"`
}

// UnmarshalJSON decodes MarketFlipsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketFlipsResponse) UnmarshalJSON(data []byte) error {
	type plain MarketFlipsResponse
	aux := struct {
		*plain
		Total  lenientNumber[int64] `json:"total"`
		Limit  lenientNumber[int64] `json:"limit"`
		Offset lenientNumber[int64] `json:"offset"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assign(&r.Total)
	aux.Limit.assign(&r.Limit)
	aux.Offset.assign(&r.Offset)
	return softTypeError(err)
}

// Snapshot: Market snapshot for a geography
//
// Demographics, economy, housing, lending, hazard and market context for exactly one geography: a
// county, census tract, CBSA or ZIP.
//
// HTTP: GET /api/v1/market/snapshot
func (s *MarketService) Snapshot(ctx context.Context, params *MarketSnapshotParams, opts ...RequestOption) (*MarketSnapshotResponse, error) {
	var out MarketSnapshotResponse
	if err := s.client.do(ctx, buildMarketSnapshotRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildMarketSnapshotRequest(params *MarketSnapshotParams) *apiRequest {
	req := newRequest("GET", "/api/v1/market/snapshot")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "tract", params.Tract)
		addQuery(req.query, "cbsa", params.CBSA)
		addQuery(req.query, "zip", params.Zip)
	}
	return req
}

// MarketSnapshotParams holds the query, header and JSON-body parameters of
// [MarketService.Snapshot]. Pass nil when you need none.
type MarketSnapshotParams struct {
	// 5-digit county FIPS.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// 11-digit census tract GEOID.
	Tract *string `query:"tract" json:"-"`

	// CBSA code.
	CBSA *string `query:"cbsa" json:"-"`

	// 5-digit ZIP.
	Zip *string `query:"zip" json:"-"`
}

// MarketSnapshotResponse: Market snapshot for a geography
type MarketSnapshotResponse struct {
	Geo              MarketSnapshotResponseGeo              `json:"geo"`
	Demographics     MarketSnapshotResponseDemographics     `json:"demographics"`
	Economy          MarketSnapshotResponseEconomy          `json:"economy"`
	Housing          MarketSnapshotResponseHousing          `json:"housing"`
	Lending          MarketSnapshotResponseLending          `json:"lending"`
	Hazard           MarketSnapshotResponseHazard           `json:"hazard"`
	Healthcare       MarketSnapshotResponseHealthcare       `json:"healthcare"`
	Market           map[string]any                         `json:"market,omitempty"`
	MarketHistory    []json.RawMessage                      `json:"market_history"`
	MarketProvenance MarketSnapshotResponseMarketProvenance `json:"market_provenance"`
	GeneratedAt      *string                                `json:"generated_at,omitempty"`
	Guards           []string                               `json:"_guards"`
}

// MarketSnapshotResponseGeo is generated from the OpenAPI spec.
type MarketSnapshotResponseGeo struct {
	Scope            *string   `json:"scope,omitempty"`
	Value            *float64  `json:"value,omitempty"`
	StateFIPS        string    `json:"state_fips"`
	CountyFIPS       string    `json:"county_fips"`
	CountyName       *string   `json:"county_name,omitempty"`
	State            *string   `json:"state,omitempty"`
	CBSACode         *string   `json:"cbsa_code,omitempty"`
	CensusTract      *string   `json:"census_tract,omitempty"`
	Zip5             *string   `json:"zip5,omitempty"`
	GeoBasis         string    `json:"geo_basis"`
	GeoBasisWithheld []*string `json:"geo_basis_withheld"`
}

// UnmarshalJSON decodes MarketSnapshotResponseGeo, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseGeo) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseGeo
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// MarketSnapshotResponseDemographics is generated from the OpenAPI spec.
type MarketSnapshotResponseDemographics struct {
	AcsMedianHhIncome         *float64                                                 `json:"acs_median_hh_income,omitempty"`
	AcsMedianHomeValue        *float64                                                 `json:"acs_median_home_value,omitempty"`
	AcsMedianRent             *float64                                                 `json:"acs_median_rent,omitempty"`
	AcsMedianYearBuilt        *float64                                                 `json:"acs_median_year_built,omitempty"`
	AcsBachelorsPlusPct       *float64                                                 `json:"acs_bachelors_plus_pct,omitempty"`
	AcsOwnerOccupiedPct       *float64                                                 `json:"acs_owner_occupied_pct,omitempty"`
	AcsPovertyPct             *float64                                                 `json:"acs_poverty_pct,omitempty"`
	AcsVacantHousingPct       *float64                                                 `json:"acs_vacant_housing_pct,omitempty"`
	AcsAge65plusPct           *float64                                                 `json:"acs_age_65plus_pct,omitempty"`
	AcsBroadbandPct           *float64                                                 `json:"acs_broadband_pct,omitempty"`
	AcsMeanCommuteMinutes     *float64                                                 `json:"acs_mean_commute_minutes,omitempty"`
	AcsLongCommutePct         *float64                                                 `json:"acs_long_commute_pct,omitempty"`
	TractPopulation           *string                                                  `json:"tract_population,omitempty"`
	TractMedianHomeValue      *string                                                  `json:"tract_median_home_value,omitempty"`
	TractMedianRent           *string                                                  `json:"tract_median_rent,omitempty"`
	TractOwnerOccupiedPct     *string                                                  `json:"tract_owner_occupied_pct,omitempty"`
	TractVacancyRate          *string                                                  `json:"tract_vacancy_rate,omitempty"`
	TractPovertyRate          *string                                                  `json:"tract_poverty_rate,omitempty"`
	TractCollegeEducatedPct   *string                                                  `json:"tract_college_educated_pct,omitempty"`
	TractAvgIncome            *string                                                  `json:"tract_avg_income,omitempty"`
	DemographicsBasis         string                                                   `json:"demographics_basis"`
	DemographicsBasisWithheld []*string                                                `json:"demographics_basis_withheld"`
	DemographicsBasisGrain    MarketSnapshotResponseDemographicsDemographicsBasisGrain `json:"demographics_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseDemographics, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseDemographics) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseDemographics
	aux := struct {
		*plain
		AcsMedianHhIncome     lenientNumber[float64] `json:"acs_median_hh_income"`
		AcsMedianHomeValue    lenientNumber[float64] `json:"acs_median_home_value"`
		AcsMedianRent         lenientNumber[float64] `json:"acs_median_rent"`
		AcsMedianYearBuilt    lenientNumber[float64] `json:"acs_median_year_built"`
		AcsBachelorsPlusPct   lenientNumber[float64] `json:"acs_bachelors_plus_pct"`
		AcsOwnerOccupiedPct   lenientNumber[float64] `json:"acs_owner_occupied_pct"`
		AcsPovertyPct         lenientNumber[float64] `json:"acs_poverty_pct"`
		AcsVacantHousingPct   lenientNumber[float64] `json:"acs_vacant_housing_pct"`
		AcsAge65plusPct       lenientNumber[float64] `json:"acs_age_65plus_pct"`
		AcsBroadbandPct       lenientNumber[float64] `json:"acs_broadband_pct"`
		AcsMeanCommuteMinutes lenientNumber[float64] `json:"acs_mean_commute_minutes"`
		AcsLongCommutePct     lenientNumber[float64] `json:"acs_long_commute_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AcsMedianHhIncome.assignPtr(&r.AcsMedianHhIncome)
	aux.AcsMedianHomeValue.assignPtr(&r.AcsMedianHomeValue)
	aux.AcsMedianRent.assignPtr(&r.AcsMedianRent)
	aux.AcsMedianYearBuilt.assignPtr(&r.AcsMedianYearBuilt)
	aux.AcsBachelorsPlusPct.assignPtr(&r.AcsBachelorsPlusPct)
	aux.AcsOwnerOccupiedPct.assignPtr(&r.AcsOwnerOccupiedPct)
	aux.AcsPovertyPct.assignPtr(&r.AcsPovertyPct)
	aux.AcsVacantHousingPct.assignPtr(&r.AcsVacantHousingPct)
	aux.AcsAge65plusPct.assignPtr(&r.AcsAge65plusPct)
	aux.AcsBroadbandPct.assignPtr(&r.AcsBroadbandPct)
	aux.AcsMeanCommuteMinutes.assignPtr(&r.AcsMeanCommuteMinutes)
	aux.AcsLongCommutePct.assignPtr(&r.AcsLongCommutePct)
	return softTypeError(err)
}

// MarketSnapshotResponseDemographicsDemographicsBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseDemographicsDemographicsBasisGrain struct {
	AcsMedianHhIncome       *string `json:"acs_median_hh_income,omitempty"`
	AcsMedianHomeValue      *string `json:"acs_median_home_value,omitempty"`
	AcsMedianRent           *string `json:"acs_median_rent,omitempty"`
	AcsMedianYearBuilt      *string `json:"acs_median_year_built,omitempty"`
	AcsBachelorsPlusPct     *string `json:"acs_bachelors_plus_pct,omitempty"`
	AcsOwnerOccupiedPct     *string `json:"acs_owner_occupied_pct,omitempty"`
	AcsPovertyPct           *string `json:"acs_poverty_pct,omitempty"`
	AcsVacantHousingPct     *string `json:"acs_vacant_housing_pct,omitempty"`
	AcsAge65plusPct         *string `json:"acs_age_65plus_pct,omitempty"`
	AcsBroadbandPct         *string `json:"acs_broadband_pct,omitempty"`
	AcsMeanCommuteMinutes   *string `json:"acs_mean_commute_minutes,omitempty"`
	AcsLongCommutePct       *string `json:"acs_long_commute_pct,omitempty"`
	TractPopulation         *string `json:"tract_population,omitempty"`
	TractMedianHomeValue    *string `json:"tract_median_home_value,omitempty"`
	TractMedianRent         *string `json:"tract_median_rent,omitempty"`
	TractOwnerOccupiedPct   *string `json:"tract_owner_occupied_pct,omitempty"`
	TractVacancyRate        *string `json:"tract_vacancy_rate,omitempty"`
	TractPovertyRate        *string `json:"tract_poverty_rate,omitempty"`
	TractCollegeEducatedPct *string `json:"tract_college_educated_pct,omitempty"`
	TractAvgIncome          *string `json:"tract_avg_income,omitempty"`
}

// MarketSnapshotResponseEconomy is generated from the OpenAPI spec.
type MarketSnapshotResponseEconomy struct {
	BeaGdp2024Thousands            *float64                                       `json:"bea_gdp_2024_thousands,omitempty"`
	BeaGdp2023Thousands            *float64                                       `json:"bea_gdp_2023_thousands,omitempty"`
	BeaGdp5yGrowthPct              *float64                                       `json:"bea_gdp_5y_growth_pct,omitempty"`
	BeaGdp10yGrowthPct             *float64                                       `json:"bea_gdp_10y_growth_pct,omitempty"`
	BeaGdp20yGrowthPct             *float64                                       `json:"bea_gdp_20y_growth_pct,omitempty"`
	BeaPcpi2024                    *float64                                       `json:"bea_pcpi_2024,omitempty"`
	BeaPcpi5yGrowthPct             *float64                                       `json:"bea_pcpi_5y_growth_pct,omitempty"`
	BeaPersonalIncomeThousands2024 *float64                                       `json:"bea_personal_income_thousands_2024,omitempty"`
	BeaPopulation2024              *float64                                       `json:"bea_population_2024,omitempty"`
	CountyEmployment               *float64                                       `json:"county_employment,omitempty"`
	CountyUnemploymentRate         *float64                                       `json:"county_unemployment_rate,omitempty"`
	CountyTotalEmployees           *float64                                       `json:"county_total_employees,omitempty"`
	CountyTotalEstablishments      *float64                                       `json:"county_total_establishments,omitempty"`
	CountyMedianIncomeIrs          *int64                                         `json:"county_median_income_irs,omitempty"`
	CountyAffordabilityRatio       *float64                                       `json:"county_affordability_ratio,omitempty"`
	CountyAffordabilityRating      *string                                        `json:"county_affordability_rating,omitempty"`
	LodesJobsTotal                 *float64                                       `json:"lodes_jobs_total,omitempty"`
	LodesJobsHighWage              *float64                                       `json:"lodes_jobs_high_wage,omitempty"`
	LodesJobsMidWage               *float64                                       `json:"lodes_jobs_mid_wage,omitempty"`
	LodesJobsLowWage               *float64                                       `json:"lodes_jobs_low_wage,omitempty"`
	LodesJobsHealthcare            *float64                                       `json:"lodes_jobs_healthcare,omitempty"`
	LodesJobsManufacturing         *float64                                       `json:"lodes_jobs_manufacturing,omitempty"`
	LodesJobsRetail                *float64                                       `json:"lodes_jobs_retail,omitempty"`
	LodesJobsEducation             *float64                                       `json:"lodes_jobs_education,omitempty"`
	LodesJobsHospitality           *float64                                       `json:"lodes_jobs_hospitality,omitempty"`
	TractTotalJobs                 *string                                        `json:"tract_total_jobs,omitempty"`
	TractJobsDensityPerSqmi        *string                                        `json:"tract_jobs_density_per_sqmi,omitempty"`
	TractHighWagePct               *string                                        `json:"tract_high_wage_pct,omitempty"`
	TractHealthcareJobsPct         *string                                        `json:"tract_healthcare_jobs_pct,omitempty"`
	EconomyBasis                   string                                         `json:"economy_basis"`
	EconomyBasisWithheld           []*string                                      `json:"economy_basis_withheld"`
	EconomyBasisGrain              MarketSnapshotResponseEconomyEconomyBasisGrain `json:"economy_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseEconomy, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseEconomy) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseEconomy
	aux := struct {
		*plain
		BeaGdp2024Thousands            lenientNumber[float64] `json:"bea_gdp_2024_thousands"`
		BeaGdp2023Thousands            lenientNumber[float64] `json:"bea_gdp_2023_thousands"`
		BeaGdp5yGrowthPct              lenientNumber[float64] `json:"bea_gdp_5y_growth_pct"`
		BeaGdp10yGrowthPct             lenientNumber[float64] `json:"bea_gdp_10y_growth_pct"`
		BeaGdp20yGrowthPct             lenientNumber[float64] `json:"bea_gdp_20y_growth_pct"`
		BeaPcpi2024                    lenientNumber[float64] `json:"bea_pcpi_2024"`
		BeaPcpi5yGrowthPct             lenientNumber[float64] `json:"bea_pcpi_5y_growth_pct"`
		BeaPersonalIncomeThousands2024 lenientNumber[float64] `json:"bea_personal_income_thousands_2024"`
		BeaPopulation2024              lenientNumber[float64] `json:"bea_population_2024"`
		CountyEmployment               lenientNumber[float64] `json:"county_employment"`
		CountyUnemploymentRate         lenientNumber[float64] `json:"county_unemployment_rate"`
		CountyTotalEmployees           lenientNumber[float64] `json:"county_total_employees"`
		CountyTotalEstablishments      lenientNumber[float64] `json:"county_total_establishments"`
		CountyMedianIncomeIrs          lenientNumber[int64]   `json:"county_median_income_irs"`
		CountyAffordabilityRatio       lenientNumber[float64] `json:"county_affordability_ratio"`
		LodesJobsTotal                 lenientNumber[float64] `json:"lodes_jobs_total"`
		LodesJobsHighWage              lenientNumber[float64] `json:"lodes_jobs_high_wage"`
		LodesJobsMidWage               lenientNumber[float64] `json:"lodes_jobs_mid_wage"`
		LodesJobsLowWage               lenientNumber[float64] `json:"lodes_jobs_low_wage"`
		LodesJobsHealthcare            lenientNumber[float64] `json:"lodes_jobs_healthcare"`
		LodesJobsManufacturing         lenientNumber[float64] `json:"lodes_jobs_manufacturing"`
		LodesJobsRetail                lenientNumber[float64] `json:"lodes_jobs_retail"`
		LodesJobsEducation             lenientNumber[float64] `json:"lodes_jobs_education"`
		LodesJobsHospitality           lenientNumber[float64] `json:"lodes_jobs_hospitality"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.BeaGdp2024Thousands.assignPtr(&r.BeaGdp2024Thousands)
	aux.BeaGdp2023Thousands.assignPtr(&r.BeaGdp2023Thousands)
	aux.BeaGdp5yGrowthPct.assignPtr(&r.BeaGdp5yGrowthPct)
	aux.BeaGdp10yGrowthPct.assignPtr(&r.BeaGdp10yGrowthPct)
	aux.BeaGdp20yGrowthPct.assignPtr(&r.BeaGdp20yGrowthPct)
	aux.BeaPcpi2024.assignPtr(&r.BeaPcpi2024)
	aux.BeaPcpi5yGrowthPct.assignPtr(&r.BeaPcpi5yGrowthPct)
	aux.BeaPersonalIncomeThousands2024.assignPtr(&r.BeaPersonalIncomeThousands2024)
	aux.BeaPopulation2024.assignPtr(&r.BeaPopulation2024)
	aux.CountyEmployment.assignPtr(&r.CountyEmployment)
	aux.CountyUnemploymentRate.assignPtr(&r.CountyUnemploymentRate)
	aux.CountyTotalEmployees.assignPtr(&r.CountyTotalEmployees)
	aux.CountyTotalEstablishments.assignPtr(&r.CountyTotalEstablishments)
	aux.CountyMedianIncomeIrs.assignPtr(&r.CountyMedianIncomeIrs)
	aux.CountyAffordabilityRatio.assignPtr(&r.CountyAffordabilityRatio)
	aux.LodesJobsTotal.assignPtr(&r.LodesJobsTotal)
	aux.LodesJobsHighWage.assignPtr(&r.LodesJobsHighWage)
	aux.LodesJobsMidWage.assignPtr(&r.LodesJobsMidWage)
	aux.LodesJobsLowWage.assignPtr(&r.LodesJobsLowWage)
	aux.LodesJobsHealthcare.assignPtr(&r.LodesJobsHealthcare)
	aux.LodesJobsManufacturing.assignPtr(&r.LodesJobsManufacturing)
	aux.LodesJobsRetail.assignPtr(&r.LodesJobsRetail)
	aux.LodesJobsEducation.assignPtr(&r.LodesJobsEducation)
	aux.LodesJobsHospitality.assignPtr(&r.LodesJobsHospitality)
	return softTypeError(err)
}

// MarketSnapshotResponseEconomyEconomyBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseEconomyEconomyBasisGrain struct {
	CountyEmployment          *string `json:"county_employment,omitempty"`
	CountyUnemploymentRate    *string `json:"county_unemployment_rate,omitempty"`
	CountyTotalEmployees      *string `json:"county_total_employees,omitempty"`
	CountyTotalEstablishments *string `json:"county_total_establishments,omitempty"`
	LodesJobsTotal            *string `json:"lodes_jobs_total,omitempty"`
	LodesJobsHighWage         *string `json:"lodes_jobs_high_wage,omitempty"`
	LodesJobsMidWage          *string `json:"lodes_jobs_mid_wage,omitempty"`
	LodesJobsLowWage          *string `json:"lodes_jobs_low_wage,omitempty"`
	LodesJobsHealthcare       *string `json:"lodes_jobs_healthcare,omitempty"`
	LodesJobsManufacturing    *string `json:"lodes_jobs_manufacturing,omitempty"`
	LodesJobsRetail           *string `json:"lodes_jobs_retail,omitempty"`
	LodesJobsEducation        *string `json:"lodes_jobs_education,omitempty"`
	LodesJobsHospitality      *string `json:"lodes_jobs_hospitality,omitempty"`
	TractTotalJobs            *string `json:"tract_total_jobs,omitempty"`
	TractJobsDensityPerSqmi   *string `json:"tract_jobs_density_per_sqmi,omitempty"`
	TractHighWagePct          *string `json:"tract_high_wage_pct,omitempty"`
	TractHealthcareJobsPct    *string `json:"tract_healthcare_jobs_pct,omitempty"`
}

// MarketSnapshotResponseHousing is generated from the OpenAPI spec.
type MarketSnapshotResponseHousing struct {
	BpsTotalUnits2024    *float64                                       `json:"bps_total_units_2024,omitempty"`
	BpsTotalUnits2023    *float64                                       `json:"bps_total_units_2023,omitempty"`
	BpsTotalValue2024    *float64                                       `json:"bps_total_value_2024,omitempty"`
	BpsSfUnits2024       *float64                                       `json:"bps_sf_units_2024,omitempty"`
	BpsSfValue2024       *float64                                       `json:"bps_sf_value_2024,omitempty"`
	BpsMfUnits2024       *float64                                       `json:"bps_mf_units_2024,omitempty"`
	BpsMfValue2024       *float64                                       `json:"bps_mf_value_2024,omitempty"`
	BpsYoyUnitGrowthPct  *float64                                       `json:"bps_yoy_unit_growth_pct,omitempty"`
	FhfaHpiLatest        *float64                                       `json:"fhfa_hpi_latest,omitempty"`
	FhfaHpiYear          *int64                                         `json:"fhfa_hpi_year,omitempty"`
	FhfaHpi1yChangePct   *float64                                       `json:"fhfa_hpi_1y_change_pct,omitempty"`
	FhfaHpi5yChangePct   *float64                                       `json:"fhfa_hpi_5y_change_pct,omitempty"`
	FhfaHpi10yChangePct  *float64                                       `json:"fhfa_hpi_10y_change_pct,omitempty"`
	FhfaHpi20yChangePct  *float64                                       `json:"fhfa_hpi_20y_change_pct,omitempty"`
	HousingBasis         string                                         `json:"housing_basis"`
	HousingBasisWithheld []*string                                      `json:"housing_basis_withheld"`
	HousingBasisGrain    MarketSnapshotResponseHousingHousingBasisGrain `json:"housing_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseHousing, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseHousing) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseHousing
	aux := struct {
		*plain
		BpsTotalUnits2024   lenientNumber[float64] `json:"bps_total_units_2024"`
		BpsTotalUnits2023   lenientNumber[float64] `json:"bps_total_units_2023"`
		BpsTotalValue2024   lenientNumber[float64] `json:"bps_total_value_2024"`
		BpsSfUnits2024      lenientNumber[float64] `json:"bps_sf_units_2024"`
		BpsSfValue2024      lenientNumber[float64] `json:"bps_sf_value_2024"`
		BpsMfUnits2024      lenientNumber[float64] `json:"bps_mf_units_2024"`
		BpsMfValue2024      lenientNumber[float64] `json:"bps_mf_value_2024"`
		BpsYoyUnitGrowthPct lenientNumber[float64] `json:"bps_yoy_unit_growth_pct"`
		FhfaHpiLatest       lenientNumber[float64] `json:"fhfa_hpi_latest"`
		FhfaHpiYear         lenientNumber[int64]   `json:"fhfa_hpi_year"`
		FhfaHpi1yChangePct  lenientNumber[float64] `json:"fhfa_hpi_1y_change_pct"`
		FhfaHpi5yChangePct  lenientNumber[float64] `json:"fhfa_hpi_5y_change_pct"`
		FhfaHpi10yChangePct lenientNumber[float64] `json:"fhfa_hpi_10y_change_pct"`
		FhfaHpi20yChangePct lenientNumber[float64] `json:"fhfa_hpi_20y_change_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.BpsTotalUnits2024.assignPtr(&r.BpsTotalUnits2024)
	aux.BpsTotalUnits2023.assignPtr(&r.BpsTotalUnits2023)
	aux.BpsTotalValue2024.assignPtr(&r.BpsTotalValue2024)
	aux.BpsSfUnits2024.assignPtr(&r.BpsSfUnits2024)
	aux.BpsSfValue2024.assignPtr(&r.BpsSfValue2024)
	aux.BpsMfUnits2024.assignPtr(&r.BpsMfUnits2024)
	aux.BpsMfValue2024.assignPtr(&r.BpsMfValue2024)
	aux.BpsYoyUnitGrowthPct.assignPtr(&r.BpsYoyUnitGrowthPct)
	aux.FhfaHpiLatest.assignPtr(&r.FhfaHpiLatest)
	aux.FhfaHpiYear.assignPtr(&r.FhfaHpiYear)
	aux.FhfaHpi1yChangePct.assignPtr(&r.FhfaHpi1yChangePct)
	aux.FhfaHpi5yChangePct.assignPtr(&r.FhfaHpi5yChangePct)
	aux.FhfaHpi10yChangePct.assignPtr(&r.FhfaHpi10yChangePct)
	aux.FhfaHpi20yChangePct.assignPtr(&r.FhfaHpi20yChangePct)
	return softTypeError(err)
}

// MarketSnapshotResponseHousingHousingBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseHousingHousingBasisGrain struct {
	FhfaHpiLatest       *string `json:"fhfa_hpi_latest,omitempty"`
	FhfaHpiYear         *string `json:"fhfa_hpi_year,omitempty"`
	FhfaHpi1yChangePct  *string `json:"fhfa_hpi_1y_change_pct,omitempty"`
	FhfaHpi5yChangePct  *string `json:"fhfa_hpi_5y_change_pct,omitempty"`
	FhfaHpi10yChangePct *string `json:"fhfa_hpi_10y_change_pct,omitempty"`
	FhfaHpi20yChangePct *string `json:"fhfa_hpi_20y_change_pct,omitempty"`
}

// MarketSnapshotResponseLending is generated from the OpenAPI spec.
type MarketSnapshotResponseLending struct {
	HmdaOrigCount              *float64                                       `json:"hmda_orig_count,omitempty"`
	HmdaOrigVolumeThousands    *float64                                       `json:"hmda_orig_volume_thousands,omitempty"`
	HmdaAvgLoanAmountThousands *float64                                       `json:"hmda_avg_loan_amount_thousands,omitempty"`
	HmdaAvgInterestRate        *float64                                       `json:"hmda_avg_interest_rate,omitempty"`
	HmdaAvgLtv                 *float64                                       `json:"hmda_avg_ltv,omitempty"`
	HmdaConvCount              *float64                                       `json:"hmda_conv_count,omitempty"`
	HmdaFhaCount               *float64                                       `json:"hmda_fha_count,omitempty"`
	HmdaVaCount                *float64                                       `json:"hmda_va_count,omitempty"`
	HmdaUsdaCount              *float64                                       `json:"hmda_usda_count,omitempty"`
	TractLoanOriginations      *string                                        `json:"tract_loan_originations,omitempty"`
	TractAvgLoanAmount         *string                                        `json:"tract_avg_loan_amount,omitempty"`
	TractAvgInterestRate       *string                                        `json:"tract_avg_interest_rate,omitempty"`
	TractFhaPct                *string                                        `json:"tract_fha_pct,omitempty"`
	TractInvestorPct           *string                                        `json:"tract_investor_pct,omitempty"`
	TractCreditRiskTier        *string                                        `json:"tract_credit_risk_tier,omitempty"`
	TractFloodClaims           *string                                        `json:"tract_flood_claims,omitempty"`
	TractFloodLossRatio        *string                                        `json:"tract_flood_loss_ratio,omitempty"`
	LendingBasis               string                                         `json:"lending_basis"`
	LendingBasisWithheld       []*string                                      `json:"lending_basis_withheld"`
	LendingBasisGrain          MarketSnapshotResponseLendingLendingBasisGrain `json:"lending_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseLending, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseLending) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseLending
	aux := struct {
		*plain
		HmdaOrigCount              lenientNumber[float64] `json:"hmda_orig_count"`
		HmdaOrigVolumeThousands    lenientNumber[float64] `json:"hmda_orig_volume_thousands"`
		HmdaAvgLoanAmountThousands lenientNumber[float64] `json:"hmda_avg_loan_amount_thousands"`
		HmdaAvgInterestRate        lenientNumber[float64] `json:"hmda_avg_interest_rate"`
		HmdaAvgLtv                 lenientNumber[float64] `json:"hmda_avg_ltv"`
		HmdaConvCount              lenientNumber[float64] `json:"hmda_conv_count"`
		HmdaFhaCount               lenientNumber[float64] `json:"hmda_fha_count"`
		HmdaVaCount                lenientNumber[float64] `json:"hmda_va_count"`
		HmdaUsdaCount              lenientNumber[float64] `json:"hmda_usda_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.HmdaOrigCount.assignPtr(&r.HmdaOrigCount)
	aux.HmdaOrigVolumeThousands.assignPtr(&r.HmdaOrigVolumeThousands)
	aux.HmdaAvgLoanAmountThousands.assignPtr(&r.HmdaAvgLoanAmountThousands)
	aux.HmdaAvgInterestRate.assignPtr(&r.HmdaAvgInterestRate)
	aux.HmdaAvgLtv.assignPtr(&r.HmdaAvgLtv)
	aux.HmdaConvCount.assignPtr(&r.HmdaConvCount)
	aux.HmdaFhaCount.assignPtr(&r.HmdaFhaCount)
	aux.HmdaVaCount.assignPtr(&r.HmdaVaCount)
	aux.HmdaUsdaCount.assignPtr(&r.HmdaUsdaCount)
	return softTypeError(err)
}

// MarketSnapshotResponseLendingLendingBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseLendingLendingBasisGrain struct {
	HmdaOrigCount              *string `json:"hmda_orig_count,omitempty"`
	HmdaOrigVolumeThousands    *string `json:"hmda_orig_volume_thousands,omitempty"`
	HmdaAvgLoanAmountThousands *string `json:"hmda_avg_loan_amount_thousands,omitempty"`
	HmdaAvgInterestRate        *string `json:"hmda_avg_interest_rate,omitempty"`
	HmdaAvgLtv                 *string `json:"hmda_avg_ltv,omitempty"`
	HmdaConvCount              *string `json:"hmda_conv_count,omitempty"`
	HmdaFhaCount               *string `json:"hmda_fha_count,omitempty"`
	HmdaVaCount                *string `json:"hmda_va_count,omitempty"`
	HmdaUsdaCount              *string `json:"hmda_usda_count,omitempty"`
	TractLoanOriginations      *string `json:"tract_loan_originations,omitempty"`
	TractAvgLoanAmount         *string `json:"tract_avg_loan_amount,omitempty"`
	TractAvgInterestRate       *string `json:"tract_avg_interest_rate,omitempty"`
	TractFhaPct                *string `json:"tract_fha_pct,omitempty"`
	TractInvestorPct           *string `json:"tract_investor_pct,omitempty"`
	TractCreditRiskTier        *string `json:"tract_credit_risk_tier,omitempty"`
	TractFloodClaims           *string `json:"tract_flood_claims,omitempty"`
	TractFloodLossRatio        *string `json:"tract_flood_loss_ratio,omitempty"`
}

// MarketSnapshotResponseHazard is generated from the OpenAPI spec.
type MarketSnapshotResponseHazard struct {
	FemaDisasterCount         *int64                                       `json:"fema_disaster_count,omitempty"`
	FemaDisasterCount10y      *int64                                       `json:"fema_disaster_count_10y,omitempty"`
	FemaFloodCount            *int64                                       `json:"fema_flood_count,omitempty"`
	FemaFireCount             *int64                                       `json:"fema_fire_count,omitempty"`
	FemaHurricaneCount        *int64                                       `json:"fema_hurricane_count,omitempty"`
	FemaTornadoCount          *int64                                       `json:"fema_tornado_count,omitempty"`
	FemaEarthquakeCount       *int64                                       `json:"fema_earthquake_count,omitempty"`
	FemaSevStormCount         *int64                                       `json:"fema_sev_storm_count,omitempty"`
	FemaBiologicalCount       *int64                                       `json:"fema_biological_count,omitempty"`
	FemaIaDeclarations        *float64                                     `json:"fema_ia_declarations,omitempty"`
	FemaPaDeclarations        *float64                                     `json:"fema_pa_declarations,omitempty"`
	FemaPolicyCount           *float64                                     `json:"fema_policy_count,omitempty"`
	FemaTopIncidentType       *string                                      `json:"fema_top_incident_type,omitempty"`
	FemaLatestDeclarationDate *string                                      `json:"fema_latest_declaration_date,omitempty"`
	CountyViolentCrimeRate    *float64                                     `json:"county_violent_crime_rate,omitempty"`
	CountyPropertyCrimeRate   *float64                                     `json:"county_property_crime_rate,omitempty"`
	HazardBasis               string                                       `json:"hazard_basis"`
	HazardBasisWithheld       []*string                                    `json:"hazard_basis_withheld"`
	HazardBasisGrain          MarketSnapshotResponseHazardHazardBasisGrain `json:"hazard_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseHazard, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseHazard) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseHazard
	aux := struct {
		*plain
		FemaDisasterCount       lenientNumber[int64]   `json:"fema_disaster_count"`
		FemaDisasterCount10y    lenientNumber[int64]   `json:"fema_disaster_count_10y"`
		FemaFloodCount          lenientNumber[int64]   `json:"fema_flood_count"`
		FemaFireCount           lenientNumber[int64]   `json:"fema_fire_count"`
		FemaHurricaneCount      lenientNumber[int64]   `json:"fema_hurricane_count"`
		FemaTornadoCount        lenientNumber[int64]   `json:"fema_tornado_count"`
		FemaEarthquakeCount     lenientNumber[int64]   `json:"fema_earthquake_count"`
		FemaSevStormCount       lenientNumber[int64]   `json:"fema_sev_storm_count"`
		FemaBiologicalCount     lenientNumber[int64]   `json:"fema_biological_count"`
		FemaIaDeclarations      lenientNumber[float64] `json:"fema_ia_declarations"`
		FemaPaDeclarations      lenientNumber[float64] `json:"fema_pa_declarations"`
		FemaPolicyCount         lenientNumber[float64] `json:"fema_policy_count"`
		CountyViolentCrimeRate  lenientNumber[float64] `json:"county_violent_crime_rate"`
		CountyPropertyCrimeRate lenientNumber[float64] `json:"county_property_crime_rate"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.FemaDisasterCount.assignPtr(&r.FemaDisasterCount)
	aux.FemaDisasterCount10y.assignPtr(&r.FemaDisasterCount10y)
	aux.FemaFloodCount.assignPtr(&r.FemaFloodCount)
	aux.FemaFireCount.assignPtr(&r.FemaFireCount)
	aux.FemaHurricaneCount.assignPtr(&r.FemaHurricaneCount)
	aux.FemaTornadoCount.assignPtr(&r.FemaTornadoCount)
	aux.FemaEarthquakeCount.assignPtr(&r.FemaEarthquakeCount)
	aux.FemaSevStormCount.assignPtr(&r.FemaSevStormCount)
	aux.FemaBiologicalCount.assignPtr(&r.FemaBiologicalCount)
	aux.FemaIaDeclarations.assignPtr(&r.FemaIaDeclarations)
	aux.FemaPaDeclarations.assignPtr(&r.FemaPaDeclarations)
	aux.FemaPolicyCount.assignPtr(&r.FemaPolicyCount)
	aux.CountyViolentCrimeRate.assignPtr(&r.CountyViolentCrimeRate)
	aux.CountyPropertyCrimeRate.assignPtr(&r.CountyPropertyCrimeRate)
	return softTypeError(err)
}

// MarketSnapshotResponseHazardHazardBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseHazardHazardBasisGrain struct {
	FemaPolicyCount         *string `json:"fema_policy_count,omitempty"`
	CountyViolentCrimeRate  *string `json:"county_violent_crime_rate,omitempty"`
	CountyPropertyCrimeRate *string `json:"county_property_crime_rate,omitempty"`
}

// MarketSnapshotResponseHealthcare is generated from the OpenAPI spec.
type MarketSnapshotResponseHealthcare struct {
	CmsHospCount            *float64                                             `json:"cms_hosp_count,omitempty"`
	CmsHospAvgStars         *float64                                             `json:"cms_hosp_avg_stars,omitempty"`
	CmsHosp45Star           *float64                                             `json:"cms_hosp_4_5_star,omitempty"`
	CmsHospWithEr           *float64                                             `json:"cms_hosp_with_er,omitempty"`
	CmsNhCount              *float64                                             `json:"cms_nh_count,omitempty"`
	CmsNhBeds               *float64                                             `json:"cms_nh_beds,omitempty"`
	CmsNhAvgStars           *float64                                             `json:"cms_nh_avg_stars,omitempty"`
	CmsNh45Star             *float64                                             `json:"cms_nh_4_5_star,omitempty"`
	CmsNh12Star             *float64                                             `json:"cms_nh_1_2_star,omitempty"`
	CmsHospiceCount         *float64                                             `json:"cms_hospice_count,omitempty"`
	CmsHhCount              *float64                                             `json:"cms_hh_count,omitempty"`
	HealthcareBasis         string                                               `json:"healthcare_basis"`
	HealthcareBasisWithheld []*string                                            `json:"healthcare_basis_withheld"`
	HealthcareBasisGrain    MarketSnapshotResponseHealthcareHealthcareBasisGrain `json:"healthcare_basis_grain"`
}

// UnmarshalJSON decodes MarketSnapshotResponseHealthcare, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseHealthcare) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseHealthcare
	aux := struct {
		*plain
		CmsHospCount    lenientNumber[float64] `json:"cms_hosp_count"`
		CmsHospAvgStars lenientNumber[float64] `json:"cms_hosp_avg_stars"`
		CmsHosp45Star   lenientNumber[float64] `json:"cms_hosp_4_5_star"`
		CmsHospWithEr   lenientNumber[float64] `json:"cms_hosp_with_er"`
		CmsNhCount      lenientNumber[float64] `json:"cms_nh_count"`
		CmsNhBeds       lenientNumber[float64] `json:"cms_nh_beds"`
		CmsNhAvgStars   lenientNumber[float64] `json:"cms_nh_avg_stars"`
		CmsNh45Star     lenientNumber[float64] `json:"cms_nh_4_5_star"`
		CmsNh12Star     lenientNumber[float64] `json:"cms_nh_1_2_star"`
		CmsHospiceCount lenientNumber[float64] `json:"cms_hospice_count"`
		CmsHhCount      lenientNumber[float64] `json:"cms_hh_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CmsHospCount.assignPtr(&r.CmsHospCount)
	aux.CmsHospAvgStars.assignPtr(&r.CmsHospAvgStars)
	aux.CmsHosp45Star.assignPtr(&r.CmsHosp45Star)
	aux.CmsHospWithEr.assignPtr(&r.CmsHospWithEr)
	aux.CmsNhCount.assignPtr(&r.CmsNhCount)
	aux.CmsNhBeds.assignPtr(&r.CmsNhBeds)
	aux.CmsNhAvgStars.assignPtr(&r.CmsNhAvgStars)
	aux.CmsNh45Star.assignPtr(&r.CmsNh45Star)
	aux.CmsNh12Star.assignPtr(&r.CmsNh12Star)
	aux.CmsHospiceCount.assignPtr(&r.CmsHospiceCount)
	aux.CmsHhCount.assignPtr(&r.CmsHhCount)
	return softTypeError(err)
}

// MarketSnapshotResponseHealthcareHealthcareBasisGrain is generated from the OpenAPI spec.
type MarketSnapshotResponseHealthcareHealthcareBasisGrain struct {
	CmsHospCount    *string `json:"cms_hosp_count,omitempty"`
	CmsHospAvgStars *string `json:"cms_hosp_avg_stars,omitempty"`
	CmsHosp45Star   *string `json:"cms_hosp_4_5_star,omitempty"`
	CmsHospWithEr   *string `json:"cms_hosp_with_er,omitempty"`
	CmsNhCount      *string `json:"cms_nh_count,omitempty"`
	CmsNhBeds       *string `json:"cms_nh_beds,omitempty"`
	CmsNhAvgStars   *string `json:"cms_nh_avg_stars,omitempty"`
	CmsNh45Star     *string `json:"cms_nh_4_5_star,omitempty"`
	CmsNh12Star     *string `json:"cms_nh_1_2_star,omitempty"`
	CmsHospiceCount *string `json:"cms_hospice_count,omitempty"`
	CmsHhCount      *string `json:"cms_hh_count,omitempty"`
}

// MarketSnapshotResponseMarketProvenance is generated from the OpenAPI spec.
type MarketSnapshotResponseMarketProvenance struct {
	Dataset               *string  `json:"dataset,omitempty"`
	PeriodUpperBound      *string  `json:"period_upper_bound,omitempty"`
	ReturnedRecords       *float64 `json:"returned_records,omitempty"`
	OldestRecordRefresh   *string  `json:"oldest_record_refresh,omitempty"`
	NewestRecordRefresh   *string  `json:"newest_record_refresh,omitempty"`
	RecordsWithoutRefresh *float64 `json:"records_without_refresh,omitempty"`
	SourceVintage         *string  `json:"source_vintage,omitempty"`
	Basis                 *string  `json:"basis,omitempty"`
}

// UnmarshalJSON decodes MarketSnapshotResponseMarketProvenance, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSnapshotResponseMarketProvenance) UnmarshalJSON(data []byte) error {
	type plain MarketSnapshotResponseMarketProvenance
	aux := struct {
		*plain
		ReturnedRecords       lenientNumber[float64] `json:"returned_records"`
		RecordsWithoutRefresh lenientNumber[float64] `json:"records_without_refresh"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ReturnedRecords.assignPtr(&r.ReturnedRecords)
	aux.RecordsWithoutRefresh.assignPtr(&r.RecordsWithoutRefresh)
	return softTypeError(err)
}
