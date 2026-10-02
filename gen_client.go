// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

// Client is the PropRaven API client. Create it with [NewClient]; each
// field groups the operations of one API namespace (x-sdk-group).
type Client struct {
	// Account holds the account operations.
	Account *AccountService
	// CMBS holds the cmbs operations.
	CMBS *CMBSService
	// Cohorts holds the cohorts operations.
	Cohorts *CohortsService
	// Coverage holds the coverage operations.
	Coverage *CoverageService
	// Credits holds the credits operations.
	Credits *CreditsService
	// Crime holds the crime operations.
	Crime *CrimeService
	// Deals holds the deals operations.
	Deals *DealsService
	// Freshness holds the freshness operations.
	Freshness *FreshnessService
	// Leads holds the leads operations.
	Leads *LeadsService
	// Lookup holds the lookup operations.
	Lookup *LookupService
	// Market holds the market operations.
	Market *MarketService
	// Owners holds the owners operations.
	Owners *OwnersService
	// Parcels holds the parcels operations.
	Parcels *ParcelsService
	// Search holds the search operations.
	Search *SearchService
	// Storefront holds the storefront operations.
	Storefront *StorefrontService
	// Traffic holds the traffic operations.
	Traffic *TrafficService
	// Verify holds the verify operations.
	Verify *VerifyService
	// Watch holds the watch operations.
	Watch *WatchService
	// Webhooks holds the webhooks operations.
	Webhooks *WebhooksService

	core *clientCore
}

func (c *Client) initServices() {
	c.Account = &AccountService{client: c}
	c.CMBS = &CMBSService{client: c}
	c.Cohorts = &CohortsService{client: c}
	c.Coverage = &CoverageService{client: c}
	c.Credits = &CreditsService{client: c}
	c.Crime = &CrimeService{client: c}
	c.Deals = &DealsService{client: c}
	c.Freshness = &FreshnessService{client: c}
	c.Leads = &LeadsService{client: c}
	c.Lookup = &LookupService{client: c}
	c.Market = &MarketService{client: c}
	c.Owners = &OwnersService{client: c}
	c.Parcels = &ParcelsService{client: c}
	c.Search = &SearchService{client: c}
	c.Storefront = &StorefrontService{client: c}
	c.Traffic = &TrafficService{client: c}
	c.Verify = &VerifyService{client: c}
	c.Watch = &WatchService{client: c}
	c.Webhooks = &WebhooksService{client: c}
}
