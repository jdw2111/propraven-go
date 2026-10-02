// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// WebhooksService groups the webhooks operations. Use it as client.Webhooks.
type WebhooksService struct {
	client *Client
}

// List: List webhook endpoints
//
// Returns all webhook endpoints for the calling account, plus the per-tier quota.
//
// HTTP: GET /api/v1/webhooks
func (s *WebhooksService) List(ctx context.Context, params *WebhooksListParams, opts ...RequestOption) (*WebhooksListResponse, error) {
	var out WebhooksListResponse
	if err := s.client.do(ctx, buildWebhooksListRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksListRequest(params *WebhooksListParams) *apiRequest {
	req := newRequest("GET", "/api/v1/webhooks")
	return req
}

// WebhooksListParams holds the query, header and JSON-body parameters of [WebhooksService.List].
// Pass nil when you need none.
type WebhooksListParams struct {
}

// WebhooksListResponse: List webhook endpoints
type WebhooksListResponse struct {
	Webhooks []Webhook                `json:"webhooks"`
	Quota    WebhookQuota             `json:"quota"`
	Tier     WebhooksListResponseTier `json:"tier"`
}

// WebhooksListResponseTier is generated from the OpenAPI spec. It is a string; the
// WebhooksListResponseTier* constants list the documented values.
type WebhooksListResponseTier = string

// Documented values of WebhooksListResponseTier.
const (
	WebhooksListResponseTierFree    WebhooksListResponseTier = "free"
	WebhooksListResponseTierStarter WebhooksListResponseTier = "starter"
	WebhooksListResponseTierPro     WebhooksListResponseTier = "pro"
	WebhooksListResponseTierScale   WebhooksListResponseTier = "scale"
	WebhooksListResponseTierAPI100k WebhooksListResponseTier = "api_100k"
)

// Create: Create a webhook endpoint
//
// Creates a new webhook subscription. The returned `secret` is shown ONCE — store it server-side
// and use it to verify every incoming delivery via the `X-PropRaven-Signature` header (HMAC-SHA256
// over `<unix_ms>.<raw_body>`). Reject deliveries where `|now - t| > 5min`.
//
// HTTP: POST /api/v1/webhooks
func (s *WebhooksService) Create(ctx context.Context, params *WebhooksCreateParams, opts ...RequestOption) (*WebhooksCreateResponse, error) {
	var out WebhooksCreateResponse
	if err := s.client.do(ctx, buildWebhooksCreateRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksCreateRequest(params *WebhooksCreateParams) *apiRequest {
	req := newRequest("POST", "/api/v1/webhooks")
	if params != nil {
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// WebhooksCreateParams holds the query, header and JSON-body parameters of
// [WebhooksService.Create]. Pass nil when you need none.
type WebhooksCreateParams struct {
	// Customer endpoint. https:// only.
	//
	// Required (JSON body).
	URL string `json:"url"`

	// Event types to subscribe to. NOTE: only parcel.sold is live in v1.0; others 501.
	//
	// Required (JSON body).
	EventTypes []WebhooksCreateParamsEventTypes `json:"event_types"`

	// Required (JSON body).
	FilterKind WebhooksCreateParamsFilterKind `json:"filter_kind"`

	// Required (JSON body).
	FilterValue WebhookFilter `json:"filter_value"`

	// Optional human-readable label for your dashboard.
	Description *string `json:"description,omitempty"`
}

// WebhooksCreateParamsEventTypes is generated from the OpenAPI spec. It is a string; the
// WebhooksCreateParamsEventTypes* constants list the documented values.
type WebhooksCreateParamsEventTypes = string

// Documented values of WebhooksCreateParamsEventTypes.
const (
	WebhooksCreateParamsEventTypesParcelSold         WebhooksCreateParamsEventTypes = "parcel.sold"
	WebhooksCreateParamsEventTypesParcelPermitFiled  WebhooksCreateParamsEventTypes = "parcel.permit_filed"
	WebhooksCreateParamsEventTypesParcelOwnerChanged WebhooksCreateParamsEventTypes = "parcel.owner_changed"
)

// WebhooksCreateParamsFilterKind is generated from the OpenAPI spec. It is a string; the
// WebhooksCreateParamsFilterKind* constants list the documented values.
type WebhooksCreateParamsFilterKind = string

// Documented values of WebhooksCreateParamsFilterKind.
const (
	WebhooksCreateParamsFilterKindParcelIDs  WebhooksCreateParamsFilterKind = "parcel_ids"
	WebhooksCreateParamsFilterKindStateFIPS  WebhooksCreateParamsFilterKind = "state_fips"
	WebhooksCreateParamsFilterKindCountyFIPS WebhooksCreateParamsFilterKind = "county_fips"
)

// WebhooksCreateResponse: Create a webhook endpoint
type WebhooksCreateResponse = WebhookCreated

// Get: Get a single webhook endpoint
//
// Returns the full endpoint record (without the secret).
//
// HTTP: GET /api/v1/webhooks/{id}
func (s *WebhooksService) Get(ctx context.Context, id string, params *WebhooksGetParams, opts ...RequestOption) (*WebhooksGetResponse, error) {
	var out WebhooksGetResponse
	if err := s.client.do(ctx, buildWebhooksGetRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksGetRequest(id string, params *WebhooksGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/webhooks/"+pathParam(id))
	return req
}

// WebhooksGetParams holds the query, header and JSON-body parameters of [WebhooksService.Get].
// Pass nil when you need none.
type WebhooksGetParams struct {
}

// WebhooksGetResponse: Get a single webhook endpoint
type WebhooksGetResponse = Webhook

// Delete: Soft-disable a webhook endpoint
//
// Marks the endpoint inactive. Delivery history is preserved. The endpoint can no longer receive
// new events but past deliveries remain queryable via the deliveries route.
//
// HTTP: DELETE /api/v1/webhooks/{id}
func (s *WebhooksService) Delete(ctx context.Context, id string, params *WebhooksDeleteParams, opts ...RequestOption) (*WebhooksDeleteResponse, error) {
	var out WebhooksDeleteResponse
	if err := s.client.do(ctx, buildWebhooksDeleteRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksDeleteRequest(id string, params *WebhooksDeleteParams) *apiRequest {
	req := newRequest("DELETE", "/api/v1/webhooks/"+pathParam(id))
	return req
}

// WebhooksDeleteParams holds the query, header and JSON-body parameters of
// [WebhooksService.Delete]. Pass nil when you need none.
type WebhooksDeleteParams struct {
}

// WebhooksDeleteResponse: Soft-disable a webhook endpoint
type WebhooksDeleteResponse struct {
	Deleted *string `json:"deleted,omitempty"`
}

// Deliveries: Recent delivery attempts for a webhook
//
// Returns the last 100 delivery attempts for an endpoint — useful for debugging signature
// mismatches, retry visibility, and dead-letter inspection.
//
// HTTP: GET /api/v1/webhooks/{id}/deliveries
func (s *WebhooksService) Deliveries(ctx context.Context, id string, params *WebhooksDeliveriesParams, opts ...RequestOption) (*WebhooksDeliveriesResponse, error) {
	var out WebhooksDeliveriesResponse
	if err := s.client.do(ctx, buildWebhooksDeliveriesRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksDeliveriesRequest(id string, params *WebhooksDeliveriesParams) *apiRequest {
	req := newRequest("GET", "/api/v1/webhooks/"+pathParam(id)+"/deliveries")
	return req
}

// WebhooksDeliveriesParams holds the query, header and JSON-body parameters of
// [WebhooksService.Deliveries]. Pass nil when you need none.
type WebhooksDeliveriesParams struct {
}

// WebhooksDeliveriesResponse: Recent delivery attempts for a webhook
type WebhooksDeliveriesResponse struct {
	Deliveries []WebhookDelivery `json:"deliveries,omitempty"`
}

// RetryDelivery: Re-queue a failed webhook delivery
//
// Moves a `pending`, `failed` or `dead_lettered` delivery back to `pending` for immediate
// redelivery. A delivered one is a 409.
//
// HTTP: POST /api/v1/webhooks/{id}/deliveries/{deliveryId}/retry
func (s *WebhooksService) RetryDelivery(ctx context.Context, id string, deliveryID string, params *WebhooksRetryDeliveryParams, opts ...RequestOption) (*WebhooksRetryDeliveryResponse, error) {
	var out WebhooksRetryDeliveryResponse
	if err := s.client.do(ctx, buildWebhooksRetryDeliveryRequest(id, deliveryID, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildWebhooksRetryDeliveryRequest(id string, deliveryID string, params *WebhooksRetryDeliveryParams) *apiRequest {
	req := newRequest("POST", "/api/v1/webhooks/"+pathParam(id)+"/deliveries/"+pathParam(deliveryID)+"/retry")
	return req
}

// WebhooksRetryDeliveryParams holds the query, header and JSON-body parameters of
// [WebhooksService.RetryDelivery]. Pass nil when you need none.
type WebhooksRetryDeliveryParams struct {
}

// WebhooksRetryDeliveryResponse: Re-queue a failed webhook delivery
type WebhooksRetryDeliveryResponse = json.RawMessage
