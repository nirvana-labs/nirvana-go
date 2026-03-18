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

// ClusterAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterAvailabilityService] method instead.
type ClusterAvailabilityService struct {
	Options []option.RequestOption
}

// NewClusterAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterAvailabilityService(opts ...option.RequestOption) (r ClusterAvailabilityService) {
	r = ClusterAvailabilityService{}
	r.Options = opts
	return
}

// Check if an NKS cluster can be created
func (r *ClusterAvailabilityService) New(ctx context.Context, body ClusterAvailabilityNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/nks/clusters/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Check if an NKS cluster can be updated
func (r *ClusterAvailabilityService) Update(ctx context.Context, clusterID string, body ClusterAvailabilityUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/availability", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

type ClusterAvailabilityNewParams struct {
	// Name of the Cluster.
	Name string `json:"name" api:"required"`
	// Project ID to create the Cluster in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "us-wdc-1", "eu-frk-1",
	// "ap-sin-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// ID of the VPC to use for the Cluster.
	VPCID string `json:"vpc_id" api:"required"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterAvailabilityUpdateParams struct {
	// Name of the Cluster.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Cluster.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterAvailabilityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
