// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// LPWithdrawQuoteService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLPWithdrawQuoteService] method instead.
type LPWithdrawQuoteService struct {
	Options []option.RequestOption
}

// NewLPWithdrawQuoteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLPWithdrawQuoteService(opts ...option.RequestOption) (r LPWithdrawQuoteService) {
	r = LPWithdrawQuoteService{}
	r.Options = opts
	return
}

// Simulate withdrawing liquidity from a specific existing LP position
func (r *LPWithdrawQuoteService) New(ctx context.Context, body LPWithdrawQuoteNewParams, opts ...option.RequestOption) (res *LPWithdrawQuoteNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lp/withdraw_quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LPWithdrawQuoteNewResponse struct {
	// An LP quote
	Items LPQuote `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPWithdrawQuoteNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LPWithdrawQuoteNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPWithdrawQuoteNewParams struct {
	// An EVM address
	Account Account `json:"account,required"`
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	Asset AssetIDOrAddressEVMOrAssetSymbol `json:"asset,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A LP pool ID, represented as a TypeID with `lp_pool` prefix
	LPPoolID string `json:"lp_pool_id,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Uniswap position specifier
	Specifier LPWithdrawQuoteNewParamsSpecifier `json:"specifier,omitzero"`
	paramObj
}

func (r LPWithdrawQuoteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LPWithdrawQuoteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPWithdrawQuoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Uniswap position specifier
//
// The property PositionNFT is required.
type LPWithdrawQuoteNewParamsSpecifier struct {
	// A NFT
	PositionNFT NFTParam `json:"position_nft,omitzero,required"`
	paramObj
}

func (r LPWithdrawQuoteNewParamsSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow LPWithdrawQuoteNewParamsSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPWithdrawQuoteNewParamsSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
