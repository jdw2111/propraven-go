// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
)

// IntelligenceService groups the intelligence operations. Use it as client.Intelligence.
type IntelligenceService struct {
	client *Client
}

// Signals: Get evidence-backed property signals
//
// Requires API-key or first-party session authentication and a current account in the default-off
// server cohort. Valid current membership does not require a new paid subscription. Existing API
// quotas still apply. Responses are private, no-store. Unknown and repeated query parameters are
// rejected. This contract does not indicate source activation, deployment or an SDK release.
// Retained objects are scoped to the current account and their creator; no caller-supplied
// account/user grant is accepted. Current source rights are checked for every read/use,
// independently of archived grants. Assessment composition, where qualified, is not market value.
// Other groups remain explicitly unavailable without approved adapters. Exact values are
// numerator/denominator strings; historical runs exclude later-learned evidence.
//
// HTTP: GET /api/v1/parcels/{id}/signals
func (s *IntelligenceService) Signals(ctx context.Context, id string, params *IntelligenceSignalsParams, opts ...RequestOption) (*IntelligenceSignalsResponse, error) {
	var out IntelligenceSignalsResponse
	if err := s.client.do(ctx, buildIntelligenceSignalsRequest(id, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildIntelligenceSignalsRequest(id string, params *IntelligenceSignalsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/parcels/"+pathParam(id)+"/signals")
	if params != nil {
		addQuery(req.query, "as_of", params.AsOf)
		addQuery(req.query, "knowledge_cutoff", params.KnowledgeCutoff)
		addQuery(req.query, "use", params.Use)
	}
	return req
}

// IntelligenceSignalsParams holds the query, header and JSON-body parameters of
// [IntelligenceService.Signals]. Pass nil when you need none.
type IntelligenceSignalsParams struct {
	// Effective-time cutoff; explicit UTC timestamp. Defaults to the service capture clock. Future
	// values rejected.
	AsOf *IntelligenceInstant `query:"as_of" json:"-"`

	// Only observations known by this time contribute. Defaults to the same capture clock; original
	// knowledge is never backdated from a publication or vintage.
	KnowledgeCutoff *IntelligenceInstant `query:"knowledge_cutoff" json:"-"`

	// Requested use; checked against current source rights. Agent use is not an external send.
	Use *IntelligenceSignalsParamsUse `query:"use" json:"-"`
}

// IntelligenceSignalsParamsUse is generated from the OpenAPI spec. It is a string; the
// IntelligenceSignalsParamsUse* constants list the documented values.
type IntelligenceSignalsParamsUse = string

// Documented values of IntelligenceSignalsParamsUse.
const (
	IntelligenceSignalsParamsUseDisplay IntelligenceSignalsParamsUse = "display"
	IntelligenceSignalsParamsUseAi      IntelligenceSignalsParamsUse = "ai"
)

// IntelligenceSignalsResponse: Get evidence-backed property signals
type IntelligenceSignalsResponse = IntelligenceRun

// Run: Read an owned retained run and evidence
//
// Requires API-key or first-party session authentication and a current account in the default-off
// server cohort. Valid current membership does not require a new paid subscription. Existing API
// quotas still apply. Responses are private, no-store. Unknown and repeated query parameters are
// rejected. This contract does not indicate source activation, deployment or an SDK release.
// Retained objects are scoped to the current account and their creator; no caller-supplied
// account/user grant is accepted. Current source rights are checked for every read/use,
// independently of archived grants. Withdrawn rights or changed source semantics may withhold a
// saved result; its immutable archived payload is not rewritten.
//
// HTTP: GET /api/v1/intelligence/runs/{runId}
func (s *IntelligenceService) Run(ctx context.Context, runID string, params *IntelligenceRunParams, opts ...RequestOption) (*IntelligenceRunResponse, error) {
	var out IntelligenceRunResponse
	if err := s.client.do(ctx, buildIntelligenceRunRequest(runID, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildIntelligenceRunRequest(runID string, params *IntelligenceRunParams) *apiRequest {
	req := newRequest("GET", "/api/v1/intelligence/runs/"+pathParam(runID))
	if params != nil {
		addQuery(req.query, "use", params.Use)
	}
	return req
}

// IntelligenceRunParams holds the query, header and JSON-body parameters of
// [IntelligenceService.Run]. Pass nil when you need none.
type IntelligenceRunParams struct {
	// Requested use; checked against current source rights. Agent use is not an external send.
	Use *IntelligenceRunParamsUse `query:"use" json:"-"`
}

// IntelligenceRunParamsUse is generated from the OpenAPI spec. It is a string; the
// IntelligenceRunParamsUse* constants list the documented values.
type IntelligenceRunParamsUse = string

// Documented values of IntelligenceRunParamsUse.
const (
	IntelligenceRunParamsUseDisplay IntelligenceRunParamsUse = "display"
	IntelligenceRunParamsUseExport  IntelligenceRunParamsUse = "export"
	IntelligenceRunParamsUseAi      IntelligenceRunParamsUse = "ai"
)

// IntelligenceRunResponse: Read an owned retained run and evidence
type IntelligenceRunResponse = IntelligenceRunDetail

// CreateScenario: Save an explicit named residual scenario
//
// Requires API-key or first-party session authentication and a current account in the default-off
// server cohort. Valid current membership does not require a new paid subscription. Existing API
// quotas still apply. Responses are private, no-store. No query parameters are used by this
// operation; its inputs are the strict JSON body. This contract does not indicate source
// activation, deployment or an SDK release. Retained objects are scoped to the current account and
// their creator; no caller-supplied account/user grant is accepted. Current source rights are
// checked for every read/use, independently of archived grants. Creates an immutable
// base/downside/upside revision linked to an owned readable run. A parent revision must belong to
// the same run and creator. Formula uses fixed-dollar profit and purchase-independent carry; no
// live values/default costs are invented. Negative residuals remain explicit. The JSON object is
// strict, bounded to 16,384 bytes, and costs must contain each of the five buckets once.
//
// HTTP: POST /api/v1/intelligence/scenarios
func (s *IntelligenceService) CreateScenario(ctx context.Context, params *IntelligenceCreateScenarioParams, opts ...RequestOption) (*IntelligenceCreateScenarioResponse, error) {
	var out IntelligenceCreateScenarioResponse
	if err := s.client.do(ctx, buildIntelligenceCreateScenarioRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildIntelligenceCreateScenarioRequest(params *IntelligenceCreateScenarioParams) *apiRequest {
	req := newRequest("POST", "/api/v1/intelligence/scenarios")
	if params != nil {
		req.body = params
	} else {
		req.body = struct{}{}
	}
	return req
}

// IntelligenceCreateScenarioParams holds the query, header and JSON-body parameters of
// [IntelligenceService.CreateScenario]. Pass nil when you need none.
type IntelligenceCreateScenarioParams struct {
	// Required (JSON body).
	RunID string `json:"run_id"`

	// Required (JSON body).
	Label            IntelligenceCreateScenarioParamsLabel `json:"label"`
	ParentRevisionID *string                               `json:"parent_revision_id,omitempty"`

	// Required (JSON body).
	Assumptions IntelligenceCreateScenarioParamsAssumptions `json:"assumptions"`
}

// IntelligenceCreateScenarioParamsLabel is generated from the OpenAPI spec. It is a string; the
// IntelligenceCreateScenarioParamsLabel* constants list the documented values.
type IntelligenceCreateScenarioParamsLabel = string

// Documented values of IntelligenceCreateScenarioParamsLabel.
const (
	IntelligenceCreateScenarioParamsLabelBase     IntelligenceCreateScenarioParamsLabel = "base"
	IntelligenceCreateScenarioParamsLabelDownside IntelligenceCreateScenarioParamsLabel = "downside"
	IntelligenceCreateScenarioParamsLabelUpside   IntelligenceCreateScenarioParamsLabel = "upside"
)

// IntelligenceCreateScenarioParamsAssumptions is generated from the OpenAPI spec.
type IntelligenceCreateScenarioParamsAssumptions struct {
	Currency              string                                             `json:"currency"`
	GrossCompletedSale    string                                             `json:"gross_completed_sale"`
	SellingCosts          string                                             `json:"selling_costs"`
	Costs                 []IntelligenceCreateScenarioParamsAssumptionsCosts `json:"costs"`
	RequiredProfitDollars string                                             `json:"required_profit_dollars"`
	FixedAcquisitionCosts string                                             `json:"fixed_acquisition_costs"`
	AcquisitionCostRate   string                                             `json:"acquisition_cost_rate"`
	ProfitMode            string                                             `json:"profit_mode"`
	CarryMode             string                                             `json:"carry_mode"`
	InputSource           string                                             `json:"input_source"`
}

// IntelligenceCreateScenarioParamsAssumptionsCosts is generated from the OpenAPI spec.
type IntelligenceCreateScenarioParamsAssumptionsCosts struct {
	Bucket   IntelligenceCreateScenarioParamsAssumptionsCostsBucket `json:"bucket"`
	Amount   string                                                 `json:"amount"`
	Currency string                                                 `json:"currency"`
}

// IntelligenceCreateScenarioParamsAssumptionsCostsBucket is generated from the OpenAPI spec. It is
// a string; the IntelligenceCreateScenarioParamsAssumptionsCostsBucket* constants list the
// documented values.
type IntelligenceCreateScenarioParamsAssumptionsCostsBucket = string

// Documented values of IntelligenceCreateScenarioParamsAssumptionsCostsBucket.
const (
	IntelligenceCreateScenarioParamsAssumptionsCostsBucketHard         IntelligenceCreateScenarioParamsAssumptionsCostsBucket = "hard"
	IntelligenceCreateScenarioParamsAssumptionsCostsBucketSoft         IntelligenceCreateScenarioParamsAssumptionsCostsBucket = "soft"
	IntelligenceCreateScenarioParamsAssumptionsCostsBucketContingency  IntelligenceCreateScenarioParamsAssumptionsCostsBucket = "contingency"
	IntelligenceCreateScenarioParamsAssumptionsCostsBucketCarry        IntelligenceCreateScenarioParamsAssumptionsCostsBucket = "carry"
	IntelligenceCreateScenarioParamsAssumptionsCostsBucketOtherNonland IntelligenceCreateScenarioParamsAssumptionsCostsBucket = "other_nonland"
)

// IntelligenceCreateScenarioResponse: Save an explicit named residual scenario
type IntelligenceCreateScenarioResponse = IntelligenceScenarioRevision

// Handoff: Prepare an owned structured investigation handoff
//
// Requires API-key or first-party session authentication and a current account in the default-off
// server cohort. Valid current membership does not require a new paid subscription. Existing API
// quotas still apply. Responses are private, no-store. Unknown and repeated query parameters are
// rejected. This contract does not indicate source activation, deployment or an SDK release.
// Retained objects are scoped to the current account and their creator; no caller-supplied
// account/user grant is accepted. Current source rights are checked for every read/use,
// independently of archived grants. Returns the same retained calculations, evidence and optional
// saved scenario, with observations separated from user assumptions. The scenario must belong to
// the same run/creator. Delivery is structured_only_not_sent; no agent/model is started.
//
// HTTP: GET /api/v1/intelligence/runs/{runId}/handoff
func (s *IntelligenceService) Handoff(ctx context.Context, runID string, params *IntelligenceHandoffParams, opts ...RequestOption) (*IntelligenceHandoffResponse, error) {
	var out IntelligenceHandoffResponse
	if err := s.client.do(ctx, buildIntelligenceHandoffRequest(runID, params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildIntelligenceHandoffRequest(runID string, params *IntelligenceHandoffParams) *apiRequest {
	req := newRequest("GET", "/api/v1/intelligence/runs/"+pathParam(runID)+"/handoff")
	if params != nil {
		addQuery(req.query, "use", params.Use)
		addQuery(req.query, "scenario_id", params.ScenarioID)
	}
	return req
}

// IntelligenceHandoffParams holds the query, header and JSON-body parameters of
// [IntelligenceService.Handoff]. Pass nil when you need none.
type IntelligenceHandoffParams struct {
	// Requested use; checked against current source rights. Agent use is not an external send.
	Use *IntelligenceHandoffParamsUse `query:"use" json:"-"`

	// Optional saved scenario revision belonging to the same run and creator.
	ScenarioID *IntelligenceRetainedID `query:"scenario_id" json:"-"`
}

// IntelligenceHandoffParamsUse is generated from the OpenAPI spec. It is a string; the
// IntelligenceHandoffParamsUse* constants list the documented values.
type IntelligenceHandoffParamsUse = string

// Documented values of IntelligenceHandoffParamsUse.
const (
	IntelligenceHandoffParamsUseDisplay IntelligenceHandoffParamsUse = "display"
	IntelligenceHandoffParamsUseExport  IntelligenceHandoffParamsUse = "export"
	IntelligenceHandoffParamsUseAi      IntelligenceHandoffParamsUse = "ai"
)

// IntelligenceHandoffResponse: Prepare an owned structured investigation handoff
type IntelligenceHandoffResponse = IntelligenceHandoff
