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
)

// BuyBuyService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuyBuyService] method instead.
type BuyBuyService struct {
	Options []option.RequestOption
}

// NewBuyBuyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuyBuyService(opts ...option.RequestOption) (r BuyBuyService) {
	r = BuyBuyService{}
	r.Options = opts
	return
}

// Buy an asset with another asset
func (r *BuyBuyService) New(ctx context.Context, body BuyBuyNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/buy/buy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BuyBuyNewParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// An arbitrary precision decimal represented as a string
	ReceiveAmount Decimal `json:"receive_amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	ReceiveAsset AssetIDOrAddressEVMOrAssetSymbol `json:"receive_asset,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	SpendAsset AssetIDOrAddressEVMOrAssetSymbol `json:"spend_asset,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An arbitrary precision decimal represented as a string
	Slippage param.Opt[string] `json:"slippage,omitzero"`
	paramObj
}

func (r BuyBuyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BuyBuyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuyBuyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
