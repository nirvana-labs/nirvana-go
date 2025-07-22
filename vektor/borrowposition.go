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

// BorrowPositionService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBorrowPositionService] method instead.
type BorrowPositionService struct {
	Options []option.RequestOption
}

// NewBorrowPositionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBorrowPositionService(opts ...option.RequestOption) (r BorrowPositionService) {
	r = BorrowPositionService{}
	r.Options = opts
	return
}

// Get info on borrowed positions
func (r *BorrowPositionService) List(ctx context.Context, body BorrowPositionListParams, opts ...option.RequestOption) (res *BorrowPositionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get info on borrowed positions
func (r *BorrowPositionService) ListHistorical(ctx context.Context, body BorrowPositionListHistoricalParams, opts ...option.RequestOption) (res *BorrowPositionListHistoricalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/positions/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BorrowPositionListResponse struct {
	// A list of borrow positions
	Items []BorrowPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowPositionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowPositionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowPositionListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                     `json:"historical,required"`
	Items      []BorrowPositionListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowPositionListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowPositionListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowPositionListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of borrow positions
	Items []BorrowPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowPositionListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BorrowPositionListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowPositionListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow positions by venues and assets
	OfBorrowPositionsByVenuesAssetsRequestInput *BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow positions by IDs
	OfBorrowPositionsByIDsRequestInput *BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInput `json:",inline"`

	paramObj
}

func (u BorrowPositionListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBorrowPositionsByVenuesAssetsRequestInput, u.OfBorrowPositionsByIDsRequestInput)
}
func (r *BorrowPositionListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow positions by venues and assets
//
// The properties Accounts, Assets, Blockchain, Venues are required.
type BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInput struct {
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
	At BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BorrowPositionListParamsBodyBorrowPositionsByVenuesAssetsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Parameters for querying borrow positions by IDs
//
// The properties Blockchain, MarketIDs are required.
type BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInput struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of lend/borrow market IDs
	MarketIDs []LendBorrowMarketID `json:"market_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BorrowPositionListParamsBodyBorrowPositionsByIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type BorrowPositionListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow positions by venues and assets
	OfHistoricalBorrowPositionsByVenuesAssetsRequest *BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByVenuesAssetsRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow positions by IDs
	OfHistoricalBorrowPositionsByIDsRequest *BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByIDsRequest `json:",inline"`

	paramObj
}

func (u BorrowPositionListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalBorrowPositionsByVenuesAssetsRequest, u.OfHistoricalBorrowPositionsByIDsRequest)
}
func (r *BorrowPositionListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow positions by venues and assets
//
// The properties Accounts, Assets, Blockchain, From, To, Venues are required.
type BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByVenuesAssetsRequest struct {
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

func (r BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByVenuesAssetsRequest) MarshalJSON() (data []byte, err error) {
	type shadow BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByVenuesAssetsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByVenuesAssetsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow positions by IDs
//
// The properties Blockchain, From, MarketIDs, To are required.
type BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByIDsRequest struct {
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
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero"`
	paramObj
}

func (r BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowPositionListHistoricalParamsBodyHistoricalBorrowPositionsByIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
