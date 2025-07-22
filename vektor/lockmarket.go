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

// LockMarketService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockMarketService] method instead.
type LockMarketService struct {
	Options []option.RequestOption
}

// NewLockMarketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockMarketService(opts ...option.RequestOption) (r LockMarketService) {
	r = LockMarketService{}
	r.Options = opts
	return
}

// Get a list of markets where assets can be locked
func (r *LockMarketService) List(ctx context.Context, body LockMarketListParams, opts ...option.RequestOption) (res *LockMarketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lock/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LockMarketListResponse struct {
	// A list of lock markets
	Items []LockMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LockMarketListResponse) RawJSON() string { return r.JSON.raw }
func (r *LockMarketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LockMarketListParams struct {
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	paramObj
}

func (r LockMarketListParams) MarshalJSON() (data []byte, err error) {
	type shadow LockMarketListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LockMarketListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
