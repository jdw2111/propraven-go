// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/jdw2111/propraven-go/internal/apijson"
	"github.com/jdw2111/propraven-go/internal/apiquery"
	"github.com/jdw2111/propraven-go/internal/requestconfig"
	"github.com/jdw2111/propraven-go/option"
	"github.com/jdw2111/propraven-go/packages/param"
	"github.com/jdw2111/propraven-go/packages/respjson"
)

// Owner search, profiles, and portfolios.
//
// V1OwnerService contains methods and other services that help with interacting
// with the propraven API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OwnerService] method instead.
type V1OwnerService struct {
	options []option.RequestOption
}

// NewV1OwnerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1OwnerService(opts ...option.RequestOption) (r V1OwnerService) {
	r = V1OwnerService{}
	r.options = opts
	return
}

// Retrieve an owner's portfolio with aggregated summary statistics and property
// breakdown.
func (r *V1OwnerService) GetPortfolioSummary(ctx context.Context, name string, opts ...option.RequestOption) (res *V1OwnerGetPortfolioSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/owners/%s/portfolio", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a specific owner profile by name, including property count, total
// assessed value, entity type, and states.
func (r *V1OwnerService) GetProfile(ctx context.Context, name string, opts ...option.RequestOption) (res *Owner, err error) {
	opts = slices.Concat(r.options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/owners/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the list of properties owned by a specific owner.
func (r *V1OwnerService) GetProperties(ctx context.Context, name string, query V1OwnerGetPropertiesParams, opts ...option.RequestOption) (res *V1OwnerGetPropertiesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/owners/%s/properties", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns up to 100 most-recent deed events where the named owner is either
// grantor or grantee. Useful for building an owner's transaction timeline across
// their portfolio.
func (r *V1OwnerService) GetTransactions(ctx context.Context, name string, opts ...option.RequestOption) (res *V1OwnerGetTransactionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/owners/%s/transactions", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search for property owners by name. Returns owner profiles with property counts
// and portfolio values.
func (r *V1OwnerService) SearchOwners(ctx context.Context, query V1OwnerSearchOwnersParams, opts ...option.RequestOption) (res *V1OwnerSearchOwnersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/owners/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Owner struct {
	// Any of "individual", "corporation", "llc", "trust", "government", "other".
	EntityType         OwnerEntityType `json:"entity_type"`
	OwnerName          string          `json:"owner_name"`
	PropertyCount      int64           `json:"property_count"`
	States             []string        `json:"states"`
	TotalAssessedValue float64         `json:"total_assessed_value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityType         respjson.Field
		OwnerName          respjson.Field
		PropertyCount      respjson.Field
		States             respjson.Field
		TotalAssessedValue respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Owner) RawJSON() string { return r.JSON.raw }
func (r *Owner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OwnerEntityType string

const (
	OwnerEntityTypeIndividual  OwnerEntityType = "individual"
	OwnerEntityTypeCorporation OwnerEntityType = "corporation"
	OwnerEntityTypeLlc         OwnerEntityType = "llc"
	OwnerEntityTypeTrust       OwnerEntityType = "trust"
	OwnerEntityTypeGovernment  OwnerEntityType = "government"
	OwnerEntityTypeOther       OwnerEntityType = "other"
)

type V1OwnerGetPortfolioSummaryResponse struct {
	// Any of "individual", "corporation", "llc", "trust", "government", "other".
	EntityType V1OwnerGetPortfolioSummaryResponseEntityType `json:"entity_type"`
	OwnerName  string                                       `json:"owner_name"`
	Properties []Parcel                                     `json:"properties"`
	Summary    V1OwnerGetPortfolioSummaryResponseSummary    `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityType  respjson.Field
		OwnerName   respjson.Field
		Properties  respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerGetPortfolioSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerGetPortfolioSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerGetPortfolioSummaryResponseEntityType string

const (
	V1OwnerGetPortfolioSummaryResponseEntityTypeIndividual  V1OwnerGetPortfolioSummaryResponseEntityType = "individual"
	V1OwnerGetPortfolioSummaryResponseEntityTypeCorporation V1OwnerGetPortfolioSummaryResponseEntityType = "corporation"
	V1OwnerGetPortfolioSummaryResponseEntityTypeLlc         V1OwnerGetPortfolioSummaryResponseEntityType = "llc"
	V1OwnerGetPortfolioSummaryResponseEntityTypeTrust       V1OwnerGetPortfolioSummaryResponseEntityType = "trust"
	V1OwnerGetPortfolioSummaryResponseEntityTypeGovernment  V1OwnerGetPortfolioSummaryResponseEntityType = "government"
	V1OwnerGetPortfolioSummaryResponseEntityTypeOther       V1OwnerGetPortfolioSummaryResponseEntityType = "other"
)

type V1OwnerGetPortfolioSummaryResponseSummary struct {
	AvgAssessedValue   float64          `json:"avg_assessed_value"`
	Counties           int64            `json:"counties"`
	PropertyCount      int64            `json:"property_count"`
	States             []string         `json:"states"`
	TotalAssessedValue float64          `json:"total_assessed_value"`
	ZoningBreakdown    map[string]int64 `json:"zoning_breakdown"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgAssessedValue   respjson.Field
		Counties           respjson.Field
		PropertyCount      respjson.Field
		States             respjson.Field
		TotalAssessedValue respjson.Field
		ZoningBreakdown    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerGetPortfolioSummaryResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerGetPortfolioSummaryResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerGetPropertiesResponse struct {
	Data   []Parcel `json:"data"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
	Total  int64    `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerGetPropertiesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerGetPropertiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerGetTransactionsResponse struct {
	Count int64                                `json:"count"`
	Data  []V1OwnerGetTransactionsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerGetTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerGetTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerGetTransactionsResponseData struct {
	DocumentNumber string `json:"document_number"`
	// Recorded document type (Warranty Deed, Quit Claim, etc.).
	DocumentType    string    `json:"document_type" api:"nullable"`
	GranteeName     string    `json:"grantee_name" api:"nullable"`
	GrantorName     string    `json:"grantor_name" api:"nullable"`
	PropertyAddress string    `json:"property_address" api:"nullable"`
	RecordingDate   time.Time `json:"recording_date" format:"date"`
	SaleDate        time.Time `json:"sale_date" format:"date"`
	// USD. Null when state is non-disclosure.
	SalePrice float64 `json:"sale_price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentNumber  respjson.Field
		DocumentType    respjson.Field
		GranteeName     respjson.Field
		GrantorName     respjson.Field
		PropertyAddress respjson.Field
		RecordingDate   respjson.Field
		SaleDate        respjson.Field
		SalePrice       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerGetTransactionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerGetTransactionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerSearchOwnersResponse struct {
	Data []Owner `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OwnerSearchOwnersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OwnerSearchOwnersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OwnerGetPropertiesParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OwnerGetPropertiesParams]'s query parameters as
// `url.Values`.
func (r V1OwnerGetPropertiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OwnerSearchOwnersParams struct {
	// Search query for owner name.
	Q     string           `query:"q" api:"required" json:"-"`
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Minimum number of properties owned.
	MinProperties param.Opt[int64] `query:"min_properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OwnerSearchOwnersParams]'s query parameters as
// `url.Values`.
func (r V1OwnerSearchOwnersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
