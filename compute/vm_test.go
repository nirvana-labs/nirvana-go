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
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Compute.VMs.New(context.TODO(), compute.VMNewParams{
		BootVolume: nirvana.F(compute.VMNewParamsBootVolume{
			Size: nirvana.F(int64(100)),
		}),
		CPU: nirvana.F(compute.CPUParam{
			Cores: nirvana.F(int64(2)),
		}),
		Name:         nirvana.F("my-vm"),
		NeedPublicIP: nirvana.F(true),
		OSImageName:  nirvana.F("noble-2024-12-06"),
		Ports:        nirvana.F([]string{"22", "80", "443"}),
		Ram: nirvana.F(compute.RamParam{
			Size: nirvana.F(int64(2)),
		}),
		Region:        nirvana.F(shared.RegionNameAmsterdam),
		SourceAddress: nirvana.F("0.0.0.0/0"),
		SSHKey: nirvana.F(compute.SSHKeyParam{
			PublicKey: nirvana.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"),
		}),
		DataVolumes: nirvana.F([]compute.VMNewParamsDataVolume{{
			Size: nirvana.F(int64(100)),
			Type: nirvana.F(compute.StorageTypeNvme),
		}}),
		SubnetID: nirvana.F("123e4567-e89b-12d3-a456-426614174000"),
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
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Compute.VMs.Update(
		context.TODO(),
		"vm_id",
		compute.VMUpdateParams{
			BootVolume: nirvana.F(compute.VMUpdateParamsBootVolume{
				Size: nirvana.F(int64(100)),
			}),
			CPU: nirvana.F(compute.CPUParam{
				Cores: nirvana.F(int64(2)),
			}),
			DataVolumes: nirvana.F([]compute.VMUpdateParamsDataVolume{{
				Size: nirvana.F(int64(100)),
				Type: nirvana.F(compute.StorageTypeNvme),
			}}),
			Ram: nirvana.F(compute.RamParam{
				Size: nirvana.F(int64(2)),
			}),
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
		option.WithAuthToken("My Auth Token"),
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
		option.WithAuthToken("My Auth Token"),
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
		option.WithAuthToken("My Auth Token"),
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
