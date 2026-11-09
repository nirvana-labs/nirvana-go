// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go/v2"
	"github.com/nirvana-labs/nirvana-go/v2/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
	"github.com/nirvana-labs/nirvana-go/v2/storage"
)

func TestBucketAvailabilityNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Storage.Buckets.Availability.New(context.TODO(), storage.BucketAvailabilityNewParams{
		Name:      "my-bucket",
		ProjectID: "123e4567-e89b-12d3-a456-426614174000",
		Region:    shared.RegionNameUsSva2,
		Tags:      []string{"production", "ethereum"},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
