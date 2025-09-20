// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// BridgeQuoteService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeQuoteService] method instead.
type BridgeQuoteService struct {
	Options []option.RequestOption
}

// NewBridgeQuoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBridgeQuoteService(opts ...option.RequestOption) (r BridgeQuoteService) {
	r = BridgeQuoteService{}
	r.Options = opts
	return
}

// Get quotes for bridging an exact amount of an asset to another blockchain
func (r *BridgeQuoteService) List(ctx context.Context, body BridgeQuoteListParams, opts ...option.RequestOption) (res *BridgeQuoteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/bridge/quotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BridgeQuoteListResponse struct {
	// A list of bridge quotes
	Items []BridgeQuote `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeQuoteListResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeQuoteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeQuoteListParams struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// An asset symbol
	Asset AssetSymbol `json:"asset,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	FromBlockchain BlockchainIDOrBlockchainSymbol `json:"from_blockchain,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	ToBlockchain BlockchainIDOrBlockchainSymbol `json:"to_blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r BridgeQuoteListParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeQuoteListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeQuoteListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
