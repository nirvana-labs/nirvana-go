// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sandbox_clusters

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/v3/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v3/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/shared"
)

// CostService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCostService] method instead.
type CostService struct {
	Options []option.RequestOption
}

// NewCostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCostService(opts ...option.RequestOption) (r CostService) {
	r = CostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed Sandbox Cluster.
func (r *CostService) New(ctx context.Context, body CostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sandbox_clusters/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CostNewParams struct {
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

func (r CostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
