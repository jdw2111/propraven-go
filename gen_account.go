// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
)

// AccountService groups the account operations. Use it as client.Account.
type AccountService struct {
	client *Client
}

// Usage: Current-period usage and quota
//
// Returns the calling key's current-period API usage, included allotment, remaining calls,
// configured per-minute and per-day rate limits, and the hard-cap status. Per-user (all of a
// user's API keys roll up to the same monthly counter, since they share a Stripe subscription).
//
// HTTP: GET /api/v1/account/usage
func (s *AccountService) Usage(ctx context.Context, params *AccountUsageParams, opts ...RequestOption) (*AccountUsageResponse, error) {
	var out AccountUsageResponse
	if err := s.client.do(ctx, buildAccountUsageRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildAccountUsageRequest(params *AccountUsageParams) *apiRequest {
	req := newRequest("GET", "/api/v1/account/usage")
	return req
}

// AccountUsageParams holds the query, header and JSON-body parameters of [AccountService.Usage].
// Pass nil when you need none.
type AccountUsageParams struct {
}

// AccountUsageResponse: Current-period usage and quota
type AccountUsageResponse = AccountUsage
