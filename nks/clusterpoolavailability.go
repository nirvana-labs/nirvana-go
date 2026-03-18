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
)

// ClusterPoolAvailabilityService contains methods and other services that help
// with interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPoolAvailabilityService] method instead.
type ClusterPoolAvailabilityService struct {
	Options []option.RequestOption
}

// NewClusterPoolAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterPoolAvailabilityService(opts ...option.RequestOption) (r ClusterPoolAvailabilityService) {
	r = ClusterPoolAvailabilityService{}
	r.Options = opts
	return
}

// Check if a node pool can be created in an NKS cluster
func (r *ClusterPoolAvailabilityService) New(ctx context.Context, clusterID string, body ClusterPoolAvailabilityNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/availability", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Check if an NKS node pool can be updated
func (r *ClusterPoolAvailabilityService) Update(ctx context.Context, clusterID string, poolID string, body ClusterPoolAvailabilityUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/availability", clusterID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

type ClusterPoolAvailabilityNewParams struct {
	// Name of the node pool.
	Name string `json:"name" api:"required"`
	// Node configuration.
	NodeConfig NKSNodeConfigParam `json:"node_config,omitzero" api:"required"`
	// Number of nodes. Must be between 1 and 100.
	NodeCount int64 `json:"node_count" api:"required"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolAvailabilityUpdateParams struct {
	// Name of the node pool.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolAvailabilityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
