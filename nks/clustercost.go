// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ClusterCostService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterCostService] method instead.
type ClusterCostService struct {
	Options []option.RequestOption
}

// NewClusterCostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClusterCostService(opts ...option.RequestOption) (r ClusterCostService) {
	r = ClusterCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed NKS cluster.
func (r *ClusterCostService) New(ctx context.Context, body ClusterCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/nks/clusters/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed NKS cluster update plus a diff
// against the current state.
func (r *ClusterCostService) Update(ctx context.Context, clusterID string, body ClusterCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/cost", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ClusterCostNewParams struct {
	// Whether to enable autoscaling for the Cluster.
	Autoscaling bool `json:"autoscaling" api:"required"`
	// Kubernetes version for the Cluster.
	KubernetesVersion string `json:"kubernetes_version" api:"required"`
	// Name of the Cluster.
	Name string `json:"name" api:"required"`
	// Project ID to create the Cluster in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// ID of the VPC to use for the Cluster.
	VPCID string `json:"vpc_id" api:"required"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterCostUpdateParams struct {
	// Whether to enable autoscaling for the Cluster.
	Autoscaling param.Opt[bool] `json:"autoscaling,omitzero"`
	// Name of the Cluster.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
