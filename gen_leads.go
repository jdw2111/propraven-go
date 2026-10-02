// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// LeadsService groups the leads operations. Use it as client.Leads.
type LeadsService struct {
	client *Client
}

// Find: Lead feed (paid, priced per lead) — with a FREE preview
//
// The Machine Storefront's lead feed: the qualified target list for one SIGNAL in one STATE,
// delivered as lead records and priced PER LEAD.
//
// Each lead carries its canonical_id (state:county:APN), owner, assessed value, that signal's own
// strength fields (years held, land/improvement ratio, flip profit, portfolio size, ...), a
// deterministic lead_score (1-100) and provenance {as_of, source}.
//
// PRICE: per_lead = clamp($0.25 x S(signal strength) x V(asset-value tier), $0.05, $1.00); total =
// min(count x per_lead, $20). You pay for the leads DELIVERED -- min(matching rows, limit) -- and
// an empty result is never charged for. The exact total is advertised in the 402's
// accepts[0].maxAmountRequired (USDC atomic units, 6 decimals).
//
// FREE PREVIEW: add preview=true for the exact count, the exact quote and up to three MASKED
// sample leads (APN truncated to state:county, house number stripped, owner name and every other
// people field withheld). No payment, no API key required; anonymous callers are IP-throttled at
// the free tier. The anonymous security alternative (`{}`) applies to the preview ONLY.
//
// ACCOUNT REQUIRED FOR DELIVERY: a lead is an owner's name and mailing address by area -- people
// data, delivered to PropRaven ACCOUNTS only. A paid (non-preview) pull must carry an API key
// (`Authorization: Bearer pz_...`), an MCP OAuth token or a signed-in session on EVERY rail.
// Payment alone (an x402 `X-PAYMENT` header or a prepaid `X-CREDIT-TOKEN`) is not an account:
// without one the call is refused with HTTP 401 `code: "account_required"`, `reason:
// "people_data_requires_account"`, before any payment is verified or any credit drawn.
//
// PAID ACCESS (preview omitted), for an account, requires ONE of: (a) x402 pay-per-call -- send a
// base64 signed x402 PaymentPayload in the `X-PAYMENT` header together with your credentials; on a
// successful build the leads are returned and the on-chain settlement receipt is in the
// `X-PAYMENT-RESPONSE` response header. (b) A prepaid `X-CREDIT-TOKEN` balance. (c) A genuine PAID
// PropRaven subscription entitlement (lead feeds are included). Being merely authenticated is NOT
// sufficient. (d) Anything else -> HTTP 402 whose `accepts` array carries the exact payment
// requirements for this pull.
//
// OWNER CONTACT: each delivered lead carries `owner_contact` -- the owner's best mailing address,
// read from ONE address column family of the record with its ZIP, flagged mail_ready.
// `mail_ready=true` narrows the pull to parcels whose record carries a complete one-family mailing
// address (parcel-grain signals only). Each delivery is recorded in PropRaven's people-data access
// log before it is returned; if the log is unavailable the delivery is refused (503) and nothing
// is charged.
//
// BOUNDED READS: every read carries a statement timeout (10 s with `mail_ready`, 25 s otherwise).
// A pull too broad to finish returns 503 `code: "lead_read_timeout"` telling you to narrow it (add
// `county`); nothing is charged.
//
// NOTE on `distressed`: an assessment-derived cohort (a structure on the books assessed at a
// nominal value, on land that carries real value). It is NOT a pre-foreclosure, tax-lien or
// lis-pendens feed.
//
// HTTP: GET /api/v1/leads/find
func (s *LeadsService) Find(ctx context.Context, params *LeadsFindParams, opts ...RequestOption) (*LeadsFindResponse, error) {
	var out LeadsFindResponse
	if err := s.client.do(ctx, buildLeadsFindRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildLeadsFindRequest(params *LeadsFindParams) *apiRequest {
	req := newRequest("GET", "/api/v1/leads/find")
	if params != nil {
		addQuery(req.query, "signal", params.Signal)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "county", params.County)
		addQuery(req.query, "zip", params.Zip)
		addQuery(req.query, "value_min", params.ValueMin)
		addQuery(req.query, "value_max", params.ValueMax)
		addQuery(req.query, "limit", params.Limit)
		addQuery(req.query, "preview", params.Preview)
		addQuery(req.query, "mail_ready", params.MailReady)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// LeadsFindParams holds the query, header and JSON-body parameters of [LeadsService.Find]. Pass
// nil when you need none.
type LeadsFindParams struct {
	// Which qualified cohort to pull from. Priced by strength: absentee (1.0x, widest) < long_hold
	// (1.1x) < entity_owned (1.15x) < portfolio_owner (1.25x) < high_land_ratio (1.4x) < flip (1.6x) <
	// distressed (1.9x).
	//
	// Required.
	Signal *LeadsFindParamsSignal `query:"signal" json:"-"`

	// REQUIRED 2-letter USPS state code. Every pull is pruned to one state partition.
	//
	// Required.
	State *string `query:"state" json:"-"`

	// 3-digit within-state code ("183") or 5-digit state+county ("37183").
	County *string `query:"county" json:"-"`

	// 5-digit ZIP. Supported only on the long_hold and entity_owned cohorts (the other source
	// relations carry no ZIP column) -- a zip on any other signal returns 400.
	Zip *string `query:"zip" json:"-"`

	// Minimum assessed value (sell price for the flip cohort), USD.
	ValueMin *int64 `query:"value_min" json:"-"`

	// Maximum assessed value (sell price for the flip cohort), USD.
	ValueMax *int64 `query:"value_max" json:"-"`

	// How many leads to buy. Default 25, max 200. You pay for min(matching rows, limit).
	Limit *int64 `query:"limit" json:"-"`

	// true -> the FREE preview (count + quote + up to three masked sample leads, no payment). Omit or
	// false -> the paid call.
	Preview *bool `query:"preview" json:"-"`

	// true -> only parcels whose record carries a complete mailing address (street, city, state,
	// 5-digit ZIP) in ONE column family. Parcel-grain signals only (400 on portfolio_owner). Each
	// delivered lead's `owner_contact.mail_ready` is the final word. Narrow with `county` in large
	// states: the filter checks every candidate.
	MailReady *bool `query:"mail_ready" json:"-"`

	// x402 payment, sent TOGETHER with your account credentials (Authorization: Bearer pz_...) -- a
	// payment alone is not an account and is refused with 401 before it is verified: a base64-encoded
	// signed x402 PaymentPayload (EIP-3009 transferWithAuthorization over USDC on Base). The signed
	// amount must equal this pull's quoted maxAmountRequired (see the 402 body, or call with
	// preview=true first). Ignored on a preview call, which is free.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// LeadsFindParamsSignal is generated from the OpenAPI spec. It is a string; the
// LeadsFindParamsSignal* constants list the documented values.
type LeadsFindParamsSignal = string

// Documented values of LeadsFindParamsSignal.
const (
	LeadsFindParamsSignalAbsentee       LeadsFindParamsSignal = "absentee"
	LeadsFindParamsSignalLongHold       LeadsFindParamsSignal = "long_hold"
	LeadsFindParamsSignalEntityOwned    LeadsFindParamsSignal = "entity_owned"
	LeadsFindParamsSignalPortfolioOwner LeadsFindParamsSignal = "portfolio_owner"
	LeadsFindParamsSignalHighLandRatio  LeadsFindParamsSignal = "high_land_ratio"
	LeadsFindParamsSignalFlip           LeadsFindParamsSignal = "flip"
	LeadsFindParamsSignalDistressed     LeadsFindParamsSignal = "distressed"
)

// LeadsFindResponse: Lead feed (paid, priced per lead) — with a FREE preview
type LeadsFindResponse = json.RawMessage
