// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// DealsService groups the deals operations. Use it as client.Deals.
type DealsService struct {
	client *Client
}

// Absentee: Find absentee owners
//
// Retrieve parcels owned by absentee owners, useful for off-market deal sourcing.
//
// HTTP: GET /api/v1/deals/absentee
func (s *DealsService) Absentee(ctx context.Context, params *DealsAbsenteeParams, opts ...RequestOption) (*DealsAbsenteeResponse, error) {
	var out DealsAbsenteeResponse
	if err := s.client.do(ctx, buildDealsAbsenteeRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// AbsenteeIter iterates every item of [DealsService.Absentee] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) AbsenteeIter(ctx context.Context, params *DealsAbsenteeParams, iter IterOptions, opts ...RequestOption) *Iter[DealsAbsenteeResponseData] {
	var p DealsAbsenteeParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[DealsAbsenteeResponseData](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsAbsenteeRequest(&q)
	}, opts)
}

func buildDealsAbsenteeRequest(params *DealsAbsenteeParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/absentee")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "out_of_state", params.OutOfState)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsAbsenteeParams holds the query, header and JSON-body parameters of [DealsService.Absentee].
// Pass nil when you need none.
type DealsAbsenteeParams struct {
	// Filter by county FIPS code.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// Filter by state FIPS code.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Minimum assessed value.
	MinValue *float64 `query:"min_value" json:"-"`

	// Only return owners whose mailing address is in a different state.
	OutOfState *bool  `query:"out_of_state" json:"-"`
	Limit      *int64 `query:"limit" json:"-"`
	Offset     *int64 `query:"offset" json:"-"`
}

// DealsAbsenteeResponse: Find absentee owners
type DealsAbsenteeResponse struct {
	Data   []DealsAbsenteeResponseData `json:"data"`
	Total  int64                       `json:"total"`
	Limit  int64                       `json:"limit"`
	Offset int64                       `json:"offset"`
}

// UnmarshalJSON decodes DealsAbsenteeResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsAbsenteeResponse) UnmarshalJSON(data []byte) error {
	type plain DealsAbsenteeResponse
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

// DealsAbsenteeResponseData is generated from the OpenAPI spec.
type DealsAbsenteeResponseData struct {
	CountyFIPS         string   `json:"county_fips"`
	StateFIPS          string   `json:"state_fips"`
	ParcelID           string   `json:"parcel_id"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	PropertyAddress    *string  `json:"property_address,omitempty"`
	PropertyCity       *string  `json:"property_city,omitempty"`
	PropertyState      *string  `json:"property_state,omitempty"`
	OwnerAddress       *string  `json:"owner_address,omitempty"`
	OwnerCity          *string  `json:"owner_city,omitempty"`
	OwnerState         *string  `json:"owner_state,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	LastSaleDate       *string  `json:"last_sale_date,omitempty"`
	LastSalePrice      *float64 `json:"last_sale_price,omitempty"`
	IsOutOfState       *bool    `json:"is_out_of_state,omitempty"`
}

// UnmarshalJSON decodes DealsAbsenteeResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsAbsenteeResponseData) UnmarshalJSON(data []byte) error {
	type plain DealsAbsenteeResponseData
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		LastSalePrice      lenientNumber[float64] `json:"last_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	return softTypeError(err)
}

// Flips: Find property flips
//
// Retrieve recently flipped properties. Use ?view=flippers to get a ranked list of top flippers
// instead.
//
// HTTP: GET /api/v1/deals/flips
func (s *DealsService) Flips(ctx context.Context, params *DealsFlipsParams, opts ...RequestOption) (*DealsFlipsResponse, error) {
	var out DealsFlipsResponse
	if err := s.client.do(ctx, buildDealsFlipsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// FlipsIter iterates every item of [DealsService.Flips] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) FlipsIter(ctx context.Context, params *DealsFlipsParams, iter IterOptions, opts ...RequestOption) *Iter[DealsFlipsResponseData] {
	var p DealsFlipsParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[DealsFlipsResponseData](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsFlipsRequest(&q)
	}, opts)
}

func buildDealsFlipsRequest(params *DealsFlipsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/flips")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "flip_tier", params.FlipTier)
		addQuery(req.query, "min_profit", params.MinProfit)
		addQuery(req.query, "view", params.View)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsFlipsParams holds the query, header and JSON-body parameters of [DealsService.Flips]. Pass
// nil when you need none.
type DealsFlipsParams struct {
	// Filter by county FIPS code.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// Filter by state FIPS code.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Filter by hold time between the two sales: QUICK_FLIP (< 180 days), SHORT_HOLD (180–364 days),
	// MEDIUM_HOLD (365 days to 24 months). Case-insensitive. The former spellings quick / standard /
	// long are accepted as deprecated aliases. Any other value is a 400.
	FlipTier *DealsFlipsParamsFlipTier `query:"flip_tier" json:"-"`

	// Minimum estimated profit.
	MinProfit *float64 `query:"min_profit" json:"-"`

	// Set to 'flippers' to return a ranked list of top flippers instead of individual flips.
	View   *DealsFlipsParamsView `query:"view" json:"-"`
	Limit  *int64                `query:"limit" json:"-"`
	Offset *int64                `query:"offset" json:"-"`
}

// DealsFlipsParamsFlipTier is generated from the OpenAPI spec. It is a string; the
// DealsFlipsParamsFlipTier* constants list the documented values.
type DealsFlipsParamsFlipTier = string

// Documented values of DealsFlipsParamsFlipTier.
const (
	DealsFlipsParamsFlipTierQuickFlip  DealsFlipsParamsFlipTier = "QUICK_FLIP"
	DealsFlipsParamsFlipTierShortHold  DealsFlipsParamsFlipTier = "SHORT_HOLD"
	DealsFlipsParamsFlipTierMediumHold DealsFlipsParamsFlipTier = "MEDIUM_HOLD"
)

// DealsFlipsParamsView is generated from the OpenAPI spec. It is a string; the
// DealsFlipsParamsView* constants list the documented values.
type DealsFlipsParamsView = string

// Documented values of DealsFlipsParamsView.
const (
	DealsFlipsParamsViewFlippers DealsFlipsParamsView = "flippers"
)

// DealsFlipsResponse: Find property flips
type DealsFlipsResponse struct {
	Data   []DealsFlipsResponseData `json:"data"`
	Total  int64                    `json:"total"`
	Limit  int64                    `json:"limit"`
	Offset int64                    `json:"offset"`
}

// UnmarshalJSON decodes DealsFlipsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsFlipsResponse) UnmarshalJSON(data []byte) error {
	type plain DealsFlipsResponse
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

// DealsFlipsResponseData is generated from the OpenAPI spec.
type DealsFlipsResponseData struct {
	CountyFIPS string   `json:"county_fips"`
	StateFIPS  string   `json:"state_fips"`
	ParcelID   string   `json:"parcel_id"`
	Address    *string  `json:"address,omitempty"`
	City       *string  `json:"city,omitempty"`
	State      *string  `json:"state,omitempty"`
	BuyDate    *string  `json:"buy_date,omitempty"`
	BuyPrice   *float64 `json:"buy_price,omitempty"`
	BuyerName  *string  `json:"buyer_name,omitempty"`
	SellDate   *string  `json:"sell_date,omitempty"`
	SellPrice  *float64 `json:"sell_price,omitempty"`
	SellerName *string  `json:"seller_name,omitempty"`
	HoldDays   *int64   `json:"hold_days,omitempty"`
	Profit     *float64 `json:"profit,omitempty"`
	ProfitPct  *float64 `json:"profit_pct,omitempty"`
	FlipTier   *string  `json:"flip_tier,omitempty"`
}

// UnmarshalJSON decodes DealsFlipsResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsFlipsResponseData) UnmarshalJSON(data []byte) error {
	type plain DealsFlipsResponseData
	aux := struct {
		*plain
		BuyPrice  lenientNumber[float64] `json:"buy_price"`
		SellPrice lenientNumber[float64] `json:"sell_price"`
		HoldDays  lenientNumber[int64]   `json:"hold_days"`
		Profit    lenientNumber[float64] `json:"profit"`
		ProfitPct lenientNumber[float64] `json:"profit_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.BuyPrice.assignPtr(&r.BuyPrice)
	aux.SellPrice.assignPtr(&r.SellPrice)
	aux.HoldDays.assignPtr(&r.HoldDays)
	aux.Profit.assignPtr(&r.Profit)
	aux.ProfitPct.assignPtr(&r.ProfitPct)
	return softTypeError(err)
}

// Contractors: Search contractors by permit activity
//
// Returns contractor profiles aggregated from 45M+ building permits. Each profile includes permit
// count, jurisdictions worked, total declared permit value, and activity dates. Use to identify
// active contractors in a market or find a specific contractor by name.
//
// HTTP: GET /api/v1/deals/contractors
func (s *DealsService) Contractors(ctx context.Context, params *DealsContractorsParams, opts ...RequestOption) (*DealsContractorsResponse, error) {
	var out DealsContractorsResponse
	if err := s.client.do(ctx, buildDealsContractorsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// ContractorsIter iterates every item of [DealsService.Contractors] across pages (offset
// pagination over "data"): it advances offset by the page size and stops on a short page, when
// offset reaches total, or when has_more is false. Use [IterOptions] for the page size and an item
// cap.
func (s *DealsService) ContractorsIter(ctx context.Context, params *DealsContractorsParams, iter IterOptions, opts ...RequestOption) *Iter[Contractor] {
	var p DealsContractorsParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[Contractor](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsContractorsRequest(&q)
	}, opts)
}

func buildDealsContractorsRequest(params *DealsContractorsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/contractors")
	if params != nil {
		addQuery(req.query, "search", params.Search)
		addQuery(req.query, "min_permits", params.MinPermits)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsContractorsParams holds the query, header and JSON-body parameters of
// [DealsService.Contractors]. Pass nil when you need none.
type DealsContractorsParams struct {
	// Contractor name search (case-insensitive substring).
	Search *string `query:"search" json:"-"`

	// Minimum permit count to include.
	MinPermits *int64 `query:"min_permits" json:"-"`

	// 2-letter state filter.
	State *string `query:"state" json:"-"`

	// Minimum total declared permit value, USD.
	MinValue *int64 `query:"min_value" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsContractorsResponse: Search contractors by permit activity
type DealsContractorsResponse struct {
	Data          []Contractor                          `json:"data"`
	Total         int64                                 `json:"total"`
	Limit         int64                                 `json:"limit"`
	Offset        int64                                 `json:"offset"`
	SourceQuality DealsContractorsResponseSourceQuality `json:"source_quality"`
}

// UnmarshalJSON decodes DealsContractorsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsContractorsResponse) UnmarshalJSON(data []byte) error {
	type plain DealsContractorsResponse
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

// DealsContractorsResponseSourceQuality is generated from the OpenAPI spec.
type DealsContractorsResponseSourceQuality struct {
	GeographyStatus          string `json:"geography_status"`
	ProfileRefreshObservedAt string `json:"profile_refresh_observed_at"`
	QualityCheckedAt         string `json:"quality_checked_at"`
	Scope                    string `json:"scope"`
	Reason                   string `json:"reason"`
}

// Entities: Find entity-owned parcels (LLC, Corp, Trust, LP)
//
// Returns parcels owned by legal entities identified from owner-name pattern matching across 221M+
// parcels. Pass `top=true` to get aggregated entity rankings instead of per-parcel rows. One of
// `county_fips`, `state_fips`, `search`, or `top` is required.
//
// HTTP: GET /api/v1/deals/entities
func (s *DealsService) Entities(ctx context.Context, params *DealsEntitiesParams, opts ...RequestOption) (*DealsEntitiesResponse, error) {
	var out DealsEntitiesResponse
	if err := s.client.do(ctx, buildDealsEntitiesRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// EntitiesIter iterates every item of [DealsService.Entities] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) EntitiesIter(ctx context.Context, params *DealsEntitiesParams, iter IterOptions, opts ...RequestOption) *Iter[json.RawMessage] {
	var p DealsEntitiesParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[json.RawMessage](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsEntitiesRequest(&q)
	}, opts)
}

func buildDealsEntitiesRequest(params *DealsEntitiesParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/entities")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "entity_type", params.EntityType)
		addQuery(req.query, "search", params.Search)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "zoning", params.Zoning)
		addQuery(req.query, "top", params.Top)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsEntitiesParams holds the query, header and JSON-body parameters of [DealsService.Entities].
// Pass nil when you need none.
type DealsEntitiesParams struct {
	// 5-digit county FIPS filter.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// 2-digit state FIPS filter.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Filter by entity classification. Case-insensitive; any other value is a 400.
	EntityType *DealsEntitiesParamsEntityType `query:"entity_type" json:"-"`

	// Owner-name substring search.
	Search *string `query:"search" json:"-"`

	// Minimum assessed value, USD.
	MinValue *int64 `query:"min_value" json:"-"`

	// Zoning substring filter.
	Zoning *string `query:"zoning" json:"-"`

	// If true, returns aggregated entity rankings with summary stats instead of per-parcel rows.
	Top *bool `query:"top" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsEntitiesParamsEntityType is generated from the OpenAPI spec. It is a string; the
// DealsEntitiesParamsEntityType* constants list the documented values.
type DealsEntitiesParamsEntityType = string

// Documented values of DealsEntitiesParamsEntityType.
const (
	DealsEntitiesParamsEntityTypeLLC         DealsEntitiesParamsEntityType = "LLC"
	DealsEntitiesParamsEntityTypeCorp        DealsEntitiesParamsEntityType = "CORP"
	DealsEntitiesParamsEntityTypeTrust       DealsEntitiesParamsEntityType = "TRUST"
	DealsEntitiesParamsEntityTypeLp          DealsEntitiesParamsEntityType = "LP"
	DealsEntitiesParamsEntityTypeLtd         DealsEntitiesParamsEntityType = "LTD"
	DealsEntitiesParamsEntityTypeAssociation DealsEntitiesParamsEntityType = "ASSOCIATION"
	DealsEntitiesParamsEntityTypeOtherEntity DealsEntitiesParamsEntityType = "OTHER_ENTITY"
)

// DealsEntitiesResponse: Find entity-owned parcels (LLC, Corp, Trust, LP)
type DealsEntitiesResponse = json.RawMessage

// HighLandRatio: Find parcels with high land-to-improvement ratio
//
// Returns parcels where land value significantly exceeds improvement value — a signal for
// redevelopment, teardown, or assemblage opportunities. `county_fips` or `state_fips` is required.
//
// HTTP: GET /api/v1/deals/high-land-ratio
func (s *DealsService) HighLandRatio(ctx context.Context, params *DealsHighLandRatioParams, opts ...RequestOption) (*DealsHighLandRatioResponse, error) {
	var out DealsHighLandRatioResponse
	if err := s.client.do(ctx, buildDealsHighLandRatioRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// HighLandRatioIter iterates every item of [DealsService.HighLandRatio] across pages (offset
// pagination over "data"): it advances offset by the page size and stops on a short page, when
// offset reaches total, or when has_more is false. Use [IterOptions] for the page size and an item
// cap.
func (s *DealsService) HighLandRatioIter(ctx context.Context, params *DealsHighLandRatioParams, iter IterOptions, opts ...RequestOption) *Iter[HighLandRatioParcel] {
	var p DealsHighLandRatioParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[HighLandRatioParcel](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsHighLandRatioRequest(&q)
	}, opts)
}

func buildDealsHighLandRatioRequest(params *DealsHighLandRatioParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/high-land-ratio")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "min_ratio", params.MinRatio)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "zoning", params.Zoning)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsHighLandRatioParams holds the query, header and JSON-body parameters of
// [DealsService.HighLandRatio]. Pass nil when you need none.
type DealsHighLandRatioParams struct {
	// 5-digit county FIPS filter.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// 2-digit state FIPS filter.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Minimum land/improvement ratio.
	MinRatio *float64 `query:"min_ratio" json:"-"`

	// Minimum land assessed value, USD.
	MinValue *int64 `query:"min_value" json:"-"`

	// Zoning substring filter.
	Zoning *string `query:"zoning" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsHighLandRatioResponse: Find parcels with high land-to-improvement ratio
type DealsHighLandRatioResponse struct {
	Data   []HighLandRatioParcel `json:"data"`
	Total  int64                 `json:"total"`
	Limit  int64                 `json:"limit"`
	Offset int64                 `json:"offset"`
}

// UnmarshalJSON decodes DealsHighLandRatioResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsHighLandRatioResponse) UnmarshalJSON(data []byte) error {
	type plain DealsHighLandRatioResponse
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

// Lenders: Search lender profiles
//
// Returns lender profiles aggregated from deed/mortgage transactions. Includes mortgage count,
// total volume, geographic spread, and a national rank.
//
// HTTP: GET /api/v1/deals/lenders
func (s *DealsService) Lenders(ctx context.Context, params *DealsLendersParams, opts ...RequestOption) (*DealsLendersResponse, error) {
	var out DealsLendersResponse
	if err := s.client.do(ctx, buildDealsLendersRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// LendersIter iterates every item of [DealsService.Lenders] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) LendersIter(ctx context.Context, params *DealsLendersParams, iter IterOptions, opts ...RequestOption) *Iter[Lender] {
	var p DealsLendersParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[Lender](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsLendersRequest(&q)
	}, opts)
}

func buildDealsLendersRequest(params *DealsLendersParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/lenders")
	if params != nil {
		addQuery(req.query, "search", params.Search)
		addQuery(req.query, "min_mortgages", params.MinMortgages)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsLendersParams holds the query, header and JSON-body parameters of [DealsService.Lenders].
// Pass nil when you need none.
type DealsLendersParams struct {
	// Lender name substring search.
	Search *string `query:"search" json:"-"`

	// Minimum mortgage count.
	MinMortgages *int64 `query:"min_mortgages" json:"-"`

	// 2-letter state filter.
	State *string `query:"state" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsLendersResponse: Search lender profiles
type DealsLendersResponse struct {
	Data   []Lender `json:"data,omitempty"`
	Total  *int64   `json:"total,omitempty"`
	Limit  *int64   `json:"limit,omitempty"`
	Offset *int64   `json:"offset,omitempty"`
}

// UnmarshalJSON decodes DealsLendersResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsLendersResponse) UnmarshalJSON(data []byte) error {
	type plain DealsLendersResponse
	aux := struct {
		*plain
		Total  lenientNumber[int64] `json:"total"`
		Limit  lenientNumber[int64] `json:"limit"`
		Offset lenientNumber[int64] `json:"offset"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assignPtr(&r.Total)
	aux.Limit.assignPtr(&r.Limit)
	aux.Offset.assignPtr(&r.Offset)
	return softTypeError(err)
}

// LongHold: Find long-held parcels (10+ years)
//
// Returns parcels not sold in `min_years` or more. Long-hold owners are often motivated sellers
// — estate planning, deferred maintenance, life changes. `county_fips` or `state_fips` is
// required.
//
// HTTP: GET /api/v1/deals/long-hold
func (s *DealsService) LongHold(ctx context.Context, params *DealsLongHoldParams, opts ...RequestOption) (*DealsLongHoldResponse, error) {
	var out DealsLongHoldResponse
	if err := s.client.do(ctx, buildDealsLongHoldRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// LongHoldIter iterates every item of [DealsService.LongHold] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) LongHoldIter(ctx context.Context, params *DealsLongHoldParams, iter IterOptions, opts ...RequestOption) *Iter[LongHoldParcel] {
	var p DealsLongHoldParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[LongHoldParcel](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsLongHoldRequest(&q)
	}, opts)
}

func buildDealsLongHoldRequest(params *DealsLongHoldParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/long-hold")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "min_years", params.MinYears)
		addQuery(req.query, "hold_tier", params.HoldTier)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsLongHoldParams holds the query, header and JSON-body parameters of [DealsService.LongHold].
// Pass nil when you need none.
type DealsLongHoldParams struct {
	// 5-digit county FIPS filter.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// 2-digit state FIPS filter.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Minimum years held.
	MinYears *int64 `query:"min_years" json:"-"`

	// Filter by hold-period tier. Case-insensitive; any other value is a 400.
	HoldTier *DealsLongHoldParamsHoldTier `query:"hold_tier" json:"-"`

	// Minimum assessed value, USD.
	MinValue *int64 `query:"min_value" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsLongHoldParamsHoldTier is generated from the OpenAPI spec. It is a string; the
// DealsLongHoldParamsHoldTier* constants list the documented values.
type DealsLongHoldParamsHoldTier = string

// Documented values of DealsLongHoldParamsHoldTier.
const (
	DealsLongHoldParamsHoldTierV1015yr   DealsLongHoldParamsHoldTier = "10-15yr"
	DealsLongHoldParamsHoldTierV1520yr   DealsLongHoldParamsHoldTier = "15-20yr"
	DealsLongHoldParamsHoldTierV2030yr   DealsLongHoldParamsHoldTier = "20-30yr"
	DealsLongHoldParamsHoldTierV30yrPlus DealsLongHoldParamsHoldTier = "30yr+"
)

// DealsLongHoldResponse: Find long-held parcels (10+ years)
type DealsLongHoldResponse struct {
	Data   []LongHoldParcel `json:"data"`
	Total  int64            `json:"total"`
	Limit  int64            `json:"limit"`
	Offset int64            `json:"offset"`
}

// UnmarshalJSON decodes DealsLongHoldResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsLongHoldResponse) UnmarshalJSON(data []byte) error {
	type plain DealsLongHoldResponse
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

// Market: County-quarter transaction summary or affordability index
//
// Default: returns county/quarter transaction summaries. Pass `view=affordability` to retrieve the
// home-affordability index instead (price-to-income ratios + rating).
//
// HTTP: GET /api/v1/deals/market
func (s *DealsService) Market(ctx context.Context, params *DealsMarketParams, opts ...RequestOption) (*DealsMarketResponse, error) {
	var out DealsMarketResponse
	if err := s.client.do(ctx, buildDealsMarketRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// MarketIter iterates every item of [DealsService.Market] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *DealsService) MarketIter(ctx context.Context, params *DealsMarketParams, iter IterOptions, opts ...RequestOption) *Iter[json.RawMessage] {
	var p DealsMarketParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[json.RawMessage](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsMarketRequest(&q)
	}, opts)
}

func buildDealsMarketRequest(params *DealsMarketParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/market")
	if params != nil {
		addQuery(req.query, "view", params.View)
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "state_fips", params.StateFIPS)
		addQuery(req.query, "year", params.Year)
		addQuery(req.query, "rating", params.Rating)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsMarketParams holds the query, header and JSON-body parameters of [DealsService.Market].
// Pass nil when you need none.
type DealsMarketParams struct {
	// Switch to the affordability-index dataset.
	View *DealsMarketParamsView `query:"view" json:"-"`

	// 5-digit county FIPS filter.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// 2-digit state FIPS filter.
	StateFIPS *string `query:"state_fips" json:"-"`

	// Year filter.
	Year *string `query:"year" json:"-"`

	// Affordability-rating filter (only meaningful with view=affordability). Case-insensitive; any
	// other value is a 400.
	Rating *DealsMarketParamsRating `query:"rating" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsMarketParamsView is generated from the OpenAPI spec. It is a string; the
// DealsMarketParamsView* constants list the documented values.
type DealsMarketParamsView = string

// Documented values of DealsMarketParamsView.
const (
	DealsMarketParamsViewAffordability DealsMarketParamsView = "affordability"
)

// DealsMarketParamsRating is generated from the OpenAPI spec. It is a string; the
// DealsMarketParamsRating* constants list the documented values.
type DealsMarketParamsRating = string

// Documented values of DealsMarketParamsRating.
const (
	DealsMarketParamsRatingAffordable    DealsMarketParamsRating = "AFFORDABLE"
	DealsMarketParamsRatingModerate      DealsMarketParamsRating = "MODERATE"
	DealsMarketParamsRatingExpensive     DealsMarketParamsRating = "EXPENSIVE"
	DealsMarketParamsRatingVeryExpensive DealsMarketParamsRating = "VERY_EXPENSIVE"
)

// DealsMarketResponse: County-quarter transaction summary or affordability index
type DealsMarketResponse = json.RawMessage

// PortfolioOwners: Find portfolio investors (owners of 2+ properties)
//
// Returns portfolio owners ranked by property count and total assessed value. Useful for finding
// institutional buyers, small landlords, or specific investor families.
//
// HTTP: GET /api/v1/deals/portfolio-owners
func (s *DealsService) PortfolioOwners(ctx context.Context, params *DealsPortfolioOwnersParams, opts ...RequestOption) (*DealsPortfolioOwnersResponse, error) {
	var out DealsPortfolioOwnersResponse
	if err := s.client.do(ctx, buildDealsPortfolioOwnersRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// PortfolioOwnersIter iterates every item of [DealsService.PortfolioOwners] across pages (offset
// pagination over "data"): it advances offset by the page size and stops on a short page, when
// offset reaches total, or when has_more is false. Use [IterOptions] for the page size and an item
// cap.
func (s *DealsService) PortfolioOwnersIter(ctx context.Context, params *DealsPortfolioOwnersParams, iter IterOptions, opts ...RequestOption) *Iter[PortfolioOwner] {
	var p DealsPortfolioOwnersParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[PortfolioOwner](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildDealsPortfolioOwnersRequest(&q)
	}, opts)
}

func buildDealsPortfolioOwnersRequest(params *DealsPortfolioOwnersParams) *apiRequest {
	req := newRequest("GET", "/api/v1/deals/portfolio-owners")
	if params != nil {
		addQuery(req.query, "min_properties", params.MinProperties)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "min_value", params.MinValue)
		addQuery(req.query, "search", params.Search)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// DealsPortfolioOwnersParams holds the query, header and JSON-body parameters of
// [DealsService.PortfolioOwners]. Pass nil when you need none.
type DealsPortfolioOwnersParams struct {
	// Minimum properties owned.
	MinProperties *int64 `query:"min_properties" json:"-"`

	// 2-letter owner mailing state filter.
	State *string `query:"state" json:"-"`

	// Minimum total portfolio assessed value, USD.
	MinValue *int64 `query:"min_value" json:"-"`

	// Owner name substring search.
	Search *string `query:"search" json:"-"`

	// Page size, max 500.
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset.
	Offset *int64 `query:"offset" json:"-"`
}

// DealsPortfolioOwnersResponse: Find portfolio investors (owners of 2+ properties)
type DealsPortfolioOwnersResponse struct {
	Data   []PortfolioOwner `json:"data"`
	Total  int64            `json:"total"`
	Limit  int64            `json:"limit"`
	Offset int64            `json:"offset"`
}

// UnmarshalJSON decodes DealsPortfolioOwnersResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DealsPortfolioOwnersResponse) UnmarshalJSON(data []byte) error {
	type plain DealsPortfolioOwnersResponse
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
