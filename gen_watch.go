// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
)

// WatchService groups the watch operations. Use it as client.Watch.
type WatchService struct {
	client *Client
}

// List: List your watches
//
// List active watches for the X-CREDIT-TOKEN. Free.
//
// HTTP: GET /api/v1/watch
func (s *WatchService) List(ctx context.Context, params *WatchListParams, opts ...RequestOption) (*WatchListResponse, error) {
	var out WatchListResponse
	if err := s.client.do(ctx, buildWatchListRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWatchListRequest(params *WatchListParams) *apiRequest {
	req := newRequest("GET", "/api/v1/watch")
	if params != nil {
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
	}
	return req
}

// WatchListParams holds the query, header and JSON-body parameters of [WatchService.List]. Pass
// nil when you need none.
type WatchListParams struct {
	// The pzc_ credit token (identity + wallet).
	//
	// Required.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`
}

// WatchListResponse: List your watches
type WatchListResponse = map[string]any

// Create: Create a watch (free)
//
// Create a watch tied to your credit token: a filter over parcels/geography + change event types.
// It reports changes going FORWARD; poll it to receive (and pay per delta for) new changes.
// Managing watches is free. Body: { filter, event_types?, name? } where filter is {
// parcel_ids:[...] } | { canonical_ids:[...] } | { state_fips } | { county_fips, state_fips } and
// event_types is a subset of parcel.sold / parcel.owner_changed / parcel.permit_filed (default
// all). Parcel ids (1–500) are canonical `state_fips:county_fips:parcel_id` ids, parcel UUIDs,
// or bare parcel_ids that name exactly ONE served parcel; a bare id that names several parcels is
// refused (422, candidates listed). The watch is stored and matched on canonical ids;
// `parcel_resolution` reports how each input resolved.
//
// HTTP: POST /api/v1/watch
func (s *WatchService) Create(ctx context.Context, params *WatchCreateParams, opts ...RequestOption) (*WatchCreateResponse, error) {
	var out WatchCreateResponse
	if err := s.client.do(ctx, buildWatchCreateRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWatchCreateRequest(params *WatchCreateParams) *apiRequest {
	req := newRequest("POST", "/api/v1/watch")
	if params != nil {
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// WatchCreateParams holds the query, header and JSON-body parameters of [WatchService.Create].
// Pass nil when you need none.
type WatchCreateParams struct {
	// The pzc_ credit token.
	//
	// Required.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`
	Name        *string `json:"name,omitempty"`

	// Default: all three.
	EventTypes []WatchCreateParamsEventTypes `json:"event_types,omitempty"`

	// Exactly one shape: `{parcel_ids: [...]}` / `{canonical_ids: [...]}` (1-500 ids), `{state_fips}`
	// or `{state_fips, county_fips}`.
	//
	// Required (JSON body).
	Filter WatchCreateParamsFilter `json:"filter"`
}

// WatchCreateParamsEventTypes is generated from the OpenAPI spec. It is a string; the
// WatchCreateParamsEventTypes* constants list the documented values.
type WatchCreateParamsEventTypes = string

// Documented values of WatchCreateParamsEventTypes.
const (
	WatchCreateParamsEventTypesParcelSold         WatchCreateParamsEventTypes = "parcel.sold"
	WatchCreateParamsEventTypesParcelOwnerChanged WatchCreateParamsEventTypes = "parcel.owner_changed"
	WatchCreateParamsEventTypesParcelPermitFiled  WatchCreateParamsEventTypes = "parcel.permit_filed"
)

// WatchCreateParamsFilter: Exactly one shape: `{parcel_ids: [...]}` / `{canonical_ids: [...]}`
// (1-500 ids), `{state_fips}` or `{state_fips, county_fips}`.
type WatchCreateParamsFilter struct {
	ParcelIDs    []string `json:"parcel_ids,omitempty"`
	CanonicalIDs []string `json:"canonical_ids,omitempty"`
	StateFIPS    *string  `json:"state_fips,omitempty"`
	CountyFIPS   *string  `json:"county_fips,omitempty"`
}

// WatchCreateResponse: Create a watch (free)
type WatchCreateResponse = map[string]any

// Poll: Poll a watch for new changes (priced per delta)
//
// Poll new parcel_deeds / parcel_permits changes matching the watch since its per-source cursor.
// preview=true returns the pending count + exact price + a masked sample (FREE, no cursor move). A
// paid poll debits the credit balance PER DELTA (clamp($0.05 x event-strength, $0.02, $0.20),
// capped $20/poll), advances the cursor, and returns the full deltas. A poll with no new changes
// is free. Insufficient balance -> 402. Parcel watches match on the full parcel identity; each
// delta carries `canonical_id` + `match_basis`. A watch created before 2026-09-22 whose stored
// bare ids are ambiguous reports them in `identity_issues` (not matched) on every poll. Deed
// parties (grantor / grantee names and addresses, prior / new owner) are people data, delivered to
// accounts only: a poll paid with a credit token and no account (no API key, no signed-in session)
// receives the deltas with those fields set to null and a top-level `people_fields` marker (see
// PeopleFieldsWithheld). The events themselves (what changed, where, when, for how much) are
// unchanged.
//
// HTTP: GET /api/v1/watch/{id}
func (s *WatchService) Poll(ctx context.Context, id string, params *WatchPollParams, opts ...RequestOption) (*WatchPollResponse, error) {
	var out WatchPollResponse
	if err := s.client.do(ctx, buildWatchPollRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWatchPollRequest(id string, params *WatchPollParams) *apiRequest {
	req := newRequest("GET", "/api/v1/watch/"+pathParam(id))
	if params != nil {
		addQuery(req.query, "preview", params.Preview)
		addQuery(req.query, "limit", params.Limit)
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
	}
	return req
}

// WatchPollParams holds the query, header and JSON-body parameters of [WatchService.Poll]. Pass
// nil when you need none.
type WatchPollParams struct {
	// FREE count + price + masked sample; no cursor move.
	Preview *bool `query:"preview" json:"-"`

	// Max deltas per source (default 50, max 200).
	Limit *int64 `query:"limit" json:"-"`

	// The pzc_ credit token.
	//
	// Required.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`
}

// WatchPollResponse: Poll a watch for new changes (priced per delta)
type WatchPollResponse struct {
	Deltas       []map[string]any      `json:"deltas,omitempty"`
	PeopleFields *PeopleFieldsWithheld `json:"people_fields,omitempty"`
}

// Delete: Delete a watch
//
// Deactivate a watch. Free.
//
// HTTP: DELETE /api/v1/watch/{id}
func (s *WatchService) Delete(ctx context.Context, id string, params *WatchDeleteParams, opts ...RequestOption) (*WatchDeleteResponse, error) {
	var out WatchDeleteResponse
	if err := s.client.do(ctx, buildWatchDeleteRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWatchDeleteRequest(id string, params *WatchDeleteParams) *apiRequest {
	req := newRequest("DELETE", "/api/v1/watch/"+pathParam(id))
	if params != nil {
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
	}
	return req
}

// WatchDeleteParams holds the query, header and JSON-body parameters of [WatchService.Delete].
// Pass nil when you need none.
type WatchDeleteParams struct {
	// The pzc_ credit token.
	//
	// Required.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`
}

// WatchDeleteResponse: Delete a watch
type WatchDeleteResponse = map[string]any
