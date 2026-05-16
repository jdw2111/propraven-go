// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"errors"
	"fmt"
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

// Parcel lookup, owner details, permits, deeds, and risk data.
//
// V1ParcelService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ParcelService] method instead.
type V1ParcelService struct {
	options []option.RequestOption
}

// NewV1ParcelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ParcelService(opts ...option.RequestOption) (r V1ParcelService) {
	r = V1ParcelService{}
	r.options = opts
	return
}

// Retrieve a single parcel by its composite ID (county_fips:parcel_id).
func (r *V1ParcelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Parcel, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve deed transactions and transfer history for a parcel.
func (r *V1ParcelService) GetDeeds(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ParcelGetDeedsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/deeds", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns parcel polygons inside a bounding box as a GeoJSON FeatureCollection.
// Only served at zoom ≥ 14 to limit data volume — coarser bbox returns an empty
// collection. Each feature's properties include parcel_id, county_fips,
// owner_name, assessed value, and basic attributes for rendering popups.
func (r *V1ParcelService) GetGeojson(ctx context.Context, query V1ParcelGetGeojsonParams, opts ...option.RequestOption) (res *V1ParcelGetGeojsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/parcels/geojson"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the owner of a parcel along with their portfolio summary and list of
// properties.
func (r *V1ParcelService) GetOwner(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ParcelGetOwnerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/owner", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve building and construction permits associated with a parcel.
func (r *V1ParcelService) GetPermits(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ParcelGetPermitsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/permits", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the full canonical parcel record plus every enrichment: UCC liens,
// comparable sales, owner portfolio context, permits, deeds, hazard composite.
// This is the single most data-dense endpoint per parcel — designed for
// due-diligence and underwriting workflows. Heavier than parcels/{id}; cache
// aggressively when serving UIs.
func (r *V1ParcelService) GetReport(ctx context.Context, id string, query V1ParcelGetReportParams, opts ...option.RequestOption) (res *V1ParcelGetReportResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/report", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve flood, wildfire, air quality, and crime risk data for a parcel.
func (r *V1ParcelService) GetRisks(ctx context.Context, id string, opts ...option.RequestOption) (res *RiskAssessment, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/risks", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Finds the AADT (annual average daily traffic) station nearest to the given
// coordinates and returns its historical time series plus 3/5/7-year CAGRs. Useful
// for retail / CRE site selection. Search radius ~2 miles; returns empty data if
// no station is in range.
func (r *V1ParcelService) GetTrafficHistory(ctx context.Context, id string, query V1ParcelGetTrafficHistoryParams, opts ...option.RequestOption) (res *V1ParcelGetTrafficHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/parcels/%s/traffic-history", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Deed struct {
	DeedType       string    `json:"deed_type" api:"nullable"`
	DocumentNumber string    `json:"document_number"`
	GranteeName    string    `json:"grantee_name"`
	GrantorName    string    `json:"grantor_name"`
	RecordingDate  time.Time `json:"recording_date" format:"date"`
	SaleDate       time.Time `json:"sale_date" api:"nullable" format:"date"`
	SalePrice      float64   `json:"sale_price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeedType       respjson.Field
		DocumentNumber respjson.Field
		GranteeName    respjson.Field
		GrantorName    respjson.Field
		RecordingDate  respjson.Field
		SaleDate       respjson.Field
		SalePrice      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Deed) RawJSON() string { return r.JSON.raw }
func (r *Deed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Parcel struct {
	AbsenteeOwner bool    `json:"absentee_owner"`
	Acreage       float64 `json:"acreage" api:"nullable"`
	Address       string  `json:"address"`
	AssessedValue float64 `json:"assessed_value" api:"nullable"`
	City          string  `json:"city"`
	// 5-digit county FIPS code.
	CountyFips string `json:"county_fips"`
	CountyName string `json:"county_name"`
	// Crime score from 0 (low) to 100 (high).
	CrimeScore float64 `json:"crime_score" api:"nullable"`
	// Composite deal opportunity score from 0 to 100.
	DealScore        float64   `json:"deal_score" api:"nullable"`
	ImprovementValue float64   `json:"improvement_value" api:"nullable"`
	LandUse          string    `json:"land_use" api:"nullable"`
	LandValue        float64   `json:"land_value" api:"nullable"`
	LastSaleDate     time.Time `json:"last_sale_date" api:"nullable" format:"date"`
	LastSalePrice    float64   `json:"last_sale_price" api:"nullable"`
	Latitude         float64   `json:"latitude" api:"nullable"`
	Longitude        float64   `json:"longitude" api:"nullable"`
	MailingAddress   string    `json:"mailing_address" api:"nullable"`
	MarketValue      float64   `json:"market_value" api:"nullable"`
	OwnerName        string    `json:"owner_name"`
	// Any of "individual", "corporation", "llc", "trust", "government", "other".
	OwnerType ParcelOwnerType `json:"owner_type"`
	// County-assigned parcel identifier.
	ParcelID  string `json:"parcel_id"`
	StateAbbr string `json:"state_abbr"`
	StateFips string `json:"state_fips"`
	YearBuilt int64  `json:"year_built" api:"nullable"`
	Zip       string `json:"zip"`
	Zoning    string `json:"zoning" api:"nullable"`
	// Any of "residential", "commercial", "industrial", "agricultural", "mixed",
	// "other".
	ZoningCategory ParcelZoningCategory `json:"zoning_category" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbsenteeOwner    respjson.Field
		Acreage          respjson.Field
		Address          respjson.Field
		AssessedValue    respjson.Field
		City             respjson.Field
		CountyFips       respjson.Field
		CountyName       respjson.Field
		CrimeScore       respjson.Field
		DealScore        respjson.Field
		ImprovementValue respjson.Field
		LandUse          respjson.Field
		LandValue        respjson.Field
		LastSaleDate     respjson.Field
		LastSalePrice    respjson.Field
		Latitude         respjson.Field
		Longitude        respjson.Field
		MailingAddress   respjson.Field
		MarketValue      respjson.Field
		OwnerName        respjson.Field
		OwnerType        respjson.Field
		ParcelID         respjson.Field
		StateAbbr        respjson.Field
		StateFips        respjson.Field
		YearBuilt        respjson.Field
		Zip              respjson.Field
		Zoning           respjson.Field
		ZoningCategory   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Parcel) RawJSON() string { return r.JSON.raw }
func (r *Parcel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParcelOwnerType string

const (
	ParcelOwnerTypeIndividual  ParcelOwnerType = "individual"
	ParcelOwnerTypeCorporation ParcelOwnerType = "corporation"
	ParcelOwnerTypeLlc         ParcelOwnerType = "llc"
	ParcelOwnerTypeTrust       ParcelOwnerType = "trust"
	ParcelOwnerTypeGovernment  ParcelOwnerType = "government"
	ParcelOwnerTypeOther       ParcelOwnerType = "other"
)

type ParcelZoningCategory string

const (
	ParcelZoningCategoryResidential  ParcelZoningCategory = "residential"
	ParcelZoningCategoryCommercial   ParcelZoningCategory = "commercial"
	ParcelZoningCategoryIndustrial   ParcelZoningCategory = "industrial"
	ParcelZoningCategoryAgricultural ParcelZoningCategory = "agricultural"
	ParcelZoningCategoryMixed        ParcelZoningCategory = "mixed"
	ParcelZoningCategoryOther        ParcelZoningCategory = "other"
)

type Permit struct {
	Contractor    string    `json:"contractor" api:"nullable"`
	Description   string    `json:"description"`
	EstimatedCost float64   `json:"estimated_cost" api:"nullable"`
	IssuedDate    time.Time `json:"issued_date" api:"nullable" format:"date"`
	PermitNumber  string    `json:"permit_number"`
	// Any of "issued", "pending", "approved", "expired", "completed", "denied".
	Status PermitStatus `json:"status"`
	Type   string       `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contractor    respjson.Field
		Description   respjson.Field
		EstimatedCost respjson.Field
		IssuedDate    respjson.Field
		PermitNumber  respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Permit) RawJSON() string { return r.JSON.raw }
func (r *Permit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PermitStatus string

const (
	PermitStatusIssued    PermitStatus = "issued"
	PermitStatusPending   PermitStatus = "pending"
	PermitStatusApproved  PermitStatus = "approved"
	PermitStatusExpired   PermitStatus = "expired"
	PermitStatusCompleted PermitStatus = "completed"
	PermitStatusDenied    PermitStatus = "denied"
)

type RiskAssessment struct {
	AirQuality RiskAssessmentAirQuality `json:"air_quality"`
	Crime      RiskAssessmentCrime      `json:"crime"`
	FloodZone  RiskAssessmentFloodZone  `json:"flood_zone"`
	Wildfire   RiskAssessmentWildfire   `json:"wildfire"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AirQuality  respjson.Field
		Crime       respjson.Field
		FloodZone   respjson.Field
		Wildfire    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskAssessment) RawJSON() string { return r.JSON.raw }
func (r *RiskAssessment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RiskAssessmentAirQuality struct {
	// Any of "good", "moderate", "unhealthy_sensitive", "unhealthy", "very_unhealthy",
	// "hazardous".
	Category string `json:"category" api:"nullable"`
	// Median Air Quality Index value.
	MedianAqi float64 `json:"median_aqi" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		MedianAqi   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskAssessmentAirQuality) RawJSON() string { return r.JSON.raw }
func (r *RiskAssessmentAirQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RiskAssessmentCrime struct {
	// Crime score from 0 (low) to 100 (high).
	Score float64 `json:"score" api:"nullable"`
	// Any of "very_low", "low", "moderate", "high", "very_high".
	Tier string `json:"tier" api:"nullable"`
	// Any of "decreasing", "stable", "increasing".
	Trend string `json:"trend" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Score       respjson.Field
		Tier        respjson.Field
		Trend       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskAssessmentCrime) RawJSON() string { return r.JSON.raw }
func (r *RiskAssessmentCrime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RiskAssessmentFloodZone struct {
	Description  string `json:"description" api:"nullable"`
	InFloodplain bool   `json:"in_floodplain"`
	// FEMA flood zone designation.
	Zone string `json:"zone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		InFloodplain respjson.Field
		Zone         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskAssessmentFloodZone) RawJSON() string { return r.JSON.raw }
func (r *RiskAssessmentFloodZone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RiskAssessmentWildfire struct {
	// Annual burn probability as a decimal.
	BurnProbability float64 `json:"burn_probability" api:"nullable"`
	// Any of "low", "moderate", "high", "very_high", "extreme".
	RiskClass string `json:"risk_class" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BurnProbability respjson.Field
		RiskClass       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskAssessmentWildfire) RawJSON() string { return r.JSON.raw }
func (r *RiskAssessmentWildfire) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetDeedsResponse struct {
	Data []Deed `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetDeedsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetDeedsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GeoJSON FeatureCollection of parcel polygons. Each feature's properties carry
// the basic parcel summary for popup rendering.
type V1ParcelGetGeojsonResponse struct {
	Features []V1ParcelGetGeojsonResponseFeature `json:"features"`
	// Any of "FeatureCollection".
	Type V1ParcelGetGeojsonResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetGeojsonResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetGeojsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetGeojsonResponseFeature struct {
	Geometry   V1ParcelGetGeojsonResponseFeatureGeometry   `json:"geometry"`
	Properties V1ParcelGetGeojsonResponseFeatureProperties `json:"properties"`
	// Any of "Feature".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetGeojsonResponseFeature) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetGeojsonResponseFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetGeojsonResponseFeatureGeometry struct {
	// GeoJSON polygon coordinate rings: outer ring first, then any inner rings.
	Coordinates [][][]float64 `json:"coordinates"`
	// Any of "Polygon".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetGeojsonResponseFeatureGeometry) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetGeojsonResponseFeatureGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetGeojsonResponseFeatureProperties struct {
	Address            string  `json:"address" api:"nullable"`
	City               string  `json:"city" api:"nullable"`
	CountyFips         string  `json:"county_fips"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	OwnerName          string  `json:"owner_name" api:"nullable"`
	ParcelID           string  `json:"parcel_id"`
	PropertyType       string  `json:"property_type" api:"nullable"`
	State              string  `json:"state" api:"nullable"`
	TotalAssessedValue float64 `json:"total_assessed_value" api:"nullable"`
	YearBuilt          int64   `json:"year_built" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		City               respjson.Field
		CountyFips         respjson.Field
		Latitude           respjson.Field
		Longitude          respjson.Field
		OwnerName          respjson.Field
		ParcelID           respjson.Field
		PropertyType       respjson.Field
		State              respjson.Field
		TotalAssessedValue respjson.Field
		YearBuilt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetGeojsonResponseFeatureProperties) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetGeojsonResponseFeatureProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetGeojsonResponseType string

const (
	V1ParcelGetGeojsonResponseTypeFeatureCollection V1ParcelGetGeojsonResponseType = "FeatureCollection"
)

type V1ParcelGetOwnerResponse struct {
	// Any of "individual", "corporation", "llc", "trust", "government", "other".
	EntityType       V1ParcelGetOwnerResponseEntityType       `json:"entity_type"`
	MailingAddress   string                                   `json:"mailing_address"`
	OwnerName        string                                   `json:"owner_name"`
	PortfolioSummary V1ParcelGetOwnerResponsePortfolioSummary `json:"portfolio_summary"`
	Properties       []Parcel                                 `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityType       respjson.Field
		MailingAddress   respjson.Field
		OwnerName        respjson.Field
		PortfolioSummary respjson.Field
		Properties       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetOwnerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetOwnerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetOwnerResponseEntityType string

const (
	V1ParcelGetOwnerResponseEntityTypeIndividual  V1ParcelGetOwnerResponseEntityType = "individual"
	V1ParcelGetOwnerResponseEntityTypeCorporation V1ParcelGetOwnerResponseEntityType = "corporation"
	V1ParcelGetOwnerResponseEntityTypeLlc         V1ParcelGetOwnerResponseEntityType = "llc"
	V1ParcelGetOwnerResponseEntityTypeTrust       V1ParcelGetOwnerResponseEntityType = "trust"
	V1ParcelGetOwnerResponseEntityTypeGovernment  V1ParcelGetOwnerResponseEntityType = "government"
	V1ParcelGetOwnerResponseEntityTypeOther       V1ParcelGetOwnerResponseEntityType = "other"
)

type V1ParcelGetOwnerResponsePortfolioSummary struct {
	PropertyCount      int64    `json:"property_count"`
	States             []string `json:"states"`
	TotalAssessedValue float64  `json:"total_assessed_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PropertyCount      respjson.Field
		States             respjson.Field
		TotalAssessedValue respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetOwnerResponsePortfolioSummary) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetOwnerResponsePortfolioSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetPermitsResponse struct {
	Data []Permit `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetPermitsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetPermitsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive parcel report — every PropRaven data point joined per parcel.
// Subsections are populated on a best-effort basis; missing data is omitted rather
// than nulled.
type V1ParcelGetReportResponse struct {
	ComparableSales []V1ParcelGetReportResponseComparableSale `json:"comparable_sales"`
	Deeds           []Deed                                    `json:"deeds"`
	Parcel          Parcel                                    `json:"parcel"`
	Permits         []Permit                                  `json:"permits"`
	Risk            RiskAssessment                            `json:"risk"`
	UccLiens        []V1ParcelGetReportResponseUccLien        `json:"ucc_liens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparableSales respjson.Field
		Deeds           respjson.Field
		Parcel          respjson.Field
		Permits         respjson.Field
		Risk            respjson.Field
		UccLiens        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetReportResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetReportResponseComparableSale struct {
	CompBaths       float64   `json:"comp_baths" api:"nullable"`
	CompBeds        int64     `json:"comp_beds" api:"nullable"`
	CompParcelID    string    `json:"comp_parcel_id"`
	CompSaleDate    time.Time `json:"comp_sale_date" format:"date"`
	CompSalePrice   float64   `json:"comp_sale_price" api:"nullable"`
	CompSqft        int64     `json:"comp_sqft" api:"nullable"`
	DistanceMiles   float64   `json:"distance_miles" api:"nullable"`
	SimilarityScore float64   `json:"similarity_score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompBaths       respjson.Field
		CompBeds        respjson.Field
		CompParcelID    respjson.Field
		CompSaleDate    respjson.Field
		CompSalePrice   respjson.Field
		CompSqft        respjson.Field
		DistanceMiles   respjson.Field
		SimilarityScore respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetReportResponseComparableSale) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetReportResponseComparableSale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetReportResponseUccLien struct {
	DebtorName      string    `json:"debtor_name" api:"nullable"`
	FilingDate      time.Time `json:"filing_date" format:"date"`
	FilingID        string    `json:"filing_id"`
	FilingStatus    string    `json:"filing_status" api:"nullable"`
	LapseDate       time.Time `json:"lapse_date" format:"date"`
	MatchConfidence float64   `json:"match_confidence"`
	MatchMethod     string    `json:"match_method" api:"nullable"`
	SecuredParty    string    `json:"secured_party" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DebtorName      respjson.Field
		FilingDate      respjson.Field
		FilingID        respjson.Field
		FilingStatus    respjson.Field
		LapseDate       respjson.Field
		MatchConfidence respjson.Field
		MatchMethod     respjson.Field
		SecuredParty    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetReportResponseUccLien) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetReportResponseUccLien) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetTrafficHistoryResponse struct {
	// Year-keyed historical counts (newest last).
	History []V1ParcelGetTrafficHistoryResponseHistory `json:"history"`
	Station V1ParcelGetTrafficHistoryResponseStation   `json:"station"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		History     respjson.Field
		Station     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetTrafficHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetTrafficHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetTrafficHistoryResponseHistory struct {
	Aadt int64 `json:"aadt" api:"nullable"`
	Year int64 `json:"year"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aadt        respjson.Field
		Year        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetTrafficHistoryResponseHistory) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetTrafficHistoryResponseHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetTrafficHistoryResponseStation struct {
	ID string `json:"id"`
	// Most recent AADT count.
	AadtCurrent     int64   `json:"aadt_current" api:"nullable"`
	AadtYear        int64   `json:"aadt_year" api:"nullable"`
	Cagr3yr         float64 `json:"cagr_3yr" api:"nullable"`
	Cagr5yr         float64 `json:"cagr_5yr" api:"nullable"`
	Cagr7yr         float64 `json:"cagr_7yr" api:"nullable"`
	FunctionalClass string  `json:"functional_class" api:"nullable"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	RouteName       string  `json:"route_name" api:"nullable"`
	StationID       string  `json:"station_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AadtCurrent     respjson.Field
		AadtYear        respjson.Field
		Cagr3yr         respjson.Field
		Cagr5yr         respjson.Field
		Cagr7yr         respjson.Field
		FunctionalClass respjson.Field
		Latitude        respjson.Field
		Longitude       respjson.Field
		RouteName       respjson.Field
		StationID       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ParcelGetTrafficHistoryResponseStation) RawJSON() string { return r.JSON.raw }
func (r *V1ParcelGetTrafficHistoryResponseStation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ParcelGetGeojsonParams struct {
	// Bounding box `west,south,east,north`.
	Bbox string `query:"bbox" api:"required" json:"-"`
	// Map zoom level. Below 14 returns an empty collection.
	Zoom int64 `query:"zoom" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1ParcelGetGeojsonParams]'s query parameters as
// `url.Values`.
func (r V1ParcelGetGeojsonParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ParcelGetReportParams struct {
	// 5-digit county FIPS. Strongly recommended when passing a county-local id.
	CountyFips param.Opt[string] `query:"county_fips,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ParcelGetReportParams]'s query parameters as
// `url.Values`.
func (r V1ParcelGetReportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ParcelGetTrafficHistoryParams struct {
	// Latitude.
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude.
	Lng float64 `query:"lng" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1ParcelGetTrafficHistoryParams]'s query parameters as
// `url.Values`.
func (r V1ParcelGetTrafficHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
