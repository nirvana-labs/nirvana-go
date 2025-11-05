// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
)

// FlexBlockchainService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFlexBlockchainService] method instead.
type FlexBlockchainService struct {
	Options []option.RequestOption
}

// NewFlexBlockchainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFlexBlockchainService(opts ...option.RequestOption) (r FlexBlockchainService) {
	r = FlexBlockchainService{}
	r.Options = opts
	return
}

// List all Flex Blockchains
func (r *FlexBlockchainService) List(ctx context.Context, query FlexBlockchainListParams, opts ...option.RequestOption) (res *pagination.Cursor[FlexBlockchain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/rpc_nodes/flex/blockchains"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all Flex Blockchains
func (r *FlexBlockchainService) ListAutoPaging(ctx context.Context, query FlexBlockchainListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[FlexBlockchain] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

type FlexBlockchainListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FlexBlockchainListParams]'s query parameters as
// `url.Values`.
func (r FlexBlockchainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
