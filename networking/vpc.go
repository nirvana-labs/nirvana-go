// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
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
	Options      []option.RequestOption
	Availability VPCAvailabilityService
}

// NewVPCService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVPCService(opts ...option.RequestOption) (r VPCService) {
	r = VPCService{}
	r.Options = opts
	r.Availability = NewVPCAvailabilityService(opts...)
	return
}

// Create a VPC
func (r *VPCService) New(ctx context.Context, body VPCNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/networking/vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a VPC
func (r *VPCService) Update(ctx context.Context, vpcID string, body VPCUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VPCs
func (r *VPCService) List(ctx context.Context, query VPCListParams, opts ...option.RequestOption) (res *pagination.Cursor[VPC], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/networking/vpcs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all VPCs
func (r *VPCService) ListAutoPaging(ctx context.Context, query VPCListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VPC] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a VPC
func (r *VPCService) Delete(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	// Unique identifier for the Subnet.
	ID string `json:"id,required"`
	// CIDR block for the Subnet.
	CIDR string `json:"cidr,required"`
	// When the Subnet was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the Subnet.
	Name string `json:"name,required"`
	// When the Subnet was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CIDR        respjson.Field
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
	// IDs of the Firewall Rules associated with the VPC.
	FirewallRuleIDs []string `json:"firewall_rule_ids,required"`
	// Name of the VPC.
	Name string `json:"name,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Subnet of the VPC.
	Subnet Subnet `json:"subnet,required"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,required"`
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
		Tags            respjson.Field
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
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
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
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// Name of the subnet to create.
	SubnetName string `json:"subnet_name,required"`
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
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
	// Tags to attach to the VPC.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VPCUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VPCUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VPCUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VPCListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VPCListParams]'s query parameters as `url.Values`.
func (r VPCListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
