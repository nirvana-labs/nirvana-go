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

// LendLendService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendLendService] method instead.
type LendLendService struct {
	Options []option.RequestOption
}

// NewLendLendService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLendLendService(opts ...option.RequestOption) (r LendLendService) {
	r = LendLendService{}
	r.Options = opts
	return
}

// Lend an asset
func (r *LendLendService) New(ctx context.Context, body LendLendNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lend/lend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendLendNewParams struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	Asset AssetIDOrAddressEVMOrAssetSymbol `json:"asset,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	paramObj
}

func (r LendLendNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LendLendNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendLendNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
