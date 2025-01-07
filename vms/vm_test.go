// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
	"github.com/nirvana-labs/nirvana-go/vms"
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
	_, err := client.VMs.New(context.TODO(), vms.VMNewParams{
		CPU: nirvana.F(vms.CPUParam{
			Cores: nirvana.F(int64(2)),
		}),
		Name:         nirvana.F("my-vm"),
		NeedPublicIP: nirvana.F(true),
		OsImageID:    nirvana.F(int64(1)),
		Ports:        nirvana.F([]string{"22", "80", "443"}),
		Ram: nirvana.F(vms.RamParam{
			Size: nirvana.F(int64(2)),
			Unit: nirvana.F(vms.RamUnitGB),
		}),
		Region:        nirvana.F(shared.RegionNameAmsterdam),
		SourceAddress: nirvana.F("0.0.0.0/0"),
		SSHKey: nirvana.F(vms.SSHKeyParam{
			PublicKey: nirvana.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"),
		}),
		Storage: nirvana.F([]vms.StorageParam{{
			Size:     nirvana.F(int64(100)),
			Type:     nirvana.F(vms.StorageTypeNvme),
			Unit:     nirvana.F(vms.StorageUnitGB),
			DiskName: nirvana.F("disk_name"),
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
	_, err := client.VMs.Update(
		context.TODO(),
		"vm_id",
		vms.VMUpdateParams{
			CPU: nirvana.F(vms.CPUParam{
				Cores: nirvana.F(int64(2)),
			}),
			Ram: nirvana.F(vms.RamParam{
				Size: nirvana.F(int64(2)),
				Unit: nirvana.F(vms.RamUnitGB),
			}),
			Storage: nirvana.F([]vms.StorageParam{{
				Size:     nirvana.F(int64(100)),
				Type:     nirvana.F(vms.StorageTypeNvme),
				Unit:     nirvana.F(vms.StorageUnitGB),
				DiskName: nirvana.F("disk_name"),
			}}),
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
	_, err := client.VMs.List(context.TODO(), vms.VMListParams{
		Region: nirvana.F("region"),
	})
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
	_, err := client.VMs.Delete(context.TODO(), "vm_id")
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
	_, err := client.VMs.Get(context.TODO(), "vm_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
