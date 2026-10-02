// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// ParcelsService groups the parcels operations. Use it as client.Parcels.
type ParcelsService struct {
	client *Client
}

// Get: Get parcel by ID
//
// Retrieve a single parcel by its composite ID (county_fips:parcel_id).
//
// HTTP: GET /api/v1/parcels/{id}
func (s *ParcelsService) Get(ctx context.Context, id string, params *ParcelsGetParams, opts ...RequestOption) (*ParcelsGetResponse, error) {
	var out ParcelsGetResponse
	if err := s.client.do(ctx, buildParcelsGetRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsGetRequest(id string, params *ParcelsGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id))
	return req
}

// ParcelsGetParams holds the query, header and JSON-body parameters of [ParcelsService.Get]. Pass
// nil when you need none.
type ParcelsGetParams struct {
}

// ParcelsGetResponse: Get parcel by ID
type ParcelsGetResponse = Parcel

// Owner: Get parcel owner details and portfolio
//
// Retrieve the owner of a parcel along with their portfolio summary and list of properties.
//
// HTTP: GET /api/v1/parcels/{id}/owner
func (s *ParcelsService) Owner(ctx context.Context, id string, params *ParcelsOwnerParams, opts ...RequestOption) (*ParcelsOwnerResponse, error) {
	var out ParcelsOwnerResponse
	if err := s.client.do(ctx, buildParcelsOwnerRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsOwnerRequest(id string, params *ParcelsOwnerParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/owner")
	return req
}

// ParcelsOwnerParams holds the query, header and JSON-body parameters of [ParcelsService.Owner].
// Pass nil when you need none.
type ParcelsOwnerParams struct {
}

// ParcelsOwnerResponse: Get parcel owner details and portfolio
type ParcelsOwnerResponse struct {
	Owner   ParcelsOwnerResponseOwner   `json:"owner"`
	Contact ParcelsOwnerResponseContact `json:"contact"`

	// Observed values include: CORP.
	EntityType       *string                               `json:"entity_type,omitempty"`
	Portfolio        ParcelsOwnerResponsePortfolio         `json:"portfolio"`
	Properties       []ParcelsOwnerResponseProperties      `json:"properties"`
	OwnerName        *string                               `json:"owner_name,omitempty"`
	MailingAddress   *string                               `json:"mailing_address,omitempty"`
	PortfolioSummary *ParcelsOwnerResponsePortfolioSummary `json:"portfolio_summary,omitempty"`
}

// ParcelsOwnerResponseOwner is generated from the OpenAPI spec.
type ParcelsOwnerResponseOwner struct {
	OwnerName    *string `json:"owner_name,omitempty"`
	OwnerAddress *string `json:"owner_address,omitempty"`
	OwnerCity    *string `json:"owner_city,omitempty"`
	OwnerState   *string `json:"owner_state,omitempty"`
}

// ParcelsOwnerResponseContact is generated from the OpenAPI spec.
type ParcelsOwnerResponseContact struct {
	OwnerName             *string                                        `json:"owner_name,omitempty"`
	OwnerNameStatus       *string                                        `json:"owner_name_status,omitempty"`
	OwnerRoles            []json.RawMessage                              `json:"owner_roles"`
	OwnerNameProvenance   ParcelsOwnerResponseContactOwnerNameProvenance `json:"owner_name_provenance"`
	CoOwners              []json.RawMessage                              `json:"co_owners"`
	CoOwnerStatus         *string                                        `json:"co_owner_status,omitempty"`
	Mailing               ParcelsOwnerResponseContactMailing             `json:"mailing"`
	MailingAlternates     []json.RawMessage                              `json:"mailing_alternates"`
	Entity                ParcelsOwnerResponseContactEntity              `json:"entity"`
	Phones                []json.RawMessage                              `json:"phones"`
	PhoneStatus           *string                                        `json:"phone_status,omitempty"`
	Emails                []json.RawMessage                              `json:"emails"`
	NonePublished         []*string                                      `json:"none_published"`
	HiddenLowConfidence   *float64                                       `json:"hidden_low_confidence,omitempty"`
	PeopleOnPermits       []json.RawMessage                              `json:"people_on_permits"`
	PeopleOnPermitsStatus *string                                        `json:"people_on_permits_status,omitempty"`
	OtherAddresses        []ParcelsOwnerResponseContactOtherAddresses    `json:"other_addresses"`
	OtherAddressesStatus  *string                                        `json:"other_addresses_status,omitempty"`
	OtherAddressesScope   ParcelsOwnerResponseContactOtherAddressesScope `json:"other_addresses_scope"`
}

// UnmarshalJSON decodes ParcelsOwnerResponseContact, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponseContact) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponseContact
	aux := struct {
		*plain
		HiddenLowConfidence lenientNumber[float64] `json:"hidden_low_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.HiddenLowConfidence.assignPtr(&r.HiddenLowConfidence)
	return softTypeError(err)
}

// ParcelsOwnerResponseContactOwnerNameProvenance is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOwnerNameProvenance struct {
	Source    ParcelsOwnerResponseContactOwnerNameProvenanceSource `json:"source"`
	AsOf      *string                                              `json:"as_of,omitempty"`
	AsOfBasis string                                               `json:"as_of_basis"`
	Grade     *string                                              `json:"grade,omitempty"`
}

// ParcelsOwnerResponseContactOwnerNameProvenanceSource is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOwnerNameProvenanceSource struct {
	Authority *string `json:"authority,omitempty"`
	Dataset   *string `json:"dataset,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// ParcelsOwnerResponseContactMailing is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactMailing struct {
	Basis              *string                                  `json:"basis,omitempty"`
	Line1              *string                                  `json:"line1,omitempty"`
	City               *string                                  `json:"city,omitempty"`
	State              *string                                  `json:"state,omitempty"`
	Zip5               *string                                  `json:"zip5,omitempty"`
	Zip4               *string                                  `json:"zip4,omitempty"`
	MailReady          *bool                                    `json:"mail_ready,omitempty"`
	PoBox              *bool                                    `json:"po_box,omitempty"`
	EqualsSitus        *bool                                    `json:"equals_situs,omitempty"`
	ZipConflict        *bool                                    `json:"zip_conflict,omitempty"`
	ParcelsCiting      *float64                                 `json:"parcels_citing,omitempty"`
	ParcelsCitingBasis string                                   `json:"parcels_citing_basis"`
	Label              *string                                  `json:"label,omitempty"`
	Source             ParcelsOwnerResponseContactMailingSource `json:"source"`
	AsOf               *string                                  `json:"as_of,omitempty"`
	AsOfBasis          string                                   `json:"as_of_basis"`
	Grade              *string                                  `json:"grade,omitempty"`
}

// UnmarshalJSON decodes ParcelsOwnerResponseContactMailing, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponseContactMailing) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponseContactMailing
	aux := struct {
		*plain
		ParcelsCiting lenientNumber[float64] `json:"parcels_citing"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelsCiting.assignPtr(&r.ParcelsCiting)
	return softTypeError(err)
}

// ParcelsOwnerResponseContactMailingSource is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactMailingSource struct {
	Authority *string `json:"authority,omitempty"`
	Dataset   *string `json:"dataset,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// ParcelsOwnerResponseContactEntity is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactEntity struct {
	EntityType       *string                                 `json:"entity_type,omitempty"`
	Status           *string                                 `json:"status,omitempty"`
	StateOfFormation *string                                 `json:"state_of_formation,omitempty"`
	FormationDate    *string                                 `json:"formation_date,omitempty"`
	RegisteredAgent  *string                                 `json:"registered_agent,omitempty"`
	Officers         []json.RawMessage                       `json:"officers"`
	Source           ParcelsOwnerResponseContactEntitySource `json:"source"`
	AsOf             *string                                 `json:"as_of,omitempty"`
	AsOfBasis        *string                                 `json:"as_of_basis,omitempty"`
	Grade            *string                                 `json:"grade,omitempty"`
}

// ParcelsOwnerResponseContactEntitySource is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactEntitySource struct {
	Authority *string `json:"authority,omitempty"`
	Dataset   *string `json:"dataset,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// ParcelsOwnerResponseContactOtherAddresses is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOtherAddresses struct {
	ParcelID          string                                             `json:"parcel_id"`
	State             *string                                            `json:"state,omitempty"`
	County            *string                                            `json:"county,omitempty"`
	Kind              *string                                            `json:"kind,omitempty"`
	Address           ParcelsOwnerResponseContactOtherAddressesAddress   `json:"address"`
	SameAs            *string                                            `json:"same_as,omitempty"`
	Grade             *string                                            `json:"grade,omitempty"`
	Link              *string                                            `json:"link,omitempty"`
	Basis             *string                                            `json:"basis,omitempty"`
	LabelNote         *string                                            `json:"label_note,omitempty"`
	Evidence          []json.RawMessage                                  `json:"evidence"`
	MailMerge         *bool                                              `json:"mail_merge,omitempty"`
	OwnerNameOnRecord *string                                            `json:"owner_name_on_record,omitempty"`
	OwnerRoles        []json.RawMessage                                  `json:"owner_roles"`
	CitedBy           []ParcelsOwnerResponseContactOtherAddressesCitedBy `json:"cited_by"`
	ParcelsCiting     *float64                                           `json:"parcels_citing,omitempty"`
	Source            ParcelsOwnerResponseContactOtherAddressesSource    `json:"source"`
	AsOf              *string                                            `json:"as_of,omitempty"`
	AsOfBasis         string                                             `json:"as_of_basis"`
}

// UnmarshalJSON decodes ParcelsOwnerResponseContactOtherAddresses, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponseContactOtherAddresses) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponseContactOtherAddresses
	aux := struct {
		*plain
		ParcelsCiting lenientNumber[float64] `json:"parcels_citing"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelsCiting.assignPtr(&r.ParcelsCiting)
	return softTypeError(err)
}

// ParcelsOwnerResponseContactOtherAddressesAddress is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOtherAddressesAddress struct {
	Line1     *string `json:"line1,omitempty"`
	City      *string `json:"city,omitempty"`
	State     *string `json:"state,omitempty"`
	Zip5      *string `json:"zip5,omitempty"`
	Zip4      *string `json:"zip4,omitempty"`
	Label     *string `json:"label,omitempty"`
	MailReady *bool   `json:"mail_ready,omitempty"`
	PoBox     *bool   `json:"po_box,omitempty"`
}

// ParcelsOwnerResponseContactOtherAddressesCitedBy is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOtherAddressesCitedBy struct {
	ParcelID string  `json:"parcel_id"`
	State    *string `json:"state,omitempty"`
	County   *string `json:"county,omitempty"`
}

// ParcelsOwnerResponseContactOtherAddressesSource is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOtherAddressesSource struct {
	Authority *string `json:"authority,omitempty"`
	Dataset   *string `json:"dataset,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// ParcelsOwnerResponseContactOtherAddressesScope is generated from the OpenAPI spec.
type ParcelsOwnerResponseContactOtherAddressesScope struct {
	NameBasis              string                                                            `json:"name_basis"`
	Spellings              *float64                                                          `json:"spellings,omitempty"`
	ParcelsRead            *float64                                                          `json:"parcels_read,omitempty"`
	Capped                 *bool                                                             `json:"capped,omitempty"`
	CopiesSkipped          *float64                                                          `json:"copies_skipped,omitempty"`
	ParcelsLinked          *float64                                                          `json:"parcels_linked,omitempty"`
	PossibleFound          *float64                                                          `json:"possible_found,omitempty"`
	PossibleListed         *float64                                                          `json:"possible_listed,omitempty"`
	PossibleCap            *float64                                                          `json:"possible_cap,omitempty"`
	PossibleWithheldReason *string                                                           `json:"possible_withheld_reason,omitempty"`
	ListedCapped           *bool                                                             `json:"listed_capped,omitempty"`
	EvidenceProviders      []ParcelsOwnerResponseContactOtherAddressesScopeEvidenceProviders `json:"evidence_providers"`
	CoOwnerLink            *string                                                           `json:"co_owner_link,omitempty"`
}

// UnmarshalJSON decodes ParcelsOwnerResponseContactOtherAddressesScope, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponseContactOtherAddressesScope) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponseContactOtherAddressesScope
	aux := struct {
		*plain
		Spellings      lenientNumber[float64] `json:"spellings"`
		ParcelsRead    lenientNumber[float64] `json:"parcels_read"`
		CopiesSkipped  lenientNumber[float64] `json:"copies_skipped"`
		ParcelsLinked  lenientNumber[float64] `json:"parcels_linked"`
		PossibleFound  lenientNumber[float64] `json:"possible_found"`
		PossibleListed lenientNumber[float64] `json:"possible_listed"`
		PossibleCap    lenientNumber[float64] `json:"possible_cap"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Spellings.assignPtr(&r.Spellings)
	aux.ParcelsRead.assignPtr(&r.ParcelsRead)
	aux.CopiesSkipped.assignPtr(&r.CopiesSkipped)
	aux.ParcelsLinked.assignPtr(&r.ParcelsLinked)
	aux.PossibleFound.assignPtr(&r.PossibleFound)
	aux.PossibleListed.assignPtr(&r.PossibleListed)
	aux.PossibleCap.assignPtr(&r.PossibleCap)
	return softTypeError(err)
}

// ParcelsOwnerResponseContactOtherAddressesScopeEvidenceProviders is generated from the OpenAPI
// spec.
type ParcelsOwnerResponseContactOtherAddressesScopeEvidenceProviders struct {
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

// ParcelsOwnerResponsePortfolio is generated from the OpenAPI spec.
type ParcelsOwnerResponsePortfolio struct {
	PropertyCount      *int64   `json:"property_count,omitempty"`
	StateCount         *int64   `json:"state_count,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	TotalAcreage       *float64 `json:"total_acreage,omitempty"`
	StatesList         *string  `json:"states_list,omitempty"`
	CountyCount        *int64   `json:"county_count,omitempty"`
	PortfolioRank      *int64   `json:"portfolio_rank,omitempty"`
}

// UnmarshalJSON decodes ParcelsOwnerResponsePortfolio, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponsePortfolio) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponsePortfolio
	aux := struct {
		*plain
		PropertyCount      lenientNumber[int64]   `json:"property_count"`
		StateCount         lenientNumber[int64]   `json:"state_count"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		TotalAcreage       lenientNumber[float64] `json:"total_acreage"`
		CountyCount        lenientNumber[int64]   `json:"county_count"`
		PortfolioRank      lenientNumber[int64]   `json:"portfolio_rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PropertyCount.assignPtr(&r.PropertyCount)
	aux.StateCount.assignPtr(&r.StateCount)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.TotalAcreage.assignPtr(&r.TotalAcreage)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.PortfolioRank.assignPtr(&r.PortfolioRank)
	return softTypeError(err)
}

// ParcelsOwnerResponseProperties is generated from the OpenAPI spec.
type ParcelsOwnerResponseProperties struct {
	ParcelID           string   `json:"parcel_id"`
	CountyFIPS         string   `json:"county_fips"`
	StateFIPS          string   `json:"state_fips"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	LastSalePrice      *float64 `json:"last_sale_price,omitempty"`
	LastSaleDate       *string  `json:"last_sale_date,omitempty"`
	PropertyType       *string  `json:"property_type,omitempty"`
	BuildingSqft       *float64 `json:"building_sqft,omitempty"`
	LotSizeAcres       *float64 `json:"lot_size_acres,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
}

// UnmarshalJSON decodes ParcelsOwnerResponseProperties, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponseProperties) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponseProperties
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		LastSalePrice      lenientNumber[float64] `json:"last_sale_price"`
		BuildingSqft       lenientNumber[float64] `json:"building_sqft"`
		LotSizeAcres       lenientNumber[float64] `json:"lot_size_acres"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	aux.BuildingSqft.assignPtr(&r.BuildingSqft)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	return softTypeError(err)
}

// ParcelsOwnerResponsePortfolioSummary is generated from the OpenAPI spec.
type ParcelsOwnerResponsePortfolioSummary struct {
	PropertyCount      *int64   `json:"property_count,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	States             []string `json:"states,omitempty"`
}

// UnmarshalJSON decodes ParcelsOwnerResponsePortfolioSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOwnerResponsePortfolioSummary) UnmarshalJSON(data []byte) error {
	type plain ParcelsOwnerResponsePortfolioSummary
	aux := struct {
		*plain
		PropertyCount      lenientNumber[int64]   `json:"property_count"`
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PropertyCount.assignPtr(&r.PropertyCount)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	return softTypeError(err)
}

// Permits: Get parcel permits
//
// Retrieve building and construction permits associated with a parcel.
//
// HTTP: GET /api/v1/parcels/{id}/permits
func (s *ParcelsService) Permits(ctx context.Context, id string, params *ParcelsPermitsParams, opts ...RequestOption) (*ParcelsPermitsResponse, error) {
	var out ParcelsPermitsResponse
	if err := s.client.do(ctx, buildParcelsPermitsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsPermitsRequest(id string, params *ParcelsPermitsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/permits")
	if params != nil {
		addQuery(req.query, "shape", params.Shape)
	}
	return req
}

// ParcelsPermitsParams holds the query, header and JSON-body parameters of
// [ParcelsService.Permits]. Pass nil when you need none.
type ParcelsPermitsParams struct {
	// Body shape. Omit for the default bare array of up to 100 permit rows (newest first); `envelope`
	// returns `{ data, permit_count, permit_count_basis, truncated, row_cap }`, where `permit_count`
	// is the parcel's true count when `permit_count_basis` is `exact` and the size of the capped
	// window (a floor) when `capped`, and `truncated` is true when the parcel has more permits than
	// `data` carries.
	Shape *ParcelsPermitsParamsShape `query:"shape" json:"-"`
}

// ParcelsPermitsParamsShape is generated from the OpenAPI spec. It is a string; the
// ParcelsPermitsParamsShape* constants list the documented values.
type ParcelsPermitsParamsShape = string

// Documented values of ParcelsPermitsParamsShape.
const (
	ParcelsPermitsParamsShapeEnvelope ParcelsPermitsParamsShape = "envelope"
)

// ParcelsPermitsResponse: Get parcel permits
//
// It is one of: []Permit (Array), ParcelsPermitsResponseObject (Object). Exactly one variant field
// is set after decoding; Raw keeps the original JSON.
type ParcelsPermitsResponse struct {
	// Array is set when the value is a JSON array.
	Array []Permit
	// Object is set when the value is a JSON object.
	Object *ParcelsPermitsResponseObject
	// Raw is the undecoded JSON value.
	Raw json.RawMessage
}

// UnmarshalJSON decodes whichever variant matches the JSON kind.
func (u *ParcelsPermitsResponse) UnmarshalJSON(b []byte) error {
	*u = ParcelsPermitsResponse{Raw: append(json.RawMessage(nil), b...)}
	switch jsonKindOf(b) {
	case 'a':
		return json.Unmarshal(b, &u.Array)
	case 'o':
		return json.Unmarshal(b, &u.Object)
	}
	return nil
}

// MarshalJSON encodes the set variant (or Raw).
func (u ParcelsPermitsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case u.Array != nil:
		return json.Marshal(u.Array)
	case u.Object != nil:
		return json.Marshal(u.Object)
	case u.Raw != nil:
		return u.Raw, nil
	}
	return []byte("null"), nil
}

// ParcelsPermitsResponseObject is generated from the OpenAPI spec.
type ParcelsPermitsResponseObject struct {
	Data []Permit `json:"data"`

	// The parcel's permit count (exact) or the capped window size (a floor) — see
	// permit_count_basis.
	PermitCount      int64                                        `json:"permit_count"`
	PermitCountBasis ParcelsPermitsResponseObjectPermitCountBasis `json:"permit_count_basis"`

	// True when the parcel has more permits than `data` carries.
	Truncated bool `json:"truncated"`

	// The row cap `data` is bounded by (100).
	RowCap int64 `json:"row_cap"`
}

// UnmarshalJSON decodes ParcelsPermitsResponseObject, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsPermitsResponseObject) UnmarshalJSON(data []byte) error {
	type plain ParcelsPermitsResponseObject
	aux := struct {
		*plain
		PermitCount lenientNumber[int64] `json:"permit_count"`
		RowCap      lenientNumber[int64] `json:"row_cap"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PermitCount.assign(&r.PermitCount)
	aux.RowCap.assign(&r.RowCap)
	return softTypeError(err)
}

// ParcelsPermitsResponseObjectPermitCountBasis is generated from the OpenAPI spec. It is a string;
// the ParcelsPermitsResponseObjectPermitCountBasis* constants list the documented values.
type ParcelsPermitsResponseObjectPermitCountBasis = string

// Documented values of ParcelsPermitsResponseObjectPermitCountBasis.
const (
	ParcelsPermitsResponseObjectPermitCountBasisExact  ParcelsPermitsResponseObjectPermitCountBasis = "exact"
	ParcelsPermitsResponseObjectPermitCountBasisCapped ParcelsPermitsResponseObjectPermitCountBasis = "capped"
)

// Deeds: Get parcel deed history
//
// Retrieve deed transactions and transfer history for a parcel.
//
// HTTP: GET /api/v1/parcels/{id}/deeds
func (s *ParcelsService) Deeds(ctx context.Context, id string, params *ParcelsDeedsParams, opts ...RequestOption) (*ParcelsDeedsResponse, error) {
	var out ParcelsDeedsResponse
	if err := s.client.do(ctx, buildParcelsDeedsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsDeedsRequest(id string, params *ParcelsDeedsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/deeds")
	if params != nil {
		addQuery(req.query, "shape", params.Shape)
	}
	return req
}

// ParcelsDeedsParams holds the query, header and JSON-body parameters of [ParcelsService.Deeds].
// Pass nil when you need none.
type ParcelsDeedsParams struct {
	// Body shape. Omit for the default bare array of deed rows; `envelope` returns the typed envelope
	// with a `status` header and the frozen `known_deed_count`.
	Shape *ParcelsDeedsParamsShape `query:"shape" json:"-"`
}

// ParcelsDeedsParamsShape is generated from the OpenAPI spec. It is a string; the
// ParcelsDeedsParamsShape* constants list the documented values.
type ParcelsDeedsParamsShape = string

// Documented values of ParcelsDeedsParamsShape.
const (
	ParcelsDeedsParamsShapeEnvelope ParcelsDeedsParamsShape = "envelope"
)

// ParcelsDeedsResponse: Get parcel deed history
//
// It is one of: []Deed (Array), ParcelsDeedsResponseObject (Object). Exactly one variant field is
// set after decoding; Raw keeps the original JSON.
type ParcelsDeedsResponse struct {
	// Array is set when the value is a JSON array.
	Array []Deed
	// Object is set when the value is a JSON object.
	Object *ParcelsDeedsResponseObject
	// Raw is the undecoded JSON value.
	Raw json.RawMessage
}

// UnmarshalJSON decodes whichever variant matches the JSON kind.
func (u *ParcelsDeedsResponse) UnmarshalJSON(b []byte) error {
	*u = ParcelsDeedsResponse{Raw: append(json.RawMessage(nil), b...)}
	switch jsonKindOf(b) {
	case 'a':
		return json.Unmarshal(b, &u.Array)
	case 'o':
		return json.Unmarshal(b, &u.Object)
	}
	return nil
}

// MarshalJSON encodes the set variant (or Raw).
func (u ParcelsDeedsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case u.Array != nil:
		return json.Marshal(u.Array)
	case u.Object != nil:
		return json.Marshal(u.Object)
	case u.Raw != nil:
		return u.Raw, nil
	}
	return []byte("null"), nil
}

// ParcelsDeedsResponseObject: Returned only when `shape=envelope`.
type ParcelsDeedsResponseObject struct {
	Data   []Deed                           `json:"data"`
	Status ParcelsDeedsResponseObjectStatus `json:"status"`

	// parcels_serving.deed_count when the parcel carries one; NOT a count of the returned rows.
	KnownDeedCount *int64 `json:"known_deed_count,omitempty"`

	// Where known_deed_count comes from: the frozen 2026-04 weld, or the identity gate's reason when
	// the parcel_id collides across counties — parcel_id_collides_bare_id_joined (rollup withheld,
	// count null), parcel_id_collides_unmeasured (served with the doubt visible),
	// collision_check_unavailable (withheld under DQ_IDENTITY_GATE_FAIL_CLOSED=1).
	KnownDeedCountBasis *string `json:"known_deed_count_basis,omitempty"`
	Note                *string `json:"note,omitempty"`

	// The identity probe's verdict on the rollup (present on rollup-derived envelopes; absent on the
	// county-scoped per_deed branch).
	JoinKeyBasis *ParcelsDeedsResponseObjectJoinKeyBasis `json:"join_key_basis,omitempty"`

	// Present when the rollup was withheld on the deeds-relation probe: the parcel_id is unique in
	// parcels_serving (join_key_basis state_county_parcel_id) but parcel_deeds carries rows for it
	// under another (state, county), so the frozen rollup was welded on the bare id from another
	// parcel's deeds.
	RollupProbe *ParcelsDeedsResponseObjectRollupProbe `json:"rollup_probe,omitempty"`
}

// UnmarshalJSON decodes ParcelsDeedsResponseObject, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsDeedsResponseObject) UnmarshalJSON(data []byte) error {
	type plain ParcelsDeedsResponseObject
	aux := struct {
		*plain
		KnownDeedCount lenientNumber[int64] `json:"known_deed_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.KnownDeedCount.assignPtr(&r.KnownDeedCount)
	return softTypeError(err)
}

// ParcelsDeedsResponseObjectStatus is generated from the OpenAPI spec. It is a string; the
// ParcelsDeedsResponseObjectStatus* constants list the documented values.
type ParcelsDeedsResponseObjectStatus = string

// Documented values of ParcelsDeedsResponseObjectStatus.
const (
	ParcelsDeedsResponseObjectStatusPerDeed      ParcelsDeedsResponseObjectStatus = "per_deed"
	ParcelsDeedsResponseObjectStatusRollupEvents ParcelsDeedsResponseObjectStatus = "rollup_events"
	ParcelsDeedsResponseObjectStatusSummaryOnly  ParcelsDeedsResponseObjectStatus = "summary_only"
	ParcelsDeedsResponseObjectStatusNone         ParcelsDeedsResponseObjectStatus = "none"
)

// ParcelsDeedsResponseObjectJoinKeyBasis: The identity probe's verdict on the rollup (present on
// rollup-derived envelopes; absent on the county-scoped per_deed branch).
//
// It is a string; the ParcelsDeedsResponseObjectJoinKeyBasis* constants list the documented
// values.
type ParcelsDeedsResponseObjectJoinKeyBasis = string

// Documented values of ParcelsDeedsResponseObjectJoinKeyBasis.
const (
	ParcelsDeedsResponseObjectJoinKeyBasisStateCountyParcelID ParcelsDeedsResponseObjectJoinKeyBasis = "state_county_parcel_id"
	ParcelsDeedsResponseObjectJoinKeyBasisParcelIDCollides    ParcelsDeedsResponseObjectJoinKeyBasis = "parcel_id_collides"
	ParcelsDeedsResponseObjectJoinKeyBasisUnchecked           ParcelsDeedsResponseObjectJoinKeyBasis = "unchecked"
)

// ParcelsDeedsResponseObjectRollupProbe: Present when the rollup was withheld on the
// deeds-relation probe: the parcel_id is unique in parcels_serving (join_key_basis
// state_county_parcel_id) but parcel_deeds carries rows for it under another (state, county), so
// the frozen rollup was welded on the bare id from another parcel's deeds.
//
// It is a string; the ParcelsDeedsResponseObjectRollupProbe* constants list the documented values.
type ParcelsDeedsResponseObjectRollupProbe = string

// Documented values of ParcelsDeedsResponseObjectRollupProbe.
const (
	ParcelsDeedsResponseObjectRollupProbeDeedsRelation ParcelsDeedsResponseObjectRollupProbe = "deeds_relation"
)

// Risks: Get parcel risk assessment
//
// Retrieve flood, wildfire, air quality, and crime risk data for a parcel.
//
// HTTP: GET /api/v1/parcels/{id}/risks
func (s *ParcelsService) Risks(ctx context.Context, id string, params *ParcelsRisksParams, opts ...RequestOption) (*ParcelsRisksResponse, error) {
	var out ParcelsRisksResponse
	if err := s.client.do(ctx, buildParcelsRisksRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsRisksRequest(id string, params *ParcelsRisksParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/risks")
	return req
}

// ParcelsRisksParams holds the query, header and JSON-body parameters of [ParcelsService.Risks].
// Pass nil when you need none.
type ParcelsRisksParams struct {
}

// ParcelsRisksResponse: Get parcel risk assessment
type ParcelsRisksResponse = RiskAssessment

// Geojson: Parcel polygons as GeoJSON for a bounding box
//
// Returns parcel polygons inside a bounding box as a GeoJSON FeatureCollection. Only served at
// zoom ≥ 14 to limit data volume — coarser bbox returns an empty collection. Each feature's
// properties include parcel_id, county_fips, owner_name, assessed value, and basic attributes for
// rendering popups.
//
// HTTP: GET /api/v1/parcels/geojson
func (s *ParcelsService) Geojson(ctx context.Context, params *ParcelsGeojsonParams, opts ...RequestOption) (*ParcelsGeojsonResponse, error) {
	var out ParcelsGeojsonResponse
	if err := s.client.do(ctx, buildParcelsGeojsonRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsGeojsonRequest(params *ParcelsGeojsonParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/geojson")
	if params != nil {
		addQuery(req.query, "bbox", params.Bbox)
		addQuery(req.query, "zoom", params.Zoom)
	}
	return req
}

// ParcelsGeojsonParams holds the query, header and JSON-body parameters of
// [ParcelsService.Geojson]. Pass nil when you need none.
type ParcelsGeojsonParams struct {
	// Bounding box `west,south,east,north`.
	//
	// Required.
	Bbox *string `query:"bbox" json:"-"`

	// Map zoom level. Below 14 returns an empty collection.
	//
	// Required.
	Zoom *int64 `query:"zoom" json:"-"`
}

// ParcelsGeojsonResponse: Parcel polygons as GeoJSON for a bounding box
type ParcelsGeojsonResponse = ParcelGeoJSON

// Report: Parcel dossier (paid, provenance-first)
//
// The Machine Storefront's per-parcel dossier: a single provenance-first JSON payload carrying
// every POPULATED field for the parcel as `{name, value}` (by default: no `sections`/`fields`
// params delivers everything the quote was priced on), plus the real deeds / comparable-sales /
// permits sub-tables. The per-field receipts (source, as_of, confidence, coverage tier) are
// opt-in: `?include_provenance=true` adds a `provenance` map keyed by field name. `?sections=`
// (identity, valuation, owner, hazard, permits, deeds, market, demographics, `core`, `all`) and
// `?fields=` narrow the delivered field list; anything narrowed out is NAMED in
// `meta.projection.omitted_fields`, and the price never changes with the projection. A ~50 KB soft
// cap applies to the `fields` portion only. The GeoJSON boundary is held out as a
// separately-priced add-on the base payload omits.
//
// PRICE (value-tiered, per parcel): price = clamp($5 x V(asset value) x R(data richness) x
// F(freshness), $2, $20). The exact amount for a given parcel is quoted, before payment, by GET
// /api/v1/storefront/availability?parcel_id=... and is what the 402 advertises in
// accepts[0].maxAmountRequired (USDC atomic units, 6 decimals).
//
// ACCESS requires ONE of: (a) x402 pay-per-call -- send a base64 signed x402 PaymentPayload in the
// `X-PAYMENT` header; on a successful build the dossier is returned and the on-chain settlement
// receipt is in the `X-PAYMENT-RESPONSE` response header. No account is needed for the parcel
// record, but PEOPLE DATA (owner names, owner mailing addresses, entity principals, deed and sale
// party names and addresses) is delivered to accounts only: a wallet-only or credit-token buyer
// receives the dossier with those fields set to null and a top-level `people_fields` marker (see
// PeopleFieldsWithheld), and the price is computed on exactly that body, so it never counts a
// field the buyer does not receive. Send your API key with the payment to receive them. (b) A
// genuine PAID PropRaven subscription entitlement -- the dossier is served on the subscription
// invoice. Being merely authenticated is NOT sufficient: a free-tier key, or a self-service
// first-party key with no paid plan, receives a 402. (c) Anything else -> HTTP 402 whose `accepts`
// array carries the exact x402 payment requirements for this parcel.
//
// HTTP: GET /api/v1/parcels/{id}/report
func (s *ParcelsService) Report(ctx context.Context, id string, params *ParcelsReportParams, opts ...RequestOption) (*ParcelsReportResponse, error) {
	var out ParcelsReportResponse
	if err := s.client.do(ctx, buildParcelsReportRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsReportRequest(id string, params *ParcelsReportParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/report")
	if params != nil {
		addQuery(req.query, "county_fips", params.CountyFIPS)
		addQuery(req.query, "sections", params.Sections)
		addQuery(req.query, "fields", params.Fields)
		addQuery(req.query, "include_provenance", params.IncludeProvenance)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// ParcelsReportParams holds the query, header and JSON-body parameters of [ParcelsService.Report].
// Pass nil when you need none.
type ParcelsReportParams struct {
	// 5-digit county FIPS. Strongly recommended when passing a county-local id.
	CountyFIPS *string `query:"county_fips" json:"-"`

	// Comma-separated report sections to deliver in `fields`: identity, valuation, owner, hazard,
	// permits, deeds, market, demographics, `core` (the first four) or `all`. Omit (with no `fields`)
	// for `all` — every populated field the quote was priced on. The compact core is delivered only
	// on an explicit `core`.
	Sections *string `query:"sections" json:"-"`

	// Comma-separated parcels_serving column names to deliver on top of `sections`.
	Fields *string `query:"fields" json:"-"`

	// `true` attaches the per-field receipts map (`provenance`, keyed by field name: source, as_of,
	// confidence, coverage, tier). Off by default — the receipts are epoch catalog metadata and
	// roughly triple the payload.
	IncludeProvenance *bool `query:"include_provenance" json:"-"`

	// x402 payment: a base64-encoded signed x402 PaymentPayload (EIP-3009 transferWithAuthorization
	// over USDC on Base). Present it to pay per call with no API key; the signed amount must equal
	// this parcel's quoted maxAmountRequired (see the 402 body or
	// /storefront/availability?parcel_id=). Omit it to be served only if your key holds a paid
	// subscription entitlement; otherwise you receive a 402 carrying the payment requirements.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// ParcelsReportResponse: Parcel dossier (paid, provenance-first)
type ParcelsReportResponse = ParcelDossier

// CompPack: Comp pack (paid, priced per pack) — with a FREE preview
//
// The Machine Storefront's comp pack: a subject valuation INDICATED BY comparable sales, wrapped
// with the comps that prove it. Answers the underwriting question "what is this worth, and which
// sales prove it?".
//
// The comps are the SAME precomputed comparable_sales the free GET /api/v1/parcels/{id}/comps
// route serves; the subject's valuation fields come from parcels_serving. The indicated value is
// derived transparently (median comp $/sqft x subject sqft when both are known, else the median
// comp sale price) with an interquartile range, and is reconstructable from the comps in the same
// payload -- never a black-box AVM.
//
// PRICE: per pack = clamp($2 x V(subject value) x Q(comp support), $1, $20). Q ramps on the comp
// count with a small bonus for high median similarity. A subject with ZERO precomputed comps has
// no evidence to support a number and is returned free, never charged. The exact price is
// advertised in the 402's accepts[0].maxAmountRequired (USDC atomic units, 6 decimals).
//
// FREE PREVIEW: add preview=true for the subject summary, the comp count + median similarity, the
// exact price, and up to three MASKED comps (parcel/APN withheld, sale price rounded, date to the
// year). The precise indicated value and the unmasked comps are the paid product.
//
// PAID ACCESS (preview omitted) requires ONE of: (a) x402 pay-per-call via a base64 signed x402
// PaymentPayload in the `X-PAYMENT` header (receipt in `X-PAYMENT-RESPONSE`); (b) a genuine PAID
// PropRaven subscription entitlement; (c) anything else -> HTTP 402 whose `accepts` carries the
// exact requirements.
//
// HTTP: GET /api/v1/parcels/{id}/comp-pack
func (s *ParcelsService) CompPack(ctx context.Context, id string, params *ParcelsCompPackParams, opts ...RequestOption) (*ParcelsCompPackResponse, error) {
	var out ParcelsCompPackResponse
	if err := s.client.do(ctx, buildParcelsCompPackRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsCompPackRequest(id string, params *ParcelsCompPackParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/comp-pack")
	if params != nil {
		addQuery(req.query, "n", params.N)
		addQuery(req.query, "radius", params.Radius)
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// ParcelsCompPackParams holds the query, header and JSON-body parameters of
// [ParcelsService.CompPack]. Pass nil when you need none.
type ParcelsCompPackParams struct {
	// How many comps back the pack.
	N *int64 `query:"n" json:"-"`

	// Optional post-filter: keep only precomputed comps within this many miles.
	Radius *float64 `query:"radius" json:"-"`

	// FREE try-before-buy: subject summary, comp count, exact price and three masked comps. No
	// payment.
	Preview *bool `query:"preview" json:"-"`

	// Base64-encoded x402 PaymentPayload (EIP-3009 signed). Present it to pay per call for the full
	// pack.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// ParcelsCompPackResponse: Comp pack (paid, priced per pack) — with a FREE preview
type ParcelsCompPackResponse struct {
	Subject     ParcelsCompPackResponseSubject  `json:"subject"`
	CompCount   int64                           `json:"comp_count"`
	RadiusMiles *float64                        `json:"radius_miles,omitempty"`
	Quote       ParcelsCompPackResponseQuote    `json:"quote"`
	Preview     bool                            `json:"preview"`
	Sample      []ParcelsCompPackResponseSample `json:"sample"`
	Note        string                          `json:"note"`
}

// UnmarshalJSON decodes ParcelsCompPackResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompPackResponse) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompPackResponse
	aux := struct {
		*plain
		CompCount   lenientNumber[int64]   `json:"comp_count"`
		RadiusMiles lenientNumber[float64] `json:"radius_miles"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CompCount.assign(&r.CompCount)
	aux.RadiusMiles.assignPtr(&r.RadiusMiles)
	return softTypeError(err)
}

// ParcelsCompPackResponseSubject is generated from the OpenAPI spec.
type ParcelsCompPackResponseSubject struct {
	CanonicalID        string   `json:"canonical_id"`
	ParcelID           string   `json:"parcel_id"`
	StateFIPS          string   `json:"state_fips"`
	CountyFIPS         string   `json:"county_fips"`
	TotalAssessedValue int64    `json:"total_assessed_value"`
	MarketValue        float64  `json:"market_value"`
	BuildingSqft       *float64 `json:"building_sqft,omitempty"`
	PropertyType       string   `json:"property_type"`
	LastSalePrice      *float64 `json:"last_sale_price,omitempty"`
	LastSaleDate       string   `json:"last_sale_date"`
}

// UnmarshalJSON decodes ParcelsCompPackResponseSubject, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompPackResponseSubject) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompPackResponseSubject
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		MarketValue        lenientNumber[float64] `json:"market_value"`
		BuildingSqft       lenientNumber[float64] `json:"building_sqft"`
		LastSalePrice      lenientNumber[float64] `json:"last_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assign(&r.TotalAssessedValue)
	aux.MarketValue.assign(&r.MarketValue)
	aux.BuildingSqft.assignPtr(&r.BuildingSqft)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	return softTypeError(err)
}

// ParcelsCompPackResponseQuote is generated from the OpenAPI spec.
type ParcelsCompPackResponseQuote struct {
	Price           ParcelsCompPackResponseQuotePrice     `json:"price"`
	PriceUsd        float64                               `json:"price_usd"`
	PriceAtomicUsdc string                                `json:"price_atomic_usdc"`
	Asset           string                                `json:"asset"`
	CompCount       int64                                 `json:"comp_count"`
	Breakdown       ParcelsCompPackResponseQuoteBreakdown `json:"breakdown"`
	Pay             []*string                             `json:"pay"`
	Note            string                                `json:"note"`
}

// UnmarshalJSON decodes ParcelsCompPackResponseQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompPackResponseQuote) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompPackResponseQuote
	aux := struct {
		*plain
		PriceUsd  lenientNumber[float64] `json:"price_usd"`
		CompCount lenientNumber[int64]   `json:"comp_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PriceUsd.assign(&r.PriceUsd)
	aux.CompCount.assign(&r.CompCount)
	return softTypeError(err)
}

// ParcelsCompPackResponseQuotePrice is generated from the OpenAPI spec.
type ParcelsCompPackResponseQuotePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// ParcelsCompPackResponseQuoteBreakdown is generated from the OpenAPI spec.
type ParcelsCompPackResponseQuoteBreakdown struct {
	Base     float64 `json:"base"`
	V        float64 `json:"V"`
	Q        float64 `json:"Q"`
	PriceRaw float64 `json:"price_raw"`
	Capped   bool    `json:"capped"`
}

// UnmarshalJSON decodes ParcelsCompPackResponseQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompPackResponseQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompPackResponseQuoteBreakdown
	aux := struct {
		*plain
		Base     lenientNumber[float64] `json:"base"`
		V        lenientNumber[float64] `json:"V"`
		Q        lenientNumber[float64] `json:"Q"`
		PriceRaw lenientNumber[float64] `json:"price_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Base.assign(&r.Base)
	aux.V.assign(&r.V)
	aux.Q.assign(&r.Q)
	aux.PriceRaw.assign(&r.PriceRaw)
	return softTypeError(err)
}

// ParcelsCompPackResponseSample is generated from the OpenAPI spec.
type ParcelsCompPackResponseSample struct {
	CompParcelID        *string  `json:"comp_parcel_id,omitempty"`
	CompAPN             *string  `json:"comp_apn,omitempty"`
	CompAddress         *string  `json:"comp_address,omitempty"`
	CompSalePriceApprox *float64 `json:"comp_sale_price_approx,omitempty"`
	CompSaleYear        *int64   `json:"comp_sale_year,omitempty"`
	CompSqft            *float64 `json:"comp_sqft,omitempty"`
	CompYearBuilt       *int64   `json:"comp_year_built,omitempty"`
	SimilarityScore     *float64 `json:"similarity_score,omitempty"`
	DistanceMiles       *float64 `json:"distance_miles,omitempty"`
	Rank                *int64   `json:"rank,omitempty"`
}

// UnmarshalJSON decodes ParcelsCompPackResponseSample, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompPackResponseSample) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompPackResponseSample
	aux := struct {
		*plain
		CompSalePriceApprox lenientNumber[float64] `json:"comp_sale_price_approx"`
		CompSaleYear        lenientNumber[int64]   `json:"comp_sale_year"`
		CompSqft            lenientNumber[float64] `json:"comp_sqft"`
		CompYearBuilt       lenientNumber[int64]   `json:"comp_year_built"`
		SimilarityScore     lenientNumber[float64] `json:"similarity_score"`
		DistanceMiles       lenientNumber[float64] `json:"distance_miles"`
		Rank                lenientNumber[int64]   `json:"rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CompSalePriceApprox.assignPtr(&r.CompSalePriceApprox)
	aux.CompSaleYear.assignPtr(&r.CompSaleYear)
	aux.CompSqft.assignPtr(&r.CompSqft)
	aux.CompYearBuilt.assignPtr(&r.CompYearBuilt)
	aux.SimilarityScore.assignPtr(&r.SimilarityScore)
	aux.DistanceMiles.assignPtr(&r.DistanceMiles)
	aux.Rank.assignPtr(&r.Rank)
	return softTypeError(err)
}

// RiskScore: Risk score (paid, priced per assessment) — with a FREE preview
//
// The Machine Storefront's risk score: a multi-hazard risk assessment for one parcel, anchored on
// FEMA's National Risk Index COMPOSITE (nri_risk_score 0-100 + rating) with the flood / seismic /
// windstorm / wildfire / air-quality / crime breakdown that supports it. Answers "what could go
// wrong with this asset?". The headline is FEMA's own methodology, not an invented weighting.
//
// Reuses the SAME panel the free GET /api/v1/parcels/{id}/risks route serves (getParcelRisksData)
// plus the NRI composite + flood detail from parcels_serving.
//
// PRICE: per assessment = clamp($0.60 x V(asset value) x C(hazard coverage), $0.20, $20). C ramps
// on how many independent hazard layers resolved. A parcel with ZERO layers is returned free,
// never charged. The cheapest paid SKU. The exact price is in the 402's
// accepts[0].maxAmountRequired (USDC atomic units).
//
// FREE PREVIEW: add preview=true for the subject, WHICH hazard layers resolved, and the exact
// price. The precise NRI score and the hazard breakdown are the paid product.
//
// PAID ACCESS: (a) x402 via a base64 signed PaymentPayload in `X-PAYMENT` (receipt in
// `X-PAYMENT-RESPONSE`); (b) a genuine PAID subscription entitlement; (c) else HTTP 402 with the
// exact requirements.
//
// HTTP: GET /api/v1/parcels/{id}/risk-score
func (s *ParcelsService) RiskScore(ctx context.Context, id string, params *ParcelsRiskScoreParams, opts ...RequestOption) (*ParcelsRiskScoreResponse, error) {
	var out ParcelsRiskScoreResponse
	if err := s.client.do(ctx, buildParcelsRiskScoreRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsRiskScoreRequest(id string, params *ParcelsRiskScoreParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/risk-score")
	if params != nil {
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// ParcelsRiskScoreParams holds the query, header and JSON-body parameters of
// [ParcelsService.RiskScore]. Pass nil when you need none.
type ParcelsRiskScoreParams struct {
	// FREE try-before-buy: subject, resolved hazard layers, and the exact price. No payment.
	Preview *bool `query:"preview" json:"-"`

	// Base64-encoded x402 PaymentPayload (EIP-3009 signed). Present it to pay per call.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// ParcelsRiskScoreResponse: Risk score (paid, priced per assessment) — with a FREE preview
type ParcelsRiskScoreResponse struct {
	Subject          ParcelsRiskScoreResponseSubject `json:"subject"`
	HazardLayers     []*string                       `json:"hazard_layers"`
	HazardLayerCount int64                           `json:"hazard_layer_count"`
	Quote            ParcelsRiskScoreResponseQuote   `json:"quote"`
	Preview          bool                            `json:"preview"`
	Note             string                          `json:"note"`
}

// UnmarshalJSON decodes ParcelsRiskScoreResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsRiskScoreResponse) UnmarshalJSON(data []byte) error {
	type plain ParcelsRiskScoreResponse
	aux := struct {
		*plain
		HazardLayerCount lenientNumber[int64] `json:"hazard_layer_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.HazardLayerCount.assign(&r.HazardLayerCount)
	return softTypeError(err)
}

// ParcelsRiskScoreResponseSubject is generated from the OpenAPI spec.
type ParcelsRiskScoreResponseSubject struct {
	CanonicalID        string `json:"canonical_id"`
	ParcelID           string `json:"parcel_id"`
	StateFIPS          string `json:"state_fips"`
	CountyFIPS         string `json:"county_fips"`
	TotalAssessedValue int64  `json:"total_assessed_value"`
}

// UnmarshalJSON decodes ParcelsRiskScoreResponseSubject, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsRiskScoreResponseSubject) UnmarshalJSON(data []byte) error {
	type plain ParcelsRiskScoreResponseSubject
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[int64] `json:"total_assessed_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assign(&r.TotalAssessedValue)
	return softTypeError(err)
}

// ParcelsRiskScoreResponseQuote is generated from the OpenAPI spec.
type ParcelsRiskScoreResponseQuote struct {
	Price           ParcelsRiskScoreResponseQuotePrice     `json:"price"`
	PriceUsd        float64                                `json:"price_usd"`
	PriceAtomicUsdc string                                 `json:"price_atomic_usdc"`
	Asset           string                                 `json:"asset"`
	HazardLayers    float64                                `json:"hazard_layers"`
	Breakdown       ParcelsRiskScoreResponseQuoteBreakdown `json:"breakdown"`
	Pay             []*string                              `json:"pay"`
	Note            string                                 `json:"note"`
}

// UnmarshalJSON decodes ParcelsRiskScoreResponseQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsRiskScoreResponseQuote) UnmarshalJSON(data []byte) error {
	type plain ParcelsRiskScoreResponseQuote
	aux := struct {
		*plain
		PriceUsd     lenientNumber[float64] `json:"price_usd"`
		HazardLayers lenientNumber[float64] `json:"hazard_layers"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PriceUsd.assign(&r.PriceUsd)
	aux.HazardLayers.assign(&r.HazardLayers)
	return softTypeError(err)
}

// ParcelsRiskScoreResponseQuotePrice is generated from the OpenAPI spec.
type ParcelsRiskScoreResponseQuotePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// ParcelsRiskScoreResponseQuoteBreakdown is generated from the OpenAPI spec.
type ParcelsRiskScoreResponseQuoteBreakdown struct {
	Base     float64 `json:"base"`
	V        float64 `json:"V"`
	C        float64 `json:"C"`
	PriceRaw float64 `json:"price_raw"`
	Capped   bool    `json:"capped"`
}

// UnmarshalJSON decodes ParcelsRiskScoreResponseQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsRiskScoreResponseQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain ParcelsRiskScoreResponseQuoteBreakdown
	aux := struct {
		*plain
		Base     lenientNumber[float64] `json:"base"`
		V        lenientNumber[float64] `json:"V"`
		C        lenientNumber[float64] `json:"C"`
		PriceRaw lenientNumber[float64] `json:"price_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Base.assign(&r.Base)
	aux.V.assign(&r.V)
	aux.C.assign(&r.C)
	aux.PriceRaw.assign(&r.PriceRaw)
	return softTypeError(err)
}

// TrafficHistory: Nearest traffic station + AADT history
//
// Finds the AADT (annual average daily traffic) station nearest to the given coordinates and
// returns its historical time series plus 3/5/7-year CAGRs. Useful for retail / CRE site
// selection. Search radius ~2 miles; returns empty data if no station is in range.
//
// HTTP: GET /api/v1/parcels/{id}/traffic-history
func (s *ParcelsService) TrafficHistory(ctx context.Context, id string, params *ParcelsTrafficHistoryParams, opts ...RequestOption) (*ParcelsTrafficHistoryResponse, error) {
	var out ParcelsTrafficHistoryResponse
	if err := s.client.do(ctx, buildParcelsTrafficHistoryRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsTrafficHistoryRequest(id string, params *ParcelsTrafficHistoryParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/traffic-history")
	if params != nil {
		addQuery(req.query, "lat", params.Lat)
		addQuery(req.query, "lng", params.Lng)
	}
	return req
}

// ParcelsTrafficHistoryParams holds the query, header and JSON-body parameters of
// [ParcelsService.TrafficHistory]. Pass nil when you need none.
type ParcelsTrafficHistoryParams struct {
	// Latitude.
	//
	// Required.
	Lat *float64 `query:"lat" json:"-"`

	// Longitude.
	//
	// Required.
	Lng *float64 `query:"lng" json:"-"`
}

// ParcelsTrafficHistoryResponse: Nearest traffic station + AADT history
type ParcelsTrafficHistoryResponse = TrafficStationHistory

// Batch: Fetch up to 100 parcels by (state, county, parcel) tuple
//
// Card-projection rows for up to 100 unique `(state_fips, county_fips, parcel_id)` tuples.
// `county_fips` is the 3-digit within-state code (a 5-digit value is refused with a 400 naming the
// fix). Tuples with no match are listed in `missing`.
//
// HTTP: POST /api/v1/parcels/batch
func (s *ParcelsService) Batch(ctx context.Context, params *ParcelsBatchParams, opts ...RequestOption) (*ParcelsBatchResponse, error) {
	var out ParcelsBatchResponse
	if err := s.client.do(ctx, buildParcelsBatchRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsBatchRequest(params *ParcelsBatchParams) *apiRequest {
	req := newRequest("POST", "/api/v1/parcels/batch")
	if params != nil {
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// ParcelsBatchParams holds the query, header and JSON-body parameters of [ParcelsService.Batch].
// Pass nil when you need none.
type ParcelsBatchParams struct {
	// Required (JSON body).
	Tuples []ParcelsBatchParamsTuples `json:"tuples"`
}

// ParcelsBatchParamsTuples is generated from the OpenAPI spec.
type ParcelsBatchParamsTuples struct {
	StateFIPS  string `json:"state_fips"`
	CountyFIPS string `json:"county_fips"`
	ParcelID   string `json:"parcel_id"`
}

// ParcelsBatchResponse: Fetch up to 100 parcels by (state, county, parcel) tuple
type ParcelsBatchResponse struct {
	Rows    []ParcelsBatchResponseRows `json:"rows"`
	Missing []json.RawMessage          `json:"missing"`
}

// ParcelsBatchResponseRows is generated from the OpenAPI spec.
type ParcelsBatchResponseRows struct {
	ID                       string   `json:"id"`
	CountyFIPS               string   `json:"county_fips"`
	StateFIPS                string   `json:"state_fips"`
	ParcelID                 string   `json:"parcel_id"`
	Address                  *string  `json:"address,omitempty"`
	NormalizedAddress        *string  `json:"normalized_address,omitempty"`
	City                     *string  `json:"city,omitempty"`
	State                    *string  `json:"state,omitempty"`
	Zip                      *string  `json:"zip,omitempty"`
	Zip5                     *string  `json:"zip5,omitempty"`
	ZipPlus4                 *string  `json:"zip_plus4,omitempty"`
	Latitude                 *float64 `json:"latitude,omitempty"`
	Longitude                *float64 `json:"longitude,omitempty"`
	LandUseCode              *string  `json:"land_use_code,omitempty"`
	LandUseDesc              *string  `json:"land_use_desc,omitempty"`
	TotalAssessedValue       *float64 `json:"total_assessed_value,omitempty"`
	LandAssessedValue        *float64 `json:"land_assessed_value,omitempty"`
	ImprovementAssessedValue *float64 `json:"improvement_assessed_value,omitempty"`
	LastSalePrice            *float64 `json:"last_sale_price,omitempty"`
	LastSaleDate             *string  `json:"last_sale_date,omitempty"`
	MarketValue              *float64 `json:"market_value,omitempty"`
	AvmValue                 *float64 `json:"avm_value,omitempty"`
	AvmConfidence            *string  `json:"avm_confidence,omitempty"`
	AvmMethod                *string  `json:"avm_method,omitempty"`
	TaxAmount                *float64 `json:"tax_amount,omitempty"`
	TaxYear                  *int64   `json:"tax_year,omitempty"`
	DealScore                *float64 `json:"deal_score,omitempty"`
	PricePerSqft             *float64 `json:"price_per_sqft,omitempty"`
	BuildingSqft             *float64 `json:"building_sqft,omitempty"`
	YearBuilt                *int64   `json:"year_built,omitempty"`
	LotSizeAcres             *float64 `json:"lot_size_acres,omitempty"`
	LotSizeSqft              *float64 `json:"lot_size_sqft,omitempty"`
	Bedrooms                 *int64   `json:"bedrooms,omitempty"`
	Bathrooms                *float64 `json:"bathrooms,omitempty"`
	Stories                  *float64 `json:"stories,omitempty"`
	Units                    *int64   `json:"units,omitempty"`
	UnitCount                *int64   `json:"unit_count,omitempty"`
	ConstructionType         *string  `json:"construction_type,omitempty"`
	OwnerName                *string  `json:"owner_name,omitempty"`
	OwnerAddress             *string  `json:"owner_address,omitempty"`
	OwnerCity                *string  `json:"owner_city,omitempty"`
	OwnerState               *string  `json:"owner_state,omitempty"`
	OwnerZip                 *string  `json:"owner_zip,omitempty"`
	OwnershipType            *string  `json:"ownership_type,omitempty"`
	OwnerEntityType          *string  `json:"owner_entity_type,omitempty"`
	EntityType               *string  `json:"entity_type,omitempty"`
	IsEntityOwned            *bool    `json:"is_entity_owned,omitempty"`
	IsAbsentee               *bool    `json:"is_absentee,omitempty"`
	IsPeAggregator           *bool    `json:"is_pe_aggregator,omitempty"`
	OwnerOccupiedFlag        *bool    `json:"owner_occupied_flag,omitempty"`
	DataQualityScore         *float64 `json:"data_quality_score,omitempty"`
	FloodZone                *string  `json:"flood_zone,omitempty"`
	IsSfha                   *bool    `json:"is_sfha,omitempty"`
	IsOpportunityZone        *bool    `json:"is_opportunity_zone,omitempty"`
	IsJustice40              *bool    `json:"is_justice40,omitempty"`
	IsFlip                   *bool    `json:"is_flip,omitempty"`
	CrimeScore               *float64 `json:"crime_score,omitempty"`
	CrimeTier                *float64 `json:"crime_tier,omitempty"`
	DeedCount                *int64   `json:"deed_count,omitempty"`
	PermitCount              *int64   `json:"permit_count,omitempty"`
	PermitCount12mo          *int64   `json:"permit_count_12mo,omitempty"`
	Zoning                   *string  `json:"zoning,omitempty"`
	ZoningCodeRaw            *string  `json:"zoning_code_raw,omitempty"`
	CountyName               *string  `json:"county_name,omitempty"`
	PropertyType             *string  `json:"property_type,omitempty"`
}

// UnmarshalJSON decodes ParcelsBatchResponseRows, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsBatchResponseRows) UnmarshalJSON(data []byte) error {
	type plain ParcelsBatchResponseRows
	aux := struct {
		*plain
		Latitude                 lenientNumber[float64] `json:"latitude"`
		Longitude                lenientNumber[float64] `json:"longitude"`
		TotalAssessedValue       lenientNumber[float64] `json:"total_assessed_value"`
		LandAssessedValue        lenientNumber[float64] `json:"land_assessed_value"`
		ImprovementAssessedValue lenientNumber[float64] `json:"improvement_assessed_value"`
		LastSalePrice            lenientNumber[float64] `json:"last_sale_price"`
		MarketValue              lenientNumber[float64] `json:"market_value"`
		AvmValue                 lenientNumber[float64] `json:"avm_value"`
		TaxAmount                lenientNumber[float64] `json:"tax_amount"`
		TaxYear                  lenientNumber[int64]   `json:"tax_year"`
		DealScore                lenientNumber[float64] `json:"deal_score"`
		PricePerSqft             lenientNumber[float64] `json:"price_per_sqft"`
		BuildingSqft             lenientNumber[float64] `json:"building_sqft"`
		YearBuilt                lenientNumber[int64]   `json:"year_built"`
		LotSizeAcres             lenientNumber[float64] `json:"lot_size_acres"`
		LotSizeSqft              lenientNumber[float64] `json:"lot_size_sqft"`
		Bedrooms                 lenientNumber[int64]   `json:"bedrooms"`
		Bathrooms                lenientNumber[float64] `json:"bathrooms"`
		Stories                  lenientNumber[float64] `json:"stories"`
		Units                    lenientNumber[int64]   `json:"units"`
		UnitCount                lenientNumber[int64]   `json:"unit_count"`
		DataQualityScore         lenientNumber[float64] `json:"data_quality_score"`
		CrimeScore               lenientNumber[float64] `json:"crime_score"`
		CrimeTier                lenientNumber[float64] `json:"crime_tier"`
		DeedCount                lenientNumber[int64]   `json:"deed_count"`
		PermitCount              lenientNumber[int64]   `json:"permit_count"`
		PermitCount12mo          lenientNumber[int64]   `json:"permit_count_12mo"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LandAssessedValue.assignPtr(&r.LandAssessedValue)
	aux.ImprovementAssessedValue.assignPtr(&r.ImprovementAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	aux.MarketValue.assignPtr(&r.MarketValue)
	aux.AvmValue.assignPtr(&r.AvmValue)
	aux.TaxAmount.assignPtr(&r.TaxAmount)
	aux.TaxYear.assignPtr(&r.TaxYear)
	aux.DealScore.assignPtr(&r.DealScore)
	aux.PricePerSqft.assignPtr(&r.PricePerSqft)
	aux.BuildingSqft.assignPtr(&r.BuildingSqft)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.LotSizeSqft.assignPtr(&r.LotSizeSqft)
	aux.Bedrooms.assignPtr(&r.Bedrooms)
	aux.Bathrooms.assignPtr(&r.Bathrooms)
	aux.Stories.assignPtr(&r.Stories)
	aux.Units.assignPtr(&r.Units)
	aux.UnitCount.assignPtr(&r.UnitCount)
	aux.DataQualityScore.assignPtr(&r.DataQualityScore)
	aux.CrimeScore.assignPtr(&r.CrimeScore)
	aux.CrimeTier.assignPtr(&r.CrimeTier)
	aux.DeedCount.assignPtr(&r.DeedCount)
	aux.PermitCount.assignPtr(&r.PermitCount)
	aux.PermitCount12mo.assignPtr(&r.PermitCount12mo)
	return softTypeError(err)
}

// Comps: Comparable sales for a parcel
//
// Precomputed comparable sales near the parcel (`tier` names the comp source).
//
// HTTP: GET /api/v1/parcels/{id}/comps
func (s *ParcelsService) Comps(ctx context.Context, id string, params *ParcelsCompsParams, opts ...RequestOption) (*ParcelsCompsResponse, error) {
	var out ParcelsCompsResponse
	if err := s.client.do(ctx, buildParcelsCompsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsCompsRequest(id string, params *ParcelsCompsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/comps")
	if params != nil {
		addQuery(req.query, "n", params.N)
		addQuery(req.query, "radius", params.Radius)
	}
	return req
}

// ParcelsCompsParams holds the query, header and JSON-body parameters of [ParcelsService.Comps].
// Pass nil when you need none.
type ParcelsCompsParams struct {
	// Number of comps (1-25).
	N *int64 `query:"n" json:"-"`

	// Search radius in miles.
	Radius *float64 `query:"radius" json:"-"`
}

// ParcelsCompsResponse: Comparable sales for a parcel
type ParcelsCompsResponse struct {
	Subject        ParcelsCompsResponseSubject        `json:"subject"`
	Tier           string                             `json:"tier"`
	RadiusMiles    *float64                           `json:"radius_miles,omitempty"`
	Count          int64                              `json:"count"`
	Comps          []ParcelsCompsResponseComps        `json:"comps"`
	ProvenanceGate ParcelsCompsResponseProvenanceGate `json:"provenance_gate"`
}

// UnmarshalJSON decodes ParcelsCompsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompsResponse) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompsResponse
	aux := struct {
		*plain
		RadiusMiles lenientNumber[float64] `json:"radius_miles"`
		Count       lenientNumber[int64]   `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RadiusMiles.assignPtr(&r.RadiusMiles)
	aux.Count.assign(&r.Count)
	return softTypeError(err)
}

// ParcelsCompsResponseSubject is generated from the OpenAPI spec.
type ParcelsCompsResponseSubject struct {
	CanonicalID string `json:"canonical_id"`
	ParcelID    string `json:"parcel_id"`
	StateFIPS   string `json:"state_fips"`
	CountyFIPS  string `json:"county_fips"`
}

// ParcelsCompsResponseComps is generated from the OpenAPI spec.
type ParcelsCompsResponseComps struct {
	CompParcelID        *string  `json:"comp_parcel_id,omitempty"`
	CompAPN             *string  `json:"comp_apn,omitempty"`
	CompSalePrice       *float64 `json:"comp_sale_price,omitempty"`
	CompSaleDate        *string  `json:"comp_sale_date,omitempty"`
	CompAddress         *string  `json:"comp_address,omitempty"`
	CompSqft            *float64 `json:"comp_sqft,omitempty"`
	CompYearBuilt       *int64   `json:"comp_year_built,omitempty"`
	CompBeds            *float64 `json:"comp_beds,omitempty"`
	CompBaths           *float64 `json:"comp_baths,omitempty"`
	SimilarityScore     *float64 `json:"similarity_score,omitempty"`
	DistanceMiles       *float64 `json:"distance_miles,omitempty"`
	Rank                *int64   `json:"rank,omitempty"`
	SalePriceReconciled *bool    `json:"sale_price_reconciled,omitempty"`
}

// UnmarshalJSON decodes ParcelsCompsResponseComps, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompsResponseComps) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompsResponseComps
	aux := struct {
		*plain
		CompSalePrice   lenientNumber[float64] `json:"comp_sale_price"`
		CompSqft        lenientNumber[float64] `json:"comp_sqft"`
		CompYearBuilt   lenientNumber[int64]   `json:"comp_year_built"`
		CompBeds        lenientNumber[float64] `json:"comp_beds"`
		CompBaths       lenientNumber[float64] `json:"comp_baths"`
		SimilarityScore lenientNumber[float64] `json:"similarity_score"`
		DistanceMiles   lenientNumber[float64] `json:"distance_miles"`
		Rank            lenientNumber[int64]   `json:"rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CompSalePrice.assignPtr(&r.CompSalePrice)
	aux.CompSqft.assignPtr(&r.CompSqft)
	aux.CompYearBuilt.assignPtr(&r.CompYearBuilt)
	aux.CompBeds.assignPtr(&r.CompBeds)
	aux.CompBaths.assignPtr(&r.CompBaths)
	aux.SimilarityScore.assignPtr(&r.SimilarityScore)
	aux.DistanceMiles.assignPtr(&r.DistanceMiles)
	aux.Rank.assignPtr(&r.Rank)
	return softTypeError(err)
}

// ParcelsCompsResponseProvenanceGate is generated from the OpenAPI spec.
type ParcelsCompsResponseProvenanceGate struct {
	Applied         bool                                    `json:"applied"`
	Rule            string                                  `json:"rule"`
	Scope           ParcelsCompsResponseProvenanceGateScope `json:"scope"`
	CompsReconciled float64                                 `json:"comps_reconciled"`
	Reason          *string                                 `json:"reason,omitempty"`
	Note            *string                                 `json:"note,omitempty"`
}

// UnmarshalJSON decodes ParcelsCompsResponseProvenanceGate, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsCompsResponseProvenanceGate) UnmarshalJSON(data []byte) error {
	type plain ParcelsCompsResponseProvenanceGate
	aux := struct {
		*plain
		CompsReconciled lenientNumber[float64] `json:"comps_reconciled"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CompsReconciled.assign(&r.CompsReconciled)
	return softTypeError(err)
}

// ParcelsCompsResponseProvenanceGateScope is generated from the OpenAPI spec.
type ParcelsCompsResponseProvenanceGateScope struct {
	StateFIPS  string `json:"state_fips"`
	CountyFIPS string `json:"county_fips"`
}

// Occupants: Business occupants of a parcel
//
// Businesses matched to the parcel (names, brands, categories, match confidence), primary occupant
// first. At most 100 rows; `truncated` says when more exist.
//
// HTTP: GET /api/v1/parcels/{id}/occupants
func (s *ParcelsService) Occupants(ctx context.Context, id string, params *ParcelsOccupantsParams, opts ...RequestOption) (*ParcelsOccupantsResponse, error) {
	var out ParcelsOccupantsResponse
	if err := s.client.do(ctx, buildParcelsOccupantsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsOccupantsRequest(id string, params *ParcelsOccupantsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/occupants")
	return req
}

// ParcelsOccupantsParams holds the query, header and JSON-body parameters of
// [ParcelsService.Occupants]. Pass nil when you need none.
type ParcelsOccupantsParams struct {
}

// ParcelsOccupantsResponse: Business occupants of a parcel
type ParcelsOccupantsResponse struct {
	ParcelID      string                              `json:"parcel_id"`
	OccupantCount int64                               `json:"occupant_count"`
	Occupants     []ParcelsOccupantsResponseOccupants `json:"occupants"`
	Truncated     bool                                `json:"truncated"`
}

// UnmarshalJSON decodes ParcelsOccupantsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOccupantsResponse) UnmarshalJSON(data []byte) error {
	type plain ParcelsOccupantsResponse
	aux := struct {
		*plain
		OccupantCount lenientNumber[int64] `json:"occupant_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.OccupantCount.assign(&r.OccupantCount)
	return softTypeError(err)
}

// ParcelsOccupantsResponseOccupants is generated from the OpenAPI spec.
type ParcelsOccupantsResponseOccupants struct {
	OccupantID     *string  `json:"occupant_id,omitempty"`
	NameRaw        *string  `json:"name_raw,omitempty"`
	NameNorm       *string  `json:"name_norm,omitempty"`
	BrandName      *string  `json:"brand_name,omitempty"`
	BrandWikidata  *string  `json:"brand_wikidata,omitempty"`
	CategoryFsq    *string  `json:"category_fsq,omitempty"`
	Naics          *string  `json:"naics,omitempty"`
	Confidence     *float64 `json:"confidence,omitempty"`
	Status         *string  `json:"status,omitempty"`
	IsPrimary      *bool    `json:"is_primary,omitempty"`
	MatchMethod    *string  `json:"match_method,omitempty"`
	MatchDistanceM *float64 `json:"match_distance_m,omitempty"`
	SourceCount    *int64   `json:"source_count,omitempty"`
	FirstSeen      *string  `json:"first_seen,omitempty"`
	LastSeen       *string  `json:"last_seen,omitempty"`
	OccupantLat    *float64 `json:"occupant_lat,omitempty"`
	OccupantLon    *float64 `json:"occupant_lon,omitempty"`
	LuClass        *string  `json:"lu_class,omitempty"`
}

// UnmarshalJSON decodes ParcelsOccupantsResponseOccupants, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsOccupantsResponseOccupants) UnmarshalJSON(data []byte) error {
	type plain ParcelsOccupantsResponseOccupants
	aux := struct {
		*plain
		Confidence     lenientNumber[float64] `json:"confidence"`
		MatchDistanceM lenientNumber[float64] `json:"match_distance_m"`
		SourceCount    lenientNumber[int64]   `json:"source_count"`
		OccupantLat    lenientNumber[float64] `json:"occupant_lat"`
		OccupantLon    lenientNumber[float64] `json:"occupant_lon"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Confidence.assignPtr(&r.Confidence)
	aux.MatchDistanceM.assignPtr(&r.MatchDistanceM)
	aux.SourceCount.assignPtr(&r.SourceCount)
	aux.OccupantLat.assignPtr(&r.OccupantLat)
	aux.OccupantLon.assignPtr(&r.OccupantLon)
	return softTypeError(err)
}

// Violations: Code violations on a parcel
//
// Municipal code-enforcement cases matched to the parcel, with the source coverage that says
// whether an empty list means 'none recorded' or 'not collected here'.
//
// HTTP: GET /api/v1/parcels/{id}/violations
func (s *ParcelsService) Violations(ctx context.Context, id string, params *ParcelsViolationsParams, opts ...RequestOption) (*ParcelsViolationsResponse, error) {
	var out ParcelsViolationsResponse
	if err := s.client.do(ctx, buildParcelsViolationsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsViolationsRequest(id string, params *ParcelsViolationsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/violations")
	return req
}

// ParcelsViolationsParams holds the query, header and JSON-body parameters of
// [ParcelsService.Violations]. Pass nil when you need none.
type ParcelsViolationsParams struct {
}

// ParcelsViolationsResponse: Code violations on a parcel
type ParcelsViolationsResponse struct {
	Data                 []json.RawMessage                               `json:"data"`
	ViolationCount       int64                                           `json:"violation_count"`
	ViolationCountBasis  string                                          `json:"violation_count_basis"`
	Truncated            bool                                            `json:"truncated"`
	RowCap               int64                                           `json:"row_cap"`
	Coverage             string                                          `json:"coverage"`
	CoverageReason       string                                          `json:"coverage_reason"`
	Place                ParcelsViolationsResponsePlace                  `json:"place"`
	Sources              []json.RawMessage                               `json:"sources"`
	Summary              ParcelsViolationsResponseSummary                `json:"summary"`
	Withheld             ParcelsViolationsResponseWithheld               `json:"withheld"`
	CoveredJurisdictions []ParcelsViolationsResponseCoveredJurisdictions `json:"covered_jurisdictions"`
	Notes                []*string                                       `json:"notes"`
}

// UnmarshalJSON decodes ParcelsViolationsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsViolationsResponse) UnmarshalJSON(data []byte) error {
	type plain ParcelsViolationsResponse
	aux := struct {
		*plain
		ViolationCount lenientNumber[int64] `json:"violation_count"`
		RowCap         lenientNumber[int64] `json:"row_cap"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ViolationCount.assign(&r.ViolationCount)
	aux.RowCap.assign(&r.RowCap)
	return softTypeError(err)
}

// ParcelsViolationsResponsePlace is generated from the OpenAPI spec.
type ParcelsViolationsResponsePlace struct {
	PlaceGeoid *string `json:"place_geoid,omitempty"`
	PlaceName  *string `json:"place_name,omitempty"`
}

// ParcelsViolationsResponseSummary is generated from the OpenAPI spec.
type ParcelsViolationsResponseSummary struct {
	OpenCount            int64   `json:"open_count"`
	ClosedCount          int64   `json:"closed_count"`
	OtherCount           int64   `json:"other_count"`
	UnknownStatusCount   int64   `json:"unknown_status_count"`
	UnmappedStatusCount  int64   `json:"unmapped_status_count"`
	LatestIssuedDate     *string `json:"latest_issued_date,omitempty"`
	LatestStatus         *string `json:"latest_status,omitempty"`
	OldestOpenIssuedDate *string `json:"oldest_open_issued_date,omitempty"`
	HasOpenCodeViolation *bool   `json:"has_open_code_violation,omitempty"`
}

// UnmarshalJSON decodes ParcelsViolationsResponseSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsViolationsResponseSummary) UnmarshalJSON(data []byte) error {
	type plain ParcelsViolationsResponseSummary
	aux := struct {
		*plain
		OpenCount           lenientNumber[int64] `json:"open_count"`
		ClosedCount         lenientNumber[int64] `json:"closed_count"`
		OtherCount          lenientNumber[int64] `json:"other_count"`
		UnknownStatusCount  lenientNumber[int64] `json:"unknown_status_count"`
		UnmappedStatusCount lenientNumber[int64] `json:"unmapped_status_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.OpenCount.assign(&r.OpenCount)
	aux.ClosedCount.assign(&r.ClosedCount)
	aux.OtherCount.assign(&r.OtherCount)
	aux.UnknownStatusCount.assign(&r.UnknownStatusCount)
	aux.UnmappedStatusCount.assign(&r.UnmappedStatusCount)
	return softTypeError(err)
}

// ParcelsViolationsResponseWithheld is generated from the OpenAPI spec.
type ParcelsViolationsResponseWithheld struct {
	OutsidePublisherArea float64           `json:"outside_publisher_area"`
	AccountOnlyFields    []json.RawMessage `json:"account_only_fields"`
}

// UnmarshalJSON decodes ParcelsViolationsResponseWithheld, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelsViolationsResponseWithheld) UnmarshalJSON(data []byte) error {
	type plain ParcelsViolationsResponseWithheld
	aux := struct {
		*plain
		OutsidePublisherArea lenientNumber[float64] `json:"outside_publisher_area"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.OutsidePublisherArea.assign(&r.OutsidePublisherArea)
	return softTypeError(err)
}

// ParcelsViolationsResponseCoveredJurisdictions is generated from the OpenAPI spec.
type ParcelsViolationsResponseCoveredJurisdictions struct {
	SourceID   *string `json:"source_id,omitempty"`
	StateFIPS  string  `json:"state_fips"`
	CountyFIPS string  `json:"county_fips"`
	PlaceGeoid *string `json:"place_geoid,omitempty"`
	PlaceName  *string `json:"place_name,omitempty"`
	Publisher  *string `json:"publisher,omitempty"`
}

// Pois: Business parcels in a small bounding box
//
// Up to 250 parcels with a recorded business type inside `bbox`, highest assessed value first.
// Boxes larger than 0.03 square degrees return an empty list.
//
// HTTP: GET /api/v1/parcels/poi
func (s *ParcelsService) Pois(ctx context.Context, params *ParcelsPoisParams, opts ...RequestOption) (*ParcelsPoisResponse, error) {
	var out ParcelsPoisResponse
	if err := s.client.do(ctx, buildParcelsPoisRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildParcelsPoisRequest(params *ParcelsPoisParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/poi")
	if params != nil {
		addQuery(req.query, "bbox", params.Bbox)
	}
	return req
}

// ParcelsPoisParams holds the query, header and JSON-body parameters of [ParcelsService.Pois].
// Pass nil when you need none.
type ParcelsPoisParams struct {
	// `west,south,east,north` in decimal degrees.
	//
	// Required.
	Bbox *string `query:"bbox" json:"-"`
}

// ParcelsPoisResponse: Business parcels in a small bounding box
type ParcelsPoisResponse struct {
	Data []json.RawMessage `json:"data"`
}
