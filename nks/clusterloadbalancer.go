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
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ClusterLoadBalancerService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterLoadBalancerService] method instead.
type ClusterLoadBalancerService struct {
	Options []option.RequestOption
}

// NewClusterLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterLoadBalancerService(opts ...option.RequestOption) (r ClusterLoadBalancerService) {
	r = ClusterLoadBalancerService{}
	r.Options = opts
	return
}

// List all load balancers in an NKS cluster
func (r *ClusterLoadBalancerService) List(ctx context.Context, clusterID string, query ClusterLoadBalancerListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSLoadBalancer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/load_balancers", clusterID)
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

// List all load balancers in an NKS cluster
func (r *ClusterLoadBalancerService) ListAutoPaging(ctx context.Context, clusterID string, query ClusterLoadBalancerListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSLoadBalancer] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, query, opts...))
}

// Get details about an NKS load balancer
func (r *ClusterLoadBalancerService) Get(ctx context.Context, clusterID string, loadBalancerID string, opts ...option.RequestOption) (res *NKSLoadBalancer, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/load_balancers/%s", clusterID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS load balancer details.
type NKSLoadBalancer struct {
	// Unique identifier for the load balancer.
	ID string `json:"id" api:"required"`
	// Cluster this load balancer belongs to.
	ClusterID string `json:"cluster_id" api:"required"`
	// When the load balancer was first discovered.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Kubernetes namespace of the load balancer.
	Namespace string `json:"namespace" api:"required"`
	// Private IP address of the load balancer.
	PrivateIP string `json:"private_ip" api:"required"`
	// Public IP address assigned to this load balancer.
	PublicIP string `json:"public_ip" api:"required"`
	// Whether a public IP is enabled for this load balancer.
	PublicIPEnabled bool `json:"public_ip_enabled" api:"required"`
	// Kubernetes service name of the load balancer.
	ServiceName string `json:"service_name" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the load balancer was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClusterID       respjson.Field
		CreatedAt       respjson.Field
		Namespace       respjson.Field
		PrivateIP       respjson.Field
		PublicIP        respjson.Field
		PublicIPEnabled respjson.Field
		ServiceName     respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSLoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *NKSLoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSLoadBalancerList struct {
	Items []NKSLoadBalancer `json:"items" api:"required"`
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
func (r NKSLoadBalancerList) RawJSON() string { return r.JSON.raw }
func (r *NKSLoadBalancerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterLoadBalancerListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterLoadBalancerListParams]'s query parameters as
// `url.Values`.
func (r ClusterLoadBalancerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
