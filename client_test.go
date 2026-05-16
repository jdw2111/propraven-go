package propraven

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient_RequiresAPIKey(t *testing.T) {
	_, err := NewClient()
	if err == nil || !strings.Contains(err.Error(), "API key required") {
		t.Fatalf("expected API key required, got: %v", err)
	}
}

func TestParcelsGet_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer pz_test" {
			t.Errorf("missing/incorrect Authorization header: %q", got)
		}
		if r.URL.Path != "/v1/parcels/06037:1234" {
			t.Errorf("unexpected path: %q", r.URL.Path)
		}
		json.NewEncoder(w).Encode(Parcel{ParcelID: "06037:1234", StateFIPS: "06", CountyFIPS: "037"})
	}))
	t.Cleanup(srv.Close)

	c, err := NewClient(WithAPIKey("pz_test"), WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	p, err := c.Parcels.Get(context.Background(), "06037:1234")
	if err != nil {
		t.Fatal(err)
	}
	if p.ParcelID != "06037:1234" {
		t.Errorf("unexpected parcel: %+v", p)
	}
}

func TestErrorDecoding_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "req_abc")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Parcel not found","code":"parcel_not_found"}`))
	}))
	t.Cleanup(srv.Close)

	c, _ := NewClient(WithAPIKey("pz_test"), WithBaseURL(srv.URL))
	_, err := c.Parcels.Get(context.Background(), "missing")
	if err == nil {
		t.Fatal("expected error")
	}
	var pe *Error
	if !errors.As(err, &pe) {
		t.Fatalf("expected *propraven.Error, got %T", err)
	}
	if pe.Status != 404 || pe.Code != "parcel_not_found" || pe.RequestID != "req_abc" {
		t.Errorf("unexpected error fields: %+v", pe)
	}
	if !IsNotFound(err) {
		t.Error("IsNotFound returned false")
	}
}

func TestErrorDecoding_RateLimited(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":"rate limit"}`))
	}))
	t.Cleanup(srv.Close)

	c, _ := NewClient(WithAPIKey("pz_test"), WithBaseURL(srv.URL))
	_, err := c.Parcels.Get(context.Background(), "anything")
	if !IsRateLimited(err) {
		t.Fatalf("IsRateLimited false, err=%v", err)
	}
}

func TestParcelsList_PassesQuery(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state_fips") != "06" {
			t.Errorf("missing state_fips: %v", r.URL.RawQuery)
		}
		if r.URL.Query().Get("limit") != "25" {
			t.Errorf("missing limit: %v", r.URL.RawQuery)
		}
		json.NewEncoder(w).Encode(ParcelPage{Data: []Parcel{{ParcelID: "06037:1"}}, NextCursor: ""})
	}))
	t.Cleanup(srv.Close)

	c, _ := NewClient(WithAPIKey("pz_test"), WithBaseURL(srv.URL))
	page, err := c.Parcels.List(context.Background(), ListOptions{StateFIPS: "06", Limit: 25})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Data) != 1 {
		t.Errorf("expected 1 parcel, got %d", len(page.Data))
	}
}
