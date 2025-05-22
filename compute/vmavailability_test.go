// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
)

func TestVMAvailabilityNewWithOptionalParams(t *testing.T) {
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
	err := client.Compute.VMs.Availability.New(context.TODO(), compute.VMAvailabilityNewParams{
		BootVolume: compute.VMAvailabilityNewParamsBootVolume{
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
		DataVolumes: []compute.VMAvailabilityNewParamsDataVolume{{
			Name: "my-data-volume",
			Size: 100,
		}},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVMAvailabilityUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Compute.VMs.Availability.Update(
		context.TODO(),
		"vm_id",
		compute.VMAvailabilityUpdateParams{
			CPUConfig: compute.CPUConfigParam{
				Vcpu: 2,
			},
			MemoryConfig: compute.MemoryConfigParam{
				Size: 2,
			},
			Name:            nirvana.String("my-vm"),
			PublicIPEnabled: nirvana.Bool(true),
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
