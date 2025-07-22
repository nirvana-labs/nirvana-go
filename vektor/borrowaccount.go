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

// BorrowAccountService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBorrowAccountService] method instead.
type BorrowAccountService struct {
	Options []option.RequestOption
}

// NewBorrowAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBorrowAccountService(opts ...option.RequestOption) (r BorrowAccountService) {
	r = BorrowAccountService{}
	r.Options = opts
	return
}

// Get account-level borrow info on specified accounts
func (r *BorrowAccountService) List(ctx context.Context, body BorrowAccountListParams, opts ...option.RequestOption) (res *BorrowAccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get account-level borrow info on specified accounts
func (r *BorrowAccountService) ListHistorical(ctx context.Context, body BorrowAccountListHistoricalParams, opts ...option.RequestOption) (res *BorrowAccountListHistoricalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/borrow/accounts/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BorrowAccountListResponse struct {
	// A list of borrow accounts
	Items []BorrowAccount `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowAccountListHistoricalResponse struct {
	// A range of blockstamps
	Historical OnChainHistoricalRange                    `json:"historical,required"`
	Items      []BorrowAccountListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowAccountListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *BorrowAccountListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowAccountListHistoricalResponseItem struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of borrow accounts
	Items []BorrowAccount `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowAccountListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *BorrowAccountListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BorrowAccountListParams struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	// Either a ISO8601 timestamp or a block number
	At BorrowAccountListParamsAtUnion `json:"at,omitzero"`
	paramObj
}

func (r BorrowAccountListParams) MarshalJSON() (data []byte, err error) {
	type shadow BorrowAccountListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowAccountListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BorrowAccountListParamsAtUnion struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u BorrowAccountListParamsAtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *BorrowAccountListParamsAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BorrowAccountListParamsAtUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type BorrowAccountListHistoricalParams struct {
	// A list of accounts. Currently only EVM addresses are supported.
	Accounts []Account `json:"accounts,omitzero,required"`
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

func (r BorrowAccountListHistoricalParams) MarshalJSON() (data []byte, err error) {
	type shadow BorrowAccountListHistoricalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BorrowAccountListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
