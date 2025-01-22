// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nirvana_test

import (
	"context"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
	"github.com/nirvana-labs/nirvana-go/vms"
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
	operation, err := client.VMs.New(context.TODO(), vms.VMNewParams{
		BootVolume: nirvana.F(vms.VMNewParamsBootVolume{
			Size: nirvana.F(int64(100)),
		}),
		CPU: nirvana.F(vms.CPUParam{
			Cores: nirvana.F(int64(2)),
		}),
		Name:         nirvana.F("my-vm"),
		NeedPublicIP: nirvana.F(true),
		OsImageID:    nirvana.F(int64(1)),
		Ports:        nirvana.F([]string{"22", "80", "443"}),
		Ram: nirvana.F(vms.RamParam{
			Size: nirvana.F(int64(2)),
		}),
		Region:        nirvana.F(shared.RegionNameAmsterdam),
		SourceAddress: nirvana.F("0.0.0.0/0"),
		SSHKey: nirvana.F(vms.SSHKeyParam{
			PublicKey: nirvana.F("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC1234567890"),
		}),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", operation.ID)
}
