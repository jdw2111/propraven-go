// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// CohortsService groups the cohorts operations. Use it as client.Cohorts.
type CohortsService struct {
	client *Client
}

// Export: Mail-merge export of one of your lists (account required; included for subscribers, per
// row otherwise)
//
// A mail-merge CSV of one of the caller's own lists: one row per MAILING ADDRESS (owner, mailing
// line1 / city / state / ZIP, the property address(es), their canonical ids, stage, source, as_of,
// as_of_basis, grade). Every address is read from ONE address column family of a parcel's record,
// with its ZIP; only mail-ready addresses become rows, and `X-Export-Parcels-Not-Mail-Ready`
// counts the rest.
//
// ACCOUNT REQUIRED: an API key or a signed-in session; anonymous callers -- including x402 /
// credit-token wallets -- get HTTP 401 `code: "account_required"` before any read. The list must
// be the caller's own (404 otherwise).
//
// PRICE: INCLUDED for a paid PropRaven subscription. Any other account pays $0.10 per mailing row,
// the quote capped at $20, from a prepaid credit balance (`X-CREDIT-TOKEN`) or per call via x402
// (`X-PAYMENT`, sent with your credentials); with neither, HTTP 402 carrying the exact quote.
// `preview=true` returns the row count and the quote, free, with no data.
//
// Each export is recorded in PropRaven's people-data access log BEFORE the CSV is returned (and
// before an x402 payment settles). If the log cannot be written the export is refused (503) and
// any charge refunded or released.
//
// HTTP: GET /api/v1/cohorts/{id}/export
func (s *CohortsService) Export(ctx context.Context, id string, params *CohortsExportParams, opts ...RequestOption) (*CohortsExportResponse, error) {
	var out CohortsExportResponse
	if err := s.client.do(ctx, buildCohortsExportRequest(id, params), opts, decodeMixed(&out.ContentType, &out.Text, func(b []byte) error {
		return json.Unmarshal(b, &out.JSON)
	})); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCohortsExportRequest(id string, params *CohortsExportParams) *apiRequest {
	req := newRequest("GET", "/api/v1/cohorts/"+pathParam(id)+"/export")
	req.accept = "text/csv, application/json"
	if params != nil {
		addQuery(req.query, "preview", params.Preview)
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// CohortsExportParams holds the query, header and JSON-body parameters of [CohortsService.Export].
// Pass nil when you need none.
type CohortsExportParams struct {
	// true -> the free row count and exact quote, no data.
	Preview *bool `query:"preview" json:"-"`

	// A prepaid credit token (pzc_...) to draw the per-row price from.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`

	// x402 payment for the quoted total, sent together with your account credentials.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// CohortsExportResponse: Mail-merge export of one of your lists (account required; included for
// subscribers, per row otherwise)
//
// The server answers text/csv or JSON; Text holds a non-JSON body and JSON a decoded JSON one.
type CohortsExportResponse struct {
	// ContentType is the response Content-Type.
	ContentType string
	// Text is the body when it is not JSON (text/csv).
	Text string
	// JSON is the decoded body when the server answered with JSON.
	JSON map[string]any
}

// List: List your saved parcel lists (cohorts)
//
// Cohorts are saved parcel lists built in the PropRaven app; export one with GET
// /cohorts/{id}/export.
//
// HTTP: GET /api/v1/cohorts
func (s *CohortsService) List(ctx context.Context, params *CohortsListParams, opts ...RequestOption) (*CohortsListResponse, error) {
	var out CohortsListResponse
	if err := s.client.do(ctx, buildCohortsListRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCohortsListRequest(params *CohortsListParams) *apiRequest {
	req := newRequest("GET", "/api/v1/cohorts")
	return req
}

// CohortsListParams holds the query, header and JSON-body parameters of [CohortsService.List].
// Pass nil when you need none.
type CohortsListParams struct {
}

// CohortsListResponse: List your saved parcel lists (cohorts)
type CohortsListResponse struct {
	Cohorts []CohortsListResponseCohorts `json:"cohorts"`
}

// CohortsListResponseCohorts is generated from the OpenAPI spec.
type CohortsListResponseCohorts struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	ParcelCount *int64  `json:"parcel_count,omitempty"`
	TotalValue  *int64  `json:"total_value,omitempty"`
}

// UnmarshalJSON decodes CohortsListResponseCohorts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CohortsListResponseCohorts) UnmarshalJSON(data []byte) error {
	type plain CohortsListResponseCohorts
	aux := struct {
		*plain
		ParcelCount lenientNumber[int64] `json:"parcel_count"`
		TotalValue  lenientNumber[int64] `json:"total_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelCount.assignPtr(&r.ParcelCount)
	aux.TotalValue.assignPtr(&r.TotalValue)
	return softTypeError(err)
}
