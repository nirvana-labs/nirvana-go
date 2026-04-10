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

// ClusterService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterService] method instead.
type ClusterService struct {
	Options                []option.RequestOption
	Availability           ClusterAvailabilityService
	PersistentVolumeClaims ClusterPersistentVolumeClaimService
	Kubeconfig             ClusterKubeconfigService
	Controllers            ClusterControllerService
	LoadBalancers          ClusterLoadBalancerService
	Pools                  ClusterPoolService
}

// NewClusterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClusterService(opts ...option.RequestOption) (r ClusterService) {
	r = ClusterService{}
	r.Options = opts
	r.Availability = NewClusterAvailabilityService(opts...)
	r.PersistentVolumeClaims = NewClusterPersistentVolumeClaimService(opts...)
	r.Kubeconfig = NewClusterKubeconfigService(opts...)
	r.Controllers = NewClusterControllerService(opts...)
	r.LoadBalancers = NewClusterLoadBalancerService(opts...)
	r.Pools = NewClusterPoolService(opts...)
	return
}

// Create an NKS Cluster
func (r *ClusterService) New(ctx context.Context, body ClusterNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/nks/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an NKS cluster
func (r *ClusterService) Update(ctx context.Context, clusterID string, body ClusterUpdateParams, opts ...option.RequestOption) (res *NKSCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all NKS clusters
func (r *ClusterService) List(ctx context.Context, query ClusterListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSCluster], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/nks/clusters"
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

// List all NKS clusters
func (r *ClusterService) ListAutoPaging(ctx context.Context, query ClusterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSCluster] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an NKS cluster
func (r *ClusterService) Delete(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get details about an NKS cluster
func (r *ClusterService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *NKSCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS Cluster details.
type NKSCluster struct {
	// Unique identifier for the Cluster.
	ID string `json:"id" api:"required"`
	// When the Cluster was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the Cluster.
	Name string `json:"name" api:"required"`
	// IDs of pools belonging to this Cluster.
	PoolIDs []string `json:"pool_ids" api:"required"`
	// Private IP (VIP) of the Cluster.
	PrivateIP string `json:"private_ip" api:"required"`
	// Project ID the Cluster belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Public IP of the Cluster.
	PublicIP string `json:"public_ip" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "ap-sin-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Tags attached to the Cluster.
	Tags []string `json:"tags" api:"required"`
	// When the Cluster was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the VPC the Cluster is in.
	VPCID string `json:"vpc_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		PoolIDs     respjson.Field
		PrivateIP   respjson.Field
		ProjectID   respjson.Field
		PublicIP    respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		VPCID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSCluster) RawJSON() string { return r.JSON.raw }
func (r *NKSCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSClusterList struct {
	Items []NKSCluster `json:"items" api:"required"`
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
func (r NKSClusterList) RawJSON() string { return r.JSON.raw }
func (r *NKSClusterList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterNewParams struct {
	// Name of the Cluster.
	Name string `json:"name" api:"required"`
	// Project ID to create the Cluster in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "ap-sin-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// ID of the VPC to use for the Cluster.
	VPCID string `json:"vpc_id" api:"required"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterUpdateParams struct {
	// Name of the Cluster.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id" api:"required" json:"-"`
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterListParams]'s query parameters as `url.Values`.
func (r ClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
