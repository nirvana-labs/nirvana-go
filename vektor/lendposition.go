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

// LendPositionService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendPositionService] method instead.
type LendPositionService struct {
	Options []option.RequestOption
}

// NewLendPositionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLendPositionService(opts ...option.RequestOption) (r LendPositionService) {
	r = LendPositionService{}
	r.Options = opts
	return
}

// Get info on lending positions
func (r *LendPositionService) List(ctx context.Context, body LendPositionListParams, opts ...option.RequestOption) (res *LendPositionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lend/positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get info on lending positions
func (r *LendPositionService) ListHistorical(ctx context.Context, body LendPositionListHistoricalParams, opts ...option.RequestOption) (res *LendPositionListHistoricalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lend/positions/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendPositionListResponse struct {
	// A list of lend positions
	Items []LendPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendPositionListResponse) RawJSON() string { return r.JSON.raw }
func (r *LendPositionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendPositionListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                   `json:"historical,required"`
	Items      []LendPositionListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendPositionListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *LendPositionListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendPositionListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of lend positions
	Items []LendPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendPositionListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LendPositionListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendPositionListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend positions by venues and assets
	OfLendPositionsByVenuesAssetsRequestInput *LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend positions by IDs
	OfLendPositionsByIDsRequestInput *LendPositionListParamsBodyLendPositionsByIDsRequestInput `json:",inline"`

	paramObj
}

func (u LendPositionListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLendPositionsByVenuesAssetsRequestInput, u.OfLendPositionsByIDsRequestInput)
}
func (r *LendPositionListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend positions by venues and assets
//
// The properties Accounts, Assets, Blockchain, Venues are required.
type LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInput struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LendPositionListParamsBodyLendPositionsByVenuesAssetsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Parameters for querying lend positions by IDs
//
// The properties Accounts, Blockchain, MarketIDs are required.
type LendPositionListParamsBodyLendPositionsByIDsRequestInput struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of lend/borrow market IDs
	MarketIDs []LendBorrowMarketID `json:"market_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LendPositionListParamsBodyLendPositionsByIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LendPositionListParamsBodyLendPositionsByIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LendPositionListParamsBodyLendPositionsByIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendPositionListParamsBodyLendPositionsByIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LendPositionListParamsBodyLendPositionsByIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LendPositionListParamsBodyLendPositionsByIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LendPositionListParamsBodyLendPositionsByIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LendPositionListParamsBodyLendPositionsByIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type LendPositionListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend positions by venues and assets
	OfHistoricalLendPositionsByVenuesAssetsRequest *LendPositionListHistoricalParamsBodyHistoricalLendPositionsByVenuesAssetsRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend positions by IDs
	OfHistoricalLendPositionsByIDsRequest *LendPositionListHistoricalParamsBodyHistoricalLendPositionsByIDsRequest `json:",inline"`

	paramObj
}

func (u LendPositionListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalLendPositionsByVenuesAssetsRequest, u.OfHistoricalLendPositionsByIDsRequest)
}
func (r *LendPositionListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend positions by venues and assets
//
// The properties Accounts, Assets, Blockchain, From, To, Venues are required.
type LendPositionListHistoricalParamsBodyHistoricalLendPositionsByVenuesAssetsRequest struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// Either a ISO8601 timestamp or a block number
	From TimestampOrBlockNumberUnionParam `json:"from,omitzero,required"`
	// Either a ISO8601 timestamp or a block number
	To TimestampOrBlockNumberUnionParam `json:"to,omitzero,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r LendPositionListHistoricalParamsBodyHistoricalLendPositionsByVenuesAssetsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LendPositionListHistoricalParamsBodyHistoricalLendPositionsByVenuesAssetsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendPositionListHistoricalParamsBodyHistoricalLendPositionsByVenuesAssetsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend positions by IDs
//
// The properties Accounts, Blockchain, From, MarketIDs, To are required.
type LendPositionListHistoricalParamsBodyHistoricalLendPositionsByIDsRequest struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// Either a ISO8601 timestamp or a block number
	From TimestampOrBlockNumberUnionParam `json:"from,omitzero,required"`
	// A list of lend/borrow market IDs
	MarketIDs []LendBorrowMarketID `json:"market_ids,omitzero,required"`
	// Either a ISO8601 timestamp or a block number
	To TimestampOrBlockNumberUnionParam `json:"to,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r LendPositionListHistoricalParamsBodyHistoricalLendPositionsByIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LendPositionListHistoricalParamsBodyHistoricalLendPositionsByIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendPositionListHistoricalParamsBodyHistoricalLendPositionsByIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
