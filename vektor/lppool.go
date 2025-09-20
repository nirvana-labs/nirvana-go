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

// LPPoolService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLPPoolService] method instead.
type LPPoolService struct {
	Options []option.RequestOption
}

// NewLPPoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLPPoolService(opts ...option.RequestOption) (r LPPoolService) {
	r = LPPoolService{}
	r.Options = opts
	return
}

// Get a list of LP pools
func (r *LPPoolService) List(ctx context.Context, body LPPoolListParams, opts ...option.RequestOption) (res *LPPoolListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lp/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of LP pools
func (r *LPPoolService) ListHistorical(ctx context.Context, body LPPoolListHistoricalParams, opts ...option.RequestOption) (res *LPPoolListHistoricalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/lp/pools/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LPPoolListResponse struct {
	// A list of liquidity pools
	Items []LPPool `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPoolListResponse) RawJSON() string { return r.JSON.raw }
func (r *LPPoolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPoolListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange             `json:"historical,required"`
	Items      []LPPoolListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPoolListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *LPPoolListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPoolListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of liquidity pools
	Items []LPPool `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPoolListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LPPoolListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LPPoolListParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying LP pools by venues and assets
	OfLPPoolsByVenueAssetRequestInput *LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying LP pools by IDs
	OfLPPoolsByIDsRequestInput *LPPoolListParamsBodyLPPoolsByIDsRequestInput `json:",inline"`

	paramObj
}

func (u LPPoolListParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLPPoolsByVenueAssetRequestInput, u.OfLPPoolsByIDsRequestInput)
}
func (r *LPPoolListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying LP pools by venues and assets
//
// The properties Assets, Blockchain, Venues are required.
type LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput struct {
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPoolListParamsBodyLPPoolsByVenueAssetRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LPPoolListParamsBodyLPPoolsByVenueAssetRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Parameters for querying LP pools by IDs
//
// The properties Blockchain, LPPoolIDs are required.
type LPPoolListParamsBodyLPPoolsByIDsRequestInput struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of LP pool IDs
	LPPoolIDs []string `json:"lp_pool_ids,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At LPPoolListParamsBodyLPPoolsByIDsRequestInputAtUnion `json:"at,omitzero"`
	paramObj
}

func (r LPPoolListParamsBodyLPPoolsByIDsRequestInput) MarshalJSON() (data []byte, err error) {
	type shadow LPPoolListParamsBodyLPPoolsByIDsRequestInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPoolListParamsBodyLPPoolsByIDsRequestInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LPPoolListParamsBodyLPPoolsByIDsRequestInputAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u LPPoolListParamsBodyLPPoolsByIDsRequestInputAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *LPPoolListParamsBodyLPPoolsByIDsRequestInputAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LPPoolListParamsBodyLPPoolsByIDsRequestInputAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type LPPoolListHistoricalParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying LP pools by venues and assets
	OfHistoricalLPPoolsByVenueAssetRequest *LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Parameters for querying LP pools by IDs
	OfHistoricalLPPoolsByIDsRequest *LPPoolListHistoricalParamsBodyHistoricalLPPoolsByIDsRequest `json:",inline"`

	paramObj
}

func (u LPPoolListHistoricalParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHistoricalLPPoolsByVenueAssetRequest, u.OfHistoricalLPPoolsByIDsRequest)
}
func (r *LPPoolListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying LP pools by venues and assets
//
// The properties Assets, Blockchain, From, To, Venues are required.
type LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest struct {
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

func (r LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest) MarshalJSON() (data []byte, err error) {
	type shadow LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPoolListHistoricalParamsBodyHistoricalLPPoolsByVenueAssetRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for querying LP pools by IDs
//
// The properties Blockchain, From, LPPoolIDs, To are required.
type LPPoolListHistoricalParamsBodyHistoricalLPPoolsByIDsRequest struct {
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
	paramObj
}

func (r LPPoolListHistoricalParamsBodyHistoricalLPPoolsByIDsRequest) MarshalJSON() (data []byte, err error) {
	type shadow LPPoolListHistoricalParamsBodyHistoricalLPPoolsByIDsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LPPoolListHistoricalParamsBodyHistoricalLPPoolsByIDsRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
