// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// VerifyService groups the verify operations. Use it as client.Verify.
type VerifyService struct {
	client *Client
}

// Get: Verify facts for one parcel
//
// Verify catalogued FACTS: one (parcel, field) assertion per lookup, batch (POST body { lookups:
// [{ parcel_id, fields }] }, up to 1000 parcels) or single (GET ?parcel_id&fields). Priced per
// lookup, bulk-discounted ($0.02/$0.012/$0.006), capped $20. Pay from a prepaid credit balance
// (X-CREDIT-TOKEN) or per call via x402. preview=true returns count+price free; empty/all-invalid
// is free. Fields whitelisted (is_sfha, flood_zone, nri_risk_score, owner_occupied, is_absentee,
// owner_name, assessed_value, market_value, building_sqft, year_built, property_type, zoning,
// lot_size_acres, last_sale_date/price, address/city/state/zip); unknown parcel -> found:false.
// `owner_name` is people data: verifying it requires an account (API key or signed-in session). A
// caller without one -- anonymous, wallet-only x402 or credit token -- that asks for owner_name is
// refused with HTTP 401 `code: "account_required"` before the quote or any payment; every other
// field is unaffected.
//
// HTTP: GET /api/v1/verify
func (s *VerifyService) Get(ctx context.Context, params *VerifyGetParams, opts ...RequestOption) (*VerifyGetResponse, error) {
	var out VerifyGetResponse
	if err := s.client.do(ctx, buildVerifyGetRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildVerifyGetRequest(params *VerifyGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/verify")
	if params != nil {
		addQuery(req.query, "parcel_id", params.ParcelID)
		addQuery(req.query, "fields", params.Fields)
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// VerifyGetParams holds the query, header and JSON-body parameters of [VerifyService.Get]. Pass
// nil when you need none.
type VerifyGetParams struct {
	// Canonical state:county:parcel.
	//
	// Required.
	ParcelID *string `query:"parcel_id" json:"-"`

	// Comma-separated field aliases.
	//
	// Required.
	Fields *string `query:"fields" json:"-"`

	// FREE count + price.
	Preview *string `query:"preview" json:"-"`

	// Prepaid credit token.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`

	// Base64 x402 PaymentPayload.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// VerifyGetResponse: Verify facts for one parcel
type VerifyGetResponse struct {
	LookupCount int64                       `json:"lookup_count"`
	ParcelCount int64                       `json:"parcel_count"`
	Quote       VerifyGetResponseQuote      `json:"quote"`
	Provenance  VerifyGetResponseProvenance `json:"provenance"`
	Preview     bool                        `json:"preview"`
	Note        string                      `json:"note"`
}

// UnmarshalJSON decodes VerifyGetResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyGetResponse) UnmarshalJSON(data []byte) error {
	type plain VerifyGetResponse
	aux := struct {
		*plain
		LookupCount lenientNumber[int64] `json:"lookup_count"`
		ParcelCount lenientNumber[int64] `json:"parcel_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LookupCount.assign(&r.LookupCount)
	aux.ParcelCount.assign(&r.ParcelCount)
	return softTypeError(err)
}

// VerifyGetResponseQuote is generated from the OpenAPI spec.
type VerifyGetResponseQuote struct {
	PerLookup       VerifyGetResponseQuotePerLookup `json:"per_lookup"`
	Total           VerifyGetResponseQuoteTotal     `json:"total"`
	TotalUsd        float64                         `json:"total_usd"`
	TotalAtomicUsdc string                          `json:"total_atomic_usdc"`
	Asset           string                          `json:"asset"`
	LookupCount     int64                           `json:"lookup_count"`
	Breakdown       VerifyGetResponseQuoteBreakdown `json:"breakdown"`
	Pay             []*string                       `json:"pay"`
	Note            string                          `json:"note"`
}

// UnmarshalJSON decodes VerifyGetResponseQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyGetResponseQuote) UnmarshalJSON(data []byte) error {
	type plain VerifyGetResponseQuote
	aux := struct {
		*plain
		TotalUsd    lenientNumber[float64] `json:"total_usd"`
		LookupCount lenientNumber[int64]   `json:"lookup_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalUsd.assign(&r.TotalUsd)
	aux.LookupCount.assign(&r.LookupCount)
	return softTypeError(err)
}

// VerifyGetResponseQuotePerLookup is generated from the OpenAPI spec.
type VerifyGetResponseQuotePerLookup struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// VerifyGetResponseQuoteTotal is generated from the OpenAPI spec.
type VerifyGetResponseQuoteTotal struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// VerifyGetResponseQuoteBreakdown is generated from the OpenAPI spec.
type VerifyGetResponseQuoteBreakdown struct {
	PerLookupUsd float64 `json:"per_lookup_usd"`
	TotalRaw     float64 `json:"total_raw"`
	Capped       bool    `json:"capped"`
}

// UnmarshalJSON decodes VerifyGetResponseQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyGetResponseQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain VerifyGetResponseQuoteBreakdown
	aux := struct {
		*plain
		PerLookupUsd lenientNumber[float64] `json:"per_lookup_usd"`
		TotalRaw     lenientNumber[float64] `json:"total_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PerLookupUsd.assign(&r.PerLookupUsd)
	aux.TotalRaw.assign(&r.TotalRaw)
	return softTypeError(err)
}

// VerifyGetResponseProvenance is generated from the OpenAPI spec.
type VerifyGetResponseProvenance struct {
	Source           string                                      `json:"source"`
	SourceDatasets   []*string                                   `json:"source_datasets"`
	AsOf             *string                                     `json:"as_of,omitempty"`
	AsOfBasis        *string                                     `json:"as_of_basis,omitempty"`
	ServingEpoch     *string                                     `json:"serving_epoch,omitempty"`
	FreshnessStatus  string                                      `json:"freshness_status"`
	CatalogReference VerifyGetResponseProvenanceCatalogReference `json:"catalog_reference"`
	Note             string                                      `json:"note"`
}

// VerifyGetResponseProvenanceCatalogReference is generated from the OpenAPI spec.
type VerifyGetResponseProvenanceCatalogReference struct {
	CatalogVersion     string                                                 `json:"catalog_version"`
	CatalogGeneratedAt string                                                 `json:"catalog_generated_at"`
	CatalogSeal        VerifyGetResponseProvenanceCatalogReferenceCatalogSeal `json:"catalog_seal"`
}

// VerifyGetResponseProvenanceCatalogReferenceCatalogSeal is generated from the OpenAPI spec.
type VerifyGetResponseProvenanceCatalogReferenceCatalogSeal struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
	Covers    string `json:"covers"`
}

// Batch: Verify facts (batch, paid per lookup) - FREE preview
//
// Verify catalogued FACTS: one (parcel, field) assertion per lookup, batch (POST body { lookups:
// [{ parcel_id, fields }] }, up to 1000 parcels) or single (GET ?parcel_id&fields). Priced per
// lookup, bulk-discounted ($0.02/$0.012/$0.006), capped $20. Pay from a prepaid credit balance
// (X-CREDIT-TOKEN) or per call via x402. preview=true returns count+price free; empty/all-invalid
// is free. Fields whitelisted (is_sfha, flood_zone, nri_risk_score, owner_occupied, is_absentee,
// owner_name, assessed_value, market_value, building_sqft, year_built, property_type, zoning,
// lot_size_acres, last_sale_date/price, address/city/state/zip); unknown parcel -> found:false.
// `owner_name` is people data: verifying it requires an account (API key or signed-in session). A
// caller without one -- anonymous, wallet-only x402 or credit token -- that asks for owner_name is
// refused with HTTP 401 `code: "account_required"` before the quote or any payment; every other
// field is unaffected.
//
// HTTP: POST /api/v1/verify
func (s *VerifyService) Batch(ctx context.Context, params *VerifyBatchParams, opts ...RequestOption) (*VerifyBatchResponse, error) {
	var out VerifyBatchResponse
	if err := s.client.do(ctx, buildVerifyBatchRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildVerifyBatchRequest(params *VerifyBatchParams) *apiRequest {
	req := newRequest("POST", "/api/v1/verify")
	if params != nil {
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
		setHeader(req.header, "X-PAYMENT", params.Payment)
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// VerifyBatchParams holds the query, header and JSON-body parameters of [VerifyService.Batch].
// Pass nil when you need none.
type VerifyBatchParams struct {
	// FREE count + price.
	Preview *string `query:"preview" json:"-"`

	// Prepaid credit token to debit the batch.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`

	// Base64 x402 PaymentPayload.
	Payment *string `header:"X-PAYMENT" json:"-"`

	// Required (JSON body).
	Lookups []VerifyBatchParamsLookups `json:"lookups"`
}

// VerifyBatchParamsLookups is generated from the OpenAPI spec.
type VerifyBatchParamsLookups struct {
	ParcelID string   `json:"parcel_id"`
	Fields   []string `json:"fields"`
}

// VerifyBatchResponse: Verify facts (batch, paid per lookup) - FREE preview
type VerifyBatchResponse struct {
	LookupCount int64                         `json:"lookup_count"`
	ParcelCount int64                         `json:"parcel_count"`
	Quote       VerifyBatchResponseQuote      `json:"quote"`
	Provenance  VerifyBatchResponseProvenance `json:"provenance"`
	Preview     bool                          `json:"preview"`
	Note        string                        `json:"note"`
}

// UnmarshalJSON decodes VerifyBatchResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyBatchResponse) UnmarshalJSON(data []byte) error {
	type plain VerifyBatchResponse
	aux := struct {
		*plain
		LookupCount lenientNumber[int64] `json:"lookup_count"`
		ParcelCount lenientNumber[int64] `json:"parcel_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LookupCount.assign(&r.LookupCount)
	aux.ParcelCount.assign(&r.ParcelCount)
	return softTypeError(err)
}

// VerifyBatchResponseQuote is generated from the OpenAPI spec.
type VerifyBatchResponseQuote struct {
	PerLookup       VerifyBatchResponseQuotePerLookup `json:"per_lookup"`
	Total           VerifyBatchResponseQuoteTotal     `json:"total"`
	TotalUsd        float64                           `json:"total_usd"`
	TotalAtomicUsdc string                            `json:"total_atomic_usdc"`
	Asset           string                            `json:"asset"`
	LookupCount     int64                             `json:"lookup_count"`
	Breakdown       VerifyBatchResponseQuoteBreakdown `json:"breakdown"`
	Pay             []*string                         `json:"pay"`
	Note            string                            `json:"note"`
}

// UnmarshalJSON decodes VerifyBatchResponseQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyBatchResponseQuote) UnmarshalJSON(data []byte) error {
	type plain VerifyBatchResponseQuote
	aux := struct {
		*plain
		TotalUsd    lenientNumber[float64] `json:"total_usd"`
		LookupCount lenientNumber[int64]   `json:"lookup_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalUsd.assign(&r.TotalUsd)
	aux.LookupCount.assign(&r.LookupCount)
	return softTypeError(err)
}

// VerifyBatchResponseQuotePerLookup is generated from the OpenAPI spec.
type VerifyBatchResponseQuotePerLookup struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// VerifyBatchResponseQuoteTotal is generated from the OpenAPI spec.
type VerifyBatchResponseQuoteTotal struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// VerifyBatchResponseQuoteBreakdown is generated from the OpenAPI spec.
type VerifyBatchResponseQuoteBreakdown struct {
	PerLookupUsd float64 `json:"per_lookup_usd"`
	TotalRaw     float64 `json:"total_raw"`
	Capped       bool    `json:"capped"`
}

// UnmarshalJSON decodes VerifyBatchResponseQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *VerifyBatchResponseQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain VerifyBatchResponseQuoteBreakdown
	aux := struct {
		*plain
		PerLookupUsd lenientNumber[float64] `json:"per_lookup_usd"`
		TotalRaw     lenientNumber[float64] `json:"total_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PerLookupUsd.assign(&r.PerLookupUsd)
	aux.TotalRaw.assign(&r.TotalRaw)
	return softTypeError(err)
}

// VerifyBatchResponseProvenance is generated from the OpenAPI spec.
type VerifyBatchResponseProvenance struct {
	Source           string                                        `json:"source"`
	SourceDatasets   []*string                                     `json:"source_datasets"`
	AsOf             *string                                       `json:"as_of,omitempty"`
	AsOfBasis        *string                                       `json:"as_of_basis,omitempty"`
	ServingEpoch     *string                                       `json:"serving_epoch,omitempty"`
	FreshnessStatus  string                                        `json:"freshness_status"`
	CatalogReference VerifyBatchResponseProvenanceCatalogReference `json:"catalog_reference"`
	Note             string                                        `json:"note"`
}

// VerifyBatchResponseProvenanceCatalogReference is generated from the OpenAPI spec.
type VerifyBatchResponseProvenanceCatalogReference struct {
	CatalogVersion     string                                                   `json:"catalog_version"`
	CatalogGeneratedAt string                                                   `json:"catalog_generated_at"`
	CatalogSeal        VerifyBatchResponseProvenanceCatalogReferenceCatalogSeal `json:"catalog_seal"`
}

// VerifyBatchResponseProvenanceCatalogReferenceCatalogSeal is generated from the OpenAPI spec.
type VerifyBatchResponseProvenanceCatalogReferenceCatalogSeal struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
	Covers    string `json:"covers"`
}
