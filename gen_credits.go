// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
)

// CreditsService groups the credits operations. Use it as client.Credits.
type CreditsService struct {
	client *Client
}

// Topup: Fund a prepaid credit balance over x402
//
// Pay `amount` USDC once via x402 to fund a prepaid credit balance; receive a credit token
// (pzc_...) to spend on any paid endpoint via the X-CREDIT-TOKEN header — the recurring/volume
// rail, no account, no CDP/Stripe. GET with no payment returns a 402 for `amount`; present an
// existing X-CREDIT-TOKEN to top it up in place. Idempotent on the settlement tx.
//
// HTTP: GET /api/v1/storefront/credits/topup
func (s *CreditsService) Topup(ctx context.Context, params *CreditsTopupParams, opts ...RequestOption) (*CreditsTopupResponse, error) {
	var out CreditsTopupResponse
	if err := s.client.do(ctx, buildCreditsTopupRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCreditsTopupRequest(params *CreditsTopupParams) *apiRequest {
	req := newRequest("GET", "/api/v1/storefront/credits/topup")
	if params != nil {
		addQuery(req.query, "amount", params.Amount)
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
		setHeader(req.header, "X-PAYMENT", params.Payment)
	}
	return req
}

// CreditsTopupParams holds the query, header and JSON-body parameters of [CreditsService.Topup].
// Pass nil when you need none.
type CreditsTopupParams struct {
	// Whole USD to fund ($1–$1000).
	//
	// Required.
	Amount *int64 `query:"amount" json:"-"`

	// Optional existing pzc_ token to top up in place.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`

	// Base64 x402 PaymentPayload to fund the balance.
	Payment *string `header:"X-PAYMENT" json:"-"`
}

// CreditsTopupResponse: Fund a prepaid credit balance over x402
type CreditsTopupResponse = map[string]any

// Balance: Read a prepaid credit balance + ledger
//
// Return the balance and recent ledger for the credit token in the X-CREDIT-TOKEN header. No
// payment; the token is the credential (never a query param).
//
// HTTP: GET /api/v1/storefront/credits/balance
func (s *CreditsService) Balance(ctx context.Context, params *CreditsBalanceParams, opts ...RequestOption) (*CreditsBalanceResponse, error) {
	var out CreditsBalanceResponse
	if err := s.client.do(ctx, buildCreditsBalanceRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCreditsBalanceRequest(params *CreditsBalanceParams) *apiRequest {
	req := newRequest("GET", "/api/v1/storefront/credits/balance")
	if params != nil {
		setHeader(req.header, "X-CREDIT-TOKEN", params.CreditToken)
	}
	return req
}

// CreditsBalanceParams holds the query, header and JSON-body parameters of
// [CreditsService.Balance]. Pass nil when you need none.
type CreditsBalanceParams struct {
	// The pzc_ credit token.
	//
	// Required.
	CreditToken *string `header:"X-CREDIT-TOKEN" json:"-"`
}

// CreditsBalanceResponse: Read a prepaid credit balance + ledger
type CreditsBalanceResponse = map[string]any
