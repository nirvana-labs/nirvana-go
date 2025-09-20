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

// RegistryBorrowMarketService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryBorrowMarketService] method instead.
type RegistryBorrowMarketService struct {
	Options []option.RequestOption
}

// NewRegistryBorrowMarketService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegistryBorrowMarketService(opts ...option.RequestOption) (r RegistryBorrowMarketService) {
	r = RegistryBorrowMarketService{}
	r.Options = opts
	return
}

// List borrow markets in the registry, optionally filtered by blockchain, assets
// or venues
func (r *RegistryBorrowMarketService) List(ctx context.Context, body RegistryBorrowMarketListParams, opts ...option.RequestOption) (res *RegistryBorrowMarketListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/registry/borrow/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryBorrowMarketListResponse struct {
	// A list of registry data for lend borrow markets
	Items []RegistryLendBorrowMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryBorrowMarketListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryBorrowMarketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryBorrowMarketListParams struct {
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	ID param.Opt[string] `json:"id,omitzero"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain param.Opt[BlockchainID] `json:"blockchain,omitzero"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero"`
	paramObj
}

func (r RegistryBorrowMarketListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryBorrowMarketListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryBorrowMarketListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
