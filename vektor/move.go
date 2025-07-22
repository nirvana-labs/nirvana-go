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

// MoveService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMoveService] method instead.
type MoveService struct {
	Options []option.RequestOption
}

// NewMoveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMoveService(opts ...option.RequestOption) (r MoveService) {
	r = MoveService{}
	r.Options = opts
	return
}

// Move balance from one address to another
func (r *MoveService) New(ctx context.Context, body MoveNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/move/move"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MoveNewParams struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	Asset AssetIDOrAddressEVMOrAssetSymbol `json:"asset,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// An EVM address
	To Account `json:"to,required"`
	paramObj
}

func (r MoveNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MoveNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MoveNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
