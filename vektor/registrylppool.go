// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// RegistryLPPoolService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryLPPoolService] method instead.
type RegistryLPPoolService struct {
	Options []option.RequestOption
}

// NewRegistryLPPoolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryLPPoolService(opts ...option.RequestOption) (r RegistryLPPoolService) {
	r = RegistryLPPoolService{}
	r.Options = opts
	return
}

// List LP pools in the registry, optionally filtered by blockchain, assets or
// venues
func (r *RegistryLPPoolService) List(ctx context.Context, body RegistryLPPoolListParams, opts ...option.RequestOption) (res *RegistryLPPoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/registry/lp/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryLPPoolListResponse struct {
	// A list of registry data for LP pools
	Items []RegistryLPPool `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryLPPoolListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryLPPoolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryLPPoolListParams struct {
	// A LP pool ID, represented as a TypeID with `lp_pool` prefix
	ID param.Opt[string] `json:"id,omitzero"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain param.Opt[BlockchainID] `json:"blockchain,omitzero"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero"`
	paramObj
}

func (r RegistryLPPoolListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryLPPoolListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryLPPoolListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
