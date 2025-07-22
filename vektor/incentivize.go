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

// IncentivizeService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIncentivizeService] method instead.
type IncentivizeService struct {
	Options []option.RequestOption
}

// NewIncentivizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIncentivizeService(opts ...option.RequestOption) (r IncentivizeService) {
	r = IncentivizeService{}
	r.Options = opts
	return
}

// List the current incentive markets
func (r *IncentivizeService) List(ctx context.Context, body IncentivizeListParams, opts ...option.RequestOption) (res *IncentivizeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/incentivize/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IncentivizeListResponse struct {
	// A list of lp pool incentive markets
	Items []IncentivizeMarket `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncentivizeListResponse) RawJSON() string { return r.JSON.raw }
func (r *IncentivizeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IncentivizeListParams struct {
	// A list of asset IDs, EVM addresses or asset symbols
	Assets []AssetIDOrAddressEVMOrAssetSymbol `json:"assets,omitzero,required"`
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	Blockchain BlockchainIDOrBlockchainSymbol `json:"blockchain,required"`
	// A list of venue IDs or venue symbols
	Venues []VenueIDOrVenueSymbol `json:"venues,omitzero,required"`
	// An asset symbol
	QuoteAssetSymbol param.Opt[string] `json:"quote_asset_symbol,omitzero"`
	paramObj
}

func (r IncentivizeListParams) MarshalJSON() (data []byte, err error) {
	type shadow IncentivizeListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IncentivizeListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
