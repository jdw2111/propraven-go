// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// TrafficService groups the traffic operations. Use it as client.Traffic.
type TrafficService struct {
	client *Client
}

// Stations: Traffic count stations in a bounding box
//
// Up to 150 traffic count stations (AADT) inside `bbox`, busiest first.
//
// HTTP: GET /api/v1/traffic/stations
func (s *TrafficService) Stations(ctx context.Context, params *TrafficStationsParams, opts ...RequestOption) (*TrafficStationsResponse, error) {
	var out TrafficStationsResponse
	if err := s.client.do(ctx, buildTrafficStationsRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildTrafficStationsRequest(params *TrafficStationsParams) *apiRequest {
	req := newRequest("GET", "/api/v1/traffic/stations")
	if params != nil {
		addQuery(req.query, "bbox", params.Bbox)
	}
	return req
}

// TrafficStationsParams holds the query, header and JSON-body parameters of
// [TrafficService.Stations]. Pass nil when you need none.
type TrafficStationsParams struct {
	// `west,south,east,north` in decimal degrees.
	//
	// Required.
	Bbox *string `query:"bbox" json:"-"`
}

// TrafficStationsResponse: Traffic count stations in a bounding box
type TrafficStationsResponse struct {
	Data []TrafficStationsResponseData `json:"data"`
}

// TrafficStationsResponseData is generated from the OpenAPI spec.
type TrafficStationsResponseData struct {
	Lat   *float64 `json:"lat,omitempty"`
	Lng   *float64 `json:"lng,omitempty"`
	Aadt  *float64 `json:"aadt,omitempty"`
	Route *string  `json:"route,omitempty"`
}

// UnmarshalJSON decodes TrafficStationsResponseData, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationsResponseData) UnmarshalJSON(data []byte) error {
	type plain TrafficStationsResponseData
	aux := struct {
		*plain
		Lat  lenientNumber[float64] `json:"lat"`
		Lng  lenientNumber[float64] `json:"lng"`
		Aadt lenientNumber[float64] `json:"aadt"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Lat.assignPtr(&r.Lat)
	aux.Lng.assignPtr(&r.Lng)
	aux.Aadt.assignPtr(&r.Aadt)
	return softTypeError(err)
}
