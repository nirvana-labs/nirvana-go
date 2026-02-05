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

// VPCAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVPCAvailabilityService] method instead.
type VPCAvailabilityService struct {
	Options []option.RequestOption
}

// NewVPCAvailabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVPCAvailabilityService(opts ...option.RequestOption) (r VPCAvailabilityService) {
	r = VPCAvailabilityService{}
	r.Options = opts
	return
}

// Check if a VPC can be created
func (r *VPCAvailabilityService) New(ctx context.Context, body VPCAvailabilityNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "v1/networking/vpcs/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check if a VPC can be updated
func (r *VPCAvailabilityService) Update(ctx context.Context, vpcID string, body VPCAvailabilityUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/availability", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type VPCAvailabilityNewParams struct {
	// Name of the VPC.
	Name string `json:"name,required"`
	// Project ID the VPC belongs to.
	ProjectID string `json:"project_id,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "us-wdc-1", "eu-frk-1",
	// "ap-sin-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// Name of the subnet to create.
	SubnetName string `json:"subnet_name,required"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VPCAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCAvailabilityUpdateParams struct {
	// Name of the VPC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Name of the subnet to create.
	SubnetName param.Opt[string] `json:"subnet_name,omitzero"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VPCAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCAvailabilityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
