// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// CMBSService groups the cmbs operations. Use it as client.CMBS.
type CMBSService struct {
	client *Client
}

// Exposure: CMBS loan exposure for a parcel or an owner
//
// Commercial mortgage-backed security loans tied to one parcel (`id`) or to a borrower/sponsor
// portfolio (`owner`). Pass exactly one.
//
// HTTP: GET /api/v1/cmbs/exposure
func (s *CMBSService) Exposure(ctx context.Context, params *CMBSExposureParams, opts ...RequestOption) (*CMBSExposureResponse, error) {
	var out CMBSExposureResponse
	if err := s.client.do(ctx, buildCMBSExposureRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCMBSExposureRequest(params *CMBSExposureParams) *apiRequest {
	req := newRequest("GET", "/api/v1/cmbs/exposure")
	if params != nil {
		addQuery(req.query, "id", params.ID)
		addQuery(req.query, "owner", params.Owner)
	}
	return req
}

// CMBSExposureParams holds the query, header and JSON-body parameters of [CMBSService.Exposure].
// Pass nil when you need none.
type CMBSExposureParams struct {
	// Parcel id (single-parcel mode).
	ID *string `query:"id" json:"-"`

	// Borrower or sponsor name (portfolio mode).
	Owner *string `query:"owner" json:"-"`
}

// CMBSExposureResponse: CMBS loan exposure for a parcel or an owner
type CMBSExposureResponse struct {
	Mode        *string                         `json:"mode,omitempty"`
	Matched     *bool                           `json:"matched,omitempty"`
	Reason      *string                         `json:"reason,omitempty"`
	Subject     CMBSExposureResponseSubject     `json:"subject"`
	Note        *string                         `json:"note,omitempty"`
	Loans       []json.RawMessage               `json:"loans"`
	Properties  []json.RawMessage               `json:"properties"`
	DataAsOf    CMBSExposureResponseDataAsOf    `json:"data_as_of"`
	Coverage    CMBSExposureResponseCoverage    `json:"coverage"`
	FieldStatus CMBSExposureResponseFieldStatus `json:"field_status"`
}

// CMBSExposureResponseSubject is generated from the OpenAPI spec.
type CMBSExposureResponseSubject struct {
	CanonicalID string `json:"canonical_id"`
	ParcelID    string `json:"parcel_id"`
	StateFIPS   string `json:"state_fips"`
	CountyFIPS  string `json:"county_fips"`
}

// CMBSExposureResponseDataAsOf is generated from the OpenAPI spec.
type CMBSExposureResponseDataAsOf struct {
	SnapshotAssembled *string `json:"snapshot_assembled,omitempty"`
	FiguresAsOf       *string `json:"figures_as_of,omitempty"`
	RefreshedSince    *bool   `json:"refreshed_since,omitempty"`
	Statement         *string `json:"statement,omitempty"`
	Source            *string `json:"source,omitempty"`
}

// CMBSExposureResponseCoverage is generated from the OpenAPI spec.
type CMBSExposureResponseCoverage struct {
	Scope        *string                                 `json:"scope,omitempty"`
	Trusts       *float64                                `json:"trusts,omitempty"`
	Loans        *float64                                `json:"loans,omitempty"`
	Properties   *float64                                `json:"properties,omitempty"`
	CountsSource *string                                 `json:"counts_source,omitempty"`
	EdgarRoster  CMBSExposureResponseCoverageEdgarRoster `json:"edgar_roster"`
	Excluded     CMBSExposureResponseCoverageExcluded    `json:"excluded"`
	NotCovered   []*string                               `json:"not_covered"`
}

// UnmarshalJSON decodes CMBSExposureResponseCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CMBSExposureResponseCoverage) UnmarshalJSON(data []byte) error {
	type plain CMBSExposureResponseCoverage
	aux := struct {
		*plain
		Trusts     lenientNumber[float64] `json:"trusts"`
		Loans      lenientNumber[float64] `json:"loans"`
		Properties lenientNumber[float64] `json:"properties"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Trusts.assignPtr(&r.Trusts)
	aux.Loans.assignPtr(&r.Loans)
	aux.Properties.assignPtr(&r.Properties)
	return softTypeError(err)
}

// CMBSExposureResponseCoverageEdgarRoster is generated from the OpenAPI spec.
type CMBSExposureResponseCoverageEdgarRoster struct {
	ReportingMonth        *string  `json:"reporting_month,omitempty"`
	CMBSTrustsFilingAbsEe *float64 `json:"cmbs_trusts_filing_abs_ee,omitempty"`
	OfWhichServed         *float64 `json:"of_which_served,omitempty"`
	MeasuredOn            *string  `json:"measured_on,omitempty"`
}

// UnmarshalJSON decodes CMBSExposureResponseCoverageEdgarRoster, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CMBSExposureResponseCoverageEdgarRoster) UnmarshalJSON(data []byte) error {
	type plain CMBSExposureResponseCoverageEdgarRoster
	aux := struct {
		*plain
		CMBSTrustsFilingAbsEe lenientNumber[float64] `json:"cmbs_trusts_filing_abs_ee"`
		OfWhichServed         lenientNumber[float64] `json:"of_which_served"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CMBSTrustsFilingAbsEe.assignPtr(&r.CMBSTrustsFilingAbsEe)
	aux.OfWhichServed.assignPtr(&r.OfWhichServed)
	return softTypeError(err)
}

// CMBSExposureResponseCoverageExcluded is generated from the OpenAPI spec.
type CMBSExposureResponseCoverageExcluded struct {
	NonCMBSTrusts             *float64 `json:"non_cmbs_trusts,omitempty"`
	NonCMBSNote               *string  `json:"non_cmbs_note,omitempty"`
	DepositorCiks             *float64 `json:"depositor_ciks,omitempty"`
	DepositorNote             *string  `json:"depositor_note,omitempty"`
	PropertyRowsStoredAsLoans *float64 `json:"property_rows_stored_as_loans,omitempty"`
	UnverifiedCiks            *float64 `json:"unverified_ciks,omitempty"`
}

// UnmarshalJSON decodes CMBSExposureResponseCoverageExcluded, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CMBSExposureResponseCoverageExcluded) UnmarshalJSON(data []byte) error {
	type plain CMBSExposureResponseCoverageExcluded
	aux := struct {
		*plain
		NonCMBSTrusts             lenientNumber[float64] `json:"non_cmbs_trusts"`
		DepositorCiks             lenientNumber[float64] `json:"depositor_ciks"`
		PropertyRowsStoredAsLoans lenientNumber[float64] `json:"property_rows_stored_as_loans"`
		UnverifiedCiks            lenientNumber[float64] `json:"unverified_ciks"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.NonCMBSTrusts.assignPtr(&r.NonCMBSTrusts)
	aux.DepositorCiks.assignPtr(&r.DepositorCiks)
	aux.PropertyRowsStoredAsLoans.assignPtr(&r.PropertyRowsStoredAsLoans)
	aux.UnverifiedCiks.assignPtr(&r.UnverifiedCiks)
	return softTypeError(err)
}

// CMBSExposureResponseFieldStatus is generated from the OpenAPI spec.
type CMBSExposureResponseFieldStatus struct {
	IsSpeciallyServiced *string                                   `json:"is_specially_serviced,omitempty"`
	IsOnWatchlist       *string                                   `json:"is_on_watchlist,omitempty"`
	IsInterestOnly      *string                                   `json:"is_interest_only,omitempty"`
	BorrowerName        *string                                   `json:"borrower_name,omitempty"`
	SponsorName         *string                                   `json:"sponsor_name,omitempty"`
	LoanType            *string                                   `json:"loan_type,omitempty"`
	PaymentStatus       *string                                   `json:"payment_status,omitempty"`
	NoteRate            *string                                   `json:"note_rate,omitempty"`
	CurrentBalance      *string                                   `json:"current_balance,omitempty"`
	UwDscr              *string                                   `json:"uw_dscr,omitempty"`
	UwLtv               *string                                   `json:"uw_ltv,omitempty"`
	UwOccupancy         *string                                   `json:"uw_occupancy,omitempty"`
	UwNoi               *string                                   `json:"uw_noi,omitempty"`
	AppraisedValue      *string                                   `json:"appraised_value,omitempty"`
	CurrentDscr         *string                                   `json:"current_dscr,omitempty"`
	CurrentLtv          *string                                   `json:"current_ltv,omitempty"`
	CurrentOccupancy    *string                                   `json:"current_occupancy,omitempty"`
	CurrentNoi          *string                                   `json:"current_noi,omitempty"`
	CurrentDebtYield    *string                                   `json:"current_debt_yield,omitempty"`
	Properties          CMBSExposureResponseFieldStatusProperties `json:"properties"`
}

// CMBSExposureResponseFieldStatusProperties is generated from the OpenAPI spec.
type CMBSExposureResponseFieldStatusProperties struct {
	LoanID              *string `json:"loan_id,omitempty"`
	AppraisedValue      *string `json:"appraised_value,omitempty"`
	AllocatedLoanAmount *string `json:"allocated_loan_amount,omitempty"`
	CurrentNoi          *string `json:"current_noi,omitempty"`
	CurrentOccupancy    *string `json:"current_occupancy,omitempty"`
	Latitude            *string `json:"latitude,omitempty"`
	Longitude           *string `json:"longitude,omitempty"`
	MatchedParcelID     *string `json:"matched_parcel_id,omitempty"`
}
