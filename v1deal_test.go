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

func TestV1DealFindAbsenteeOwnersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindAbsenteeOwners(context.TODO(), propraven.V1DealFindAbsenteeOwnersParams{
		CountyFips: propraven.String("37183"),
		Limit:      propraven.Int(1),
		MinValue:   propraven.Float(50000),
		Offset:     propraven.Int(0),
		OutOfState: propraven.Bool(true),
		StateFips:  propraven.String("37"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealFindEntityOwnedParcelsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindEntityOwnedParcels(context.TODO(), propraven.V1DealFindEntityOwnedParcelsParams{
		CountyFips: propraven.String("37183"),
		EntityType: propraven.V1DealFindEntityOwnedParcelsParamsEntityTypeLlc,
		Limit:      propraven.Int(1),
		MinValue:   propraven.Int(0),
		Offset:     propraven.Int(0),
		Search:     propraven.String("BLACKSTONE"),
		StateFips:  propraven.String("37"),
		Top:        propraven.Bool(true),
		Zoning:     propraven.String("zoning"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealFindFlipsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindFlips(context.TODO(), propraven.V1DealFindFlipsParams{
		CountyFips: propraven.String("37183"),
		FlipTier:   propraven.V1DealFindFlipsParamsFlipTierQuick,
		Limit:      propraven.Int(1),
		MinProfit:  propraven.Float(25000),
		Offset:     propraven.Int(0),
		StateFips:  propraven.String("37"),
		View:       propraven.V1DealFindFlipsParamsViewFlippers,
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealFindHighLandRatioWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindHighLandRatio(context.TODO(), propraven.V1DealFindHighLandRatioParams{
		CountyFips: propraven.String("county_fips"),
		Limit:      propraven.Int(1),
		MinRatio:   propraven.Float(1),
		MinValue:   propraven.Int(0),
		Offset:     propraven.Int(0),
		StateFips:  propraven.String("state_fips"),
		Zoning:     propraven.String("zoning"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealFindLongHoldParcelsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindLongHoldParcels(context.TODO(), propraven.V1DealFindLongHoldParcelsParams{
		CountyFips: propraven.String("county_fips"),
		HoldTier:   propraven.V1DealFindLongHoldParcelsParamsHoldTier10_15yr,
		Limit:      propraven.Int(1),
		MinValue:   propraven.Int(0),
		MinYears:   propraven.Int(0),
		Offset:     propraven.Int(0),
		StateFips:  propraven.String("state_fips"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealFindPortfolioOwnersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.FindPortfolioOwners(context.TODO(), propraven.V1DealFindPortfolioOwnersParams{
		Limit:         propraven.Int(1),
		MinProperties: propraven.Int(0),
		MinValue:      propraven.Int(0),
		Offset:        propraven.Int(0),
		Search:        propraven.String("search"),
		State:         propraven.String("NC"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealGetMarketSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.GetMarketSummary(context.TODO(), propraven.V1DealGetMarketSummaryParams{
		CountyFips: propraven.String("county_fips"),
		Limit:      propraven.Int(1),
		Offset:     propraven.Int(0),
		Rating:     propraven.V1DealGetMarketSummaryParamsRatingAffordable,
		StateFips:  propraven.String("state_fips"),
		View:       propraven.V1DealGetMarketSummaryParamsViewAffordability,
		Year:       propraven.String("2024"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealSearchContractorsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.SearchContractors(context.TODO(), propraven.V1DealSearchContractorsParams{
		Limit:      propraven.Int(1),
		MinPermits: propraven.Int(1),
		MinValue:   propraven.Int(0),
		Offset:     propraven.Int(0),
		Search:     propraven.String("SMITH CONSTRUCTION"),
		State:      propraven.String("NC"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1DealSearchLendersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Deals.SearchLenders(context.TODO(), propraven.V1DealSearchLendersParams{
		Limit:        propraven.Int(1),
		MinMortgages: propraven.Int(0),
		Offset:       propraven.Int(0),
		Search:       propraven.String("WELLS FARGO"),
		State:        propraven.String("NC"),
	})
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
