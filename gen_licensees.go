// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// LicenseesService groups the licensees operations. Use it as client.Licensees.
type LicenseesService struct {
	client *Client
}

// Firms: Search licensed firms in a place
//
// Firms licensed by the issuing state boards (FL DBPR and DFS, CA DRE and the Board of
// Accountancy, NY DOS, CT DCP, VA DPOR) in one city or zip, grouped by name and by the issuer's
// parent-license link, with their distinct street locations. `min_locations=3` finds firms with at
// least three licensed locations there. Requires an account; a person-shaped firm is people data
// and each response serving one is logged. 503 `licensee_layer_unavailable` until the layer is
// loaded.
//
// HTTP: GET /api/v1/licensees/firms
func (s *LicenseesService) Firms(ctx context.Context, params *LicenseesFirmsParams, opts ...RequestOption) (*LicenseesFirmsResponse, error) {
	var out LicenseesFirmsResponse
	if err := s.client.do(ctx, buildLicenseesFirmsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildLicenseesFirmsRequest(params *LicenseesFirmsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/licensees/firms")
	if params != nil {
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "profession", params.Profession)
		addQuery(req.query, "city", params.City)
		addQuery(req.query, "zip", params.Zip)
		addQuery(req.query, "status", params.Status)
		addQuery(req.query, "min_locations", params.MinLocations)
		addQuery(req.query, "limit", params.Limit)
	}
	return req
}

// LicenseesFirmsParams holds the query, header and JSON-body parameters of
// [LicenseesService.Firms]. Pass nil when you need none.
type LicenseesFirmsParams struct {
	// A state the licensee layer serves (Texas licensee files are a QA witness and never served).
	//
	// Required.
	State *LicenseesFirmsParamsState `query:"state" json:"-"`

	// License profession.
	Profession *LicenseesFirmsParamsProfession `query:"profession" json:"-"`

	// City as the licensee address publishes it (case-insensitive). A city or a zip is required.
	City *string `query:"city" json:"-"`

	// 5-digit ZIP. A city or a zip is required.
	Zip *string `query:"zip" json:"-"`

	// License status as published, mapped.
	Status *LicenseesFirmsParamsStatus `query:"status" json:"-"`

	// Minimum distinct locations per firm in the place.
	MinLocations *int64 `query:"min_locations" json:"-"`

	// Firms returned (largest first).
	Limit *int64 `query:"limit" json:"-"`
}

// LicenseesFirmsParamsState is generated from the OpenAPI spec. It is a string; the
// LicenseesFirmsParamsState* constants list the documented values.
type LicenseesFirmsParamsState = string

// Documented values of LicenseesFirmsParamsState.
const (
	LicenseesFirmsParamsStateFl LicenseesFirmsParamsState = "FL"
	LicenseesFirmsParamsStateCa LicenseesFirmsParamsState = "CA"
	LicenseesFirmsParamsStateNy LicenseesFirmsParamsState = "NY"
	LicenseesFirmsParamsStateCt LicenseesFirmsParamsState = "CT"
	LicenseesFirmsParamsStateVa LicenseesFirmsParamsState = "VA"
)

// LicenseesFirmsParamsProfession is generated from the OpenAPI spec. It is a string; the
// LicenseesFirmsParamsProfession* constants list the documented values.
type LicenseesFirmsParamsProfession = string

// Documented values of LicenseesFirmsParamsProfession.
const (
	LicenseesFirmsParamsProfessionRealEstate LicenseesFirmsParamsProfession = "real_estate"
	LicenseesFirmsParamsProfessionInsurance  LicenseesFirmsParamsProfession = "insurance"
	LicenseesFirmsParamsProfessionCpa        LicenseesFirmsParamsProfession = "cpa"
	LicenseesFirmsParamsProfessionCam        LicenseesFirmsParamsProfession = "cam"
)

// LicenseesFirmsParamsStatus is generated from the OpenAPI spec. It is a string; the
// LicenseesFirmsParamsStatus* constants list the documented values.
type LicenseesFirmsParamsStatus = string

// Documented values of LicenseesFirmsParamsStatus.
const (
	LicenseesFirmsParamsStatusActive     LicenseesFirmsParamsStatus = "active"
	LicenseesFirmsParamsStatusInactive   LicenseesFirmsParamsStatus = "inactive"
	LicenseesFirmsParamsStatusDelinquent LicenseesFirmsParamsStatus = "delinquent"
	LicenseesFirmsParamsStatusVoid       LicenseesFirmsParamsStatus = "void"
	LicenseesFirmsParamsStatusExpired    LicenseesFirmsParamsStatus = "expired"
	LicenseesFirmsParamsStatusOther      LicenseesFirmsParamsStatus = "other"
	LicenseesFirmsParamsStatusAny        LicenseesFirmsParamsStatus = "any"
)

// LicenseesFirmsResponse: Search licensed firms in a place
type LicenseesFirmsResponse struct {
	Query     map[string]any                `json:"query"`
	FirmCount int64                         `json:"firm_count"`
	FirmTotal int64                         `json:"firm_total"`
	Truncated bool                          `json:"truncated"`
	Firms     []LicenseesFirmsResponseFirms `json:"firms"`
	Note      string                        `json:"note"`

	// Present when person-shaped licensee fields were withheld from this caller (null values, keys
	// kept).
	PeopleFields map[string]any `json:"people_fields,omitempty"`
}

// UnmarshalJSON decodes LicenseesFirmsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LicenseesFirmsResponse) UnmarshalJSON(data []byte) error {
	type plain LicenseesFirmsResponse
	aux := struct {
		*plain
		FirmCount lenientNumber[int64] `json:"firm_count"`
		FirmTotal lenientNumber[int64] `json:"firm_total"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.FirmCount.assign(&r.FirmCount)
	aux.FirmTotal.assign(&r.FirmTotal)
	return softTypeError(err)
}

// LicenseesFirmsResponseFirms is generated from the OpenAPI spec.
type LicenseesFirmsResponseFirms struct {
	FirmKey            *string                                `json:"firm_key,omitempty"`
	Name               *string                                `json:"name,omitempty"`
	Profession         *string                                `json:"profession,omitempty"`
	LocationCount      *int64                                 `json:"location_count,omitempty"`
	LicenseCount       *int64                                 `json:"license_count,omitempty"`
	Issuers            []string                               `json:"issuers,omitempty"`
	Locations          []LicenseesFirmsResponseFirmsLocations `json:"locations,omitempty"`
	LocationsTruncated *bool                                  `json:"locations_truncated,omitempty"`
}

// UnmarshalJSON decodes LicenseesFirmsResponseFirms, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LicenseesFirmsResponseFirms) UnmarshalJSON(data []byte) error {
	type plain LicenseesFirmsResponseFirms
	aux := struct {
		*plain
		LocationCount lenientNumber[int64] `json:"location_count"`
		LicenseCount  lenientNumber[int64] `json:"license_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LocationCount.assignPtr(&r.LocationCount)
	aux.LicenseCount.assignPtr(&r.LicenseCount)
	return softTypeError(err)
}

// LicenseesFirmsResponseFirmsLocations is generated from the OpenAPI spec.
type LicenseesFirmsResponseFirmsLocations struct {
	LicenseUid      *string                                        `json:"license_uid,omitempty"`
	SourceID        *string                                        `json:"source_id,omitempty"`
	IssuerName      *string                                        `json:"issuer_name,omitempty"`
	LicenseClass    *string                                        `json:"license_class,omitempty"`
	LicenseNumber   *string                                        `json:"license_number,omitempty"`
	Status          *LicenseesFirmsResponseFirmsLocationsStatus    `json:"status,omitempty"`
	StatusRaw       *string                                        `json:"status_raw,omitempty"`
	PartyType       *LicenseesFirmsResponseFirmsLocationsPartyType `json:"party_type,omitempty"`
	PersonShaped    *bool                                          `json:"person_shaped,omitempty"`
	Dba             *string                                        `json:"dba,omitempty"`
	AddrType        *LicenseesFirmsResponseFirmsLocationsAddrType  `json:"addr_type,omitempty"`
	AddressLine     *string                                        `json:"address_line,omitempty"`
	Unit            *string                                        `json:"unit,omitempty"`
	City            *string                                        `json:"city,omitempty"`
	State           *string                                        `json:"state,omitempty"`
	Zip5            *string                                        `json:"zip5,omitempty"`
	Email           *string                                        `json:"email,omitempty"`
	Phone           *string                                        `json:"phone,omitempty"`
	MatchConfidence *float64                                       `json:"match_confidence,omitempty"`
	AsOf            *string                                        `json:"as_of,omitempty"`
	Name            *string                                        `json:"name,omitempty"`

	// Canonical parcel id the address matched, when it matched.
	ParcelID     *string `json:"parcel_id,omitempty"`
	OperatesHere *bool   `json:"operates_here,omitempty"`
}

// UnmarshalJSON decodes LicenseesFirmsResponseFirmsLocations, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LicenseesFirmsResponseFirmsLocations) UnmarshalJSON(data []byte) error {
	type plain LicenseesFirmsResponseFirmsLocations
	aux := struct {
		*plain
		MatchConfidence lenientNumber[float64] `json:"match_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MatchConfidence.assignPtr(&r.MatchConfidence)
	return softTypeError(err)
}

// LicenseesFirmsResponseFirmsLocationsStatus is generated from the OpenAPI spec. It is a string;
// the LicenseesFirmsResponseFirmsLocationsStatus* constants list the documented values.
type LicenseesFirmsResponseFirmsLocationsStatus = string

// Documented values of LicenseesFirmsResponseFirmsLocationsStatus.
const (
	LicenseesFirmsResponseFirmsLocationsStatusActive     LicenseesFirmsResponseFirmsLocationsStatus = "active"
	LicenseesFirmsResponseFirmsLocationsStatusInactive   LicenseesFirmsResponseFirmsLocationsStatus = "inactive"
	LicenseesFirmsResponseFirmsLocationsStatusDelinquent LicenseesFirmsResponseFirmsLocationsStatus = "delinquent"
	LicenseesFirmsResponseFirmsLocationsStatusVoid       LicenseesFirmsResponseFirmsLocationsStatus = "void"
	LicenseesFirmsResponseFirmsLocationsStatusExpired    LicenseesFirmsResponseFirmsLocationsStatus = "expired"
	LicenseesFirmsResponseFirmsLocationsStatusOther      LicenseesFirmsResponseFirmsLocationsStatus = "other"
)

// LicenseesFirmsResponseFirmsLocationsPartyType is generated from the OpenAPI spec. It is a
// string; the LicenseesFirmsResponseFirmsLocationsPartyType* constants list the documented values.
type LicenseesFirmsResponseFirmsLocationsPartyType = string

// Documented values of LicenseesFirmsResponseFirmsLocationsPartyType.
const (
	LicenseesFirmsResponseFirmsLocationsPartyTypeFirm   LicenseesFirmsResponseFirmsLocationsPartyType = "firm"
	LicenseesFirmsResponseFirmsLocationsPartyTypeBranch LicenseesFirmsResponseFirmsLocationsPartyType = "branch"
	LicenseesFirmsResponseFirmsLocationsPartyTypePerson LicenseesFirmsResponseFirmsLocationsPartyType = "person"
)

// LicenseesFirmsResponseFirmsLocationsAddrType is generated from the OpenAPI spec. It is a string;
// the LicenseesFirmsResponseFirmsLocationsAddrType* constants list the documented values.
type LicenseesFirmsResponseFirmsLocationsAddrType = string

// Documented values of LicenseesFirmsResponseFirmsLocationsAddrType.
const (
	LicenseesFirmsResponseFirmsLocationsAddrTypeBusiness        LicenseesFirmsResponseFirmsLocationsAddrType = "business"
	LicenseesFirmsResponseFirmsLocationsAddrTypeMailing         LicenseesFirmsResponseFirmsLocationsAddrType = "mailing"
	LicenseesFirmsResponseFirmsLocationsAddrTypeAddressOfRecord LicenseesFirmsResponseFirmsLocationsAddrType = "address_of_record"
)
