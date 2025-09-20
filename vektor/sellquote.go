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

// SellQuoteService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSellQuoteService] method instead.
type SellQuoteService struct {
	Options []option.RequestOption
}

// NewSellQuoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSellQuoteService(opts ...option.RequestOption) (r SellQuoteService) {
	r = SellQuoteService{}
	r.Options = opts
	return
}

// Get quotes for selling an exact amount of an asset at current market rate
func (r *SellQuoteService) List(ctx context.Context, body SellQuoteListParams, opts ...option.RequestOption) (res *SellQuoteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/sell/quotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SellQuoteListResponse struct {
	// A list of sell quotes
	Items []SellQuote `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SellQuoteListResponse) RawJSON() string { return r.JSON.raw }
func (r *SellQuoteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SellQuoteListParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	ReceiveAsset AssetIDOrAddressEVMOrAssetSymbol `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	SpendAmount Decimal `json:"spend_amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	SpendAsset AssetIDOrAddressEVMOrAssetSymbol `json:"spend_asset,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r SellQuoteListParams) MarshalJSON() (data []byte, err error) {
	type shadow SellQuoteListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SellQuoteListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
