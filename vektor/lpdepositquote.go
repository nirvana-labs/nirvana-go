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

// LPDepositQuoteService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLPDepositQuoteService] method instead.
type LPDepositQuoteService struct {
	Options []option.RequestOption
}

// NewLPDepositQuoteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLPDepositQuoteService(opts ...option.RequestOption) (r LPDepositQuoteService) {
	r = LPDepositQuoteService{}
	r.Options = opts
	return
}

// Simulate depositing liquidity to a specific LP pool, creating a position or
// adding to an existing one.
func (r *LPDepositQuoteService) New(ctx context.Context, body LPDepositQuoteNewParams, opts ...option.RequestOption) (res *LPDepositQuoteNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lp/deposit_quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LPDepositQuoteNewResponse struct {
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
func (r LPDepositQuoteNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LPDepositQuoteNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPDepositQuoteNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for simulating a deposit to a liquidity pool, creating a position
	OfLPPoolDepositQuoteRequestInput *LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for simulating a deposit to a liquidity pool, adding to an existing
	// position
	OfLPPositionDepositQuoteRequestInput *LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInput `json:",inline"`

	paramObj
}

func (u LPDepositQuoteNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLPPoolDepositQuoteRequestInput, u.OfLPPositionDepositQuoteRequestInput)
}
func (r *LPDepositQuoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for simulating a deposit to a liquidity pool, creating a position
//
// The properties Amount, Asset, Blockchain, LPPoolID are required.
type LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput struct {
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
	// A Uniswap V3 range. Lower and upper bounds should satisfy 0 <= `lower` <
	// `upper`. The value -1 can be used in `upper` for infinity
	Range LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange `json:"range,omitzero"`
	paramObj
}

func (r LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Uniswap V3 range. Lower and upper bounds should satisfy 0 <= `lower` <
// `upper`. The value -1 can be used in `upper` for infinity
//
// The properties Lower, Upper are required.
type LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange struct {
	// An arbitrary precision decimal represented as a string
	Lower Decimal `json:"lower,required"`
	// An arbitrary precision decimal represented as a string
	Upper Decimal `json:"upper,required"`
	paramObj
}

func (r LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange) MarshalJSON() (data []byte, err error) {
	type shadow LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPDepositQuoteNewParamsBodyLPPoolDepositQuoteRequestInputRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for simulating a deposit to a liquidity pool, adding to an existing
// position
//
// The properties Account, Amount, Asset, Blockchain, LPPoolID are required.
type LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInput struct {
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
	Specifier LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInputSpecifier `json:"specifier,omitzero"`
	paramObj
}

func (r LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Uniswap position specifier
//
// The property PositionNFT is required.
type LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInputSpecifier struct {
	// A NFT
	PositionNFT NFTParam `json:"position_nft,omitzero,required"`
	paramObj
}

func (r LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInputSpecifier) MarshalJSON() (data []byte, err error) {
	type shadow LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInputSpecifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPDepositQuoteNewParamsBodyLPPositionDepositQuoteRequestInputSpecifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
