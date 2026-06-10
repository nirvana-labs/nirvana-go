// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
)

func TestConnectConnectionCostNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Networking.Connect.Connections.Cost.New(context.TODO(), networking.ConnectConnectionCostNewParams{
		BandwidthMbps: 50,
		CIDRs:         []string{"10.0.0.0/16"},
		Name:          "my-connect-connection",
		ProjectID:     "123e4567-e89b-12d3-a456-426614174000",
		ProviderCIDRs: []string{"172.16.0.0/16"},
		Region:        shared.RegionNameUsSva2,
		AWS: networking.ConnectConnectionAWSConfigRequestParam{
			AccountID: "523816707215",
			Region:    "us-east-1",
		},
		Tags: []string{"production", "ethereum"},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectConnectionCostUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Networking.Connect.Connections.Cost.Update(
		context.TODO(),
		"connection_id",
		networking.ConnectConnectionCostUpdateParams{
			Name: nirvana.String("my-connect-connection"),
			Tags: []string{"production", "ethereum"},
		},
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
