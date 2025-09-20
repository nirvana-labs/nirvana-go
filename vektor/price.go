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

// PriceService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPriceService] method instead.
type PriceService struct {
	Options []option.RequestOption
}

// NewPriceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPriceService(opts ...option.RequestOption) (r PriceService) {
	r = PriceService{}
	r.Options = opts
	return
}

// Get a list of asset prices
func (r *PriceService) List(ctx context.Context, body PriceListParams, opts ...option.RequestOption) (res *PriceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/prices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of asset prices
func (r *PriceService) ListHistorical(ctx context.Context, body PriceListHistoricalParams, opts ...option.RequestOption) (res *PriceListHistoricalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/prices/historical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PriceListResponse struct {
	// Array of price
	Items []Price `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceListResponse) RawJSON() string { return r.JSON.raw }
func (r *PriceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceListHistoricalResponse struct {
	// A range of timestamps
	Historical OffChainHistoricalRange           `json:"historical,required"`
	Items      []PriceListHistoricalResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Historical  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceListHistoricalResponse) RawJSON() string { return r.JSON.raw }
func (r *PriceListHistoricalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceListHistoricalResponseItem struct {
	// Array of price
	Items []Price `json:"items,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceListHistoricalResponseItem) RawJSON() string { return r.JSON.raw }
func (r *PriceListHistoricalResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceListParams struct {
	// A list of asset symbols
	AssetSymbols []AssetSymbol `json:"asset_symbols,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[AssetSymbol] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r PriceListParams) MarshalJSON() (data []byte, err error) {
	type shadow PriceListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceListHistoricalParams struct {
	// A list of asset symbols
	AssetSymbols []AssetSymbol `json:"asset_symbols,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[AssetSymbol] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r PriceListHistoricalParams) MarshalJSON() (data []byte, err error) {
	type shadow PriceListHistoricalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceListHistoricalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
