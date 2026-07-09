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

// ClusterPoolCostService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPoolCostService] method instead.
type ClusterPoolCostService struct {
	Options []option.RequestOption
}

// NewClusterPoolCostService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClusterPoolCostService(opts ...option.RequestOption) (r ClusterPoolCostService) {
	r = ClusterPoolCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed NKS node pool.
func (r *ClusterPoolCostService) New(ctx context.Context, clusterID string, body ClusterPoolCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/cost", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed NKS node pool update plus a diff
// against the current state.
func (r *ClusterPoolCostService) Update(ctx context.Context, clusterID string, poolID string, body ClusterPoolCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/pools/%s/cost", clusterID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ClusterPoolCostNewParams struct {
	// Name of the node pool.
	Name string `json:"name" api:"required"`
	// Node configuration.
	NodeConfig NKSNodePoolNodeConfigParam `json:"node_config,omitzero" api:"required"`
	// Number of nodes. Must be between 0 and 100.
	NodeCount param.Opt[int64] `json:"node_count,omitzero"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPoolCostUpdateParams struct {
	// Name of the node pool.
	Name param.Opt[string] `json:"name,omitzero"`
	// Number of nodes.
	NodeCount param.Opt[int64] `json:"node_count,omitzero"`
	// Tags to attach to the node pool.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ClusterPoolCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClusterPoolCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClusterPoolCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
