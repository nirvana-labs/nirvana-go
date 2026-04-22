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

	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ClusterPoolService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPoolService] method instead.
type ClusterPoolService struct {
	Options      []option.RequestOption
	Availability ClusterPoolAvailabilityService
	Nodes        ClusterPoolNodeService
}

// NewClusterPoolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClusterPoolService(opts ...option.RequestOption) (r ClusterPoolService) {
	r = ClusterPoolService{}
	r.Options = opts
	r.Availability = NewClusterPoolAvailabilityService(opts...)
	r.Nodes = NewClusterPoolNodeService(opts...)
	return
}

// Create a node pool in an NKS cluster
func (r *ClusterPoolService) New(ctx context.Context, clusterID string, body ClusterPoolNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an NKS node pool
func (r *ClusterPoolService) Update(ctx context.Context, clusterID string, poolID string, body ClusterPoolUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s", clusterID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all node pools in an NKS cluster
func (r *ClusterPoolService) List(ctx context.Context, clusterID string, query ClusterPoolListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSNodePool], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools", clusterID)
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

// List all node pools in an NKS cluster
func (r *ClusterPoolService) ListAutoPaging(ctx context.Context, clusterID string, query ClusterPoolListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSNodePool] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, query, opts...))
}

// Delete an NKS node pool
func (r *ClusterPoolService) Delete(ctx context.Context, clusterID string, poolID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s", clusterID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get details about an NKS node pool
func (r *ClusterPoolService) Get(ctx context.Context, clusterID string, poolID string, opts ...option.RequestOption) (res *NKSNodePool, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s", clusterID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS node pool details.
type NKSNodePool struct {
	// Unique identifier for the node pool.
	ID string `json:"id" api:"required"`
	// ID of the Cluster this node pool belongs to.
	ClusterID string `json:"cluster_id" api:"required"`
	// When the node pool was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the node pool.
	Name string `json:"name" api:"required"`
	// Node configuration.
	NodeConfig NKSNodePoolNodeConfigResponse `json:"node_config" api:"required"`
	// Number of nodes.
	NodeCount int64 `json:"node_count" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Tags attached to the node pool.
	Tags []string `json:"tags" api:"required"`
	// When the node pool was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClusterID   respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		NodeConfig  respjson.Field
		NodeCount   respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSNodePool) RawJSON() string { return r.JSON.raw }
func (r *NKSNodePool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume configuration.
//
// The properties Size, Type are required.
type NKSNodePoolBootVolumeParam struct {
	// Size of the boot volume in GB.
	Size int64 `json:"size" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type compute.VolumeType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r NKSNodePoolBootVolumeParam) MarshalJSON() (data []byte, err error) {
	type shadow NKSNodePoolBootVolumeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NKSNodePoolBootVolumeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume configuration.
type NKSNodePoolBootVolumeResponse struct {
	// Size of the boot volume in GB.
	Size int64 `json:"size" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type compute.VolumeType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Size        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSNodePoolBootVolumeResponse) RawJSON() string { return r.JSON.raw }
func (r *NKSNodePoolBootVolumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSNodePoolList struct {
	Items []NKSNodePool `json:"items" api:"required"`
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
func (r NKSNodePoolList) RawJSON() string { return r.JSON.raw }
func (r *NKSNodePoolList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Node configuration.
//
// The properties BootVolume, InstanceType are required.
type NKSNodePoolNodeConfigParam struct {
	// Boot volume configuration.
	BootVolume NKSNodePoolBootVolumeParam `json:"boot_volume,omitzero" api:"required"`
	// Instance type name used for worker nodes.
	InstanceType string `json:"instance_type" api:"required"`
	// Kubernetes labels to apply to each node in the pool. Each entry is "key=value".
	// Keys under kubernetes.io, k8s.io, and nirvanalabs.io prefixes are reserved.
	Labels []string `json:"labels,omitzero"`
	paramObj
}

func (r NKSNodePoolNodeConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow NKSNodePoolNodeConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NKSNodePoolNodeConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Node configuration.
type NKSNodePoolNodeConfigResponse struct {
	// Boot volume configuration.
	BootVolume NKSNodePoolBootVolumeResponse `json:"boot_volume" api:"required"`
	// Instance type name.
	InstanceType string `json:"instance_type" api:"required"`
	// Kubernetes labels applied to each node in the pool. Each entry is "key=value".
	Labels []string `json:"labels" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BootVolume   respjson.Field
		InstanceType respjson.Field
		Labels       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSNodePoolNodeConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *NKSNodePoolNodeConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolNewParams struct {
	// Name of the node pool.
	Name string `json:"name" api:"required"`
	// Node configuration.
	NodeConfig NKSNodePoolNodeConfigParam `json:"node_config,omitzero" api:"required"`
	// Number of nodes. Must be between 1 and 100.
	NodeCount int64 `json:"node_count" api:"required"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolUpdateParams struct {
	// Name of the node pool.
	Name param.Opt[string] `json:"name,omitzero"`
	// Number of nodes.
	NodeCount param.Opt[int64] `json:"node_count,omitzero"`
	// Partial node configuration update.
	NodeConfig ClusterPoolUpdateParamsNodeConfig `json:"node_config,omitzero"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Partial node configuration update.
type ClusterPoolUpdateParamsNodeConfig struct {
	// Kubernetes labels to apply to each node in the pool. Each entry is "key=value".
	// When provided, the list fully replaces the current labels on the pool and on
	// live nodes.
	Labels []string `json:"labels,omitzero"`
	paramObj
}

func (r ClusterPoolUpdateParamsNodeConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolUpdateParamsNodeConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolUpdateParamsNodeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterPoolListParams]'s query parameters as `url.Values`.
func (r ClusterPoolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
