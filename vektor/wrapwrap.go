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

// WrapWrapService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWrapWrapService] method instead.
type WrapWrapService struct {
	Options []option.RequestOption
}

// NewWrapWrapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWrapWrapService(opts ...option.RequestOption) (r WrapWrapService) {
	r = WrapWrapService{}
	r.Options = opts
	return
}

// Wrap the native asset
func (r *WrapWrapService) New(ctx context.Context, body WrapWrapNewParams, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/wrap/wrap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WrapWrapNewParams struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	paramObj
}

func (r WrapWrapNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WrapWrapNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WrapWrapNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
