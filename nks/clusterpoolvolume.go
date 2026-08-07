// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
)

// ClusterPoolVolumeService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPoolVolumeService] method instead.
type ClusterPoolVolumeService struct {
	Options []option.RequestOption
}

// NewClusterPoolVolumeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterPoolVolumeService(opts ...option.RequestOption) (r ClusterPoolVolumeService) {
	r = ClusterPoolVolumeService{}
	r.Options = opts
	return
}

// List all volumes attached to the nodes of an NKS node pool
func (r *ClusterPoolVolumeService) List(ctx context.Context, clusterID string, poolID string, query ClusterPoolVolumeListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSNodeVolume], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/volumes", clusterID, poolID)
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

// List all volumes attached to the nodes of an NKS node pool
func (r *ClusterPoolVolumeService) ListAutoPaging(ctx context.Context, clusterID string, poolID string, query ClusterPoolVolumeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSNodeVolume] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, poolID, query, opts...))
}

type ClusterPoolVolumeListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterPoolVolumeListParams]'s query parameters as
// `url.Values`.
func (r ClusterPoolVolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
