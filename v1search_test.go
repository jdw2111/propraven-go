// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package propraven_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jdw2111/propraven-go"
	"github.com/jdw2111/propraven-go/internal/testutil"
	"github.com/jdw2111/propraven-go/option"
)

func TestV1SearchAutocomplete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := propraven.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Search.Autocomplete(context.TODO(), propraven.V1SearchAutocompleteParams{
		Q: "xx",
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SearchExportResultsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := propraven.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Search.ExportResults(context.TODO(), propraven.V1SearchExportResultsParams{
		East:             propraven.Float(0),
		Limit:            propraven.Int(1),
		North:            propraven.Float(0),
		Order:            propraven.V1SearchExportResultsParamsOrderAsc,
		Sort:             propraven.V1SearchExportResultsParamsSortAddress,
		South:            propraven.Float(0),
		West:             propraven.Float(0),
		ZoningCategories: propraven.String("zoningCategories"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SearchFullSearchWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := propraven.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Search.FullSearch(context.TODO(), propraven.V1SearchFullSearchParams{
		Q:     "xx",
		City:  propraven.String("city"),
		Dir:   propraven.V1SearchFullSearchParamsDirAsc,
		Field: propraven.V1SearchFullSearchParamsFieldAll,
		Limit: propraven.Int(1),
		Page:  propraven.Int(1),
		Sort:  propraven.V1SearchFullSearchParamsSortAddress,
		State: propraven.String("xx"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1SearchParcelSearchWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := propraven.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Search.ParcelSearch(context.TODO(), propraven.V1SearchParcelSearchParams{
		Bounds: propraven.V1SearchParcelSearchParamsBounds{
			East:  -78.55,
			North: 35.85,
			South: 35.75,
			West:  -78.7,
		},
		Filters: propraven.V1SearchParcelSearchParamsFilters{
			AbsenteeOnly: propraven.Bool(true),
			AcreageRange: propraven.V1SearchParcelSearchParamsFiltersAcreageRange{
				Max: propraven.Float(10),
				Min: propraven.Float(0.5),
			},
			OwnerTypes: []string{"individual"},
			ValueRange: propraven.V1SearchParcelSearchParamsFiltersValueRange{
				Max: propraven.Float(500000),
				Min: propraven.Float(100000),
			},
			YearBuiltRange: propraven.V1SearchParcelSearchParamsFiltersYearBuiltRange{
				Max: propraven.Int(2024),
				Min: propraven.Int(1990),
			},
			ZoningCategories: []string{"residential", "commercial"},
		},
		Limit:  propraven.Int(1),
		Offset: propraven.Int(0),
		Order:  propraven.V1SearchParcelSearchParamsOrderAsc,
		Sort:   propraven.V1SearchParcelSearchParamsSortAssessedValue,
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
