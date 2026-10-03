// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// OwnersService groups the owners operations. Use it as client.Owners.
type OwnersService struct {
	client *Client
}

// Search: Search property owners
//
// Search for property owners by name. Returns owner profiles with property counts and portfolio
// values.
//
// HTTP: GET /api/v1/owners/search
func (s *OwnersService) Search(ctx context.Context, params *OwnersSearchParams, opts ...RequestOption) (*OwnersSearchResponse, error) {
	var out OwnersSearchResponse
	if err := s.client.do(ctx, buildOwnersSearchRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersSearchRequest(params *OwnersSearchParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/search")
	if params != nil {
		addQuery(req.query, "q", params.Q)
		addQuery(req.query, "min_properties", params.MinProperties)
		addQuery(req.query, "limit", params.Limit)
	}
	return req
}

// OwnersSearchParams holds the query, header and JSON-body parameters of [OwnersService.Search].
// Pass nil when you need none.
type OwnersSearchParams struct {
	// Search query for owner name.
	//
	// Required.
	Q *string `query:"q" json:"-"`

	// Minimum number of properties owned.
	MinProperties *int64 `query:"min_properties" json:"-"`
	Limit         *int64 `query:"limit" json:"-"`
}

// OwnersSearchResponse: Search property owners
type OwnersSearchResponse struct {
	Data  []Owner `json:"data"`
	Count int64   `json:"count"`
}

// UnmarshalJSON decodes OwnersSearchResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersSearchResponse) UnmarshalJSON(data []byte) error {
	type plain OwnersSearchResponse
	aux := struct {
		*plain
		Count lenientNumber[int64] `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assign(&r.Count)
	return softTypeError(err)
}

// Get: Get owner profile
//
// Retrieve a specific owner profile by name, including property count, total assessed value,
// entity type, and states.
//
// HTTP: GET /api/v1/owners/{name}
func (s *OwnersService) Get(ctx context.Context, name string, params *OwnersGetParams, opts ...RequestOption) (*OwnersGetResponse, error) {
	var out OwnersGetResponse
	if err := s.client.do(ctx, buildOwnersGetRequest(name, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersGetRequest(name string, params *OwnersGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/"+pathParam(name))
	return req
}

// OwnersGetParams holds the query, header and JSON-body parameters of [OwnersService.Get]. Pass
// nil when you need none.
type OwnersGetParams struct {
}

// OwnersGetResponse: Get owner profile
type OwnersGetResponse = Owner

// Properties: Get owner's properties
//
// Retrieve the list of properties owned by a specific owner.
//
// HTTP: GET /api/v1/owners/{name}/properties
func (s *OwnersService) Properties(ctx context.Context, name string, params *OwnersPropertiesParams, opts ...RequestOption) (*OwnersPropertiesResponse, error) {
	var out OwnersPropertiesResponse
	if err := s.client.do(ctx, buildOwnersPropertiesRequest(name, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

// PropertiesIter iterates every item of [OwnersService.Properties] across pages (offset pagination
// over "data"): it advances offset by the page size and stops on a short page, when offset reaches
// total, or when has_more is false. Use [IterOptions] for the page size and an item cap.
func (s *OwnersService) PropertiesIter(ctx context.Context, name string, params *OwnersPropertiesParams, iter IterOptions, opts ...RequestOption) *Iter[OwnersPropertiesResponseData] {
	var p OwnersPropertiesParams
	if params != nil {
		p = *params
	}
	return newOffsetIter[OwnersPropertiesResponseData](ctx, s.client, iter, "data", p.Limit, p.Offset, func(limit, offset int64) *apiRequest {
		q := p
		q.Limit = &limit
		q.Offset = &offset
		return buildOwnersPropertiesRequest(name, &q)
	}, opts)
}

func buildOwnersPropertiesRequest(name string, params *OwnersPropertiesParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/"+pathParam(name)+"/properties")
	if params != nil {
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
	}
	return req
}

// OwnersPropertiesParams holds the query, header and JSON-body parameters of
// [OwnersService.Properties]. Pass nil when you need none.
type OwnersPropertiesParams struct {
	Limit  *int64 `query:"limit" json:"-"`
	Offset *int64 `query:"offset" json:"-"`
}

// OwnersPropertiesResponse: Get owner's properties
type OwnersPropertiesResponse struct {
	Data   []OwnersPropertiesResponseData `json:"data"`
	Total  int64                          `json:"total"`
	Limit  int64                          `json:"limit"`
	Offset int64                          `json:"offset"`
}

// UnmarshalJSON decodes OwnersPropertiesResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPropertiesResponse) UnmarshalJSON(data []byte) error {
	type plain OwnersPropertiesResponse
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

// OwnersPropertiesResponseData is generated from the OpenAPI spec.
type OwnersPropertiesResponseData struct {
	StateFIPS          string   `json:"state_fips"`
	CountyFIPS         string   `json:"county_fips"`
	ParcelID           string   `json:"parcel_id"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	LandValue          *float64 `json:"land_value,omitempty"`
	ImprovementValue   *float64 `json:"improvement_value,omitempty"`
	LotSizeAcres       *float64 `json:"lot_size_acres,omitempty"`
	Zoning             *string  `json:"zoning,omitempty"`
	ZoningCodeRaw      *string  `json:"zoning_code_raw,omitempty"`
	LandUseDesc        *string  `json:"land_use_desc,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
}

// UnmarshalJSON decodes OwnersPropertiesResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPropertiesResponseData) UnmarshalJSON(data []byte) error {
	type plain OwnersPropertiesResponseData
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		LandValue          lenientNumber[float64] `json:"land_value"`
		ImprovementValue   lenientNumber[float64] `json:"improvement_value"`
		LotSizeAcres       lenientNumber[float64] `json:"lot_size_acres"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LandValue.assignPtr(&r.LandValue)
	aux.ImprovementValue.assignPtr(&r.ImprovementValue)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	return softTypeError(err)
}

// Portfolio: Get owner portfolio summary
//
// Retrieve an owner's portfolio with aggregated summary statistics and property breakdown.
//
// HTTP: GET /api/v1/owners/{name}/portfolio
func (s *OwnersService) Portfolio(ctx context.Context, name string, params *OwnersPortfolioParams, opts ...RequestOption) (*OwnersPortfolioResponse, error) {
	var out OwnersPortfolioResponse
	if err := s.client.do(ctx, buildOwnersPortfolioRequest(name, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersPortfolioRequest(name string, params *OwnersPortfolioParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/"+pathParam(name)+"/portfolio")
	return req
}

// OwnersPortfolioParams holds the query, header and JSON-body parameters of
// [OwnersService.Portfolio]. Pass nil when you need none.
type OwnersPortfolioParams struct {
}

// OwnersPortfolioResponse: Get owner portfolio summary
type OwnersPortfolioResponse struct {
	OwnerName  string                              `json:"owner_name"`
	Properties []OwnersPortfolioResponseProperties `json:"properties"`
	Summary    OwnersPortfolioResponseSummary      `json:"summary"`
	Match      OwnersPortfolioResponseMatch        `json:"match"`
	EntityType *OwnersPortfolioResponseEntityType  `json:"entity_type,omitempty"`
}

// OwnersPortfolioResponseProperties is generated from the OpenAPI spec.
type OwnersPortfolioResponseProperties struct {
	ParcelID                 string            `json:"parcel_id"`
	CountyFIPS               string            `json:"county_fips"`
	StateFIPS                string            `json:"state_fips"`
	Address                  *string           `json:"address,omitempty"`
	City                     *string           `json:"city,omitempty"`
	State                    *string           `json:"state,omitempty"`
	Zip                      *string           `json:"zip,omitempty"`
	Latitude                 *float64          `json:"latitude,omitempty"`
	Longitude                *float64          `json:"longitude,omitempty"`
	OwnerName                *string           `json:"owner_name,omitempty"`
	TotalAssessedValue       *float64          `json:"total_assessed_value,omitempty"`
	LandAssessedValue        *float64          `json:"land_assessed_value,omitempty"`
	ImprovementAssessedValue *float64          `json:"improvement_assessed_value,omitempty"`
	LotSizeAcres             *float64          `json:"lot_size_acres,omitempty"`
	BuildingSqft             *float64          `json:"building_sqft,omitempty"`
	YearBuilt                *int64            `json:"year_built,omitempty"`
	Zoning                   *string           `json:"zoning,omitempty"`
	ZoningCodeRaw            *string           `json:"zoning_code_raw,omitempty"`
	LastSaleDate             *string           `json:"last_sale_date,omitempty"`
	LastSalePrice            *float64          `json:"last_sale_price,omitempty"`
	PropertyType             *string           `json:"property_type,omitempty"`
	MatchBasis               string            `json:"match_basis"`
	MatchConfidence          *string           `json:"match_confidence,omitempty"`
	OwnerRoles               []json.RawMessage `json:"owner_roles"`
}

// UnmarshalJSON decodes OwnersPortfolioResponseProperties, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPortfolioResponseProperties) UnmarshalJSON(data []byte) error {
	type plain OwnersPortfolioResponseProperties
	aux := struct {
		*plain
		Latitude                 lenientNumber[float64] `json:"latitude"`
		Longitude                lenientNumber[float64] `json:"longitude"`
		TotalAssessedValue       lenientNumber[float64] `json:"total_assessed_value"`
		LandAssessedValue        lenientNumber[float64] `json:"land_assessed_value"`
		ImprovementAssessedValue lenientNumber[float64] `json:"improvement_assessed_value"`
		LotSizeAcres             lenientNumber[float64] `json:"lot_size_acres"`
		BuildingSqft             lenientNumber[float64] `json:"building_sqft"`
		YearBuilt                lenientNumber[int64]   `json:"year_built"`
		LastSalePrice            lenientNumber[float64] `json:"last_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LandAssessedValue.assignPtr(&r.LandAssessedValue)
	aux.ImprovementAssessedValue.assignPtr(&r.ImprovementAssessedValue)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.BuildingSqft.assignPtr(&r.BuildingSqft)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	return softTypeError(err)
}

// OwnersPortfolioResponseSummary is generated from the OpenAPI spec.
type OwnersPortfolioResponseSummary struct {
	Count              int64                                   `json:"count"`
	Returned           float64                                 `json:"returned"`
	Limit              int64                                   `json:"limit"`
	Offset             int64                                   `json:"offset"`
	Truncated          bool                                    `json:"truncated"`
	TotalValue         int64                                   `json:"total_value"`
	TotalAcreage       float64                                 `json:"total_acreage"`
	States             []*float64                              `json:"states"`
	ByState            []OwnersPortfolioResponseSummaryByState `json:"by_state"`
	PropertyCount      *int64                                  `json:"property_count,omitempty"`
	TotalAssessedValue *float64                                `json:"total_assessed_value,omitempty"`
	AvgAssessedValue   *float64                                `json:"avg_assessed_value,omitempty"`
	Counties           *int64                                  `json:"counties,omitempty"`
	ZoningBreakdown    map[string]int64                        `json:"zoning_breakdown,omitempty"`
}

// UnmarshalJSON decodes OwnersPortfolioResponseSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPortfolioResponseSummary) UnmarshalJSON(data []byte) error {
	type plain OwnersPortfolioResponseSummary
	aux := struct {
		*plain
		Count              lenientNumber[int64]   `json:"count"`
		Returned           lenientNumber[float64] `json:"returned"`
		Limit              lenientNumber[int64]   `json:"limit"`
		Offset             lenientNumber[int64]   `json:"offset"`
		TotalValue         lenientNumber[int64]   `json:"total_value"`
		TotalAcreage       lenientNumber[float64] `json:"total_acreage"`
		PropertyCount      lenientNumber[int64]   `json:"property_count"`
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		AvgAssessedValue   lenientNumber[float64] `json:"avg_assessed_value"`
		Counties           lenientNumber[int64]   `json:"counties"`
		ZoningBreakdown    lenientMap[int64]      `json:"zoning_breakdown"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assign(&r.Count)
	aux.Returned.assign(&r.Returned)
	aux.Limit.assign(&r.Limit)
	aux.Offset.assign(&r.Offset)
	aux.TotalValue.assign(&r.TotalValue)
	aux.TotalAcreage.assign(&r.TotalAcreage)
	aux.PropertyCount.assignPtr(&r.PropertyCount)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.AvgAssessedValue.assignPtr(&r.AvgAssessedValue)
	aux.Counties.assignPtr(&r.Counties)
	aux.ZoningBreakdown.assign(&r.ZoningBreakdown)
	return softTypeError(err)
}

// OwnersPortfolioResponseSummaryByState is generated from the OpenAPI spec.
type OwnersPortfolioResponseSummaryByState struct {
	State        *string  `json:"state,omitempty"`
	StateFIPS    string   `json:"state_fips"`
	Abbr         *string  `json:"abbr,omitempty"`
	Count        *int64   `json:"count,omitempty"`
	TotalValue   *int64   `json:"total_value,omitempty"`
	TotalAcreage *float64 `json:"total_acreage,omitempty"`
}

// UnmarshalJSON decodes OwnersPortfolioResponseSummaryByState, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPortfolioResponseSummaryByState) UnmarshalJSON(data []byte) error {
	type plain OwnersPortfolioResponseSummaryByState
	aux := struct {
		*plain
		Count        lenientNumber[int64]   `json:"count"`
		TotalValue   lenientNumber[int64]   `json:"total_value"`
		TotalAcreage lenientNumber[float64] `json:"total_acreage"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assignPtr(&r.Count)
	aux.TotalValue.assignPtr(&r.TotalValue)
	aux.TotalAcreage.assignPtr(&r.TotalAcreage)
	return softTypeError(err)
}

// OwnersPortfolioResponseMatch is generated from the OpenAPI spec.
type OwnersPortfolioResponseMatch struct {
	Method                string                                  `json:"method"`
	Confidence            string                                  `json:"confidence"`
	VerifiedCorporateLink bool                                    `json:"verified_corporate_link"`
	BasisCounts           OwnersPortfolioResponseMatchBasisCounts `json:"basis_counts"`
	BasisCountsScope      string                                  `json:"basis_counts_scope"`
	Note                  string                                  `json:"note"`
}

// OwnersPortfolioResponseMatchBasisCounts is generated from the OpenAPI spec.
type OwnersPortfolioResponseMatchBasisCounts struct {
	ExactSpelling float64 `json:"exact_spelling"`
	Variant       float64 `json:"variant"`
	TickerCurated float64 `json:"ticker_curated"`
}

// UnmarshalJSON decodes OwnersPortfolioResponseMatchBasisCounts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersPortfolioResponseMatchBasisCounts) UnmarshalJSON(data []byte) error {
	type plain OwnersPortfolioResponseMatchBasisCounts
	aux := struct {
		*plain
		ExactSpelling lenientNumber[float64] `json:"exact_spelling"`
		Variant       lenientNumber[float64] `json:"variant"`
		TickerCurated lenientNumber[float64] `json:"ticker_curated"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ExactSpelling.assign(&r.ExactSpelling)
	aux.Variant.assign(&r.Variant)
	aux.TickerCurated.assign(&r.TickerCurated)
	return softTypeError(err)
}

// OwnersPortfolioResponseEntityType is generated from the OpenAPI spec. It is a string; the
// OwnersPortfolioResponseEntityType* constants list the documented values.
type OwnersPortfolioResponseEntityType = string

// Documented values of OwnersPortfolioResponseEntityType.
const (
	OwnersPortfolioResponseEntityTypeIndividual  OwnersPortfolioResponseEntityType = "individual"
	OwnersPortfolioResponseEntityTypeCorporation OwnersPortfolioResponseEntityType = "corporation"
	OwnersPortfolioResponseEntityTypeLLC         OwnersPortfolioResponseEntityType = "llc"
	OwnersPortfolioResponseEntityTypeTrust       OwnersPortfolioResponseEntityType = "trust"
	OwnersPortfolioResponseEntityTypeGovernment  OwnersPortfolioResponseEntityType = "government"
	OwnersPortfolioResponseEntityTypeOther       OwnersPortfolioResponseEntityType = "other"
)

// Report: Owner intelligence report (paid, priced per resolution; account required) — with a
// free preview
//
// The Machine Storefront's owner intelligence report: every parcel nationwide whose
// OWNER-OF-RECORD name matches one owner NAME's realistic spellings (or a public company TICKER's
// hand-curated entity list), deduped nationally and priced PER RESOLUTION.
//
// ACCOUNT REQUIRED (every mode, including the free preview): an API key (`Authorization: Bearer
// pz_...`) or a signed-in session. Payment alone (an x402 `X-PAYMENT` header or a prepaid
// `X-CREDIT-TOKEN`) is not an account. An anonymous call is refused with HTTP 401, `code:
// "account_required"`, `reason: "owner_lookup_requires_account"`, before any read, payment
// verification or credit debit.
//
// MATCHING is spelling matching, NOT a verified corporate or beneficial-ownership link: different
// owners can share a spelling and some of one owner's spellings can be missed. Every property
// carries `match_basis` (exact_spelling = the record's owner name is the requested name after
// case/punctuation/whitespace normalization; variant = matched only through an expanded
// suffix/hyphen/N.A. spelling; ticker_curated = on the hand-curated ticker list) and
// `match_confidence: "candidate"`. The report carries: owner {query_name, ticker, entity_type,
// resolved_variants, match}, summary {count, total_assessed_value, total_acreage, states,
// by_state[]}, properties[] (each with match_basis, canonical_id, address, valuation,
// lot/building, last sale, property_type), and provenance {as_of, source}.
//
// PRICE: per resolution = clamp($1.75 x P(portfolio size) x V(portfolio value), $0.25, $20). P is
// log-scaled on the distinct parcel count the match uncovers (the dominant axis); V is log-scaled
// on the summed assessed value; a missing value floors V rather than raising the price. A match
// that uncovers zero parcels is returned free and never charged. The exact price is advertised in
// the 402's accepts[0].maxAmountRequired (USDC atomic units, 6 decimals).
//
// FREE PREVIEW: add preview=true for the summary (count, total value, states spanned, entity
// type), the exact price, and up to three MASKED sample properties (APN truncated to state:county,
// house number stripped). No payment; an account is still required.
//
// PAID ACCESS (preview omitted), for an authenticated account, requires ONE of: (a) x402
// pay-per-call — send a base64 signed x402 PaymentPayload in the `X-PAYMENT` header alongside
// your credentials; on a successful build the full portfolio is returned and the on-chain
// settlement receipt is in the `X-PAYMENT-RESPONSE` response header. (b) A prepaid
// `X-CREDIT-TOKEN` balance. (c) A genuine PAID PropRaven subscription entitlement (owner reports
// are included). Being merely authenticated is NOT sufficient. (d) Anything else -> HTTP 402 whose
// `accepts` array carries the exact payment requirements for this resolution.
//
// HTTP: GET /api/v1/owners/{name}/report
func (s *OwnersService) Report(ctx context.Context, name string, params *OwnersReportParams, opts ...RequestOption) (*OwnersReportResponse, error) {
	var out OwnersReportResponse
	if err := s.client.do(ctx, buildOwnersReportRequest(name, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersReportRequest(name string, params *OwnersReportParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/"+pathParam(name)+"/report")
	if params != nil {
		addQuery(req.query, "ticker", params.Ticker)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "offset", params.Offset)
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// OwnersReportParams holds the query, header and JSON-body parameters of [OwnersService.Report].
// Pass nil when you need none.
type OwnersReportParams struct {
	// Optional public stock ticker; when it is on PropRaven's hand-curated list, that entity list is
	// matched instead of expanding `name` (match_basis: ticker_curated — curated, not verified).
	Ticker *string `query:"ticker" json:"-"`

	// Optional 2-letter USPS code (or 2-digit FIPS) to scope the portfolio to one state.
	State *string `query:"state" json:"-"`

	// Properties per page in the paid report (the summary always covers the FULL portfolio).
	Limit *int64 `query:"limit" json:"-"`

	// Pagination offset into the property list.
	Offset *int64 `query:"offset" json:"-"`

	// FREE try-before-buy: the summary, the exact price and three masked sample properties. No
	// payment; an account (API key or session) is still required.
	Preview *bool `query:"preview" json:"-"`

	// Base64-encoded x402 PaymentPayload (EIP-3009 signed). Present it to pay per call for the full
	// report.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// OwnersReportResponse: Owner intelligence report (paid, priced per resolution; account required)
// — with a free preview
type OwnersReportResponse struct {
	Owner   OwnersReportResponseOwner    `json:"owner"`
	Filter  OwnersReportResponseFilter   `json:"filter"`
	Summary OwnersReportResponseSummary  `json:"summary"`
	Quote   OwnersReportResponseQuote    `json:"quote"`
	Preview bool                         `json:"preview"`
	Sample  []OwnersReportResponseSample `json:"sample"`
	Note    string                       `json:"note"`
}

// OwnersReportResponseOwner is generated from the OpenAPI spec.
type OwnersReportResponseOwner struct {
	QueryName        string                         `json:"query_name"`
	Ticker           *string                        `json:"ticker,omitempty"`
	EntityType       string                         `json:"entity_type"`
	ResolvedVariants []*string                      `json:"resolved_variants"`
	Match            OwnersReportResponseOwnerMatch `json:"match"`
}

// OwnersReportResponseOwnerMatch is generated from the OpenAPI spec.
type OwnersReportResponseOwnerMatch struct {
	Method                string                                    `json:"method"`
	Confidence            string                                    `json:"confidence"`
	VerifiedCorporateLink bool                                      `json:"verified_corporate_link"`
	BasisCounts           OwnersReportResponseOwnerMatchBasisCounts `json:"basis_counts"`
	BasisCountsScope      string                                    `json:"basis_counts_scope"`
	Note                  string                                    `json:"note"`
}

// OwnersReportResponseOwnerMatchBasisCounts is generated from the OpenAPI spec.
type OwnersReportResponseOwnerMatchBasisCounts struct {
	ExactSpelling float64 `json:"exact_spelling"`
	Variant       float64 `json:"variant"`
	TickerCurated float64 `json:"ticker_curated"`
}

// UnmarshalJSON decodes OwnersReportResponseOwnerMatchBasisCounts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseOwnerMatchBasisCounts) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseOwnerMatchBasisCounts
	aux := struct {
		*plain
		ExactSpelling lenientNumber[float64] `json:"exact_spelling"`
		Variant       lenientNumber[float64] `json:"variant"`
		TickerCurated lenientNumber[float64] `json:"ticker_curated"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ExactSpelling.assign(&r.ExactSpelling)
	aux.Variant.assign(&r.Variant)
	aux.TickerCurated.assign(&r.TickerCurated)
	return softTypeError(err)
}

// OwnersReportResponseFilter is generated from the OpenAPI spec.
type OwnersReportResponseFilter struct {
	StateFIPS *string `json:"state_fips,omitempty"`
}

// OwnersReportResponseSummary is generated from the OpenAPI spec.
type OwnersReportResponseSummary struct {
	Count              int64                                `json:"count"`
	TotalAssessedValue int64                                `json:"total_assessed_value"`
	TotalAcreage       float64                              `json:"total_acreage"`
	States             []*float64                           `json:"states"`
	ByState            []OwnersReportResponseSummaryByState `json:"by_state"`
}

// UnmarshalJSON decodes OwnersReportResponseSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseSummary) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseSummary
	aux := struct {
		*plain
		Count              lenientNumber[int64]   `json:"count"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		TotalAcreage       lenientNumber[float64] `json:"total_acreage"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assign(&r.Count)
	aux.TotalAssessedValue.assign(&r.TotalAssessedValue)
	aux.TotalAcreage.assign(&r.TotalAcreage)
	return softTypeError(err)
}

// OwnersReportResponseSummaryByState is generated from the OpenAPI spec.
type OwnersReportResponseSummaryByState struct {
	State        *string  `json:"state,omitempty"`
	StateFIPS    string   `json:"state_fips"`
	Abbr         *string  `json:"abbr,omitempty"`
	Count        *int64   `json:"count,omitempty"`
	TotalValue   *int64   `json:"total_value,omitempty"`
	TotalAcreage *float64 `json:"total_acreage,omitempty"`
}

// UnmarshalJSON decodes OwnersReportResponseSummaryByState, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseSummaryByState) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseSummaryByState
	aux := struct {
		*plain
		Count        lenientNumber[int64]   `json:"count"`
		TotalValue   lenientNumber[int64]   `json:"total_value"`
		TotalAcreage lenientNumber[float64] `json:"total_acreage"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assignPtr(&r.Count)
	aux.TotalValue.assignPtr(&r.TotalValue)
	aux.TotalAcreage.assignPtr(&r.TotalAcreage)
	return softTypeError(err)
}

// OwnersReportResponseQuote is generated from the OpenAPI spec.
type OwnersReportResponseQuote struct {
	Price           OwnersReportResponseQuotePrice     `json:"price"`
	PriceUsd        float64                            `json:"price_usd"`
	PriceAtomicUsdc string                             `json:"price_atomic_usdc"`
	Asset           string                             `json:"asset"`
	ParcelCount     int64                              `json:"parcel_count"`
	Band            string                             `json:"band"`
	BandLabel       string                             `json:"band_label"`
	Breakdown       OwnersReportResponseQuoteBreakdown `json:"breakdown"`
	Pay             []*string                          `json:"pay"`
	Note            string                             `json:"note"`
}

// UnmarshalJSON decodes OwnersReportResponseQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseQuote) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseQuote
	aux := struct {
		*plain
		PriceUsd    lenientNumber[float64] `json:"price_usd"`
		ParcelCount lenientNumber[int64]   `json:"parcel_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PriceUsd.assign(&r.PriceUsd)
	aux.ParcelCount.assign(&r.ParcelCount)
	return softTypeError(err)
}

// OwnersReportResponseQuotePrice is generated from the OpenAPI spec.
type OwnersReportResponseQuotePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// OwnersReportResponseQuoteBreakdown is generated from the OpenAPI spec.
type OwnersReportResponseQuoteBreakdown struct {
	Base     float64 `json:"base"`
	P        float64 `json:"P"`
	V        float64 `json:"V"`
	PriceRaw float64 `json:"price_raw"`
	Capped   bool    `json:"capped"`
}

// UnmarshalJSON decodes OwnersReportResponseQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseQuoteBreakdown
	aux := struct {
		*plain
		Base     lenientNumber[float64] `json:"base"`
		P        lenientNumber[float64] `json:"P"`
		V        lenientNumber[float64] `json:"V"`
		PriceRaw lenientNumber[float64] `json:"price_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Base.assign(&r.Base)
	aux.P.assign(&r.P)
	aux.V.assign(&r.V)
	aux.PriceRaw.assign(&r.PriceRaw)
	return softTypeError(err)
}

// OwnersReportResponseSample is generated from the OpenAPI spec.
type OwnersReportResponseSample struct {
	CanonicalID        string   `json:"canonical_id"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	PropertyType       *string  `json:"property_type,omitempty"`
	MatchBasis         string   `json:"match_basis"`
	MatchConfidence    *string  `json:"match_confidence,omitempty"`
}

// UnmarshalJSON decodes OwnersReportResponseSample, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersReportResponseSample) UnmarshalJSON(data []byte) error {
	type plain OwnersReportResponseSample
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	return softTypeError(err)
}

// Transactions: Recorded deed transactions for an owner
//
// Returns up to 100 most-recent deed events where the named owner is either grantor or grantee.
// Useful for building an owner's transaction timeline across their portfolio.
//
// HTTP: GET /api/v1/owners/{name}/transactions
func (s *OwnersService) Transactions(ctx context.Context, name string, params *OwnersTransactionsParams, opts ...RequestOption) (*OwnersTransactionsResponse, error) {
	var out OwnersTransactionsResponse
	if err := s.client.do(ctx, buildOwnersTransactionsRequest(name, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersTransactionsRequest(name string, params *OwnersTransactionsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/"+pathParam(name)+"/transactions")
	return req
}

// OwnersTransactionsParams holds the query, header and JSON-body parameters of
// [OwnersService.Transactions]. Pass nil when you need none.
type OwnersTransactionsParams struct {
}

// OwnersTransactionsResponse: Recorded deed transactions for an owner
type OwnersTransactionsResponse struct {
	Data  []OwnerTransaction `json:"data"`
	Count int64              `json:"count"`
}

// UnmarshalJSON decodes OwnersTransactionsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnersTransactionsResponse) UnmarshalJSON(data []byte) error {
	type plain OwnersTransactionsResponse
	aux := struct {
		*plain
		Count lenientNumber[int64] `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assign(&r.Count)
	return softTypeError(err)
}

// Card: Owner card -- the owner of record and their mailing contact (account required)
//
// The owner of record and how to reach them BY MAIL, for one parcel (`parcel_id`) or one owner
// name (`name`, the exact owner-of-record spelling). Returns the co-owners the same assessor
// record names (`co_owners`, `co_owner_status`: listed / none_listed / not_checked) and the best
// mailing address -- read from ONE address column family of the record, with its ZIP, flagged
// mail_ready / po_box / equals_situs, graded A-D with its source and as-of date (`as_of_basis`:
// roll_year = the record's tax year; release_vintage = the PropRaven parcel release that served a
// record with no roll year, never presented as the county's date; recording_date = a deed; null
// when there is no date) -- plus, for an owner name, the owner's distinct mailing addresses across
// up to 50 of their parcels (`parcels_citing`). `phones` are OWNER phones only: a building permit
// filed in the current owner's era published the phone for the owner role (the field name, the
// publisher, or a role column on the permit says so) and names the current owner; otherwise
// `phone_status` says none_published, or not_checked when the lookup could not run.
// `people_on_permits` separately lists applicant and contractor phones from the parcel's permits
// (role, name, permit, date, and `era`: filed in the current owner's era or before it); they are
// never the owner's phone, and a phone whose role the source does not state is never served. Every
// phone is E.164 with its extension split off and the published value kept verbatim (`phone_raw`).
// Grade-D (contradictory) addresses are hidden unless include_low_confidence=true.
//
// ACCOUNT REQUIRED: an API key, an MCP OAuth token or a signed-in session. Anonymous callers --
// including x402 / credit-token wallets -- get HTTP 401 `code: "account_required"` before any
// read.
//
// LOOKUP CAP: on an account without a paid plan each card counts as ONE lookup against the monthly
// lookup cap, spent before any read (a card that is not served is refunded). At the cap: HTTP 402
// `code: "lookup_cap_reached"` with used / limit / plan and the upgrade link. Each served card is
// recorded in PropRaven's people-data access log.
//
// HTTP: GET /api/v1/owners/card
func (s *OwnersService) Card(ctx context.Context, params *OwnersCardParams, opts ...RequestOption) (*OwnersCardResponse, error) {
	var out OwnersCardResponse
	if err := s.client.do(ctx, buildOwnersCardRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildOwnersCardRequest(params *OwnersCardParams) *apiRequest {
	req := newRequest("GET", "/api/v1/owners/card")
	if params != nil {
		addQuery(req.query, "parcel_id", params.ParcelID)
		addQuery(req.query, "name", params.Name)
		addQuery(req.query, "include_low_confidence", params.IncludeLowConfidence)
	}
	return req
}

// OwnersCardParams holds the query, header and JSON-body parameters of [OwnersService.Card]. Pass
// nil when you need none.
type OwnersCardParams struct {
	// A canonical id (state:county:parcel) or a parcel UUID. Pass this OR `name`.
	ParcelID *string `query:"parcel_id" json:"-"`

	// An owner of record, exact spelling (e.g. from a parcel lookup). Pass this OR `parcel_id`.
	Name *string `query:"name" json:"-"`

	// true -> also return grade-D (contradictory) addresses.
	IncludeLowConfidence *bool `query:"include_low_confidence" json:"-"`
}

// OwnersCardResponse: Owner card -- the owner of record and their mailing contact (account
// required)
type OwnersCardResponse = OwnerCard
