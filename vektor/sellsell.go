// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
)

// SellSellService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSellSellService] method instead.
type SellSellService struct {
	Options []option.RequestOption
}

// NewSellSellService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSellSellService(opts ...option.RequestOption) (r SellSellService) {
	r = SellSellService{}
	r.Options = opts
	return
}

// Sell an asset for another asset
func (r *SellSellService) New(ctx context.Context, body SellSellNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/sell/sell"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SellSellNewParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	ReceiveAsset AssetIDOrAddressEVMOrAssetSymbol `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	SpendAmount Decimal `json:"spend_amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	SpendAsset AssetIDOrAddressEVMOrAssetSymbol `json:"spend_asset,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An arbitrary precision decimal represented as a string
	Slippage param.Opt[string] `json:"slippage,omitzero"`
	paramObj
}

func (r SellSellNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SellSellNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SellSellNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
