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

func TestVMNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.VMs.New(context.TODO(), compute.VMNewParams{
		BootVolume: nirvana.F(compute.VMNewParamsBootVolume{
			Size: nirvana.F(int64(100)),
		}),
		CPUConfig: nirvana.F(compute.CPUConfigParam{
			Vcpu: nirvana.F(int64(2)),
		}),
		MemoryConfig: nirvana.F(compute.MemoryConfigParam{
			Size: nirvana.F(int64(2)),
		}),
		Name:            nirvana.F("my-vm"),
		OSImageName:     nirvana.F("ubuntu-noble-2025-03-04"),
		PublicIPEnabled: nirvana.F(true),
		Region:          nirvana.F(shared.RegionNameUsSea1),
		SSHKey: nirvana.F(compute.SSHKeyParam{
			PublicKey: nirvana.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDJiJabIUkXw7VrQG+yBohvhEsyoKEYvejZc4RFzV5maybqQei1punVsoe4r6gJttMM1Gr3cNr3OfepikCQAhAchw5ww94ZWqDsDYIqMrlDFbqhGTXDNzFAjeVIKptCOlz9k+7aM69YtLXJ6gFUCq1fbK9PjY+AK28UpMfKYUcyHQ== noname"),
		}),
		SubnetID: nirvana.F("123e4567-e89b-12d3-a456-426614174000"),
		DataVolumes: nirvana.F([]compute.VMNewParamsDataVolume{{
			Name: nirvana.F("my-data-volume"),
			Size: nirvana.F(int64(100)),
		}}),
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVMUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.VMs.Update(
		context.TODO(),
		"vm_id",
		compute.VMUpdateParams{
			CPUConfig: nirvana.F(compute.CPUConfigParam{
				Vcpu: nirvana.F(int64(2)),
			}),
			MemoryConfig: nirvana.F(compute.MemoryConfigParam{
				Size: nirvana.F(int64(2)),
			}),
			Name:            nirvana.F("my-vm"),
			PublicIPEnabled: nirvana.F(true),
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

func TestVMList(t *testing.T) {
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
	_, err := client.Compute.VMs.List(context.TODO())
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVMDelete(t *testing.T) {
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
	_, err := client.Compute.VMs.Delete(context.TODO(), "vm_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVMGet(t *testing.T) {
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
	_, err := client.Compute.VMs.Get(context.TODO(), "vm_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
