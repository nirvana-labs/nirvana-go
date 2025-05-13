// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// VPCService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVPCService] method instead.
type VPCService struct {
	Options []option.RequestOption
}

// NewVPCService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVPCService(opts ...option.RequestOption) (r VPCService) {
	r = VPCService{}
	r.Options = opts
	return
}

// Create a VPC
func (r *VPCService) New(ctx context.Context, body VPCNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/networking/vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a VPC
func (r *VPCService) Update(ctx context.Context, vpcID string, body VPCUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VPCs
func (r *VPCService) List(ctx context.Context, opts ...option.RequestOption) (res *VPCList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/networking/vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VPC
func (r *VPCService) Delete(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a VPC
func (r *VPCService) Get(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *VPC, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Subnet of the VPC.
type Subnet struct {
	// Unique identifier for the subnet.
	ID string `json:"id,required"`
	// CIDR block for the subnet.
	Cidr string `json:"cidr,required"`
	// When the subnet was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the subnet.
	Name string `json:"name,required"`
	// When the subnet was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cidr        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subnet) RawJSON() string { return r.JSON.raw }
func (r *Subnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VPC details.
type VPC struct {
	// Unique identifier for the VPC.
	ID string `json:"id,required"`
	// When the VPC was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// IDs of the firewall rules associated with the VPC.
	FirewallRuleIDs []string `json:"firewall_rule_ids,required"`
	// Name of the VPC.
	Name string `json:"name,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-lon-1", "eu-ams-1",
	// "eu-frk-1", "ap-sin-1", "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Subnet of the VPC.
	Subnet Subnet `json:"subnet,required"`
	// When the VPC was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FirewallRuleIDs respjson.Field
		Name            respjson.Field
		Region          respjson.Field
		Status          respjson.Field
		Subnet          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VPC) RawJSON() string { return r.JSON.raw }
func (r *VPC) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCList struct {
	Items []VPC `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VPCList) RawJSON() string { return r.JSON.raw }
func (r *VPCList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCNewParams struct {
	// Name of the VPC.
	Name string `json:"name,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-lon-1", "eu-ams-1",
	// "eu-frk-1", "ap-sin-1", "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// Name of the subnet to create.
	SubnetName string `json:"subnet_name,required"`
	paramObj
}

func (r VPCNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCUpdateParams struct {
	// Name of the VPC.
	Name param.Opt[string] `json:"name,omitzero"`
	// Name of the subnet to create.
	SubnetName param.Opt[string] `json:"subnet_name,omitzero"`
	paramObj
}

func (r VPCUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
