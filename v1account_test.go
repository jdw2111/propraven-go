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

func TestV1AccountGetUsage(t *testing.T) {
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
	_, err := client.V1.Account.GetUsage(context.TODO())
	if err != nil {
		var apierr *propraven.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
