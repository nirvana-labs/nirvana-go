// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nirvana

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/param"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// NetworkingVPCService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkingVPCService] method instead.
type NetworkingVPCService struct {
	Options []option.RequestOption
}

// NewNetworkingVPCService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkingVPCService(opts ...option.RequestOption) (r *NetworkingVPCService) {
	r = &NetworkingVPCService{}
	r.Options = opts
	return
}

// Create a VPC
func (r *NetworkingVPCService) New(ctx context.Context, body NetworkingVPCNewParams, opts ...option.RequestOption) (res *Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "networking/vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all VPCs
func (r *NetworkingVPCService) List(ctx context.Context, opts ...option.RequestOption) (res *VPCList, err error) {
	opts = append(r.Options[:], opts...)
	path := "networking/vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VPC
func (r *NetworkingVPCService) Delete(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a VPC
func (r *NetworkingVPCService) Get(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *VPC, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Subnet details.
type Subnet struct {
	ID        string     `json:"id,required"`
	Cidr      string     `json:"cidr,required"`
	CreatedAt string     `json:"created_at,required"`
	Name      string     `json:"name,required"`
	UpdatedAt string     `json:"updated_at,required"`
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
	ID            string                `json:"id,required"`
	CreatedAt     string                `json:"created_at,required"`
	FirewallRules []FirewallRule        `json:"firewall_rules,required"`
	Name          string                `json:"name,required"`
	Region        shared.RegionName     `json:"region,required"`
	Status        shared.ResourceStatus `json:"status,required"`
	// Subnet details.
	Subnet    Subnet  `json:"subnet,required"`
	UpdatedAt string  `json:"updated_at,required"`
	JSON      vpcJSON `json:"-"`
}

// vpcJSON contains the JSON metadata for the struct [VPC]
type vpcJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	FirewallRules apijson.Field
	Name          apijson.Field
	Region        apijson.Field
	Status        apijson.Field
	Subnet        apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

type NetworkingVPCNewParams struct {
	Name       param.Field[string]            `json:"name,required"`
	Region     param.Field[shared.RegionName] `json:"region,required"`
	SubnetName param.Field[string]            `json:"subnet_name,required"`
}

func (r NetworkingVPCNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
