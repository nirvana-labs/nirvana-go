// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/operations"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// ClusterPoolNodeService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPoolNodeService] method instead.
type ClusterPoolNodeService struct {
	Options []option.RequestOption
	Volumes ClusterPoolNodeVolumeService
}

// NewClusterPoolNodeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClusterPoolNodeService(opts ...option.RequestOption) (r ClusterPoolNodeService) {
	r = ClusterPoolNodeService{}
	r.Options = opts
	r.Volumes = NewClusterPoolNodeVolumeService(opts...)
	return
}

// List all nodes in an NKS node pool
func (r *ClusterPoolNodeService) List(ctx context.Context, clusterID string, poolID string, query ClusterPoolNodeListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSNode], err error) {
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
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/nodes", clusterID, poolID)
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

// List all nodes in an NKS node pool
func (r *ClusterPoolNodeService) ListAutoPaging(ctx context.Context, clusterID string, poolID string, query ClusterPoolNodeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSNode] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, poolID, query, opts...))
}

// Delete a single node from an NKS node pool
func (r *ClusterPoolNodeService) Delete(ctx context.Context, clusterID string, poolID string, nodeID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/nodes/%s", clusterID, poolID, nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get details about an NKS node
func (r *ClusterPoolNodeService) Get(ctx context.Context, clusterID string, poolID string, nodeID string, opts ...option.RequestOption) (res *NKSNode, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/nodes/%s", clusterID, poolID, nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS node details.
type NKSNode struct {
	// Unique identifier for the node.
	ID string `json:"id" api:"required"`
	// When the node was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the node.
	Name string `json:"name" api:"required"`
	// Private IP address of the node.
	PrivateIP string `json:"private_ip" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the node was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		PrivateIP   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSNode) RawJSON() string { return r.JSON.raw }
func (r *NKSNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSNodeList struct {
	Items []NKSNode `json:"items" api:"required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSNodeList) RawJSON() string { return r.JSON.raw }
func (r *NKSNodeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolNodeListParams struct {
	// Pagination cursor returned by a previous request. Only valid for the same
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by whether a private address has been assigned yet
	HasPrivateIP param.Opt[bool] `query:"has_private_ip,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the node name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by the node's private address
	PrivateIP param.Opt[string] `query:"private_ip,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at, name, status
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by node status
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "error".
	Status ClusterPoolNodeListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterPoolNodeListParams]'s query parameters as
// `url.Values`.
func (r ClusterPoolNodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by node status
type ClusterPoolNodeListParamsStatus string

const (
	ClusterPoolNodeListParamsStatusPending  ClusterPoolNodeListParamsStatus = "pending"
	ClusterPoolNodeListParamsStatusCreating ClusterPoolNodeListParamsStatus = "creating"
	ClusterPoolNodeListParamsStatusUpdating ClusterPoolNodeListParamsStatus = "updating"
	ClusterPoolNodeListParamsStatusReady    ClusterPoolNodeListParamsStatus = "ready"
	ClusterPoolNodeListParamsStatusDeleting ClusterPoolNodeListParamsStatus = "deleting"
	ClusterPoolNodeListParamsStatusError    ClusterPoolNodeListParamsStatus = "error"
)
