// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"io"
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

// Geographic and filtered parcel search.
//
// V1SearchService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SearchService] method instead.
type V1SearchService struct {
	options []option.RequestOption
}

// NewV1SearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1SearchService(opts ...option.RequestOption) (r V1SearchService) {
	r = V1SearchService{}
	r.options = opts
	return
}

// Fast prefix-matched autocomplete: returns cities/places, matching parcels, and
// Mapbox-geocoded addresses for the prefix. Three independent tiers run in
// parallel with per-tier timeouts so a slow DB query never blocks fast Mapbox
// results. Use for type-ahead UIs and address entry; for full-detail lookup use
// parcel lookup.
func (r *V1SearchService) Autocomplete(ctx context.Context, query V1SearchAutocompleteParams, opts ...option.RequestOption) (res *V1SearchAutocompleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/search/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns up to 10,000 search-matched parcels as a CSV download. Accepts the same
// geographic + attribute filters as POST /api/v1/search. Designed for spreadsheet
// / Excel workflows; for programmatic ingestion, use the JSON search endpoint and
// paginate.
func (r *V1SearchService) ExportResults(ctx context.Context, query V1SearchExportResultsParams, opts ...option.RequestOption) (res *io.Reader, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := "api/v1/search/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated search across the parcel search index by free-text query (`q`) and
// optional field/state/city filters. Distinct from POST /api/v1/search which is
// geo-bounded; this endpoint is text-anchored and works without a bounding box.
// Returns 50/page by default, 200 max.
func (r *V1SearchService) FullSearch(ctx context.Context, query V1SearchFullSearchParams, opts ...option.RequestOption) (res *V1SearchFullSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/search/full"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search parcels within geographic bounds with optional filters, sorting, and
// pagination.
func (r *V1SearchService) ParcelSearch(ctx context.Context, body V1SearchParcelSearchParams, opts ...option.RequestOption) (res *V1SearchParcelSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V1SearchAutocompleteResponse struct {
	Addresses []V1SearchAutocompleteResponseAddress  `json:"addresses"`
	Locations []V1SearchAutocompleteResponseLocation `json:"locations"`
	Parcels   []V1SearchAutocompleteResponseParcel   `json:"parcels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses   respjson.Field
		Locations   respjson.Field
		Parcels     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SearchAutocompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SearchAutocompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapbox-geocoded address suggestion. Use to disambiguate user input before
// calling /api/v1/lookup or /api/v1/parcels/{id}.
type V1SearchAutocompleteResponseAddress struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Name string  `json:"name"`
	// Any of "address".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SearchAutocompleteResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *V1SearchAutocompleteResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchAutocompleteResponseLocation struct {
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Name        string  `json:"name"`
	ParcelCount int64   `json:"parcel_count"`
	State       string  `json:"state"`
	// Any of "city".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Lat         respjson.Field
		Lng         respjson.Field
		Name        respjson.Field
		ParcelCount respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SearchAutocompleteResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *V1SearchAutocompleteResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchAutocompleteResponseParcel struct {
	Address    string `json:"address" api:"nullable"`
	City       string `json:"city" api:"nullable"`
	CountyFips string `json:"county_fips"`
	ParcelID   string `json:"parcel_id"`
	State      string `json:"state" api:"nullable"`
	// Any of "parcel".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		City        respjson.Field
		CountyFips  respjson.Field
		ParcelID    respjson.Field
		State       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SearchAutocompleteResponseParcel) RawJSON() string { return r.JSON.raw }
func (r *V1SearchAutocompleteResponseParcel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchFullSearchResponse struct {
	Page    int64    `json:"page"`
	Pages   int64    `json:"pages"`
	Results []Parcel `json:"results"`
	Total   int64    `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Pages       respjson.Field
		Results     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SearchFullSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SearchFullSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchParcelSearchResponse struct {
	Data   []Parcel `json:"data"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
	Total  int64    `json:"total"`
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
func (r V1SearchParcelSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SearchParcelSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchAutocompleteParams struct {
	// Search prefix. Minimum 2 chars.
	Q string `query:"q" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1SearchAutocompleteParams]'s query parameters as
// `url.Values`.
func (r V1SearchAutocompleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SearchExportResultsParams struct {
	// Bounding box: east longitude.
	East param.Opt[float64] `query:"east,omitzero" json:"-"`
	// Row cap. Hard max 10,000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Bounding box: north latitude.
	North param.Opt[float64] `query:"north,omitzero" json:"-"`
	// Bounding box: south latitude.
	South param.Opt[float64] `query:"south,omitzero" json:"-"`
	// Bounding box: west longitude.
	West param.Opt[float64] `query:"west,omitzero" json:"-"`
	// Comma-delimited zoning categories.
	ZoningCategories param.Opt[string] `query:"zoningCategories,omitzero" json:"-"`
	// Sort direction.
	//
	// Any of "asc", "desc".
	Order V1SearchExportResultsParamsOrder `query:"order,omitzero" json:"-"`
	// Sort column.
	//
	// Any of "address", "city", "state", "owner_name", "total_assessed_value",
	// "zoning", "year_built", "lot_size_acres", "ownership_type".
	Sort V1SearchExportResultsParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SearchExportResultsParams]'s query parameters as
// `url.Values`.
func (r V1SearchExportResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction.
type V1SearchExportResultsParamsOrder string

const (
	V1SearchExportResultsParamsOrderAsc  V1SearchExportResultsParamsOrder = "asc"
	V1SearchExportResultsParamsOrderDesc V1SearchExportResultsParamsOrder = "desc"
)

// Sort column.
type V1SearchExportResultsParamsSort string

const (
	V1SearchExportResultsParamsSortAddress            V1SearchExportResultsParamsSort = "address"
	V1SearchExportResultsParamsSortCity               V1SearchExportResultsParamsSort = "city"
	V1SearchExportResultsParamsSortState              V1SearchExportResultsParamsSort = "state"
	V1SearchExportResultsParamsSortOwnerName          V1SearchExportResultsParamsSort = "owner_name"
	V1SearchExportResultsParamsSortTotalAssessedValue V1SearchExportResultsParamsSort = "total_assessed_value"
	V1SearchExportResultsParamsSortZoning             V1SearchExportResultsParamsSort = "zoning"
	V1SearchExportResultsParamsSortYearBuilt          V1SearchExportResultsParamsSort = "year_built"
	V1SearchExportResultsParamsSortLotSizeAcres       V1SearchExportResultsParamsSort = "lot_size_acres"
	V1SearchExportResultsParamsSortOwnershipType      V1SearchExportResultsParamsSort = "ownership_type"
)

type V1SearchFullSearchParams struct {
	// Search query. Min 2 chars.
	Q string `query:"q" api:"required" json:"-"`
	// City filter.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Page size, max 200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 1-indexed page number.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// 2-letter state filter.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// Sort direction.
	//
	// Any of "asc", "desc".
	Dir V1SearchFullSearchParamsDir `query:"dir,omitzero" json:"-"`
	// Field to match against.
	//
	// Any of "all", "address", "owner_name", "city".
	Field V1SearchFullSearchParamsField `query:"field,omitzero" json:"-"`
	// Sort column.
	//
	// Any of "address", "city", "state", "owner_name", "total_value", "year_built".
	Sort V1SearchFullSearchParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SearchFullSearchParams]'s query parameters as
// `url.Values`.
func (r V1SearchFullSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction.
type V1SearchFullSearchParamsDir string

const (
	V1SearchFullSearchParamsDirAsc  V1SearchFullSearchParamsDir = "asc"
	V1SearchFullSearchParamsDirDesc V1SearchFullSearchParamsDir = "desc"
)

// Field to match against.
type V1SearchFullSearchParamsField string

const (
	V1SearchFullSearchParamsFieldAll       V1SearchFullSearchParamsField = "all"
	V1SearchFullSearchParamsFieldAddress   V1SearchFullSearchParamsField = "address"
	V1SearchFullSearchParamsFieldOwnerName V1SearchFullSearchParamsField = "owner_name"
	V1SearchFullSearchParamsFieldCity      V1SearchFullSearchParamsField = "city"
)

// Sort column.
type V1SearchFullSearchParamsSort string

const (
	V1SearchFullSearchParamsSortAddress    V1SearchFullSearchParamsSort = "address"
	V1SearchFullSearchParamsSortCity       V1SearchFullSearchParamsSort = "city"
	V1SearchFullSearchParamsSortState      V1SearchFullSearchParamsSort = "state"
	V1SearchFullSearchParamsSortOwnerName  V1SearchFullSearchParamsSort = "owner_name"
	V1SearchFullSearchParamsSortTotalValue V1SearchFullSearchParamsSort = "total_value"
	V1SearchFullSearchParamsSortYearBuilt  V1SearchFullSearchParamsSort = "year_built"
)

type V1SearchParcelSearchParams struct {
	Bounds V1SearchParcelSearchParamsBounds `json:"bounds,omitzero" api:"required"`
	// Number of results to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Number of results to skip for pagination.
	Offset  param.Opt[int64]                  `json:"offset,omitzero"`
	Filters V1SearchParcelSearchParamsFilters `json:"filters,omitzero"`
	// Any of "asc", "desc".
	Order V1SearchParcelSearchParamsOrder `json:"order,omitzero"`
	// Sort field.
	//
	// Any of "assessed_value", "sale_price", "acreage", "year_built", "deal_score".
	Sort V1SearchParcelSearchParamsSort `json:"sort,omitzero"`
	paramObj
}

func (r V1SearchParcelSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties East, North, South, West are required.
type V1SearchParcelSearchParamsBounds struct {
	East  float64 `json:"east" api:"required"`
	North float64 `json:"north" api:"required"`
	South float64 `json:"south" api:"required"`
	West  float64 `json:"west" api:"required"`
	paramObj
}

func (r V1SearchParcelSearchParamsBounds) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParamsBounds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParamsBounds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchParcelSearchParamsFilters struct {
	// Only return parcels with absentee owners.
	AbsenteeOnly param.Opt[bool] `json:"absenteeOnly,omitzero"`
	// Filter by acreage range.
	AcreageRange V1SearchParcelSearchParamsFiltersAcreageRange `json:"acreageRange,omitzero"`
	// Filter by owner entity type.
	//
	// Any of "individual", "corporation", "llc", "trust", "government", "other".
	OwnerTypes []string `json:"ownerTypes,omitzero"`
	// Filter by assessed value range.
	ValueRange V1SearchParcelSearchParamsFiltersValueRange `json:"valueRange,omitzero"`
	// Filter by year built range.
	YearBuiltRange V1SearchParcelSearchParamsFiltersYearBuiltRange `json:"yearBuiltRange,omitzero"`
	// Filter by zoning category (e.g., residential, commercial, industrial).
	ZoningCategories []string `json:"zoningCategories,omitzero"`
	paramObj
}

func (r V1SearchParcelSearchParamsFilters) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParamsFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParamsFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by acreage range.
type V1SearchParcelSearchParamsFiltersAcreageRange struct {
	Max param.Opt[float64] `json:"max,omitzero"`
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r V1SearchParcelSearchParamsFiltersAcreageRange) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParamsFiltersAcreageRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParamsFiltersAcreageRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by assessed value range.
type V1SearchParcelSearchParamsFiltersValueRange struct {
	Max param.Opt[float64] `json:"max,omitzero"`
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r V1SearchParcelSearchParamsFiltersValueRange) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParamsFiltersValueRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParamsFiltersValueRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by year built range.
type V1SearchParcelSearchParamsFiltersYearBuiltRange struct {
	Max param.Opt[int64] `json:"max,omitzero"`
	Min param.Opt[int64] `json:"min,omitzero"`
	paramObj
}

func (r V1SearchParcelSearchParamsFiltersYearBuiltRange) MarshalJSON() (data []byte, err error) {
	type shadow V1SearchParcelSearchParamsFiltersYearBuiltRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1SearchParcelSearchParamsFiltersYearBuiltRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SearchParcelSearchParamsOrder string

const (
	V1SearchParcelSearchParamsOrderAsc  V1SearchParcelSearchParamsOrder = "asc"
	V1SearchParcelSearchParamsOrderDesc V1SearchParcelSearchParamsOrder = "desc"
)

// Sort field.
type V1SearchParcelSearchParamsSort string

const (
	V1SearchParcelSearchParamsSortAssessedValue V1SearchParcelSearchParamsSort = "assessed_value"
	V1SearchParcelSearchParamsSortSalePrice     V1SearchParcelSearchParamsSort = "sale_price"
	V1SearchParcelSearchParamsSortAcreage       V1SearchParcelSearchParamsSort = "acreage"
	V1SearchParcelSearchParamsSortYearBuilt     V1SearchParcelSearchParamsSort = "year_built"
	V1SearchParcelSearchParamsSortDealScore     V1SearchParcelSearchParamsSort = "deal_score"
)
