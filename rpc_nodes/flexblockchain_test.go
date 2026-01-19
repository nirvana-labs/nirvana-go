// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/rpc_nodes"
)

func TestFlexBlockchainListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.RPCNodes.Flex.Blockchains.List(context.TODO(), rpc_nodes.FlexBlockchainListParams{
		Cursor: nirvana.String("cursor"),
		Limit:  nirvana.Int(10),
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
