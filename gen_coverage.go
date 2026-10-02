// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// CoverageService groups the coverage operations. Use it as client.Coverage.
type CoverageService struct {
	client *Client
}

// Get: Get coverage statistics
//
// Retrieve parcel coverage statistics at the state or county level. API key optional: anonymous
// callers are rate-limited per IP; a present but invalid key is a 401; a valid key is rate-limited
// at its own tier and is not metered.
//
// HTTP: GET /api/v1/coverage
func (s *CoverageService) Get(ctx context.Context, params *CoverageGetParams, opts ...RequestOption) (*CoverageGetResponse, error) {
	var out CoverageGetResponse
	if err := s.client.do(ctx, buildCoverageGetRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCoverageGetRequest(params *CoverageGetParams) *apiRequest {
	req := newRequest("GET", "/api/v1/coverage")
	if params != nil {
		addQuery(req.query, "state", params.State)
	}
	return req
}

// CoverageGetParams holds the query, header and JSON-body parameters of [CoverageService.Get].
// Pass nil when you need none.
type CoverageGetParams struct {
	// State FIPS code or abbreviation to filter coverage to a specific state and return county-level
	// breakdown.
	State *string `query:"state" json:"-"`
}

// CoverageGetResponse: Get coverage statistics
type CoverageGetResponse struct {
	Data                  []CoverageGetResponseData          `json:"data"`
	State                 *string                            `json:"state,omitempty"`
	TotalCounties         *int64                             `json:"total_counties,omitempty"`
	TotalCountiesUs       *int64                             `json:"total_counties_us,omitempty"`
	TotalPopulationUs     *int64                             `json:"total_population_us,omitempty"`
	CoveredPopulationUs   *float64                           `json:"covered_population_us,omitempty"`
	PopulationCoveragePct *float64                           `json:"population_coverage_pct,omitempty"`
	NationalCounts        *CoverageGetResponseNationalCounts `json:"national_counts,omitempty"`
	TotalParcels          *int64                             `json:"total_parcels,omitempty"`
	StatesCovered         *int64                             `json:"states_covered,omitempty"`
}

// UnmarshalJSON decodes CoverageGetResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CoverageGetResponse) UnmarshalJSON(data []byte) error {
	type plain CoverageGetResponse
	aux := struct {
		*plain
		TotalCounties         lenientNumber[int64]   `json:"total_counties"`
		TotalCountiesUs       lenientNumber[int64]   `json:"total_counties_us"`
		TotalPopulationUs     lenientNumber[int64]   `json:"total_population_us"`
		CoveredPopulationUs   lenientNumber[float64] `json:"covered_population_us"`
		PopulationCoveragePct lenientNumber[float64] `json:"population_coverage_pct"`
		TotalParcels          lenientNumber[int64]   `json:"total_parcels"`
		StatesCovered         lenientNumber[int64]   `json:"states_covered"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalCounties.assignPtr(&r.TotalCounties)
	aux.TotalCountiesUs.assignPtr(&r.TotalCountiesUs)
	aux.TotalPopulationUs.assignPtr(&r.TotalPopulationUs)
	aux.CoveredPopulationUs.assignPtr(&r.CoveredPopulationUs)
	aux.PopulationCoveragePct.assignPtr(&r.PopulationCoveragePct)
	aux.TotalParcels.assignPtr(&r.TotalParcels)
	aux.StatesCovered.assignPtr(&r.StatesCovered)
	return softTypeError(err)
}

// CoverageGetResponseData is generated from the OpenAPI spec.
type CoverageGetResponseData struct {
	StateFIPS             string   `json:"state_fips"`
	CountyFIPS            *string  `json:"county_fips,omitempty"`
	CountyName            *string  `json:"county_name,omitempty"`
	ParcelCount           *int64   `json:"parcel_count,omitempty"`
	WithAddress           *float64 `json:"with_address,omitempty"`
	WithGeocode           *float64 `json:"with_geocode,omitempty"`
	WithOwner             *float64 `json:"with_owner,omitempty"`
	WithValue             *float64 `json:"with_value,omitempty"`
	LastUpdated           *string  `json:"last_updated,omitempty"`
	Population            *float64 `json:"population,omitempty"`
	WithGeometry          *float64 `json:"with_geometry,omitempty"`
	StateAbbr             *string  `json:"state_abbr,omitempty"`
	CountyCount           *int64   `json:"county_count,omitempty"`
	TotalCounties         *int64   `json:"total_counties,omitempty"`
	TotalPopulation       *int64   `json:"total_population,omitempty"`
	CoveredPopulation     *float64 `json:"covered_population,omitempty"`
	PopulationCoveragePct *float64 `json:"population_coverage_pct,omitempty"`
	StateName             *string  `json:"state_name,omitempty"`
	GeocodedPct           *float64 `json:"geocoded_pct,omitempty"`
	OwnerPct              *float64 `json:"owner_pct,omitempty"`
	ValuePct              *float64 `json:"value_pct,omitempty"`
}

// UnmarshalJSON decodes CoverageGetResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CoverageGetResponseData) UnmarshalJSON(data []byte) error {
	type plain CoverageGetResponseData
	aux := struct {
		*plain
		ParcelCount           lenientNumber[int64]   `json:"parcel_count"`
		WithAddress           lenientNumber[float64] `json:"with_address"`
		WithGeocode           lenientNumber[float64] `json:"with_geocode"`
		WithOwner             lenientNumber[float64] `json:"with_owner"`
		WithValue             lenientNumber[float64] `json:"with_value"`
		Population            lenientNumber[float64] `json:"population"`
		WithGeometry          lenientNumber[float64] `json:"with_geometry"`
		CountyCount           lenientNumber[int64]   `json:"county_count"`
		TotalCounties         lenientNumber[int64]   `json:"total_counties"`
		TotalPopulation       lenientNumber[int64]   `json:"total_population"`
		CoveredPopulation     lenientNumber[float64] `json:"covered_population"`
		PopulationCoveragePct lenientNumber[float64] `json:"population_coverage_pct"`
		GeocodedPct           lenientNumber[float64] `json:"geocoded_pct"`
		OwnerPct              lenientNumber[float64] `json:"owner_pct"`
		ValuePct              lenientNumber[float64] `json:"value_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelCount.assignPtr(&r.ParcelCount)
	aux.WithAddress.assignPtr(&r.WithAddress)
	aux.WithGeocode.assignPtr(&r.WithGeocode)
	aux.WithOwner.assignPtr(&r.WithOwner)
	aux.WithValue.assignPtr(&r.WithValue)
	aux.Population.assignPtr(&r.Population)
	aux.WithGeometry.assignPtr(&r.WithGeometry)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.TotalCounties.assignPtr(&r.TotalCounties)
	aux.TotalPopulation.assignPtr(&r.TotalPopulation)
	aux.CoveredPopulation.assignPtr(&r.CoveredPopulation)
	aux.PopulationCoveragePct.assignPtr(&r.PopulationCoveragePct)
	aux.GeocodedPct.assignPtr(&r.GeocodedPct)
	aux.OwnerPct.assignPtr(&r.OwnerPct)
	aux.ValuePct.assignPtr(&r.ValuePct)
	return softTypeError(err)
}

// CoverageGetResponseNationalCounts is generated from the OpenAPI spec.
type CoverageGetResponseNationalCounts struct {
	ParcelCountsEpoch      string                                                  `json:"parcel_counts_epoch"`
	Parcels                float64                                                 `json:"parcels"`
	MappedLocations        float64                                                 `json:"mapped_locations"`
	GeocodedParcels        float64                                                 `json:"geocoded_parcels"`
	Rows                   float64                                                 `json:"rows"`
	ParcelCountDefinitions CoverageGetResponseNationalCountsParcelCountDefinitions `json:"parcel_count_definitions"`
}

// UnmarshalJSON decodes CoverageGetResponseNationalCounts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CoverageGetResponseNationalCounts) UnmarshalJSON(data []byte) error {
	type plain CoverageGetResponseNationalCounts
	aux := struct {
		*plain
		Parcels         lenientNumber[float64] `json:"parcels"`
		MappedLocations lenientNumber[float64] `json:"mapped_locations"`
		GeocodedParcels lenientNumber[float64] `json:"geocoded_parcels"`
		Rows            lenientNumber[float64] `json:"rows"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assign(&r.Parcels)
	aux.MappedLocations.assign(&r.MappedLocations)
	aux.GeocodedParcels.assign(&r.GeocodedParcels)
	aux.Rows.assign(&r.Rows)
	return softTypeError(err)
}

// CoverageGetResponseNationalCountsParcelCountDefinitions is generated from the OpenAPI spec.
type CoverageGetResponseNationalCountsParcelCountDefinitions struct {
	Parcels         string `json:"parcels"`
	MappedLocations string `json:"mapped_locations"`
	GeocodedParcels string `json:"geocoded_parcels"`
	Rows            string `json:"rows"`
}

// Map: County coverage map data
//
// Per-county parcel coverage for the national coverage map (compact keys). Free; IP-throttled on
// the restricted budget; cached for an hour.
//
// HTTP: GET /api/v1/coverage/map
func (s *CoverageService) Map(ctx context.Context, params *CoverageMapParams, opts ...RequestOption) (*CoverageMapResponse, error) {
	var out CoverageMapResponse
	if err := s.client.do(ctx, buildCoverageMapRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCoverageMapRequest(params *CoverageMapParams) *apiRequest {
	req := newRequest("GET", "/api/v1/coverage/map")
	return req
}

// CoverageMapParams holds the query, header and JSON-body parameters of [CoverageService.Map].
// Pass nil when you need none.
type CoverageMapParams struct {
}

// CoverageMapResponse: County coverage map data
type CoverageMapResponse struct {
	Meta     CoverageMapResponseMeta       `json:"meta"`
	Counties []CoverageMapResponseCounties `json:"counties"`
}

// CoverageMapResponseMeta is generated from the OpenAPI spec.
type CoverageMapResponseMeta struct {
	Epoch       string                        `json:"epoch"`
	GeneratedAt string                        `json:"generated_at"`
	Totals      CoverageMapResponseMetaTotals `json:"totals"`
}

// CoverageMapResponseMetaTotals is generated from the OpenAPI spec.
type CoverageMapResponseMetaTotals struct {
	Parcels      float64 `json:"parcels"`
	WithOwner    float64 `json:"with_owner"`
	WithAddress  float64 `json:"with_address"`
	WithGeometry float64 `json:"with_geometry"`
	WithValue    float64 `json:"with_value"`
	Counties     int64   `json:"counties"`
}

// UnmarshalJSON decodes CoverageMapResponseMetaTotals, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CoverageMapResponseMetaTotals) UnmarshalJSON(data []byte) error {
	type plain CoverageMapResponseMetaTotals
	aux := struct {
		*plain
		Parcels      lenientNumber[float64] `json:"parcels"`
		WithOwner    lenientNumber[float64] `json:"with_owner"`
		WithAddress  lenientNumber[float64] `json:"with_address"`
		WithGeometry lenientNumber[float64] `json:"with_geometry"`
		WithValue    lenientNumber[float64] `json:"with_value"`
		Counties     lenientNumber[int64]   `json:"counties"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assign(&r.Parcels)
	aux.WithOwner.assign(&r.WithOwner)
	aux.WithAddress.assign(&r.WithAddress)
	aux.WithGeometry.assign(&r.WithGeometry)
	aux.WithValue.assign(&r.WithValue)
	aux.Counties.assign(&r.Counties)
	return softTypeError(err)
}

// CoverageMapResponseCounties is generated from the OpenAPI spec.
type CoverageMapResponseCounties struct {
	S   *string         `json:"s,omitempty"`
	C   json.RawMessage `json:"c,omitempty"`
	Nm  *string         `json:"nm,omitempty"`
	N   *int64          `json:"n,omitempty"`
	O   *float64        `json:"o,omitempty"`
	A   *float64        `json:"a,omitempty"`
	G   *float64        `json:"g,omitempty"`
	V   *float64        `json:"v,omitempty"`
	Lat *float64        `json:"lat,omitempty"`
	Lon *float64        `json:"lon,omitempty"`
}

// UnmarshalJSON decodes CoverageMapResponseCounties, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CoverageMapResponseCounties) UnmarshalJSON(data []byte) error {
	type plain CoverageMapResponseCounties
	aux := struct {
		*plain
		N   lenientNumber[int64]   `json:"n"`
		O   lenientNumber[float64] `json:"o"`
		A   lenientNumber[float64] `json:"a"`
		G   lenientNumber[float64] `json:"g"`
		V   lenientNumber[float64] `json:"v"`
		Lat lenientNumber[float64] `json:"lat"`
		Lon lenientNumber[float64] `json:"lon"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.N.assignPtr(&r.N)
	aux.O.assignPtr(&r.O)
	aux.A.assignPtr(&r.A)
	aux.G.assignPtr(&r.G)
	aux.V.assignPtr(&r.V)
	aux.Lat.assignPtr(&r.Lat)
	aux.Lon.assignPtr(&r.Lon)
	return softTypeError(err)
}
