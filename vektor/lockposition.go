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

// LockPositionService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockPositionService] method instead.
type LockPositionService struct {
	Options []option.RequestOption
}

// NewLockPositionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockPositionService(opts ...option.RequestOption) (r LockPositionService) {
	r = LockPositionService{}
	r.Options = opts
	return
}

// Get info on locked positions
func (r *LockPositionService) List(ctx context.Context, body LockPositionListParams, opts ...option.RequestOption) (res *LockPositionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lock/positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LockPositionListResponse struct {
	// A list of lock positions
	Items []LockPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LockPositionListResponse) RawJSON() string { return r.JSON.raw }
func (r *LockPositionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LockPositionListParams struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	paramObj
}

func (r LockPositionListParams) MarshalJSON() (data []byte, err error) {
	type shadow LockPositionListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LockPositionListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
