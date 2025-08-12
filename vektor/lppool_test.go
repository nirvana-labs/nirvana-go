// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/vektor"
)

func TestLPPoolListWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Vektor.LP.Pools.List(context.TODO(), vektor.LPPoolListParams{
		OfLPPoolsByVenueAssetRequestInput: &vektor.LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput{
			Assets:     []vektor.AssetIDOrAddressEVMOrAssetSymbol{"asset_01jbz9qc00f8wr64hfe459gb7y"},
			Blockchain: "blockchain_01jbz9nsy8egar70jg79dkwmaf",
			Venues:     []vektor.VenueIDOrVenueSymbol{"venue_01jbz9qc18evw86sg8m0sj9jg5"},
			At: vektor.LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			QuoteAssetSymbol: nirvana.String("eth"),
		},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLPPoolListHistoricalWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Vektor.LP.Pools.ListHistorical(context.TODO(), vektor.LPPoolListHistoricalParams{
		OfHistoricalLPPoolsByVenueAssetRequest: &vektor.LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest{
			Assets:     []vektor.AssetIDOrAddressEVMOrAssetSymbol{"asset_01jbz9qc00f8wr64hfe459gb7y"},
			Blockchain: "blockchain_01jbz9nsy8egar70jg79dkwmaf",
			From: vektor.TimestampOrBlockNumberUnionParam{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			To: vektor.TimestampOrBlockNumberUnionParam{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			Venues:           []vektor.VenueIDOrVenueSymbol{"venue_01jbz9qc18evw86sg8m0sj9jg5"},
			QuoteAssetSymbol: nirvana.String("eth"),
		},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
