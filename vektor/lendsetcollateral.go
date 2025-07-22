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

// LendSetCollateralService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendSetCollateralService] method instead.
type LendSetCollateralService struct {
	Options []option.RequestOption
}

// NewLendSetCollateralService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLendSetCollateralService(opts ...option.RequestOption) (r LendSetCollateralService) {
	r = LendSetCollateralService{}
	r.Options = opts
	return
}

// Enable/disable a specific lending position to be used as collateral
func (r *LendSetCollateralService) New(ctx context.Context, body LendSetCollateralNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lend/set_collateral"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendSetCollateralNewParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	MarketID LendBorrowMarketID `json:"market_id,required"`
	Status   bool               `json:"status,required"`
	paramObj
}

func (r LendSetCollateralNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LendSetCollateralNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendSetCollateralNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
