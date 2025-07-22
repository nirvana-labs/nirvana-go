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

// LPPositionService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLPPositionService] method instead.
type LPPositionService struct {
	Options []option.RequestOption
}

// NewLPPositionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLPPositionService(opts ...option.RequestOption) (r LPPositionService) {
	r = LPPositionService{}
	r.Options = opts
	return
}

// Get info on LP positions
func (r *LPPositionService) List(ctx context.Context, body LPPositionListParams, opts ...option.RequestOption) (res *LPPositionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lp/positions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get info on LP positions
func (r *LPPositionService) ListHistorical(ctx context.Context, body LPPositionListHistoricalParams, opts ...option.RequestOption) (res *LPPositionListHistoricalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/lp/positions/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LPPositionListResponse struct {
	// A list of liquidity pool positions
	Items []LPPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPositionListResponse) RawJSON() string { return r.JSON.raw }
func (r *LPPositionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPositionListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                 `json:"historical,required"`
	Items      []LPPositionListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPositionListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *LPPositionListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPositionListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of liquidity pool positions
	Items []LPPosition `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPositionListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LPPositionListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPositionListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A
	// request to query positions by blockchain, venues and assets
	OfLPPositionsByVenueAssetRequestInput *LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set. A
	// request to query positions by pool ids
	OfLPPositionsByPoolIDsRequestInput *LPPositionListParamsBodyLPPositionsByPoolIDsRequestInput `json:",inline"`

	paramObj
}

func (u LPPositionListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLPPositionsByVenueAssetRequestInput, u.OfLPPositionsByPoolIDsRequestInput)
}
func (r *LPPositionListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to query positions by blockchain, venues and assets
//
// The properties Accounts, Assets, Blockchain, Venues are required.
type LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput struct {
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
	// Whether to exclude empty positions
	ExcludeZeros param.Opt[bool] `json:"exclude_zeros,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPositionListParamsBodyLPPositionsByVenueAssetRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LPPositionListParamsBodyLPPositionsByVenueAssetRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// A request to query positions by pool ids
//
// The properties Accounts, Blockchain, LPPoolIDs are required.
type LPPositionListParamsBodyLPPositionsByPoolIDsRequestInput struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of LP pool IDs
	LPPoolIDs []string `json:"lp_pool_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Whether to exclude empty positions
	ExcludeZeros param.Opt[bool] `json:"exclude_zeros,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LPPositionListParamsBodyLPPositionsByPoolIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LPPositionListParamsBodyLPPositionsByPoolIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPPositionListParamsBodyLPPositionsByPoolIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPositionListParamsBodyLPPositionsByPoolIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LPPositionListParamsBodyLPPositionsByPoolIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LPPositionListParamsBodyLPPositionsByPoolIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LPPositionListParamsBodyLPPositionsByPoolIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LPPositionListParamsBodyLPPositionsByPoolIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type LPPositionListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A
	// request to query positions by blockchain, venues and assets
	OfHistoricalLPPositionsByVenueAssetRequest *LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set. A
	// request to query positions by pool ids
	OfHistoricalLPPositionsByPoolIDsRequest *LPPositionListHistoricalParamsBodyHistoricalLPPositionsByPoolIDsRequest `json:",inline"`

	paramObj
}

func (u LPPositionListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalLPPositionsByVenueAssetRequest, u.OfHistoricalLPPositionsByPoolIDsRequest)
}
func (r *LPPositionListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to query positions by blockchain, venues and assets
//
// The properties Accounts, Assets, Blockchain, From, To, Venues are required.
type LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest struct {
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
	// Whether to exclude empty positions
	ExcludeZeros param.Opt[bool] `json:"exclude_zeros,omitzero"`
	paramObj
}

func (r LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest) MarshalJSON() (data []byte, err error) {
	type shadow LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPositionListHistoricalParamsBodyHistoricalLPPositionsByVenueAssetRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to query positions by pool ids
//
// The properties Accounts, Blockchain, From, LPPoolIDs, To are required.
type LPPositionListHistoricalParamsBodyHistoricalLPPositionsByPoolIDsRequest struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// Either a ISO8601 timestamp or a block number
	From TimestampOrBlockNumberUnionParam `json:"from,omitzero,required"`
	// A list of LP pool IDs
	LPPoolIDs []string `json:"lp_pool_ids,omitzero,required"`
	// Either a ISO8601 timestamp or a block number
	To TimestampOrBlockNumberUnionParam `json:"to,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Whether to exclude empty positions
	ExcludeZeros param.Opt[bool] `json:"exclude_zeros,omitzero"`
	paramObj
}

func (r LPPositionListHistoricalParamsBodyHistoricalLPPositionsByPoolIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LPPositionListHistoricalParamsBodyHistoricalLPPositionsByPoolIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPositionListHistoricalParamsBodyHistoricalLPPositionsByPoolIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
