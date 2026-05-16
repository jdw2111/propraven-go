// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/jdw2111/propraven-go/internal/apijson"
	"github.com/jdw2111/propraven-go/internal/requestconfig"
	"github.com/jdw2111/propraven-go/option"
	"github.com/jdw2111/propraven-go/packages/param"
	"github.com/jdw2111/propraven-go/packages/respjson"
)

// Webhook subscriptions and delivery history. Manage which events PropRaven pushes
// to your endpoints.
//
// V1WebhookService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1WebhookService] method instead.
type V1WebhookService struct {
	options []option.RequestOption
}

// NewV1WebhookService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1WebhookService(opts ...option.RequestOption) (r V1WebhookService) {
	r = V1WebhookService{}
	r.options = opts
	return
}

// Creates a new webhook subscription. The returned `secret` is shown ONCE — store
// it server-side and use it to verify every incoming delivery via the
// `X-PropRaven-Signature` header (HMAC-SHA256 over `<unix_ms>.<raw_body>`). Reject
// deliveries where `|now - t| > 5min`.
func (r *V1WebhookService) NewEndpoint(ctx context.Context, body V1WebhookNewEndpointParams, opts ...option.RequestOption) (res *V1WebhookNewEndpointResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the endpoint inactive. Delivery history is preserved. The endpoint can no
// longer receive new events but past deliveries remain queryable via the
// deliveries route.
func (r *V1WebhookService) DisableEndpoint(ctx context.Context, id string, opts ...option.RequestOption) (res *V1WebhookDisableEndpointResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns all webhook endpoints for the calling account, plus the per-tier quota.
func (r *V1WebhookService) ListEndpoints(ctx context.Context, opts ...option.RequestOption) (res *V1WebhookListEndpointsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the last 100 delivery attempts for an endpoint — useful for debugging
// signature mismatches, retry visibility, and dead-letter inspection.
func (r *V1WebhookService) GetDeliveries(ctx context.Context, id string, opts ...option.RequestOption) (res *V1WebhookGetDeliveriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/webhooks/%s/deliveries", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the full endpoint record (without the secret).
func (r *V1WebhookService) GetEndpoint(ctx context.Context, id string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Webhook struct {
	ID                  string    `json:"id" format:"uuid"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	DeliveriesAttempted int64     `json:"deliveries_attempted"`
	DeliveriesSucceeded int64     `json:"deliveries_succeeded"`
	Description         string    `json:"description"`
	DisabledAt          time.Time `json:"disabled_at" format:"date-time"`
	DisabledReason      string    `json:"disabled_reason" api:"nullable"`
	// Any of "parcel.sold", "parcel.permit_filed", "parcel.owner_changed".
	EventTypes []string `json:"event_types"`
	// Any of "parcel_ids", "state_fips", "county_fips".
	FilterKind WebhookFilterKind `json:"filter_kind"`
	// Shape varies with filter_kind. parcel_ids: explicit list. state_fips: all
	// parcels in a state. county_fips: all parcels in a county within a state.
	FilterValue    WebhookFilterUnion `json:"filter_value"`
	IsActive       bool               `json:"is_active"`
	LastDeliveryAt time.Time          `json:"last_delivery_at" format:"date-time"`
	LastSuccessAt  time.Time          `json:"last_success_at" format:"date-time"`
	// First 14 chars of the secret (whsec\_ + 8 hex). Use to identify the webhook in
	// your dashboard; full secret is shown only at create time.
	SecretPrefix string `json:"secret_prefix"`
	// Customer endpoint. Must be https://.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		DeliveriesAttempted respjson.Field
		DeliveriesSucceeded respjson.Field
		Description         respjson.Field
		DisabledAt          respjson.Field
		DisabledReason      respjson.Field
		EventTypes          respjson.Field
		FilterKind          respjson.Field
		FilterValue         respjson.Field
		IsActive            respjson.Field
		LastDeliveryAt      respjson.Field
		LastSuccessAt       respjson.Field
		SecretPrefix        respjson.Field
		URL                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookFilterKind string

const (
	WebhookFilterKindParcelIDs  WebhookFilterKind = "parcel_ids"
	WebhookFilterKindStateFips  WebhookFilterKind = "state_fips"
	WebhookFilterKindCountyFips WebhookFilterKind = "county_fips"
)

// WebhookFilterUnion contains all possible properties and values from
// [WebhookFilterParcelIDs], [WebhookFilterStateFips], [WebhookFilterObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebhookFilterUnion struct {
	// This field is from variant [WebhookFilterParcelIDs].
	ParcelIDs []string `json:"parcel_ids"`
	StateFips string   `json:"state_fips"`
	// This field is from variant [WebhookFilterObject].
	CountyFips string `json:"county_fips"`
	JSON       struct {
		ParcelIDs  respjson.Field
		StateFips  respjson.Field
		CountyFips respjson.Field
		raw        string
	} `json:"-"`
}

func (u WebhookFilterUnion) AsWebhookFilterParcelIDs() (v WebhookFilterParcelIDs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookFilterUnion) AsWebhookFilterStateFips() (v WebhookFilterStateFips) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookFilterUnion) AsWebhookFilterObject() (v WebhookFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebhookFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *WebhookFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebhookFilterUnion to a WebhookFilterUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebhookFilterUnionParam.Overrides()
func (r WebhookFilterUnion) ToParam() WebhookFilterUnionParam {
	return param.Override[WebhookFilterUnionParam](json.RawMessage(r.RawJSON()))
}

type WebhookFilterParcelIDs struct {
	// Composite parcel IDs to subscribe to. 1–1000 IDs.
	ParcelIDs []string `json:"parcel_ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParcelIDs   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookFilterParcelIDs) RawJSON() string { return r.JSON.raw }
func (r *WebhookFilterParcelIDs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookFilterStateFips struct {
	// 2-digit state FIPS — subscribe to all parcels in this state.
	StateFips string `json:"state_fips" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StateFips   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookFilterStateFips) RawJSON() string { return r.JSON.raw }
func (r *WebhookFilterStateFips) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookFilterObject struct {
	// 3-digit county FIPS (within the state) — subscribe to all parcels in this
	// county.
	CountyFips string `json:"county_fips" api:"required"`
	StateFips  string `json:"state_fips" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountyFips  respjson.Field
		StateFips   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookFilterObject) RawJSON() string { return r.JSON.raw }
func (r *WebhookFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func WebhookFilterParamOfWebhookFilterParcelIDs(parcelIDs []string) WebhookFilterUnionParam {
	var variant WebhookFilterParcelIDsParam
	variant.ParcelIDs = parcelIDs
	return WebhookFilterUnionParam{OfWebhookFilterParcelIDs: &variant}
}

func WebhookFilterParamOfWebhookFilterStateFips(stateFips string) WebhookFilterUnionParam {
	var variant WebhookFilterStateFipsParam
	variant.StateFips = stateFips
	return WebhookFilterUnionParam{OfWebhookFilterStateFips: &variant}
}

func WebhookFilterParamOfWebhookFilterObject(countyFips string, stateFips string) WebhookFilterUnionParam {
	var variant WebhookFilterObjectParam
	variant.CountyFips = countyFips
	variant.StateFips = stateFips
	return WebhookFilterUnionParam{OfWebhookFilterObject: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookFilterUnionParam struct {
	OfWebhookFilterParcelIDs *WebhookFilterParcelIDsParam `json:",omitzero,inline"`
	OfWebhookFilterStateFips *WebhookFilterStateFipsParam `json:",omitzero,inline"`
	OfWebhookFilterObject    *WebhookFilterObjectParam    `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhookFilterParcelIDs, u.OfWebhookFilterStateFips, u.OfWebhookFilterObject)
}
func (u *WebhookFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The property ParcelIDs is required.
type WebhookFilterParcelIDsParam struct {
	// Composite parcel IDs to subscribe to. 1–1000 IDs.
	ParcelIDs []string `json:"parcel_ids,omitzero" api:"required"`
	paramObj
}

func (r WebhookFilterParcelIDsParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookFilterParcelIDsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookFilterParcelIDsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property StateFips is required.
type WebhookFilterStateFipsParam struct {
	// 2-digit state FIPS — subscribe to all parcels in this state.
	StateFips string `json:"state_fips" api:"required"`
	paramObj
}

func (r WebhookFilterStateFipsParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookFilterStateFipsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookFilterStateFipsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CountyFips, StateFips are required.
type WebhookFilterObjectParam struct {
	// 3-digit county FIPS (within the state) — subscribe to all parcels in this
	// county.
	CountyFips string `json:"county_fips" api:"required"`
	StateFips  string `json:"state_fips" api:"required"`
	paramObj
}

func (r WebhookFilterObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookFilterObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookFilterObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookNewEndpointResponse struct {
	// Signature-verification reminder.
	Hint string `json:"hint" api:"required"`
	// **Shown once.** Copy and store server-side immediately. Used to sign every
	// outgoing delivery.
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hint        respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Webhook
}

// Returns the unmodified JSON received from the API
func (r V1WebhookNewEndpointResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookNewEndpointResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookDisableEndpointResponse struct {
	Deleted string `json:"deleted" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1WebhookDisableEndpointResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookDisableEndpointResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookListEndpointsResponse struct {
	Quota V1WebhookListEndpointsResponseQuota `json:"quota"`
	// Any of "free", "starter", "pro", "scale", "api_100k".
	Tier     V1WebhookListEndpointsResponseTier `json:"tier"`
	Webhooks []Webhook                          `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quota       respjson.Field
		Tier        respjson.Field
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1WebhookListEndpointsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookListEndpointsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookListEndpointsResponseQuota struct {
	// Max simultaneous active webhook endpoints on this tier.
	MaxEndpoints int64 `json:"maxEndpoints"`
	// Max event deliveries per UTC day on this tier.
	MaxEventsPerDay int64 `json:"maxEventsPerDay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxEndpoints    respjson.Field
		MaxEventsPerDay respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1WebhookListEndpointsResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookListEndpointsResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookListEndpointsResponseTier string

const (
	V1WebhookListEndpointsResponseTierFree    V1WebhookListEndpointsResponseTier = "free"
	V1WebhookListEndpointsResponseTierStarter V1WebhookListEndpointsResponseTier = "starter"
	V1WebhookListEndpointsResponseTierPro     V1WebhookListEndpointsResponseTier = "pro"
	V1WebhookListEndpointsResponseTierScale   V1WebhookListEndpointsResponseTier = "scale"
	V1WebhookListEndpointsResponseTierAPI100k V1WebhookListEndpointsResponseTier = "api_100k"
)

type V1WebhookGetDeliveriesResponse struct {
	Deliveries []V1WebhookGetDeliveriesResponseDelivery `json:"deliveries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1WebhookGetDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookGetDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookGetDeliveriesResponseDelivery struct {
	ID             string    `json:"id" format:"uuid"`
	Attempts       int64     `json:"attempts"`
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	DeadLetteredAt time.Time `json:"dead_lettered_at" format:"date-time"`
	// Deterministic event identifier — sha256(source || pk || event_type). Idempotent
	// re-deliveries share this.
	EventID         string    `json:"event_id"`
	EventOccurredAt time.Time `json:"event_occurred_at" format:"date-time"`
	// Any of "parcel.sold", "parcel.permit_filed", "parcel.owner_changed".
	EventType     string    `json:"event_type"`
	LastAttemptAt time.Time `json:"last_attempt_at" format:"date-time"`
	LastError     string    `json:"last_error" api:"nullable"`
	// Truncated to ~1KB.
	LastResponseBody   string    `json:"last_response_body" api:"nullable"`
	LastResponseStatus int64     `json:"last_response_status" api:"nullable"`
	NextAttemptAt      time.Time `json:"next_attempt_at" format:"date-time"`
	// Any of "pending", "in_flight", "succeeded", "failed", "dead_lettered".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Attempts           respjson.Field
		CreatedAt          respjson.Field
		DeadLetteredAt     respjson.Field
		EventID            respjson.Field
		EventOccurredAt    respjson.Field
		EventType          respjson.Field
		LastAttemptAt      respjson.Field
		LastError          respjson.Field
		LastResponseBody   respjson.Field
		LastResponseStatus respjson.Field
		NextAttemptAt      respjson.Field
		Status             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1WebhookGetDeliveriesResponseDelivery) RawJSON() string { return r.JSON.raw }
func (r *V1WebhookGetDeliveriesResponseDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookNewEndpointParams struct {
	// Event types to subscribe to. NOTE: only parcel.sold is live in v1.0; others 501.
	//
	// Any of "parcel.sold", "parcel.permit_filed", "parcel.owner_changed".
	EventTypes []string `json:"event_types,omitzero" api:"required"`
	// Any of "parcel_ids", "state_fips", "county_fips".
	FilterKind V1WebhookNewEndpointParamsFilterKind `json:"filter_kind,omitzero" api:"required"`
	// Shape varies with filter_kind. parcel_ids: explicit list. state_fips: all
	// parcels in a state. county_fips: all parcels in a county within a state.
	FilterValue WebhookFilterUnionParam `json:"filter_value,omitzero" api:"required"`
	// Customer endpoint. https:// only.
	URL string `json:"url" api:"required" format:"uri"`
	// Optional human-readable label for your dashboard.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r V1WebhookNewEndpointParams) MarshalJSON() (data []byte, err error) {
	type shadow V1WebhookNewEndpointParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1WebhookNewEndpointParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WebhookNewEndpointParamsFilterKind string

const (
	V1WebhookNewEndpointParamsFilterKindParcelIDs  V1WebhookNewEndpointParamsFilterKind = "parcel_ids"
	V1WebhookNewEndpointParamsFilterKindStateFips  V1WebhookNewEndpointParamsFilterKind = "state_fips"
	V1WebhookNewEndpointParamsFilterKindCountyFips V1WebhookNewEndpointParamsFilterKind = "county_fips"
)
