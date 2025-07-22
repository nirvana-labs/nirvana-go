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

func TestRegistryLendMarketListWithOptionalParams(t *testing.T) {
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
	_, err := client.Vektor.Registry.LendMarkets.List(context.TODO(), vektor.RegistryLendMarketListParams{
		ID:         nirvana.String("lend_borrow_market_01h455vb4pex5vsknk084sn02q"),
		Assets:     []vektor.AssetIDOrAddressEVMOrAssetSymbol{"asset_01jbz9qc00f8wr64hfe459gb7y"},
		Blockchain: nirvana.String("blockchain_01jbz9nsy8egar70jg79dkwmaf"),
		Venues:     []vektor.VenueIDOrVenueSymbol{"venue_01jbz9qc18evw86sg8m0sj9jg5"},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
