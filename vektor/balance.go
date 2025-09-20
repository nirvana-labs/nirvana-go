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

// BalanceService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r BalanceService) {
	r = BalanceService{}
	r.Options = opts
	return
}

// List balances for a given set of assets and addresses
func (r *BalanceService) List(ctx context.Context, body BalanceListParams, opts ...option.RequestOption) (res *BalanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/balances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List balances for a given set of assets and addresses
func (r *BalanceService) ListHistorical(ctx context.Context, body BalanceListHistoricalParams, opts ...option.RequestOption) (res *BalanceListHistoricalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/balances/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BalanceListResponse struct {
	// Array of Balance
	Items []Balance `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange              `json:"historical,required"`
	Items      []BalanceListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// Array of Balance
	Items []Balance `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BalanceListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BalanceListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceListParams struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[AssetSymbol] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At BalanceListParamsAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BalanceListParams) MarshalJSON() (data []byte, err error) {
	type shadow BalanceListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BalanceListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BalanceListParamsAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BalanceListParamsAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BalanceListParamsAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BalanceListParamsAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type BalanceListHistoricalParams struct {
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
	// An asset symbol
	QuoteAssetSymbol param.Opt[AssetSymbol] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r BalanceListHistoricalParams) MarshalJSON() (data []byte, err error) {
	type shadow BalanceListHistoricalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BalanceListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
