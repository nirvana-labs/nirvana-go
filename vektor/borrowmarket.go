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

// BorrowMarketService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBorrowMarketService] method instead.
type BorrowMarketService struct {
	Options []option.RequestOption
}

// NewBorrowMarketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBorrowMarketService(opts ...option.RequestOption) (r BorrowMarketService) {
	r = BorrowMarketService{}
	r.Options = opts
	return
}

// Get the current market rates for borrowing an asset
func (r *BorrowMarketService) List(ctx context.Context, body BorrowMarketListParams, opts ...option.RequestOption) (res *BorrowMarketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current market rates for borrowing an asset
func (r *BorrowMarketService) ListHistorical(ctx context.Context, body BorrowMarketListHistoricalParams, opts ...option.RequestOption) (res *BorrowMarketListHistoricalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/markets/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BorrowMarketListResponse struct {
	// A list of borrow markets
	Items []BorrowMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowMarketListResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowMarketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowMarketListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                   `json:"historical,required"`
	Items      []BorrowMarketListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowMarketListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowMarketListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowMarketListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of borrow markets
	Items []BorrowMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowMarketListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BorrowMarketListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowMarketListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow markets by venues and assets
	OfBorrowMarketsByVenuesAssetsRequestInput *BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow markets by IDs
	OfBorrowMarketsByIDsRequestInput *BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInput `json:",inline"`

	paramObj
}

func (u BorrowMarketListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBorrowMarketsByVenuesAssetsRequestInput, u.OfBorrowMarketsByIDsRequestInput)
}
func (r *BorrowMarketListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow markets by venues and assets
//
// The properties Assets, Blockchain, Venues are required.
type BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInput struct {
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BorrowMarketListParamsBodyBorrowMarketsByVenuesAssetsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Parameters for querying borrow markets by IDs
//
// The properties Blockchain, MarketIDs are required.
type BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInput struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of lend/borrow market IDs
	MarketIDs []LendBorrowMarketID `json:"market_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BorrowMarketListParamsBodyBorrowMarketsByIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type BorrowMarketListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow markets by venues and assets
	OfHistoricalBorrowMarketsByVenuesAssetsRequest *BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByVenuesAssetsRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying borrow markets by IDs
	OfHistoricalBorrowMarketsByIDsRequest *BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByIDsRequest `json:",inline"`

	paramObj
}

func (u BorrowMarketListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalBorrowMarketsByVenuesAssetsRequest, u.OfHistoricalBorrowMarketsByIDsRequest)
}
func (r *BorrowMarketListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow markets by venues and assets
//
// The properties Assets, Blockchain, From, To, Venues are required.
type BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByVenuesAssetsRequest struct {
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

func (r BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByVenuesAssetsRequest) MarshalJSON() (data []byte, err error) {
	type shadow BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByVenuesAssetsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByVenuesAssetsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying borrow markets by IDs
//
// The properties Blockchain, From, MarketIDs, To are required.
type BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByIDsRequest struct {
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

func (r BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowMarketListHistoricalParamsBodyHistoricalBorrowMarketsByIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
