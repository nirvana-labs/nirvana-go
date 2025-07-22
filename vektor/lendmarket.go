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

// LendMarketService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendMarketService] method instead.
type LendMarketService struct {
	Options []option.RequestOption
}

// NewLendMarketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLendMarketService(opts ...option.RequestOption) (r LendMarketService) {
	r = LendMarketService{}
	r.Options = opts
	return
}

// Get the current market rates for lending an asset
func (r *LendMarketService) List(ctx context.Context, body LendMarketListParams, opts ...option.RequestOption) (res *LendMarketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lend/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the current market rates for lending an asset
func (r *LendMarketService) ListHistorical(ctx context.Context, body LendMarketListHistoricalParams, opts ...option.RequestOption) (res *LendMarketListHistoricalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lend/markets/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LendMarketListResponse struct {
	// A list of lend markets
	Items []LendMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendMarketListResponse) RawJSON() string { return r.JSON.raw }
func (r *LendMarketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendMarketListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                 `json:"historical,required"`
	Items      []LendMarketListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendMarketListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *LendMarketListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendMarketListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of lend markets
	Items []LendMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendMarketListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LendMarketListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendMarketListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend markets by venues and assets
	OfLendMarketsByVenuesAssetsRequestInput *LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend markets by IDs
	OfLendMarketsByIDsRequestInput *LendMarketListParamsBodyLendMarketsByIDsRequestInput `json:",inline"`

	paramObj
}

func (u LendMarketListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLendMarketsByVenuesAssetsRequestInput, u.OfLendMarketsByIDsRequestInput)
}
func (r *LendMarketListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend markets by venues and assets
//
// The properties Assets, Blockchain, Venues are required.
type LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInput struct {
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LendMarketListParamsBodyLendMarketsByVenuesAssetsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Parameters for querying lend markets by IDs
//
// The properties Blockchain, MarketIDs are required.
type LendMarketListParamsBodyLendMarketsByIDsRequestInput struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of lend/borrow market IDs
	MarketIDs []LendBorrowMarketID `json:"market_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LendMarketListParamsBodyLendMarketsByIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LendMarketListParamsBodyLendMarketsByIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LendMarketListParamsBodyLendMarketsByIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendMarketListParamsBodyLendMarketsByIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LendMarketListParamsBodyLendMarketsByIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LendMarketListParamsBodyLendMarketsByIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LendMarketListParamsBodyLendMarketsByIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LendMarketListParamsBodyLendMarketsByIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type LendMarketListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend markets by venues and assets
	OfHistoricalLendMarketsByVenuesAssetsRequest *LendMarketListHistoricalParamsBodyHistoricalLendMarketsByVenuesAssetsRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying lend markets by IDs
	OfHistoricalLendMarketsByIDsRequest *LendMarketListHistoricalParamsBodyHistoricalLendMarketsByIDsRequest `json:",inline"`

	paramObj
}

func (u LendMarketListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalLendMarketsByVenuesAssetsRequest, u.OfHistoricalLendMarketsByIDsRequest)
}
func (r *LendMarketListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend markets by venues and assets
//
// The properties Assets, Blockchain, From, To, Venues are required.
type LendMarketListHistoricalParamsBodyHistoricalLendMarketsByVenuesAssetsRequest struct {
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

func (r LendMarketListHistoricalParamsBodyHistoricalLendMarketsByVenuesAssetsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LendMarketListHistoricalParamsBodyHistoricalLendMarketsByVenuesAssetsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendMarketListHistoricalParamsBodyHistoricalLendMarketsByVenuesAssetsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying lend markets by IDs
//
// The properties Blockchain, From, MarketIDs, To are required.
type LendMarketListHistoricalParamsBodyHistoricalLendMarketsByIDsRequest struct {
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

func (r LendMarketListHistoricalParamsBodyHistoricalLendMarketsByIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LendMarketListHistoricalParamsBodyHistoricalLendMarketsByIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LendMarketListHistoricalParamsBodyHistoricalLendMarketsByIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
