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
		option.WithAuthToken("My Auth Token"),
	)
	operation, err := client.Compute.VMs.New(context.TODO(), compute.VMNewParams{
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
		OSImageName:     nirvana.F("noble-2024-12-06"),
		PublicIPEnabled: nirvana.F(true),
		Region:          nirvana.F(shared.RegionNameUsSea1),
		SSHKey: nirvana.F(compute.SSHKeyParam{
			PublicKey: nirvana.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDJiJabIUkXw7VrQG+yBohvhEsyoKEYvejZc4RFzV5maybqQei1punVsoe4r6gJttMM1Gr3cNr3OfepikCQAhAchw5ww94ZWqDsDYIqMrlDFbqhGTXDNzFAjeVIKptCOlz9k+7aM69YtLXJ6gFUCq1fbK9PjY+AK28UpMfKYUcyHQ== noname"),
		}),
		SubnetID: nirvana.F("123e4567-e89b-12d3-a456-426614174000"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", operation.ID)
}
