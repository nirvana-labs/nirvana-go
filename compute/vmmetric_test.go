// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
)

func TestVMMetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.VMs.Metrics.List(
		context.TODO(),
		"vm_id",
		compute.VMMetricListParams{
			Aggregation: compute.VMMetricListParamsAggregationMean,
			EndTime:     nirvana.Time(time.Now()),
			Metric:      []string{"compute.nirvanalabs.io/vm/cpu/used_cores"},
			Period:      compute.VMMetricListParamsPeriod5m,
			StartTime:   nirvana.Time(time.Now()),
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
