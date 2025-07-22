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

// WrapUnwrapService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWrapUnwrapService] method instead.
type WrapUnwrapService struct {
	Options []option.RequestOption
}

// NewWrapUnwrapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWrapUnwrapService(opts ...option.RequestOption) (r WrapUnwrapService) {
	r = WrapUnwrapService{}
	r.Options = opts
	return
}

// Unwrap the wrapped native asset
func (r *WrapUnwrapService) New(ctx context.Context, body WrapUnwrapNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/wrap/unwrap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WrapUnwrapNewParams struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	paramObj
}

func (r WrapUnwrapNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WrapUnwrapNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WrapUnwrapNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
