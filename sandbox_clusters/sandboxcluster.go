// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sandbox_clusters

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/v3/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v3/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v3/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v3/operations"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v3/shared"
)

// SandboxClusterService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxClusterService] method instead.
type SandboxClusterService struct {
	Options []option.RequestOption
}

// NewSandboxClusterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxClusterService(opts ...option.RequestOption) (r SandboxClusterService) {
	r = SandboxClusterService{}
	r.Options = opts
	return
}

// Create a Sandbox Cluster to run sandboxes on. The infrastructure it runs on is
// provisioned in a project the platform manages.
func (r *SandboxClusterService) New(ctx context.Context, body SandboxClusterNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sandbox_clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a Sandbox Cluster's name and tags
func (r *SandboxClusterService) Update(ctx context.Context, sandboxClusterID string, body SandboxClusterUpdateParams, opts ...option.RequestOption) (res *SandboxCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxClusterID == "" {
		err = errors.New("missing required sandbox_cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_clusters/%s", sandboxClusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the Sandbox Clusters in a project
func (r *SandboxClusterService) List(ctx context.Context, query SandboxClusterListParams, opts ...option.RequestOption) (res *pagination.Cursor[SandboxCluster], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/sandbox_clusters"
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

// List the Sandbox Clusters in a project
func (r *SandboxClusterService) ListAutoPaging(ctx context.Context, query SandboxClusterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[SandboxCluster] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a Sandbox Cluster, with its sandboxes and the infrastructure they run on.
func (r *SandboxClusterService) Delete(ctx context.Context, sandboxClusterID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxClusterID == "" {
		err = errors.New("missing required sandbox_cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_clusters/%s", sandboxClusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a Sandbox Cluster
func (r *SandboxClusterService) Get(ctx context.Context, sandboxClusterID string, opts ...option.RequestOption) (res *SandboxCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxClusterID == "" {
		err = errors.New("missing required sandbox_cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_clusters/%s", sandboxClusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A cluster sandboxes run on, in one region.
type SandboxCluster struct {
	// Unique identifier for the Sandbox Cluster.
	ID string `json:"id" api:"required"`
	// When the Sandbox Cluster was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Hostname sandboxes' endpoints are served on over HTTPS, to allow in outbound
	// firewalls. Null until assigned; a recreated Sandbox Cluster gets a new one.
	GatewayHostname string `json:"gateway_hostname" api:"required"`
	// Instance type of the nodes the sandboxes run on.
	InstanceType string `json:"instance_type" api:"required"`
	// Name of the Sandbox Cluster.
	Name string `json:"name" api:"required"`
	// Number of nodes the sandboxes run on.
	NodeCount int64 `json:"node_count" api:"required"`
	// ID of the project the Sandbox Cluster is in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Tags attached to the Sandbox Cluster.
	Tags []string `json:"tags" api:"required"`
	// When the Sandbox Cluster was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		GatewayHostname respjson.Field
		InstanceType    respjson.Field
		Name            respjson.Field
		NodeCount       respjson.Field
		ProjectID       respjson.Field
		Region          respjson.Field
		Status          respjson.Field
		Tags            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxCluster) RawJSON() string { return r.JSON.raw }
func (r *SandboxCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxClusterList struct {
	Items []SandboxCluster `json:"items" api:"required"`
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
func (r SandboxClusterList) RawJSON() string { return r.JSON.raw }
func (r *SandboxClusterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxClusterNewParams struct {
	// Instance type of the nodes the sandboxes run on.
	InstanceType string `json:"instance_type" api:"required"`
	// Name of the Sandbox Cluster. Unique within its project.
	Name string `json:"name" api:"required"`
	// Number of nodes the sandboxes run on.
	NodeCount int64 `json:"node_count" api:"required"`
	// ID of the project to create the Sandbox Cluster in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Tags to attach to the Sandbox Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SandboxClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SandboxClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxClusterUpdateParams struct {
	// Name of the Sandbox Cluster. Unique within its project.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Sandbox Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SandboxClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SandboxClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxClusterListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id" api:"required" json:"-"`
	// Pagination cursor returned by a previous request. Only valid for the same sort
	// order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the Sandbox Cluster name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by region
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by tags. Repeat the parameter to require several tags; a Sandbox Cluster
	// must carry all of them.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxClusterListParams]'s query parameters as
// `url.Values`.
func (r SandboxClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
