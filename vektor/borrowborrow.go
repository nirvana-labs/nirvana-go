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

// BorrowBorrowService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBorrowBorrowService] method instead.
type BorrowBorrowService struct {
	Options []option.RequestOption
}

// NewBorrowBorrowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBorrowBorrowService(opts ...option.RequestOption) (r BorrowBorrowService) {
	r = BorrowBorrowService{}
	r.Options = opts
	return
}

// Borrow an asset
func (r *BorrowBorrowService) New(ctx context.Context, body BorrowBorrowNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/borrow/borrow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BorrowBorrowNewParams struct {
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

func (r BorrowBorrowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BorrowBorrowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowBorrowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
