// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go/v2"
	"github.com/nirvana-labs/nirvana-go/v2/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/v2/nks"
	"github.com/nirvana-labs/nirvana-go/v2/option"
)

func TestClusterControllerListWithOptionalParams(t *testing.T) {
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
	_, err := client.NKS.Clusters.Controllers.List(
		context.TODO(),
		"cluster_id",
		nks.ClusterControllerListParams{
			Cursor:       nirvana.String("cursor"),
			HasPrivateIP: nirvana.Bool(true),
			InstanceType: nirvana.String("instance_type"),
			Limit:        nirvana.Int(10),
			Name:         nirvana.String("name"),
			PrivateIP:    nirvana.String("private_ip"),
			Sort:         nirvana.String("sort"),
			Status:       nks.ClusterControllerListParamsStatusReady,
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

func TestClusterControllerGet(t *testing.T) {
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
	_, err := client.NKS.Clusters.Controllers.Get(
		context.TODO(),
		"cluster_id",
		"controller_id",
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
