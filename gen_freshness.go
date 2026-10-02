// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// FreshnessService groups the freshness operations. Use it as client.Freshness.
type FreshnessService struct {
	client *Client
}

// Get: How fresh the served parcel snapshot is
//
// Build and swap times of the served snapshot, plus a source-registry freshness proxy
// (`content_*`). Free; anonymous callers are IP-throttled.
//
// HTTP: GET /api/v1/freshness
func (s *FreshnessService) Get(ctx context.Context, params *FreshnessGetParams, opts ...RequestOption) (*FreshnessGetResponse, error) {
	var out FreshnessGetResponse
	if err := s.client.do(ctx, buildFreshnessGetRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildFreshnessGetRequest(params *FreshnessGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/freshness")
	return req
}

// FreshnessGetParams holds the query, header and JSON-body parameters of [FreshnessService.Get].
// Pass nil when you need none.
type FreshnessGetParams struct {
}

// FreshnessGetResponse: How fresh the served parcel snapshot is
type FreshnessGetResponse struct {
	ContentAsOf              string `json:"content_as_of"`
	LastEnrichedAt           string `json:"last_enriched_at"`
	ContentDateBasis         string `json:"content_date_basis"`
	ContentDateNote          string `json:"content_date_note"`
	ContentMedianDate        string `json:"content_median_date"`
	ContentMaxDate           string `json:"content_max_date"`
	ContentOldestDate        string `json:"content_oldest_date"`
	ContentSourceCount       int64  `json:"content_source_count"`
	ContentActiveSourceCount int64  `json:"content_active_source_count"`
	SwappedAt                string `json:"swapped_at"`
	UpdatedAt                string `json:"updated_at"`
	SnapshotBuiltAt          string `json:"snapshot_built_at"`
	ParcelCount              int64  `json:"parcel_count"`
}

// UnmarshalJSON decodes FreshnessGetResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FreshnessGetResponse) UnmarshalJSON(data []byte) error {
	type plain FreshnessGetResponse
	aux := struct {
		*plain
		ContentSourceCount       lenientNumber[int64] `json:"content_source_count"`
		ContentActiveSourceCount lenientNumber[int64] `json:"content_active_source_count"`
		ParcelCount              lenientNumber[int64] `json:"parcel_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ContentSourceCount.assign(&r.ContentSourceCount)
	aux.ContentActiveSourceCount.assign(&r.ContentActiveSourceCount)
	aux.ParcelCount.assign(&r.ParcelCount)
	return softTypeError(err)
}

// Datasets: Per-dataset availability and freshness
//
// HTTP: GET /api/v1/freshness/datasets
func (s *FreshnessService) Datasets(ctx context.Context, params *FreshnessDatasetsParams, opts ...RequestOption) (*FreshnessDatasetsResponse, error) {
	var out FreshnessDatasetsResponse
	if err := s.client.do(ctx, buildFreshnessDatasetsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildFreshnessDatasetsRequest(params *FreshnessDatasetsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/freshness/datasets")
	return req
}

// FreshnessDatasetsParams holds the query, header and JSON-body parameters of
// [FreshnessService.Datasets]. Pass nil when you need none.
type FreshnessDatasetsParams struct {
}

// FreshnessDatasetsResponse: Per-dataset availability and freshness
type FreshnessDatasetsResponse struct {
	ContractVersion float64                             `json:"contract_version"`
	GeneratedAt     string                              `json:"generated_at"`
	Status          string                              `json:"status"`
	Datasets        []FreshnessDatasetsResponseDatasets `json:"datasets"`
}

// UnmarshalJSON decodes FreshnessDatasetsResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FreshnessDatasetsResponse) UnmarshalJSON(data []byte) error {
	type plain FreshnessDatasetsResponse
	aux := struct {
		*plain
		ContractVersion lenientNumber[float64] `json:"contract_version"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ContractVersion.assign(&r.ContractVersion)
	return softTypeError(err)
}

// FreshnessDatasetsResponseDatasets is generated from the OpenAPI spec.
type FreshnessDatasetsResponseDatasets struct {
	Dataset         *string                                         `json:"dataset,omitempty"`
	Availability    *string                                         `json:"availability,omitempty"`
	Freshness       *string                                         `json:"freshness,omitempty"`
	FreshnessReason *string                                         `json:"freshness_reason,omitempty"`
	Refresh         FreshnessDatasetsResponseDatasetsRefresh        `json:"refresh"`
	RecordActivity  FreshnessDatasetsResponseDatasetsRecordActivity `json:"record_activity"`
	Coverage        FreshnessDatasetsResponseDatasetsCoverage       `json:"coverage"`
}

// FreshnessDatasetsResponseDatasetsRefresh is generated from the OpenAPI spec.
type FreshnessDatasetsResponseDatasetsRefresh struct {
	SourceWatermark *string  `json:"source_watermark,omitempty"`
	SourceAgeHours  *float64 `json:"source_age_hours,omitempty"`
	EvaluatedAt     *string  `json:"evaluated_at,omitempty"`
	PublishedAt     *string  `json:"published_at,omitempty"`
	Status          *string  `json:"status,omitempty"`
	WarningReason   *string  `json:"warning_reason,omitempty"`
}

// UnmarshalJSON decodes FreshnessDatasetsResponseDatasetsRefresh, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FreshnessDatasetsResponseDatasetsRefresh) UnmarshalJSON(data []byte) error {
	type plain FreshnessDatasetsResponseDatasetsRefresh
	aux := struct {
		*plain
		SourceAgeHours lenientNumber[float64] `json:"source_age_hours"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SourceAgeHours.assignPtr(&r.SourceAgeHours)
	return softTypeError(err)
}

// FreshnessDatasetsResponseDatasetsRecordActivity is generated from the OpenAPI spec.
type FreshnessDatasetsResponseDatasetsRecordActivity struct {
	LatestAt    *string  `json:"latest_at,omitempty"`
	EvaluatedAt *string  `json:"evaluated_at,omitempty"`
	Basis       *string  `json:"basis,omitempty"`
	AgeHours    *float64 `json:"age_hours,omitempty"`
	MaxAgeHours *float64 `json:"max_age_hours,omitempty"`
	Status      *string  `json:"status,omitempty"`
}

// UnmarshalJSON decodes FreshnessDatasetsResponseDatasetsRecordActivity, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FreshnessDatasetsResponseDatasetsRecordActivity) UnmarshalJSON(data []byte) error {
	type plain FreshnessDatasetsResponseDatasetsRecordActivity
	aux := struct {
		*plain
		AgeHours    lenientNumber[float64] `json:"age_hours"`
		MaxAgeHours lenientNumber[float64] `json:"max_age_hours"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AgeHours.assignPtr(&r.AgeHours)
	aux.MaxAgeHours.assignPtr(&r.MaxAgeHours)
	return softTypeError(err)
}

// FreshnessDatasetsResponseDatasetsCoverage is generated from the OpenAPI spec.
type FreshnessDatasetsResponseDatasetsCoverage struct {
	Status  *string  `json:"status,omitempty"`
	Unit    *string  `json:"unit,omitempty"`
	Covered *float64 `json:"covered,omitempty"`
	AsOf    *string  `json:"as_of,omitempty"`
}

// UnmarshalJSON decodes FreshnessDatasetsResponseDatasetsCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FreshnessDatasetsResponseDatasetsCoverage) UnmarshalJSON(data []byte) error {
	type plain FreshnessDatasetsResponseDatasetsCoverage
	aux := struct {
		*plain
		Covered lenientNumber[float64] `json:"covered"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Covered.assignPtr(&r.Covered)
	return softTypeError(err)
}
