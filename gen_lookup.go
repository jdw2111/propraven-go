// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// LookupService groups the lookup operations. Use it as client.Lookup.
type LookupService struct {
	client *Client
}

// Get: Exact parcel lookup (UUID or APN)
//
// Resolve one parcel from an exact key: a PropRaven parcel UUID or an APN. Answered from the
// search index (the `parcel` fields are a subset of GET /parcels/{id}). Anonymous callers are
// IP-throttled on the restricted budget (100/day); fuzzy address input needs a key and is answered
// by POST /lookup/batch.
//
// HTTP: GET /api/v1/lookup
func (s *LookupService) Get(ctx context.Context, params *LookupGetParams, opts ...RequestOption) (*LookupGetResponse, error) {
	var out LookupGetResponse
	if err := s.client.do(ctx, buildLookupGetRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildLookupGetRequest(params *LookupGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/lookup")
	if params != nil {
		addQuery(req.query, "q", params.Q)
	}
	return req
}

// LookupGetParams holds the query, header and JSON-body parameters of [LookupService.Get]. Pass
// nil when you need none.
type LookupGetParams struct {
	// PropRaven parcel UUID or APN (min 2 chars).
	//
	// Required.
	Q *string `query:"q" json:"-"`
}

// LookupGetResponse: Exact parcel lookup (UUID or APN)
type LookupGetResponse struct {
	Found  *bool                   `json:"found,omitempty"`
	Query  *string                 `json:"query,omitempty"`
	Parcel LookupGetResponseParcel `json:"parcel"`
}

// LookupGetResponseParcel is generated from the OpenAPI spec.
type LookupGetResponseParcel struct {
	ID                 string   `json:"id"`
	APN                *string  `json:"apn,omitempty"`
	CountyFIPS         string   `json:"county_fips"`
	StateFIPS          string   `json:"state_fips"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	LastSalePrice      *float64 `json:"last_sale_price,omitempty"`
	LandUseCode        *string  `json:"land_use_code,omitempty"`
	OwnerName          *string  `json:"owner_name,omitempty"`
}

// UnmarshalJSON decodes LookupGetResponseParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LookupGetResponseParcel) UnmarshalJSON(data []byte) error {
	type plain LookupGetResponseParcel
	aux := struct {
		*plain
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		LastSalePrice      lenientNumber[float64] `json:"last_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	return softTypeError(err)
}

// Batch: Resolve up to 500 parcel queries in one call
//
// Each query resolves independently (UUID → APN → address), so a partial batch degrades row by
// row. Answered from the search index: fields listed in `unavailable_fields` are null on every row
// of this endpoint. Each matched row is one lookup against your plan.
//
// HTTP: POST /api/v1/lookup/batch
func (s *LookupService) Batch(ctx context.Context, params *LookupBatchParams, opts ...RequestOption) (*LookupBatchResponse, error) {
	var out LookupBatchResponse
	if err := s.client.do(ctx, buildLookupBatchRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildLookupBatchRequest(params *LookupBatchParams) *apiRequest {
	req := newRequest("POST", "/api/v1/lookup/batch")
	if params != nil {
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// LookupBatchParams holds the query, header and JSON-body parameters of [LookupService.Batch].
// Pass nil when you need none.
type LookupBatchParams struct {
	// Parcel UUIDs, APNs, canonical ids or street addresses (≤ 500).
	//
	// Required (JSON body).
	Queries []string `json:"queries"`
}

// LookupBatchResponse: Resolve up to 500 parcel queries in one call
type LookupBatchResponse struct {
	Items             []LookupBatchResponseItems `json:"items"`
	LookupsCharged    float64                    `json:"lookups_charged"`
	UnavailableFields []*string                  `json:"unavailable_fields"`
	Tier              string                     `json:"tier"`
	Notice            string                     `json:"notice"`
}

// UnmarshalJSON decodes LookupBatchResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LookupBatchResponse) UnmarshalJSON(data []byte) error {
	type plain LookupBatchResponse
	aux := struct {
		*plain
		LookupsCharged lenientNumber[float64] `json:"lookups_charged"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LookupsCharged.assign(&r.LookupsCharged)
	return softTypeError(err)
}

// LookupBatchResponseItems is generated from the OpenAPI spec.
type LookupBatchResponseItems struct {
	Query     *string                         `json:"query,omitempty"`
	MatchType *string                         `json:"match_type,omitempty"`
	Parcel    *LookupBatchResponseItemsParcel `json:"parcel,omitempty"`
}

// LookupBatchResponseItemsParcel is generated from the OpenAPI spec.
type LookupBatchResponseItemsParcel struct {
	ID                 string                                 `json:"id"`
	APN                *string                                `json:"apn,omitempty"`
	Address            *string                                `json:"address,omitempty"`
	City               *string                                `json:"city,omitempty"`
	State              *string                                `json:"state,omitempty"`
	Zip                *string                                `json:"zip,omitempty"`
	CountyFIPS         string                                 `json:"county_fips"`
	StateFIPS          string                                 `json:"state_fips"`
	OwnerName          *string                                `json:"owner_name,omitempty"`
	LandValue          *float64                               `json:"land_value,omitempty"`
	ImprovementValue   *float64                               `json:"improvement_value,omitempty"`
	TotalAssessedValue *int64                                 `json:"total_assessed_value,omitempty"`
	LastSale           LookupBatchResponseItemsParcelLastSale `json:"last_sale"`

	// (Only null in observed responses.)
	Permits     json.RawMessage `json:"permits"`
	HazardScore *float64        `json:"hazard_score,omitempty"`
	Latitude    *float64        `json:"latitude,omitempty"`
	Longitude   *float64        `json:"longitude,omitempty"`
	HasGeometry *bool           `json:"has_geometry,omitempty"`
}

// UnmarshalJSON decodes LookupBatchResponseItemsParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LookupBatchResponseItemsParcel) UnmarshalJSON(data []byte) error {
	type plain LookupBatchResponseItemsParcel
	aux := struct {
		*plain
		LandValue          lenientNumber[float64] `json:"land_value"`
		ImprovementValue   lenientNumber[float64] `json:"improvement_value"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		HazardScore        lenientNumber[float64] `json:"hazard_score"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LandValue.assignPtr(&r.LandValue)
	aux.ImprovementValue.assignPtr(&r.ImprovementValue)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.HazardScore.assignPtr(&r.HazardScore)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// LookupBatchResponseItemsParcelLastSale is generated from the OpenAPI spec.
type LookupBatchResponseItemsParcelLastSale struct {
	Date  *string  `json:"date,omitempty"`
	Price *float64 `json:"price,omitempty"`
}

// UnmarshalJSON decodes LookupBatchResponseItemsParcelLastSale, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LookupBatchResponseItemsParcelLastSale) UnmarshalJSON(data []byte) error {
	type plain LookupBatchResponseItemsParcelLastSale
	aux := struct {
		*plain
		Price lenientNumber[float64] `json:"price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Price.assignPtr(&r.Price)
	return softTypeError(err)
}
