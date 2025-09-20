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

// RegistryBlockchainService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryBlockchainService] method instead.
type RegistryBlockchainService struct {
	Options []option.RequestOption
}

// NewRegistryBlockchainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistryBlockchainService(opts ...option.RequestOption) (r RegistryBlockchainService) {
	r = RegistryBlockchainService{}
	r.Options = opts
	return
}

// List supported blockchains
func (r *RegistryBlockchainService) List(ctx context.Context, body RegistryBlockchainListParams, opts ...option.RequestOption) (res *RegistryBlockchainListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/registry/blockchains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryBlockchainListResponse struct {
	// A list of blockchains
	Items []Blockchain `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryBlockchainListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryBlockchainListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryBlockchainListParams struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	ID param.Opt[string] `json:"id,omitzero"`
	// A blockchain symbol
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	paramObj
}

func (r RegistryBlockchainListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryBlockchainListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryBlockchainListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
