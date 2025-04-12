// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/param"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
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
func NewVPCService(opts ...option.RequestOption) (r *VPCService) {
	r = &VPCService{}
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
	// Time the subnet was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the subnet.
	Name string `json:"name,required"`
	// Time the subnet was updated.
	UpdatedAt time.Time  `json:"updated_at,required" format:"date-time"`
	JSON      subnetJSON `json:"-"`
}

// subnetJSON contains the JSON metadata for the struct [Subnet]
type subnetJSON struct {
	ID          apijson.Field
	Cidr        apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Subnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subnetJSON) RawJSON() string {
	return r.raw
}

// VPC details.
type VPC struct {
	// Unique identifier for the VPC.
	ID string `json:"id,required"`
	// Time the VPC was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// IDs of the firewall rules associated with the VPC.
	FirewallRuleIDs []string `json:"firewall_rule_ids,required"`
	// Name of the VPC.
	Name string `json:"name,required"`
	// Region the resource is in.
	Region shared.RegionName `json:"region,required"`
	// Status of the resource.
	Status shared.ResourceStatus `json:"status,required"`
	// Subnet of the VPC.
	Subnet Subnet `json:"subnet,required"`
	// Time the VPC was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      vpcJSON   `json:"-"`
}

// vpcJSON contains the JSON metadata for the struct [VPC]
type vpcJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	FirewallRuleIDs apijson.Field
	Name            apijson.Field
	Region          apijson.Field
	Status          apijson.Field
	Subnet          apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VPC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vpcJSON) RawJSON() string {
	return r.raw
}

type VPCList struct {
	Items []VPC       `json:"items,required"`
	JSON  vpcListJSON `json:"-"`
}

// vpcListJSON contains the JSON metadata for the struct [VPCList]
type vpcListJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VPCList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vpcListJSON) RawJSON() string {
	return r.raw
}

type VPCNewParams struct {
	// Name of the VPC.
	Name param.Field[string] `json:"name,required"`
	// Region the resource is in.
	Region param.Field[shared.RegionName] `json:"region,required"`
	// Name of the subnet to create.
	SubnetName param.Field[string] `json:"subnet_name,required"`
}

func (r VPCNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VPCUpdateParams struct {
	// Name of the VPC.
	Name param.Field[string] `json:"name"`
	// Name of the subnet to create.
	SubnetName param.Field[string] `json:"subnet_name"`
}

func (r VPCUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
