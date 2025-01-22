// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vms

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
	"github.com/nirvana-labs/nirvana-go/volumes"
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
func (r *VMService) List(ctx context.Context, opts ...option.RequestOption) (res *VMListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	Size int64   `json:"size,required"`
	JSON ramJSON `json:"-"`
}

// ramJSON contains the JSON metadata for the struct [Ram]
type ramJSON struct {
	Size        apijson.Field
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
}

func (r RamParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSH key details.
type SSHKeyParam struct {
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r SSHKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// VM details.
type VM struct {
	ID string `json:"id,required"`
	// Volume details.
	BootVolume volumes.Volume `json:"boot_volume,required"`
	// CPU details.
	CPUConfig   CPU              `json:"cpu_config,required"`
	CreatedAt   string           `json:"created_at,required"`
	DataVolumes []volumes.Volume `json:"data_volumes,required"`
	// RAM details.
	MemConfig Ram                   `json:"mem_config,required"`
	Name      string                `json:"name,required"`
	PublicIP  string                `json:"public_ip,required"`
	Region    shared.RegionName     `json:"region,required"`
	Status    shared.ResourceStatus `json:"status,required"`
	UpdatedAt string                `json:"updated_at,required"`
	// VPC details.
	VPC  vpcs.VPC `json:"vpc,required"`
	JSON vmJSON   `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [VM]
type vmJSON struct {
	ID          apijson.Field
	BootVolume  apijson.Field
	CPUConfig   apijson.Field
	CreatedAt   apijson.Field
	DataVolumes apijson.Field
	MemConfig   apijson.Field
	Name        apijson.Field
	PublicIP    apijson.Field
	Region      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	VPC         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// Boot volume create request.
	BootVolume param.Field[VMNewParamsBootVolume] `json:"boot_volume,required"`
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
	SSHKey      param.Field[SSHKeyParam]             `json:"ssh_key,required"`
	DataVolumes param.Field[[]VMNewParamsDataVolume] `json:"data_volumes"`
	SubnetID    param.Field[string]                  `json:"subnet_id"`
}

func (r VMNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boot volume create request.
type VMNewParamsBootVolume struct {
	Size param.Field[int64] `json:"size,required"`
}

func (r VMNewParamsBootVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data volume create request.
type VMNewParamsDataVolume struct {
	Size param.Field[int64] `json:"size,required"`
	// Storage type.
	Type param.Field[volumes.StorageType] `json:"type"`
}

func (r VMNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMUpdateParams struct {
	// Boot volume create request.
	BootVolume param.Field[VMUpdateParamsBootVolume] `json:"boot_volume"`
	// CPU details.
	CPU         param.Field[CPUParam]                   `json:"cpu"`
	DataVolumes param.Field[[]VMUpdateParamsDataVolume] `json:"data_volumes"`
	// RAM details.
	Ram param.Field[RamParam] `json:"ram"`
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boot volume create request.
type VMUpdateParamsBootVolume struct {
	Size param.Field[int64] `json:"size,required"`
}

func (r VMUpdateParamsBootVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data volume create request.
type VMUpdateParamsDataVolume struct {
	Size param.Field[int64] `json:"size,required"`
	// Storage type.
	Type param.Field[volumes.StorageType] `json:"type"`
}

func (r VMUpdateParamsDataVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
