// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nirvana_test

import (
	"context"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
)

func TestUsage(t *testing.T) {
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
	operation, err := client.Compute.VMs.New(context.TODO(), compute.VMNewParams{
		BootVolume: compute.VMNewParamsBootVolume{
			Size: 100,
		},
		CPUConfig: compute.CPUConfigParam{
			Vcpu: 2,
		},
		MemoryConfig: compute.MemoryConfigParam{
			Size: 2,
		},
		Name:            "my-vm",
		OSImageName:     "ubuntu-noble-2025-04-03",
		PublicIPEnabled: true,
		Region:          shared.RegionNameUsWdc1,
		SSHKey: compute.SSHKeyParam{
			PublicKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2",
		},
		SubnetID: "123e4567-e89b-12d3-a456-426614174000",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", operation.ID)
}
