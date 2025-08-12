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

func TestLPDepositQuoteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Vektor.LP.DepositQuote.New(context.TODO(), vektor.LPDepositQuoteNewParams{
		OfLPPoolDepositQuoteRequestInput: &vektor.LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput{
			Amount:           "10.0000000000000024",
			Asset:            "asset_01jbz9qc00f8wr64hfe459gb7y",
			Blockchain:       "blockchain_01jbz9nsy8egar70jg79dkwmaf",
			LPPoolID:         "lp_pool_01h455vb4pex5vsknk084sn02q",
			QuoteAssetSymbol: nirvana.String("eth"),
			Range: vektor.LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange{
				Lower: "10.0000000000000024",
				Upper: "10.0000000000000024",
			},
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
