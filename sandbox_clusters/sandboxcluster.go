// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sandbox_clusters

import (
	"context"
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

// List the organization's sandbox clusters, one per region where sandboxes are
// enabled
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

// List the organization's sandbox clusters, one per region where sandboxes are
// enabled
func (r *SandboxClusterService) ListAutoPaging(ctx context.Context, query SandboxClusterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[SandboxCluster] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Disable sandboxes in a region. This deletes the sandboxes and the infrastructure
// they run on.
func (r *SandboxClusterService) Disable(ctx context.Context, region SandboxClusterDisableParamsRegion, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/sandbox_clusters/%v", region)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Enable sandboxes in a region. This provisions the sandbox cluster they run on,
// in a project the platform manages.
func (r *SandboxClusterService) Enable(ctx context.Context, region SandboxClusterEnableParamsRegion, body SandboxClusterEnableParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/sandbox_clusters/%v", region)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Get the sandbox cluster in a region
func (r *SandboxClusterService) Get(ctx context.Context, region SandboxClusterGetParamsRegion, opts ...option.RequestOption) (res *SandboxCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/sandbox_clusters/%v", region)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An organization's sandboxes in one region, and the infrastructure they run on.
type SandboxCluster struct {
	// Unique identifier for the Sandbox Cluster.
	ID string `json:"id" api:"required"`
	// Hostname the Sandbox Cluster's sandbox API is served on, over HTTPS. Null until
	// it is assigned.
	APIHostname string `json:"api_hostname" api:"required"`
	// When the Sandbox Cluster was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Hostname sandboxes' endpoints are served on, over HTTPS. Null until it is
	// assigned.
	GatewayHostname string `json:"gateway_hostname" api:"required"`
	// Instance type of the nodes the sandboxes run on.
	InstanceType string `json:"instance_type" api:"required"`
	// Number of nodes the sandboxes run on.
	NodeCount int64 `json:"node_count" api:"required"`
	// ID of the project that holds the Sandbox Cluster's infrastructure.
	ProjectID string `json:"project_id" api:"required"`
	// Public IP address sandbox traffic reaches the Sandbox Cluster on. Null until it
	// is assigned.
	PublicIP string `json:"public_ip" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the Sandbox Cluster was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		APIHostname     respjson.Field
		CreatedAt       respjson.Field
		GatewayHostname respjson.Field
		InstanceType    respjson.Field
		NodeCount       respjson.Field
		ProjectID       respjson.Field
		PublicIP        respjson.Field
		Region          respjson.Field
		Status          respjson.Field
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

type SandboxClusterListParams struct {
	// Pagination cursor returned by a previous request. Only valid for the same sort
	// order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
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

type SandboxClusterDisableParamsRegion string

const (
	SandboxClusterDisableParamsRegionUsSva2 SandboxClusterDisableParamsRegion = "us-sva-2"
	SandboxClusterDisableParamsRegionUsWdc1 SandboxClusterDisableParamsRegion = "us-wdc-1"
)

type SandboxClusterEnableParams struct {
	// Instance type of the nodes the sandboxes run on.
	InstanceType string `json:"instance_type" api:"required"`
	// Number of nodes the sandboxes run on.
	NodeCount int64 `json:"node_count" api:"required"`
	paramObj
}

func (r SandboxClusterEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow SandboxClusterEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxClusterEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxClusterEnableParamsRegion string

const (
	SandboxClusterEnableParamsRegionUsSva2 SandboxClusterEnableParamsRegion = "us-sva-2"
	SandboxClusterEnableParamsRegionUsWdc1 SandboxClusterEnableParamsRegion = "us-wdc-1"
)

type SandboxClusterGetParamsRegion string

const (
	SandboxClusterGetParamsRegionUsSva2 SandboxClusterGetParamsRegion = "us-sva-2"
	SandboxClusterGetParamsRegionUsWdc1 SandboxClusterGetParamsRegion = "us-wdc-1"
)
