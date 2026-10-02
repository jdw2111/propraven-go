// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// SearchService groups the search operations. Use it as client.Search.
type SearchService struct {
	client *Client
}

// Parcels: Search parcels
//
// Search parcels with optional geographic bounds and attribute filters. `bounds` and `filters` are
// both optional (`filters` defaults to `{}`); the body is validated and any bad member is a 400
// naming it. `sort` applies to UNBOUNDED queries only: with `bounds`, rows come back in
// spatial-index order and the response says `sort_applied: false`. `total` is the exact number of
// matching parcels up to 10,000; beyond that it is 10,000 with `total_is_lower_bound: true`, and
// it is null if the count timed out (`total_status: "timed_out"`). Traffic-count columns are
// withheld until verified: they are returned as null with a `withhold_gate` block, and filters on
// them (minVpd, maxVpd, minVisibilityScore) are refused with 400 `filter_withheld`.
//
// HTTP: POST /api/v1/search
func (s *SearchService) Parcels(ctx context.Context, params *SearchParcelsParams, opts ...RequestOption) (*SearchParcelsResponse, error) {
	var out SearchParcelsResponse
	if err := s.client.do(ctx, buildSearchParcelsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// ParcelsIter iterates every item of [SearchService.Parcels] across pages (offset pagination over
// "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *SearchService) ParcelsIter(ctx context.Context, params *SearchParcelsParams, iter IterOptions, opts ...RequestOption) *Iter[SearchParcelsResponseData] {
	var p SearchParcelsParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[SearchParcelsResponseData](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildSearchParcelsRequest(&q)
	}, opts)
}

func buildSearchParcelsRequest(params *SearchParcelsParams) *apiRequest {
	req := newRequest("POST", "/api/v1/search")
	if params != nil {
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// SearchParcelsParams holds the query, header and JSON-body parameters of [SearchService.Parcels].
// Pass nil when you need none.
type SearchParcelsParams struct {
	// Viewport in WGS84 degrees. north > south. west > east is an antimeridian-crossing box.
	Bounds *SearchParcelsParamsBounds `json:"bounds,omitempty"`

	// All optional; null means not set. Unknown members are a 400.
	Filters *SearchParcelsParamsFilters `json:"filters,omitempty"`

	// Sort field (unbounded queries only). The first five are the primary names; the rest are accepted
	// column names. A legacy `{field, direction}` object is also accepted.
	Sort  *SearchParcelsParamsSort  `json:"sort,omitempty"`
	Order *SearchParcelsParamsOrder `json:"order,omitempty"`

	// Page size. Values above 500 are clamped to 500 (the response echoes the effective limit).
	Limit  *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
}

// SearchParcelsParamsBounds: Viewport in WGS84 degrees. north > south. west > east is an
// antimeridian-crossing box.
type SearchParcelsParamsBounds struct {
	North float64 `json:"north"`
	South float64 `json:"south"`
	East  float64 `json:"east"`
	West  float64 `json:"west"`
}

// UnmarshalJSON decodes SearchParcelsParamsBounds, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsParamsBounds) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsParamsBounds
	aux := struct {
		*plain
		North lenientNumber[float64] `json:"north"`
		South lenientNumber[float64] `json:"south"`
		East  lenientNumber[float64] `json:"east"`
		West  lenientNumber[float64] `json:"west"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.North.assign(&r.North)
	aux.South.assign(&r.South)
	aux.East.assign(&r.East)
	aux.West.assign(&r.West)
	return softTypeError(err)
}

// SearchParcelsParamsFilters: All optional; null means not set. Unknown members are a 400.
type SearchParcelsParamsFilters struct {
	// Zoning categories (Residential, Commercial, Industrial, Mixed-Use, Agricultural, …).
	// Unrecognised labels are reported in `zoning_categories_unrecognized`.
	ZoningCategories []string `json:"zoningCategories,omitempty"`

	// Owner entity type (case-insensitive).
	OwnerTypes []SearchParcelsParamsFiltersOwnerTypes `json:"ownerTypes,omitempty"`

	// Property type group, e.g. Commercial, Residential.
	PropertyType []string `json:"propertyType,omitempty"`

	// Business / use type, e.g. car_wash.
	BusinessTypes []string `json:"businessTypes,omitempty"`

	// Only parcels whose owner state differs from the parcel state.
	AbsenteeOnly *bool                                 `json:"absenteeOnly,omitempty"`
	SoldWithin   *SearchParcelsParamsFiltersSoldWithin `json:"soldWithin,omitempty"`
	FloodZone    *string                               `json:"floodZone,omitempty"`

	// Assessed value range.
	ValueRange *SearchParcelsParamsFiltersValueRange `json:"valueRange,omitempty"`

	// Year built range.
	YearBuiltRange *SearchParcelsParamsFiltersYearBuiltRange `json:"yearBuiltRange,omitempty"`

	// Lot size range in acres.
	AcreageRange  *SearchParcelsParamsFiltersAcreageRange `json:"acreageRange,omitempty"`
	MinYearsOwned *float64                                `json:"minYearsOwned,omitempty"`
	MinCrimeScore *float64                                `json:"minCrimeScore,omitempty"`
	MaxCrimeScore *float64                                `json:"maxCrimeScore,omitempty"`
	CrimeTrend    *string                                 `json:"crimeTrend,omitempty"`

	// WITHHELD — refused with 400 filter_withheld until traffic counts are verified.
	MinVpd *float64 `json:"minVpd,omitempty"`

	// WITHHELD — refused with 400 filter_withheld.
	MaxVpd *float64 `json:"maxVpd,omitempty"`

	// WITHHELD — refused with 400 filter_withheld.
	MinVisibilityScore *float64 `json:"minVisibilityScore,omitempty"`
}

// UnmarshalJSON decodes SearchParcelsParamsFilters, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsParamsFilters) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsParamsFilters
	aux := struct {
		*plain
		MinYearsOwned      lenientNumber[float64] `json:"minYearsOwned"`
		MinCrimeScore      lenientNumber[float64] `json:"minCrimeScore"`
		MaxCrimeScore      lenientNumber[float64] `json:"maxCrimeScore"`
		MinVpd             lenientNumber[float64] `json:"minVpd"`
		MaxVpd             lenientNumber[float64] `json:"maxVpd"`
		MinVisibilityScore lenientNumber[float64] `json:"minVisibilityScore"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MinYearsOwned.assignPtr(&r.MinYearsOwned)
	aux.MinCrimeScore.assignPtr(&r.MinCrimeScore)
	aux.MaxCrimeScore.assignPtr(&r.MaxCrimeScore)
	aux.MinVpd.assignPtr(&r.MinVpd)
	aux.MaxVpd.assignPtr(&r.MaxVpd)
	aux.MinVisibilityScore.assignPtr(&r.MinVisibilityScore)
	return softTypeError(err)
}

// SearchParcelsParamsFiltersOwnerTypes is generated from the OpenAPI spec. It is a string; the
// SearchParcelsParamsFiltersOwnerTypes* constants list the documented values.
type SearchParcelsParamsFiltersOwnerTypes = string

// Documented values of SearchParcelsParamsFiltersOwnerTypes.
const (
	SearchParcelsParamsFiltersOwnerTypesIndividual  SearchParcelsParamsFiltersOwnerTypes = "individual"
	SearchParcelsParamsFiltersOwnerTypesLLC         SearchParcelsParamsFiltersOwnerTypes = "llc"
	SearchParcelsParamsFiltersOwnerTypesTrust       SearchParcelsParamsFiltersOwnerTypes = "trust"
	SearchParcelsParamsFiltersOwnerTypesCorporation SearchParcelsParamsFiltersOwnerTypes = "corporation"
	SearchParcelsParamsFiltersOwnerTypesGovernment  SearchParcelsParamsFiltersOwnerTypes = "government"
	SearchParcelsParamsFiltersOwnerTypesReligious   SearchParcelsParamsFiltersOwnerTypes = "religious"
	SearchParcelsParamsFiltersOwnerTypesPartnership SearchParcelsParamsFiltersOwnerTypes = "partnership"
	SearchParcelsParamsFiltersOwnerTypesAssociation SearchParcelsParamsFiltersOwnerTypes = "association"
	SearchParcelsParamsFiltersOwnerTypesFinancial   SearchParcelsParamsFiltersOwnerTypes = "financial"
)

// SearchParcelsParamsFiltersSoldWithin is generated from the OpenAPI spec. It is a string; the
// SearchParcelsParamsFiltersSoldWithin* constants list the documented values.
type SearchParcelsParamsFiltersSoldWithin = string

// Documented values of SearchParcelsParamsFiltersSoldWithin.
const (
	SearchParcelsParamsFiltersSoldWithinV6mo  SearchParcelsParamsFiltersSoldWithin = "6mo"
	SearchParcelsParamsFiltersSoldWithinV1yr  SearchParcelsParamsFiltersSoldWithin = "1yr"
	SearchParcelsParamsFiltersSoldWithinV3yr  SearchParcelsParamsFiltersSoldWithin = "3yr"
	SearchParcelsParamsFiltersSoldWithinNever SearchParcelsParamsFiltersSoldWithin = "never"
)

// SearchParcelsParamsFiltersValueRange: Assessed value range.
type SearchParcelsParamsFiltersValueRange struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// UnmarshalJSON decodes SearchParcelsParamsFiltersValueRange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsParamsFiltersValueRange) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsParamsFiltersValueRange
	aux := struct {
		*plain
		Min lenientNumber[float64] `json:"min"`
		Max lenientNumber[float64] `json:"max"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Min.assignPtr(&r.Min)
	aux.Max.assignPtr(&r.Max)
	return softTypeError(err)
}

// SearchParcelsParamsFiltersYearBuiltRange: Year built range.
type SearchParcelsParamsFiltersYearBuiltRange struct {
	Min *int64 `json:"min,omitempty"`
	Max *int64 `json:"max,omitempty"`
}

// UnmarshalJSON decodes SearchParcelsParamsFiltersYearBuiltRange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsParamsFiltersYearBuiltRange) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsParamsFiltersYearBuiltRange
	aux := struct {
		*plain
		Min lenientNumber[int64] `json:"min"`
		Max lenientNumber[int64] `json:"max"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Min.assignPtr(&r.Min)
	aux.Max.assignPtr(&r.Max)
	return softTypeError(err)
}

// SearchParcelsParamsFiltersAcreageRange: Lot size range in acres.
type SearchParcelsParamsFiltersAcreageRange struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// UnmarshalJSON decodes SearchParcelsParamsFiltersAcreageRange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsParamsFiltersAcreageRange) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsParamsFiltersAcreageRange
	aux := struct {
		*plain
		Min lenientNumber[float64] `json:"min"`
		Max lenientNumber[float64] `json:"max"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Min.assignPtr(&r.Min)
	aux.Max.assignPtr(&r.Max)
	return softTypeError(err)
}

// SearchParcelsParamsSort: Sort field (unbounded queries only). The first five are the primary
// names; the rest are accepted column names. A legacy `{field, direction}` object is also
// accepted.
//
// It is a string; the SearchParcelsParamsSort* constants list the documented values.
type SearchParcelsParamsSort = string

// Documented values of SearchParcelsParamsSort.
const (
	SearchParcelsParamsSortAssessedValue      SearchParcelsParamsSort = "assessed_value"
	SearchParcelsParamsSortSalePrice          SearchParcelsParamsSort = "sale_price"
	SearchParcelsParamsSortAcreage            SearchParcelsParamsSort = "acreage"
	SearchParcelsParamsSortYearBuilt          SearchParcelsParamsSort = "year_built"
	SearchParcelsParamsSortDealScore          SearchParcelsParamsSort = "deal_score"
	SearchParcelsParamsSortAddress            SearchParcelsParamsSort = "address"
	SearchParcelsParamsSortCity               SearchParcelsParamsSort = "city"
	SearchParcelsParamsSortState              SearchParcelsParamsSort = "state"
	SearchParcelsParamsSortOwnerName          SearchParcelsParamsSort = "owner_name"
	SearchParcelsParamsSortTotalAssessedValue SearchParcelsParamsSort = "total_assessed_value"
	SearchParcelsParamsSortZoning             SearchParcelsParamsSort = "zoning"
	SearchParcelsParamsSortLotSizeAcres       SearchParcelsParamsSort = "lot_size_acres"
	SearchParcelsParamsSortOwnershipType      SearchParcelsParamsSort = "ownership_type"
)

// SearchParcelsParamsOrder is generated from the OpenAPI spec. It is a string; the
// SearchParcelsParamsOrder* constants list the documented values.
type SearchParcelsParamsOrder = string

// Documented values of SearchParcelsParamsOrder.
const (
	SearchParcelsParamsOrderAsc  SearchParcelsParamsOrder = "asc"
	SearchParcelsParamsOrderDesc SearchParcelsParamsOrder = "desc"
)

// SearchParcelsResponse: Search parcels
type SearchParcelsResponse struct {
	Data []SearchParcelsResponseData `json:"data"`

	// Exact matching count up to 10,000; 10,000 when capped; null when the count timed out.
	Total int64 `json:"total"`

	// True only when `total` is capped (a lower bound).
	TotalIsEstimate   bool  `json:"total_is_estimate"`
	TotalIsLowerBound bool  `json:"total_is_lower_bound"`
	HasMore           bool  `json:"has_more"`
	Limit             int64 `json:"limit"`
	Offset            int64 `json:"offset"`

	// False when `bounds` was given (sort is not applied to bounded queries) or no sort was requested.
	SortApplied bool `json:"sort_applied"`

	// Present when withheld columns (traffic counts) are in the rows; they are null.
	WithholdGate SearchParcelsResponseWithholdGate `json:"withhold_gate"`

	// Present when the count did not finish.
	TotalStatus                  *SearchParcelsResponseTotalStatus `json:"total_status,omitempty"`
	ZoningCategoriesApplied      []string                          `json:"zoning_categories_applied,omitempty"`
	ZoningCategoriesUnrecognized []string                          `json:"zoning_categories_unrecognized,omitempty"`
	ZoningUnknownEstimate        *int64                            `json:"zoning_unknown_estimate,omitempty"`
	ZoningUnknownIsEstimate      *bool                             `json:"zoning_unknown_is_estimate,omitempty"`
	ZoningFilterNote             *string                           `json:"zoning_filter_note,omitempty"`
}

// UnmarshalJSON decodes SearchParcelsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsResponse) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsResponse
	aux := struct {
		*plain
		Total                 lenientNumber[int64] `json:"total"`
		Limit                 lenientNumber[int64] `json:"limit"`
		Offset                lenientNumber[int64] `json:"offset"`
		ZoningUnknownEstimate lenientNumber[int64] `json:"zoning_unknown_estimate"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assign(&r.Total)
	aux.Limit.assign(&r.Limit)
	aux.Offset.assign(&r.Offset)
	aux.ZoningUnknownEstimate.assignPtr(&r.ZoningUnknownEstimate)
	return softTypeError(err)
}

// SearchParcelsResponseData is generated from the OpenAPI spec.
type SearchParcelsResponseData struct {
	ParcelID           string   `json:"parcel_id"`
	StateFIPS          string   `json:"state_fips"`
	CountyFIPS         string   `json:"county_fips"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	Zoning             *string  `json:"zoning,omitempty"`
	ZoningCodeRaw      *string  `json:"zoning_code_raw,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
	LotSizeAcres       *float64 `json:"lot_size_acres,omitempty"`
	OwnershipType      *string  `json:"ownership_type,omitempty"`
	LandUseDesc        *string  `json:"land_use_desc,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	DirectVpd          *float64 `json:"direct_vpd,omitempty"`
	NearbyVpd          *float64 `json:"nearby_vpd,omitempty"`
	VpdVisibilityScore *float64 `json:"vpd_visibility_score,omitempty"`
	CrimeScore         *float64 `json:"crime_score,omitempty"`
	CrimeTier          *float64 `json:"crime_tier,omitempty"`
	CrimeTrend         *string  `json:"crime_trend,omitempty"`
	Guards             []string `json:"_guards"`
}

// UnmarshalJSON decodes SearchParcelsResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *SearchParcelsResponseData) UnmarshalJSON(data []byte) error {
	type plain SearchParcelsResponseData
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
		LotSizeAcres       lenientNumber[float64] `json:"lot_size_acres"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
		DirectVpd          lenientNumber[float64] `json:"direct_vpd"`
		NearbyVpd          lenientNumber[float64] `json:"nearby_vpd"`
		VpdVisibilityScore lenientNumber[float64] `json:"vpd_visibility_score"`
		CrimeScore         lenientNumber[float64] `json:"crime_score"`
		CrimeTier          lenientNumber[float64] `json:"crime_tier"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.DirectVpd.assignPtr(&r.DirectVpd)
	aux.NearbyVpd.assignPtr(&r.NearbyVpd)
	aux.VpdVisibilityScore.assignPtr(&r.VpdVisibilityScore)
	aux.CrimeScore.assignPtr(&r.CrimeScore)
	aux.CrimeTier.assignPtr(&r.CrimeTier)
	return softTypeError(err)
}

// SearchParcelsResponseWithholdGate: Present when withheld columns (traffic counts) are in the
// rows; they are null.
type SearchParcelsResponseWithholdGate struct {
	Applied    bool                                     `json:"applied"`
	Suppressed []*string                                `json:"suppressed"`
	Reason     string                                   `json:"reason"`
	Reasons    SearchParcelsResponseWithholdGateReasons `json:"reasons"`
	WithheldOn string                                   `json:"withheld_on"`
	Note       string                                   `json:"note"`
}

// SearchParcelsResponseWithholdGateReasons is generated from the OpenAPI spec.
type SearchParcelsResponseWithholdGateReasons struct {
	DirectVpd          string `json:"direct_vpd"`
	NearbyVpd          string `json:"nearby_vpd"`
	VpdVisibilityScore string `json:"vpd_visibility_score"`
}

// SearchParcelsResponseTotalStatus: Present when the count did not finish.
//
// It is a string; the SearchParcelsResponseTotalStatus* constants list the documented values.
type SearchParcelsResponseTotalStatus = string

// Documented values of SearchParcelsResponseTotalStatus.
const (
	SearchParcelsResponseTotalStatusTimedOut SearchParcelsResponseTotalStatus = "timed_out"
)

// Autocomplete: Address / place / parcel autocomplete
//
// Fast prefix-matched autocomplete: returns cities/places, matching parcels, and Mapbox-geocoded
// addresses for the prefix. Three independent tiers run in parallel with per-tier timeouts so a
// slow DB query never blocks fast Mapbox results. Use for type-ahead UIs and address entry; for
// full-detail lookup use parcel lookup.
//
// HTTP: GET /api/v1/search/autocomplete
func (s *SearchService) Autocomplete(ctx context.Context, params *SearchAutocompleteParams, opts ...RequestOption) (*SearchAutocompleteResponse, error) {
	var out SearchAutocompleteResponse
	if err := s.client.do(ctx, buildSearchAutocompleteRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildSearchAutocompleteRequest(params *SearchAutocompleteParams) *apiRequest {
	req := newRequest("GET", "/api/v1/search/autocomplete")
	if params != nil {
		addQuery(req.query, "q", params.Q)
	}
	return req
}

// SearchAutocompleteParams holds the query, header and JSON-body parameters of
// [SearchService.Autocomplete]. Pass nil when you need none.
type SearchAutocompleteParams struct {
	// Search prefix. Minimum 2 chars.
	//
	// Required.
	Q *string `query:"q" json:"-"`
}

// SearchAutocompleteResponse: Address / place / parcel autocomplete
type SearchAutocompleteResponse = AutocompleteResult

// Export: Export search results as CSV
//
// Returns up to 10,000 search-matched parcels as a CSV download. Accepts the same geographic +
// attribute filters as POST /api/v1/search. Designed for spreadsheet / Excel workflows; for
// programmatic ingestion, use the JSON search endpoint and paginate.
//
// HTTP: GET /api/v1/search/export
func (s *SearchService) Export(ctx context.Context, params *SearchExportParams, opts ...RequestOption) (string, error) {
	var out string
	if err := s.client.do(ctx, buildSearchExportRequest(params), opts, decodeText(&out)); err != nil {
		return "", err
	}
	return out, nil
}

func buildSearchExportRequest(params *SearchExportParams) *apiRequest {
	req := newRequest("GET", "/api/v1/search/export")
	req.accept = "text/csv, application/json"
	if params != nil {
		addQuery(req.query, "north", params.North)
		addQuery(req.query, "south", params.South)
		addQuery(req.query, "east", params.East)
		addQuery(req.query, "west", params.West)
		addQuery(req.query, "zoningCategories", params.ZoningCategories)
		addQuery(req.query, "sort", params.Sort)
		addQuery(req.query, "order", params.Order)
		addQuery(req.query, "limit", params.Limit)
	}
	return req
}

// SearchExportParams holds the query, header and JSON-body parameters of [SearchService.Export].
// Pass nil when you need none.
type SearchExportParams struct {
	// Bounding box: north latitude.
	North *float64 `query:"north" json:"-"`

	// Bounding box: south latitude.
	South *float64 `query:"south" json:"-"`

	// Bounding box: east longitude.
	East *float64 `query:"east" json:"-"`

	// Bounding box: west longitude.
	West *float64 `query:"west" json:"-"`

	// Comma-delimited zoning categories.
	ZoningCategories *string `query:"zoningCategories" json:"-"`

	// Sort column.
	Sort *SearchExportParamsSort `query:"sort" json:"-"`

	// Sort direction.
	Order *SearchExportParamsOrder `query:"order" json:"-"`

	// Row cap. Hard max 10,000.
	Limit *int64 `query:"limit" json:"-"`
}

// SearchExportParamsSort is generated from the OpenAPI spec. It is a string; the
// SearchExportParamsSort* constants list the documented values.
type SearchExportParamsSort = string

// Documented values of SearchExportParamsSort.
const (
	SearchExportParamsSortAddress            SearchExportParamsSort = "address"
	SearchExportParamsSortCity               SearchExportParamsSort = "city"
	SearchExportParamsSortState              SearchExportParamsSort = "state"
	SearchExportParamsSortOwnerName          SearchExportParamsSort = "owner_name"
	SearchExportParamsSortTotalAssessedValue SearchExportParamsSort = "total_assessed_value"
	SearchExportParamsSortZoning             SearchExportParamsSort = "zoning"
	SearchExportParamsSortYearBuilt          SearchExportParamsSort = "year_built"
	SearchExportParamsSortLotSizeAcres       SearchExportParamsSort = "lot_size_acres"
	SearchExportParamsSortOwnershipType      SearchExportParamsSort = "ownership_type"
)

// SearchExportParamsOrder is generated from the OpenAPI spec. It is a string; the
// SearchExportParamsOrder* constants list the documented values.
type SearchExportParamsOrder = string

// Documented values of SearchExportParamsOrder.
const (
	SearchExportParamsOrderAsc  SearchExportParamsOrder = "asc"
	SearchExportParamsOrderDesc SearchExportParamsOrder = "desc"
)

// Full: Full paginated text + attribute search
//
// Paginated search across the parcel search index by free-text query (`q`) and optional
// field/state/city filters. Distinct from POST /api/v1/search which is geo-bounded; this endpoint
// is text-anchored and works without a bounding box. Returns 50/page by default, 200 max.
//
// HTTP: GET /api/v1/search/full
func (s *SearchService) Full(ctx context.Context, params *SearchFullParams, opts ...RequestOption) (*SearchFullResponse, error) {
	var out SearchFullResponse
	if err := s.client.do(ctx, buildSearchFullRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// FullIter iterates every item of [SearchService.Full] across pages (cursor pagination over
// "results"): it passes the previous page's "nextCursor" as "after" until the cursor is
// null/absent or hasMore is false. Use [IterOptions] for the page size and an item cap.
func (s *SearchService) FullIter(ctx context.Context, params *SearchFullParams, iter IterOptions, opts ...RequestOption) *Iter[FullSearchResultResults] {
	var p SearchFullParams
	if params != nil {
		p = *params
	}
	return newCursorIter[FullSearchResultResults](ctx, s.client, iter, "results", "nextCursor", p.Limit, func(limit *int64, cursor *string) *apiRequest {
		q := p
		if limit != nil {
			q.Limit = limit
		}
		if cursor != nil {
			q.After = cursor
		}
		return buildSearchFullRequest(&q)
	}, opts)
}

func buildSearchFullRequest(params *SearchFullParams) *apiRequest {
	req := newRequest("GET", "/api/v1/search/full")
	if params != nil {
		addQuery(req.query, "q", params.Q)
		addQuery(req.query, "field", params.Field)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "city", params.City)
		addQuery(req.query, "page", params.Page)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "sort", params.Sort)
		addQuery(req.query, "dir", params.Dir)
		addQuery(req.query, "after", params.After)
	}
	return req
}

// SearchFullParams holds the query, header and JSON-body parameters of [SearchService.Full]. Pass
// nil when you need none.
type SearchFullParams struct {
	// Search query. Min 2 chars.
	//
	// Required.
	Q *string `query:"q" json:"-"`

	// Field to match against.
	Field *SearchFullParamsField `query:"field" json:"-"`

	// 2-letter state filter.
	State *string `query:"state" json:"-"`

	// City filter.
	City *string `query:"city" json:"-"`

	// 1-indexed page number.
	Page *int64 `query:"page" json:"-"`

	// Page size, max 200.
	Limit *int64 `query:"limit" json:"-"`

	// Sort column.
	Sort *SearchFullParamsSort `query:"sort" json:"-"`

	// Sort direction.
	Dir *SearchFullParamsDir `query:"dir" json:"-"`

	// Opaque keyset cursor: pass the previous page's `nextCursor`. Preferred over `page`.
	After *string `query:"after" json:"-"`
}

// SearchFullParamsField is generated from the OpenAPI spec. It is a string; the
// SearchFullParamsField* constants list the documented values.
type SearchFullParamsField = string

// Documented values of SearchFullParamsField.
const (
	SearchFullParamsFieldAll       SearchFullParamsField = "all"
	SearchFullParamsFieldAddress   SearchFullParamsField = "address"
	SearchFullParamsFieldOwnerName SearchFullParamsField = "owner_name"
	SearchFullParamsFieldCity      SearchFullParamsField = "city"
)

// SearchFullParamsSort is generated from the OpenAPI spec. It is a string; the
// SearchFullParamsSort* constants list the documented values.
type SearchFullParamsSort = string

// Documented values of SearchFullParamsSort.
const (
	SearchFullParamsSortAddress    SearchFullParamsSort = "address"
	SearchFullParamsSortCity       SearchFullParamsSort = "city"
	SearchFullParamsSortState      SearchFullParamsSort = "state"
	SearchFullParamsSortOwnerName  SearchFullParamsSort = "owner_name"
	SearchFullParamsSortTotalValue SearchFullParamsSort = "total_value"
	SearchFullParamsSortYearBuilt  SearchFullParamsSort = "year_built"
)

// SearchFullParamsDir is generated from the OpenAPI spec. It is a string; the SearchFullParamsDir*
// constants list the documented values.
type SearchFullParamsDir = string

// Documented values of SearchFullParamsDir.
const (
	SearchFullParamsDirAsc  SearchFullParamsDir = "asc"
	SearchFullParamsDirDesc SearchFullParamsDir = "desc"
)

// SearchFullResponse: Full paginated text + attribute search
type SearchFullResponse = FullSearchResult
