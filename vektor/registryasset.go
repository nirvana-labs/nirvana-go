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

// RegistryAssetService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryAssetService] method instead.
type RegistryAssetService struct {
	Options []option.RequestOption
}

// NewRegistryAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryAssetService(opts ...option.RequestOption) (r RegistryAssetService) {
	r = RegistryAssetService{}
	r.Options = opts
	return
}

// List supported assets, optionally filtered by blockchain
func (r *RegistryAssetService) List(ctx context.Context, body RegistryAssetListParams, opts ...option.RequestOption) (res *RegistryAssetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/registry/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryAssetListResponse struct {
	// Response for multiple assets
	Items []Asset `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryAssetListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryAssetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryAssetListParams struct {
	// An asset ID, represented as a TypeID with `asset` prefix
	ID param.Opt[string] `json:"id,omitzero"`
	// An asset symbol
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain param.Opt[BlockchainID] `json:"blockchain,omitzero"`
	paramObj
}

func (r RegistryAssetListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryAssetListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryAssetListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
