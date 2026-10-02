// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// CrimeService groups the crime operations. Use it as client.Crime.
type CrimeService struct {
	client *Client
}

// Lookup: Crime score near a point
//
// The crime score of the nearest scored parcel within about 500 m of the point, or `crime: null`.
// Free; IP-throttled.
//
// HTTP: GET /api/v1/crime/lookup
func (s *CrimeService) Lookup(ctx context.Context, params *CrimeLookupParams, opts ...RequestOption) (*CrimeLookupResponse, error) {
	var out CrimeLookupResponse
	if err := s.client.do(ctx, buildCrimeLookupRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildCrimeLookupRequest(params *CrimeLookupParams) *apiRequest {
	req := newRequest("GET", "/api/v1/crime/lookup")
	if params != nil {
		addQuery(req.query, "lat", params.Lat)
		addQuery(req.query, "lng", params.Lng)
	}
	return req
}

// CrimeLookupParams holds the query, header and JSON-body parameters of [CrimeService.Lookup].
// Pass nil when you need none.
type CrimeLookupParams struct {
	// Latitude.
	//
	// Required.
	Lat *float64 `query:"lat" json:"-"`

	// Longitude.
	//
	// Required.
	Lng *float64 `query:"lng" json:"-"`
}

// CrimeLookupResponse: Crime score near a point
type CrimeLookupResponse struct {
	Crime CrimeLookupResponseCrime `json:"crime"`
}

// CrimeLookupResponseCrime is generated from the OpenAPI spec.
type CrimeLookupResponseCrime struct {
	CrimeScore              *float64 `json:"crime_score,omitempty"`
	CrimeTier               *float64 `json:"crime_tier,omitempty"`
	CrimeTrend              *string  `json:"crime_trend,omitempty"`
	CountyViolentCrimeRate  *float64 `json:"county_violent_crime_rate,omitempty"`
	CountyPropertyCrimeRate *float64 `json:"county_property_crime_rate,omitempty"`
}

// UnmarshalJSON decodes CrimeLookupResponseCrime, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CrimeLookupResponseCrime) UnmarshalJSON(data []byte) error {
	type plain CrimeLookupResponseCrime
	aux := struct {
		*plain
		CrimeScore              lenientNumber[float64] `json:"crime_score"`
		CrimeTier               lenientNumber[float64] `json:"crime_tier"`
		CountyViolentCrimeRate  lenientNumber[float64] `json:"county_violent_crime_rate"`
		CountyPropertyCrimeRate lenientNumber[float64] `json:"county_property_crime_rate"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CrimeScore.assignPtr(&r.CrimeScore)
	aux.CrimeTier.assignPtr(&r.CrimeTier)
	aux.CountyViolentCrimeRate.assignPtr(&r.CountyViolentCrimeRate)
	aux.CountyPropertyCrimeRate.assignPtr(&r.CountyPropertyCrimeRate)
	return softTypeError(err)
}
