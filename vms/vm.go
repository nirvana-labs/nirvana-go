// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/param"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
	"github.com/nirvana-labs/nirvana-go/vpcs"
)

// VMService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMService] method instead.
type VMService struct {
	Options    []option.RequestOption
	Operations *OperationService
}

// NewVMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMService(opts ...option.RequestOption) (r *VMService) {
	r = &VMService{}
	r.Options = opts
	r.Operations = NewOperationService(opts...)
	return
}

// Create a VM
func (r *VMService) New(ctx context.Context, body VMNewParams, opts ...option.RequestOption) (res *shared.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a VM
func (r *VMService) Update(ctx context.Context, vmID string, body VMUpdateParams, opts ...option.RequestOption) (res *shared.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VMs
func (r *VMService) List(ctx context.Context, query VMListParams, opts ...option.RequestOption) (res *VMListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a VM
func (r *VMService) Delete(ctx context.Context, vmID string, opts ...option.RequestOption) (res *shared.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a VM
func (r *VMService) Get(ctx context.Context, vmID string, opts ...option.RequestOption) (res *VM, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// CPU details.
type CPU struct {
	Cores int64   `json:"cores,required"`
	JSON  cpuJSON `json:"-"`
}

// cpuJSON contains the JSON metadata for the struct [CPU]
type cpuJSON struct {
	Cores       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CPU) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cpuJSON) RawJSON() string {
	return r.raw
}

// CPU details.
type CPUParam struct {
	Cores param.Field[int64] `json:"cores,required"`
}

func (r CPUParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RAM details.
type Ram struct {
	// RAM size
	Size int64 `json:"size,required"`
	// Unit (GB, MB, etc.)
	Unit RamUnit `json:"unit,required"`
	JSON ramJSON `json:"-"`
}

// ramJSON contains the JSON metadata for the struct [Ram]
type ramJSON struct {
	Size        apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ram) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ramJSON) RawJSON() string {
	return r.raw
}

// RAM details.
type RamParam struct {
	// RAM size
	Size param.Field[int64] `json:"size,required"`
	// Unit (GB, MB, etc.)
	Unit param.Field[RamUnit] `json:"unit,required"`
}

func (r RamParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unit (GB, MB, etc.)
type RamUnit string

const (
	RamUnitGB RamUnit = "GB"
)

func (r RamUnit) IsKnown() bool {
	switch r {
	case RamUnitGB:
		return true
	}
	return false
}

// SSH key details.
type SSHKeyParam struct {
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r SSHKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage details.
type Storage struct {
	// Storage size
	Size int64 `json:"size,required"`
	// Storage type.
	Type StorageType `json:"type,required"`
	// Storage unit.
	Unit StorageUnit `json:"unit,required"`
	// Disk name, used later
	DiskName string      `json:"disk_name"`
	JSON     storageJSON `json:"-"`
}

// storageJSON contains the JSON metadata for the struct [Storage]
type storageJSON struct {
	Size        apijson.Field
	Type        apijson.Field
	Unit        apijson.Field
	DiskName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Storage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storageJSON) RawJSON() string {
	return r.raw
}

// Storage details.
type StorageParam struct {
	// Storage size
	Size param.Field[int64] `json:"size,required"`
	// Storage type.
	Type param.Field[StorageType] `json:"type,required"`
	// Storage unit.
	Unit param.Field[StorageUnit] `json:"unit,required"`
	// Disk name, used later
	DiskName param.Field[string] `json:"disk_name"`
}

func (r StorageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage type.
type StorageType string

const (
	StorageTypeNvme StorageType = "nvme"
)

func (r StorageType) IsKnown() bool {
	switch r {
	case StorageTypeNvme:
		return true
	}
	return false
}

// Storage unit.
type StorageUnit string

const (
	StorageUnitGB StorageUnit = "GB"
)

func (r StorageUnit) IsKnown() bool {
	switch r {
	case StorageUnitGB:
		return true
	}
	return false
}

// VM details.
type VM struct {
	ID string `json:"id,required"`
	// CPU details.
	CPUConfig CPU    `json:"cpu_config,required"`
	CreatedAt string `json:"created_at,required"`
	// RAM details.
	MemConfig     Ram                   `json:"mem_config,required"`
	Name          string                `json:"name,required"`
	PublicIP      string                `json:"public_ip,required"`
	Region        shared.RegionName     `json:"region,required"`
	Status        shared.ResourceStatus `json:"status,required"`
	StorageConfig []Storage             `json:"storage_config,required"`
	UpdatedAt     string                `json:"updated_at,required"`
	// VPC details.
	VPC  vpcs.VPC `json:"vpc,required"`
	JSON vmJSON   `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [VM]
type vmJSON struct {
	ID            apijson.Field
	CPUConfig     apijson.Field
	CreatedAt     apijson.Field
	MemConfig     apijson.Field
	Name          apijson.Field
	PublicIP      apijson.Field
	Region        apijson.Field
	Status        apijson.Field
	StorageConfig apijson.Field
	UpdatedAt     apijson.Field
	VPC           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VM) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmJSON) RawJSON() string {
	return r.raw
}

type VMListResponse struct {
	Items []VM               `json:"items,required"`
	JSON  vmListResponseJSON `json:"-"`
}

// vmListResponseJSON contains the JSON metadata for the struct [VMListResponse]
type vmListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VMListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmListResponseJSON) RawJSON() string {
	return r.raw
}

type VMNewParams struct {
	// CPU details.
	CPU          param.Field[CPUParam] `json:"cpu,required"`
	Name         param.Field[string]   `json:"name,required"`
	NeedPublicIP param.Field[bool]     `json:"need_public_ip,required"`
	OsImageID    param.Field[int64]    `json:"os_image_id,required"`
	Ports        param.Field[[]string] `json:"ports,required"`
	// RAM details.
	Ram           param.Field[RamParam]          `json:"ram,required"`
	Region        param.Field[shared.RegionName] `json:"region,required"`
	SourceAddress param.Field[string]            `json:"source_address,required"`
	// SSH key details.
	SSHKey   param.Field[SSHKeyParam]    `json:"ssh_key,required"`
	Storage  param.Field[[]StorageParam] `json:"storage,required"`
	SubnetID param.Field[string]         `json:"subnet_id"`
}

func (r VMNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMUpdateParams struct {
	// CPU details.
	CPU param.Field[CPUParam] `json:"cpu"`
	// RAM details.
	Ram     param.Field[RamParam]       `json:"ram"`
	Storage param.Field[[]StorageParam] `json:"storage"`
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMListParams struct {
	// Region
	Region param.Field[string] `query:"region,required"`
}

// URLQuery serializes [VMListParams]'s query parameters as `url.Values`.
func (r VMListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
