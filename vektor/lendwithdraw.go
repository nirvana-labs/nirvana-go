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

// LendWithdrawService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendWithdrawService] method instead.
type LendWithdrawService struct {
	Options []option.RequestOption
}

// NewLendWithdrawService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLendWithdrawService(opts ...option.RequestOption) (r LendWithdrawService) {
	r = LendWithdrawService{}
	r.Options = opts
	return
}

// Withdraw an asset
func (r *LendWithdrawService) New(ctx context.Context, body LendWithdrawNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lend/withdraw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendWithdrawNewParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	MarketID LendBorrowMarketID `json:"market_id,required"`
	// An arbitrary precision decimal represented as a string
	Amount param.Opt[string] `json:"amount,omitzero"`
	// An asset ID, represented as a TypeID with `asset` prefix
	Asset param.Opt[AssetID] `json:"asset,omitzero"`
	paramObj
}

func (r LendWithdrawNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LendWithdrawNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendWithdrawNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
