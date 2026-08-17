// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_types_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/instance_types"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
)

func TestInstanceTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.InstanceTypes.List(context.TODO(), instance_types.InstanceTypeListParams{
		Chipset:                 nirvana.String("chipset"),
		Cursor:                  nirvana.String("cursor"),
		Family:                  nirvana.String("family"),
		Limit:                   nirvana.Int(10),
		MemoryGBMax:             nirvana.Int(0),
		MemoryGBMin:             nirvana.Int(0),
		Name:                    nirvana.String("name"),
		NetworkBandwidthGbpsMax: nirvana.Float(0),
		NetworkBandwidthGbpsMin: nirvana.Float(0),
		Region:                  nirvana.String("region"),
		Series:                  nirvana.String("series"),
		Sort:                    nirvana.String("sort"),
		VcpuMax:                 nirvana.Int(0),
		VcpuMin:                 nirvana.Int(0),
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceTypeGet(t *testing.T) {
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
	_, err := client.InstanceTypes.Get(
		context.TODO(),
		instance_types.InstanceTypeGetParamsRegionUsSva2,
		"n1-standard-8",
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
