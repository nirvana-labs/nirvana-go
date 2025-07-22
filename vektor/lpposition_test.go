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

func TestLPPositionListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
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
	_, err := client.Vektor.LP.Positions.List(context.TODO(), vektor.LPPositionListParams{
		OfLPPositionsByVenueAssetRequestInput: &vektor.LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput{
			Accounts:   []vektor.Account{"0x6b175474e89094c44da98b954eedeac495271d0f"},
			Assets:     []vektor.AssetIDOrAddressEVMOrAssetSymbol{"asset_01jbz9qc00f8wr64hfe459gb7y"},
			Blockchain: "blockchain_01jbz9nsy8egar70jg79dkwmaf",
			Venues:     []vektor.VenueIDOrVenueSymbol{"venue_01jbz9qc18evw86sg8m0sj9jg5"},
			At: vektor.LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			ExcludeZeros:     nirvana.Bool(true),
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

func TestLPPositionListHistoricalWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
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
	_, err := client.Vektor.LP.Positions.ListHistorical(context.TODO(), vektor.LPPositionListHistoricalParams{
		OfHistoricalLPPositionsByVenueAssetRequest: &vektor.LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest{
			Accounts:   []vektor.Account{"0x6b175474e89094c44da98b954eedeac495271d0f"},
			Assets:     []vektor.AssetIDOrAddressEVMOrAssetSymbol{"asset_01jbz9qc00f8wr64hfe459gb7y"},
			Blockchain: "blockchain_01jbz9nsy8egar70jg79dkwmaf",
			From: vektor.TimestampOrBlockNumberUnionParam{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			To: vektor.TimestampOrBlockNumberUnionParam{
				OfString: nirvana.String("2021-01-01T12:00:00Z"),
			},
			Venues:           []vektor.VenueIDOrVenueSymbol{"venue_01jbz9qc18evw86sg8m0sj9jg5"},
			ExcludeZeros:     nirvana.Bool(true),
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
