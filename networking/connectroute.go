// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

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

// ConnectRouteService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectRouteService] method instead.
type ConnectRouteService struct {
	Options []option.RequestOption
}

// NewConnectRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectRouteService(opts ...option.RequestOption) (r ConnectRouteService) {
	r = ConnectRouteService{}
	r.Options = opts
	return
}

// List all supported routes with regions for Connect.
func (r *ConnectRouteService) List(ctx context.Context, query ConnectRouteListParams, opts ...option.RequestOption) (res *pagination.Cursor[ConnectRoute], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/networking/connect/routes"
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

// List all supported routes with regions for Connect.
func (r *ConnectRouteService) ListAutoPaging(ctx context.Context, query ConnectRouteListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ConnectRoute] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

type ConnectRouteListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectRouteListParams]'s query parameters as `url.Values`.
func (r ConnectRouteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
