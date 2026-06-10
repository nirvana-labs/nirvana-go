// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

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

// VPCCostService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVPCCostService] method instead.
type VPCCostService struct {
	Options []option.RequestOption
}

// NewVPCCostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVPCCostService(opts ...option.RequestOption) (r VPCCostService) {
	r = VPCCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed VPC.
func (r *VPCCostService) New(ctx context.Context, body VPCCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/networking/vpcs/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed VPC update plus a diff against the
// current state.
func (r *VPCCostService) Update(ctx context.Context, vpcID string, body VPCCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/cost", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type VPCCostNewParams struct {
	// Name of the VPC.
	Name string `json:"name" api:"required"`
	// Project ID the VPC belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Name of the subnet to create.
	SubnetName string `json:"subnet_name" api:"required"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VPCCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCCostUpdateParams struct {
	// Name of the VPC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Name of the subnet to create.
	SubnetName param.Opt[string] `json:"subnet_name,omitzero"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VPCCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
