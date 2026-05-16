// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/jdw2111/propraven-go/internal/apijson"
	"github.com/jdw2111/propraven-go/internal/apiquery"
	"github.com/jdw2111/propraven-go/internal/requestconfig"
	"github.com/jdw2111/propraven-go/option"
	"github.com/jdw2111/propraven-go/packages/param"
	"github.com/jdw2111/propraven-go/packages/respjson"
)

// Deal sourcing: absentee owners, property flips.
//
// V1DealService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1DealService] method instead.
type V1DealService struct {
	options []option.RequestOption
}

// NewV1DealService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1DealService(opts ...option.RequestOption) (r V1DealService) {
	r = V1DealService{}
	r.options = opts
	return
}

// Retrieve parcels owned by absentee owners, useful for off-market deal sourcing.
func (r *V1DealService) FindAbsenteeOwners(ctx context.Context, query V1DealFindAbsenteeOwnersParams, opts ...option.RequestOption) (res *V1DealFindAbsenteeOwnersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/absentee"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns parcels owned by legal entities identified from owner-name pattern
// matching across 221M+ parcels. Pass `top=true` to get aggregated entity rankings
// instead of per-parcel rows. One of `county_fips`, `state_fips`, `search`, or
// `top` is required.
func (r *V1DealService) FindEntityOwnedParcels(ctx context.Context, query V1DealFindEntityOwnedParcelsParams, opts ...option.RequestOption) (res *V1DealFindEntityOwnedParcelsResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve recently flipped properties. Use ?view=flippers to get a ranked list of
// top flippers instead.
func (r *V1DealService) FindFlips(ctx context.Context, query V1DealFindFlipsParams, opts ...option.RequestOption) (res *V1DealFindFlipsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/flips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns parcels where land value significantly exceeds improvement value — a
// signal for redevelopment, teardown, or assemblage opportunities. `county_fips`
// or `state_fips` is required.
func (r *V1DealService) FindHighLandRatio(ctx context.Context, query V1DealFindHighLandRatioParams, opts ...option.RequestOption) (res *V1DealFindHighLandRatioResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/high-land-ratio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns parcels not sold in `min_years` or more. Long-hold owners are often
// motivated sellers — estate planning, deferred maintenance, life changes.
// `county_fips` or `state_fips` is required.
func (r *V1DealService) FindLongHoldParcels(ctx context.Context, query V1DealFindLongHoldParcelsParams, opts ...option.RequestOption) (res *V1DealFindLongHoldParcelsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/long-hold"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns portfolio owners ranked by property count and total assessed value.
// Useful for finding institutional buyers, small landlords, or specific investor
// families.
func (r *V1DealService) FindPortfolioOwners(ctx context.Context, query V1DealFindPortfolioOwnersParams, opts ...option.RequestOption) (res *V1DealFindPortfolioOwnersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/portfolio-owners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Default: returns county/quarter transaction summaries. Pass `view=affordability`
// to retrieve the home-affordability index instead (price-to-income ratios +
// rating).
func (r *V1DealService) GetMarketSummary(ctx context.Context, query V1DealGetMarketSummaryParams, opts ...option.RequestOption) (res *V1DealGetMarketSummaryResponseUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/market"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns contractor profiles aggregated from 45M+ building permits. Each profile
// includes permit count, jurisdictions worked, total declared permit value, and
// activity dates. Use to identify active contractors in a market or find a
// specific contractor by name.
func (r *V1DealService) SearchContractors(ctx context.Context, query V1DealSearchContractorsParams, opts ...option.RequestOption) (res *V1DealSearchContractorsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/contractors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns lender profiles aggregated from deed/mortgage transactions. Includes
// mortgage count, total volume, geographic spread, and a national rank.
func (r *V1DealService) SearchLenders(ctx context.Context, query V1DealSearchLendersParams, opts ...option.RequestOption) (res *V1DealSearchLendersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/deals/lenders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AffordabilityRow struct {
	// Any of "AFFORDABLE", "MODERATE", "EXPENSIVE", "VERY_EXPENSIVE".
	AffordabilityRating    AffordabilityRowAffordabilityRating `json:"affordability_rating"`
	CountyFips             string                              `json:"county_fips"`
	CountyName             string                              `json:"county_name" api:"nullable"`
	MedianHouseholdIncome  float64                             `json:"median_household_income" api:"nullable"`
	MedianSalePrice        float64                             `json:"median_sale_price" api:"nullable"`
	MonthlyPaymentEstimate float64                             `json:"monthly_payment_estimate" api:"nullable"`
	PctIncomeForHousing    float64                             `json:"pct_income_for_housing" api:"nullable"`
	PriceToIncomeRatio     float64                             `json:"price_to_income_ratio" api:"nullable"`
	StateFips              string                              `json:"state_fips"`
	Year                   int64                               `json:"year"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffordabilityRating    respjson.Field
		CountyFips             respjson.Field
		CountyName             respjson.Field
		MedianHouseholdIncome  respjson.Field
		MedianSalePrice        respjson.Field
		MonthlyPaymentEstimate respjson.Field
		PctIncomeForHousing    respjson.Field
		PriceToIncomeRatio     respjson.Field
		StateFips              respjson.Field
		Year                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AffordabilityRow) RawJSON() string { return r.JSON.raw }
func (r *AffordabilityRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AffordabilityRowAffordabilityRating string

const (
	AffordabilityRowAffordabilityRatingAffordable    AffordabilityRowAffordabilityRating = "AFFORDABLE"
	AffordabilityRowAffordabilityRatingModerate      AffordabilityRowAffordabilityRating = "MODERATE"
	AffordabilityRowAffordabilityRatingExpensive     AffordabilityRowAffordabilityRating = "EXPENSIVE"
	AffordabilityRowAffordabilityRatingVeryExpensive AffordabilityRowAffordabilityRating = "VERY_EXPENSIVE"
)

type V1DealFindAbsenteeOwnersResponse struct {
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
func (r V1DealFindAbsenteeOwnersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindAbsenteeOwnersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1DealFindEntityOwnedParcelsResponseUnion contains all possible properties and
// values from [V1DealFindEntityOwnedParcelsResponseObject],
// [V1DealFindEntityOwnedParcelsResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1DealFindEntityOwnedParcelsResponseUnion struct {
	// This field is a union of [[]V1DealFindEntityOwnedParcelsResponseObjectData],
	// [[]V1DealFindEntityOwnedParcelsResponseObject2Data]
	Data   V1DealFindEntityOwnedParcelsResponseUnionData `json:"data"`
	Limit  int64                                         `json:"limit"`
	Offset int64                                         `json:"offset"`
	Total  int64                                         `json:"total"`
	// This field is from variant [V1DealFindEntityOwnedParcelsResponseObject2].
	Summary V1DealFindEntityOwnedParcelsResponseObject2Summary `json:"summary"`
	JSON    struct {
		Data    respjson.Field
		Limit   respjson.Field
		Offset  respjson.Field
		Total   respjson.Field
		Summary respjson.Field
		raw     string
	} `json:"-"`
}

func (u V1DealFindEntityOwnedParcelsResponseUnion) AsV1DealFindEntityOwnedParcelsResponseObject() (v V1DealFindEntityOwnedParcelsResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1DealFindEntityOwnedParcelsResponseUnion) AsV1DealFindEntityOwnedParcelsResponseObject2() (v V1DealFindEntityOwnedParcelsResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1DealFindEntityOwnedParcelsResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1DealFindEntityOwnedParcelsResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1DealFindEntityOwnedParcelsResponseUnionData is an implicit subunion of
// [V1DealFindEntityOwnedParcelsResponseUnion].
// V1DealFindEntityOwnedParcelsResponseUnionData provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1DealFindEntityOwnedParcelsResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfV1DealFindEntityOwnedParcelsResponseObjectData
// OfV1DealFindEntityOwnedParcelsResponseObject2Data]
type V1DealFindEntityOwnedParcelsResponseUnionData struct {
	// This field will be present if the value is a
	// [[]V1DealFindEntityOwnedParcelsResponseObjectData] instead of an object.
	OfV1DealFindEntityOwnedParcelsResponseObjectData []V1DealFindEntityOwnedParcelsResponseObjectData `json:",inline"`
	// This field will be present if the value is a
	// [[]V1DealFindEntityOwnedParcelsResponseObject2Data] instead of an object.
	OfV1DealFindEntityOwnedParcelsResponseObject2Data []V1DealFindEntityOwnedParcelsResponseObject2Data `json:",inline"`
	JSON                                              struct {
		OfV1DealFindEntityOwnedParcelsResponseObjectData  respjson.Field
		OfV1DealFindEntityOwnedParcelsResponseObject2Data respjson.Field
		raw                                               string
	} `json:"-"`
}

func (r *V1DealFindEntityOwnedParcelsResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindEntityOwnedParcelsResponseObject struct {
	Data   []V1DealFindEntityOwnedParcelsResponseObjectData `json:"data"`
	Limit  int64                                            `json:"limit"`
	Offset int64                                            `json:"offset"`
	Total  int64                                            `json:"total"`
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
func (r V1DealFindEntityOwnedParcelsResponseObject) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindEntityOwnedParcelsResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindEntityOwnedParcelsResponseObjectData struct {
	Address    string `json:"address" api:"nullable"`
	City       string `json:"city" api:"nullable"`
	CountyFips string `json:"county_fips"`
	// Any of "LLC", "CORP", "TRUST", "LP", "LTD", "ASSOCIATION", "OTHER_ENTITY".
	EntityType         string  `json:"entity_type"`
	OwnerName          string  `json:"owner_name"`
	ParcelID           string  `json:"parcel_id"`
	State              string  `json:"state" api:"nullable"`
	StateFips          string  `json:"state_fips"`
	TotalAssessedValue float64 `json:"total_assessed_value" api:"nullable"`
	Zip                string  `json:"zip" api:"nullable"`
	Zoning             string  `json:"zoning" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		City               respjson.Field
		CountyFips         respjson.Field
		EntityType         respjson.Field
		OwnerName          respjson.Field
		ParcelID           respjson.Field
		State              respjson.Field
		StateFips          respjson.Field
		TotalAssessedValue respjson.Field
		Zip                respjson.Field
		Zoning             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindEntityOwnedParcelsResponseObjectData) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindEntityOwnedParcelsResponseObjectData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindEntityOwnedParcelsResponseObject2 struct {
	Data    []V1DealFindEntityOwnedParcelsResponseObject2Data  `json:"data"`
	Limit   int64                                              `json:"limit"`
	Offset  int64                                              `json:"offset"`
	Summary V1DealFindEntityOwnedParcelsResponseObject2Summary `json:"summary"`
	Total   int64                                              `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Summary     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindEntityOwnedParcelsResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindEntityOwnedParcelsResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindEntityOwnedParcelsResponseObject2Data struct {
	EntityType  string   `json:"entity_type"`
	OwnerName   string   `json:"owner_name"`
	ParcelCount int64    `json:"parcel_count"`
	StatesArr   []string `json:"states_arr"`
	TotalValue  float64  `json:"total_value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityType  respjson.Field
		OwnerName   respjson.Field
		ParcelCount respjson.Field
		StatesArr   respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindEntityOwnedParcelsResponseObject2Data) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindEntityOwnedParcelsResponseObject2Data) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindEntityOwnedParcelsResponseObject2Summary struct {
	CorpCount     int64 `json:"corp_count"`
	LlcCount      int64 `json:"llc_count"`
	LpCount       int64 `json:"lp_count"`
	TotalEntities int64 `json:"total_entities"`
	TotalParcels  int64 `json:"total_parcels"`
	TrustCount    int64 `json:"trust_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorpCount     respjson.Field
		LlcCount      respjson.Field
		LpCount       respjson.Field
		TotalEntities respjson.Field
		TotalParcels  respjson.Field
		TrustCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindEntityOwnedParcelsResponseObject2Summary) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindEntityOwnedParcelsResponseObject2Summary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindFlipsResponse struct {
	Data   []V1DealFindFlipsResponseData `json:"data"`
	Limit  int64                         `json:"limit"`
	Offset int64                         `json:"offset"`
	Total  int64                         `json:"total"`
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
func (r V1DealFindFlipsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindFlipsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindFlipsResponseData struct {
	Address    string    `json:"address"`
	BuyDate    time.Time `json:"buy_date" format:"date"`
	BuyPrice   float64   `json:"buy_price"`
	BuyerName  string    `json:"buyer_name"`
	CountyFips string    `json:"county_fips"`
	// Any of "quick", "standard", "long".
	FlipTier   string    `json:"flip_tier"`
	HoldDays   int64     `json:"hold_days"`
	ParcelID   string    `json:"parcel_id"`
	Profit     float64   `json:"profit"`
	SellDate   time.Time `json:"sell_date" format:"date"`
	SellPrice  float64   `json:"sell_price"`
	SellerName string    `json:"seller_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		BuyDate     respjson.Field
		BuyPrice    respjson.Field
		BuyerName   respjson.Field
		CountyFips  respjson.Field
		FlipTier    respjson.Field
		HoldDays    respjson.Field
		ParcelID    respjson.Field
		Profit      respjson.Field
		SellDate    respjson.Field
		SellPrice   respjson.Field
		SellerName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindFlipsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindFlipsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindHighLandRatioResponse struct {
	Data   []V1DealFindHighLandRatioResponseData `json:"data"`
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
func (r V1DealFindHighLandRatioResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindHighLandRatioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindHighLandRatioResponseData struct {
	Address                  string  `json:"address" api:"nullable"`
	City                     string  `json:"city" api:"nullable"`
	CountyFips               string  `json:"county_fips"`
	ImprovementAssessedValue float64 `json:"improvement_assessed_value" api:"nullable"`
	LandAssessedValue        float64 `json:"land_assessed_value" api:"nullable"`
	// land / improvement, higher = more redevelopment potential.
	LandImprovementRatio float64 `json:"land_improvement_ratio"`
	OwnerName            string  `json:"owner_name" api:"nullable"`
	ParcelID             string  `json:"parcel_id"`
	State                string  `json:"state" api:"nullable"`
	StateFips            string  `json:"state_fips"`
	TotalAssessedValue   float64 `json:"total_assessed_value" api:"nullable"`
	Zoning               string  `json:"zoning" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                  respjson.Field
		City                     respjson.Field
		CountyFips               respjson.Field
		ImprovementAssessedValue respjson.Field
		LandAssessedValue        respjson.Field
		LandImprovementRatio     respjson.Field
		OwnerName                respjson.Field
		ParcelID                 respjson.Field
		State                    respjson.Field
		StateFips                respjson.Field
		TotalAssessedValue       respjson.Field
		Zoning                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindHighLandRatioResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindHighLandRatioResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindLongHoldParcelsResponse struct {
	Data   []V1DealFindLongHoldParcelsResponseData `json:"data"`
	Limit  int64                                   `json:"limit"`
	Offset int64                                   `json:"offset"`
	Total  int64                                   `json:"total"`
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
func (r V1DealFindLongHoldParcelsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindLongHoldParcelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindLongHoldParcelsResponseData struct {
	Address    string `json:"address" api:"nullable"`
	City       string `json:"city" api:"nullable"`
	CountyFips string `json:"county_fips"`
	// Any of "10-15yr", "15-20yr", "20-30yr", "30yr+".
	HoldTier           string    `json:"hold_tier"`
	LastSaleDate       time.Time `json:"last_sale_date" format:"date"`
	OwnerName          string    `json:"owner_name" api:"nullable"`
	ParcelID           string    `json:"parcel_id"`
	State              string    `json:"state" api:"nullable"`
	StateFips          string    `json:"state_fips"`
	TotalAssessedValue float64   `json:"total_assessed_value" api:"nullable"`
	YearsHeld          float64   `json:"years_held"`
	Zip                string    `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		City               respjson.Field
		CountyFips         respjson.Field
		HoldTier           respjson.Field
		LastSaleDate       respjson.Field
		OwnerName          respjson.Field
		ParcelID           respjson.Field
		State              respjson.Field
		StateFips          respjson.Field
		TotalAssessedValue respjson.Field
		YearsHeld          respjson.Field
		Zip                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindLongHoldParcelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindLongHoldParcelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindPortfolioOwnersResponse struct {
	Data   []V1DealFindPortfolioOwnersResponseData `json:"data"`
	Limit  int64                                   `json:"limit"`
	Offset int64                                   `json:"offset"`
	Total  int64                                   `json:"total"`
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
func (r V1DealFindPortfolioOwnersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindPortfolioOwnersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindPortfolioOwnersResponseData struct {
	AvgAssessedValue    float64 `json:"avg_assessed_value" api:"nullable"`
	CountyCount         int64   `json:"county_count"`
	OwnerNameNormalized string  `json:"owner_name_normalized"`
	OwnerState          string  `json:"owner_state" api:"nullable"`
	// National rank, 1 = largest portfolio.
	PortfolioRank      int64   `json:"portfolio_rank"`
	PropertyCount      int64   `json:"property_count"`
	StateCount         int64   `json:"state_count"`
	StatesList         string  `json:"states_list" api:"nullable"`
	TotalAcreage       float64 `json:"total_acreage" api:"nullable"`
	TotalAssessedValue float64 `json:"total_assessed_value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgAssessedValue    respjson.Field
		CountyCount         respjson.Field
		OwnerNameNormalized respjson.Field
		OwnerState          respjson.Field
		PortfolioRank       respjson.Field
		PropertyCount       respjson.Field
		StateCount          respjson.Field
		StatesList          respjson.Field
		TotalAcreage        respjson.Field
		TotalAssessedValue  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealFindPortfolioOwnersResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealFindPortfolioOwnersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1DealGetMarketSummaryResponseUnion contains all possible properties and values
// from [V1DealGetMarketSummaryResponseObject],
// [V1DealGetMarketSummaryResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V1DealGetMarketSummaryResponseUnion struct {
	// This field is a union of [[]V1DealGetMarketSummaryResponseObjectData],
	// [[]AffordabilityRow]
	Data   V1DealGetMarketSummaryResponseUnionData `json:"data"`
	Limit  int64                                   `json:"limit"`
	Offset int64                                   `json:"offset"`
	Total  int64                                   `json:"total"`
	JSON   struct {
		Data   respjson.Field
		Limit  respjson.Field
		Offset respjson.Field
		Total  respjson.Field
		raw    string
	} `json:"-"`
}

func (u V1DealGetMarketSummaryResponseUnion) AsV1DealGetMarketSummaryResponseObject() (v V1DealGetMarketSummaryResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V1DealGetMarketSummaryResponseUnion) AsV1DealGetMarketSummaryResponseObject2() (v V1DealGetMarketSummaryResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V1DealGetMarketSummaryResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V1DealGetMarketSummaryResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V1DealGetMarketSummaryResponseUnionData is an implicit subunion of
// [V1DealGetMarketSummaryResponseUnion]. V1DealGetMarketSummaryResponseUnionData
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [V1DealGetMarketSummaryResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfV1DealGetMarketSummaryResponseObjectData
// OfAffordabilityRowArray]
type V1DealGetMarketSummaryResponseUnionData struct {
	// This field will be present if the value is a
	// [[]V1DealGetMarketSummaryResponseObjectData] instead of an object.
	OfV1DealGetMarketSummaryResponseObjectData []V1DealGetMarketSummaryResponseObjectData `json:",inline"`
	// This field will be present if the value is a [[]AffordabilityRow] instead of an
	// object.
	OfAffordabilityRowArray []AffordabilityRow `json:",inline"`
	JSON                    struct {
		OfV1DealGetMarketSummaryResponseObjectData respjson.Field
		OfAffordabilityRowArray                    respjson.Field
		raw                                        string
	} `json:"-"`
}

func (r *V1DealGetMarketSummaryResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealGetMarketSummaryResponseObject struct {
	Data   []V1DealGetMarketSummaryResponseObjectData `json:"data"`
	Limit  int64                                      `json:"limit"`
	Offset int64                                      `json:"offset"`
	Total  int64                                      `json:"total"`
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
func (r V1DealGetMarketSummaryResponseObject) RawJSON() string { return r.JSON.raw }
func (r *V1DealGetMarketSummaryResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealGetMarketSummaryResponseObjectData struct {
	AvgSalePrice    float64 `json:"avg_sale_price" api:"nullable"`
	CountyFips      string  `json:"county_fips"`
	CountyName      string  `json:"county_name" api:"nullable"`
	MedianSalePrice float64 `json:"median_sale_price" api:"nullable"`
	SaleCount       int64   `json:"sale_count"`
	StateFips       string  `json:"state_fips"`
	Year            int64   `json:"year"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgSalePrice    respjson.Field
		CountyFips      respjson.Field
		CountyName      respjson.Field
		MedianSalePrice respjson.Field
		SaleCount       respjson.Field
		StateFips       respjson.Field
		Year            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealGetMarketSummaryResponseObjectData) RawJSON() string { return r.JSON.raw }
func (r *V1DealGetMarketSummaryResponseObjectData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealGetMarketSummaryResponseObject2 struct {
	Data   []AffordabilityRow `json:"data"`
	Limit  int64              `json:"limit"`
	Offset int64              `json:"offset"`
	Total  int64              `json:"total"`
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
func (r V1DealGetMarketSummaryResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *V1DealGetMarketSummaryResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealSearchContractorsResponse struct {
	Data   []V1DealSearchContractorsResponseData `json:"data"`
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
func (r V1DealSearchContractorsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealSearchContractorsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealSearchContractorsResponseData struct {
	ActiveYears              float64 `json:"active_years" api:"nullable"`
	AvgPermitValue           float64 `json:"avg_permit_value"`
	ContractorLicense        string  `json:"contractor_license" api:"nullable"`
	ContractorNameNormalized string  `json:"contractor_name_normalized"`
	// National rank, 1 = highest activity.
	ContractorRank    int64     `json:"contractor_rank"`
	CountyCount       int64     `json:"county_count"`
	FirstPermitDate   time.Time `json:"first_permit_date" format:"date"`
	JurisdictionCount int64     `json:"jurisdiction_count"`
	LastPermitDate    time.Time `json:"last_permit_date" format:"date"`
	PermitCount       int64     `json:"permit_count"`
	StateCount        int64     `json:"state_count"`
	// Comma-delimited 2-letter state codes.
	StatesList       string   `json:"states_list" api:"nullable"`
	TopPermitTypes   []string `json:"top_permit_types"`
	TotalPermitValue float64  `json:"total_permit_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveYears              respjson.Field
		AvgPermitValue           respjson.Field
		ContractorLicense        respjson.Field
		ContractorNameNormalized respjson.Field
		ContractorRank           respjson.Field
		CountyCount              respjson.Field
		FirstPermitDate          respjson.Field
		JurisdictionCount        respjson.Field
		LastPermitDate           respjson.Field
		PermitCount              respjson.Field
		StateCount               respjson.Field
		StatesList               respjson.Field
		TopPermitTypes           respjson.Field
		TotalPermitValue         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealSearchContractorsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealSearchContractorsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealSearchLendersResponse struct {
	Data   []V1DealSearchLendersResponseData `json:"data"`
	Limit  int64                             `json:"limit"`
	Offset int64                             `json:"offset"`
	Total  int64                             `json:"total"`
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
func (r V1DealSearchLendersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DealSearchLendersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealSearchLendersResponseData struct {
	AvgMortgageAmount    float64   `json:"avg_mortgage_amount" api:"nullable"`
	CountyCount          int64     `json:"county_count"`
	FirstMortgageDate    time.Time `json:"first_mortgage_date" format:"date"`
	LastMortgageDate     time.Time `json:"last_mortgage_date" format:"date"`
	LenderNameNormalized string    `json:"lender_name_normalized"`
	// National rank, 1 = highest volume.
	LenderRank           int64   `json:"lender_rank"`
	MedianMortgageAmount float64 `json:"median_mortgage_amount" api:"nullable"`
	MortgageCount        int64   `json:"mortgage_count"`
	StateCount           int64   `json:"state_count"`
	StatesList           string  `json:"states_list" api:"nullable"`
	TotalMortgageVolume  float64 `json:"total_mortgage_volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgMortgageAmount    respjson.Field
		CountyCount          respjson.Field
		FirstMortgageDate    respjson.Field
		LastMortgageDate     respjson.Field
		LenderNameNormalized respjson.Field
		LenderRank           respjson.Field
		MedianMortgageAmount respjson.Field
		MortgageCount        respjson.Field
		StateCount           respjson.Field
		StatesList           respjson.Field
		TotalMortgageVolume  respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DealSearchLendersResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1DealSearchLendersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DealFindAbsenteeOwnersParams struct {
	// Filter by county FIPS code.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	Limit      param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Minimum assessed value.
	MinValue param.Opt[float64] `query:"min_value,omitzero" json:"-"`
	Offset   param.Opt[int64]   `query:"offset,omitzero" json:"-"`
	// Only return owners whose mailing address is in a different state.
	OutOfState param.Opt[bool] `query:"out_of_state,omitzero" json:"-"`
	// Filter by state FIPS code.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindAbsenteeOwnersParams]'s query parameters as
// `url.Values`.
func (r V1DealFindAbsenteeOwnersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1DealFindEntityOwnedParcelsParams struct {
	// 5-digit county FIPS filter.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum assessed value, USD.
	MinValue param.Opt[int64] `query:"min_value,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Owner-name substring search.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// 2-digit state FIPS filter.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// If true, returns aggregated entity rankings with summary stats instead of
	// per-parcel rows.
	Top param.Opt[bool] `query:"top,omitzero" json:"-"`
	// Zoning substring filter.
	Zoning param.Opt[string] `query:"zoning,omitzero" json:"-"`
	// Filter by entity classification.
	//
	// Any of "LLC", "CORP", "TRUST", "LP", "LTD", "ASSOCIATION", "OTHER_ENTITY".
	EntityType V1DealFindEntityOwnedParcelsParamsEntityType `query:"entity_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindEntityOwnedParcelsParams]'s query parameters as
// `url.Values`.
func (r V1DealFindEntityOwnedParcelsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by entity classification.
type V1DealFindEntityOwnedParcelsParamsEntityType string

const (
	V1DealFindEntityOwnedParcelsParamsEntityTypeLlc         V1DealFindEntityOwnedParcelsParamsEntityType = "LLC"
	V1DealFindEntityOwnedParcelsParamsEntityTypeCorp        V1DealFindEntityOwnedParcelsParamsEntityType = "CORP"
	V1DealFindEntityOwnedParcelsParamsEntityTypeTrust       V1DealFindEntityOwnedParcelsParamsEntityType = "TRUST"
	V1DealFindEntityOwnedParcelsParamsEntityTypeLp          V1DealFindEntityOwnedParcelsParamsEntityType = "LP"
	V1DealFindEntityOwnedParcelsParamsEntityTypeLtd         V1DealFindEntityOwnedParcelsParamsEntityType = "LTD"
	V1DealFindEntityOwnedParcelsParamsEntityTypeAssociation V1DealFindEntityOwnedParcelsParamsEntityType = "ASSOCIATION"
	V1DealFindEntityOwnedParcelsParamsEntityTypeOtherEntity V1DealFindEntityOwnedParcelsParamsEntityType = "OTHER_ENTITY"
)

type V1DealFindFlipsParams struct {
	// Filter by county FIPS code.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	Limit      param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Minimum estimated profit.
	MinProfit param.Opt[float64] `query:"min_profit,omitzero" json:"-"`
	Offset    param.Opt[int64]   `query:"offset,omitzero" json:"-"`
	// Filter by state FIPS code.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// Filter by flip speed tier (quick: <6mo, standard: 6-12mo, long: 12-24mo).
	//
	// Any of "quick", "standard", "long".
	FlipTier V1DealFindFlipsParamsFlipTier `query:"flip_tier,omitzero" json:"-"`
	// Set to 'flippers' to return a ranked list of top flippers instead of individual
	// flips.
	//
	// Any of "flippers".
	View V1DealFindFlipsParamsView `query:"view,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindFlipsParams]'s query parameters as `url.Values`.
func (r V1DealFindFlipsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by flip speed tier (quick: <6mo, standard: 6-12mo, long: 12-24mo).
type V1DealFindFlipsParamsFlipTier string

const (
	V1DealFindFlipsParamsFlipTierQuick    V1DealFindFlipsParamsFlipTier = "quick"
	V1DealFindFlipsParamsFlipTierStandard V1DealFindFlipsParamsFlipTier = "standard"
	V1DealFindFlipsParamsFlipTierLong     V1DealFindFlipsParamsFlipTier = "long"
)

// Set to 'flippers' to return a ranked list of top flippers instead of individual
// flips.
type V1DealFindFlipsParamsView string

const (
	V1DealFindFlipsParamsViewFlippers V1DealFindFlipsParamsView = "flippers"
)

type V1DealFindHighLandRatioParams struct {
	// 5-digit county FIPS filter.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum land/improvement ratio.
	MinRatio param.Opt[float64] `query:"min_ratio,omitzero" json:"-"`
	// Minimum land assessed value, USD.
	MinValue param.Opt[int64] `query:"min_value,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 2-digit state FIPS filter.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// Zoning substring filter.
	Zoning param.Opt[string] `query:"zoning,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindHighLandRatioParams]'s query parameters as
// `url.Values`.
func (r V1DealFindHighLandRatioParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1DealFindLongHoldParcelsParams struct {
	// 5-digit county FIPS filter.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum assessed value, USD.
	MinValue param.Opt[int64] `query:"min_value,omitzero" json:"-"`
	// Minimum years held.
	MinYears param.Opt[int64] `query:"min_years,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 2-digit state FIPS filter.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// Filter by hold-period tier.
	//
	// Any of "10-15yr", "15-20yr", "20-30yr", "30yr+".
	HoldTier V1DealFindLongHoldParcelsParamsHoldTier `query:"hold_tier,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindLongHoldParcelsParams]'s query parameters as
// `url.Values`.
func (r V1DealFindLongHoldParcelsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by hold-period tier.
type V1DealFindLongHoldParcelsParamsHoldTier string

const (
	V1DealFindLongHoldParcelsParamsHoldTier10_15yr V1DealFindLongHoldParcelsParamsHoldTier = "10-15yr"
	V1DealFindLongHoldParcelsParamsHoldTier15_20yr V1DealFindLongHoldParcelsParamsHoldTier = "15-20yr"
	V1DealFindLongHoldParcelsParamsHoldTier20_30yr V1DealFindLongHoldParcelsParamsHoldTier = "20-30yr"
	V1DealFindLongHoldParcelsParamsHoldTier30yr    V1DealFindLongHoldParcelsParamsHoldTier = "30yr+"
)

type V1DealFindPortfolioOwnersParams struct {
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum properties owned.
	MinProperties param.Opt[int64] `query:"min_properties,omitzero" json:"-"`
	// Minimum total portfolio assessed value, USD.
	MinValue param.Opt[int64] `query:"min_value,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Owner name substring search.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// 2-letter owner mailing state filter.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealFindPortfolioOwnersParams]'s query parameters as
// `url.Values`.
func (r V1DealFindPortfolioOwnersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1DealGetMarketSummaryParams struct {
	// 5-digit county FIPS filter.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// 2-digit state FIPS filter.
	StateFips param.Opt[string] `query:"state_fips,omitzero" json:"-"`
	// Year filter.
	Year param.Opt[string] `query:"year,omitzero" json:"-"`
	// Affordability-rating filter (only meaningful with view=affordability).
	//
	// Any of "AFFORDABLE", "MODERATE", "EXPENSIVE", "VERY_EXPENSIVE".
	Rating V1DealGetMarketSummaryParamsRating `query:"rating,omitzero" json:"-"`
	// Switch to the affordability-index dataset.
	//
	// Any of "affordability".
	View V1DealGetMarketSummaryParamsView `query:"view,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealGetMarketSummaryParams]'s query parameters as
// `url.Values`.
func (r V1DealGetMarketSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Affordability-rating filter (only meaningful with view=affordability).
type V1DealGetMarketSummaryParamsRating string

const (
	V1DealGetMarketSummaryParamsRatingAffordable    V1DealGetMarketSummaryParamsRating = "AFFORDABLE"
	V1DealGetMarketSummaryParamsRatingModerate      V1DealGetMarketSummaryParamsRating = "MODERATE"
	V1DealGetMarketSummaryParamsRatingExpensive     V1DealGetMarketSummaryParamsRating = "EXPENSIVE"
	V1DealGetMarketSummaryParamsRatingVeryExpensive V1DealGetMarketSummaryParamsRating = "VERY_EXPENSIVE"
)

// Switch to the affordability-index dataset.
type V1DealGetMarketSummaryParamsView string

const (
	V1DealGetMarketSummaryParamsViewAffordability V1DealGetMarketSummaryParamsView = "affordability"
)

type V1DealSearchContractorsParams struct {
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum permit count to include.
	MinPermits param.Opt[int64] `query:"min_permits,omitzero" json:"-"`
	// Minimum total declared permit value, USD.
	MinValue param.Opt[int64] `query:"min_value,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Contractor name search (case-insensitive substring).
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// 2-letter state filter.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealSearchContractorsParams]'s query parameters as
// `url.Values`.
func (r V1DealSearchContractorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1DealSearchLendersParams struct {
	// Page size, max 500.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum mortgage count.
	MinMortgages param.Opt[int64] `query:"min_mortgages,omitzero" json:"-"`
	// Pagination offset.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Lender name substring search.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// 2-letter state filter.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DealSearchLendersParams]'s query parameters as
// `url.Values`.
func (r V1DealSearchLendersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
