// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/nirvana-go/option"
)

func TestClusterPoolAvailabilityNewWithOptionalParams(t *testing.T) {
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
	err := client.NKS.Clusters.Pools.Availability.New(
		context.TODO(),
		"cluster_id",
		nks.ClusterPoolAvailabilityNewParams{
			Name: "my-node-pool",
			NodeConfig: nks.NKSNodePoolNodeConfigParam{
				BootVolume: nks.NKSNodePoolBootVolumeParam{
					Size: 100,
					Type: compute.VolumeTypeABS,
				},
				InstanceType: "n1-standard-8",
				Labels:       []string{"env=prod", "team=platform"},
			},
			NodeCount: 3,
			Tags:      []string{"production", "ethereum"},
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

func TestClusterPoolAvailabilityUpdateWithOptionalParams(t *testing.T) {
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
	err := client.NKS.Clusters.Pools.Availability.Update(
		context.TODO(),
		"cluster_id",
		"pool_id",
		nks.ClusterPoolAvailabilityUpdateParams{
			Name: nirvana.String("my-node-pool"),
			NodeConfig: nks.ClusterPoolAvailabilityUpdateParamsNodeConfig{
				Labels: []string{"env=prod", "team=platform"},
			},
			NodeCount: nirvana.Int(5),
			Tags:      []string{"production", "ethereum"},
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
