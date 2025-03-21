// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

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

// VMService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMService] method instead.
type VMService struct {
	Options  []option.RequestOption
	Volumes  *VMVolumeService
	OSImages *VMOSImageService
}

// NewVMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMService(opts ...option.RequestOption) (r *VMService) {
	r = &VMService{}
	r.Options = opts
	r.Volumes = NewVMVolumeService(opts...)
	r.OSImages = NewVMOSImageService(opts...)
	return
}

// Create a VM
func (r *VMService) New(ctx context.Context, body VMNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a VM
func (r *VMService) Update(ctx context.Context, vmID string, body VMUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("compute/vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VMs
func (r *VMService) List(ctx context.Context, opts ...option.RequestOption) (res *VMList, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VM
func (r *VMService) Delete(ctx context.Context, vmID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("compute/vms/%s", vmID)
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
	path := fmt.Sprintf("compute/vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// CPU configuration details.
type CPUConfig struct {
	// virtual CPUs
	Vcpu int64         `json:"vcpu,required"`
	JSON cpuConfigJSON `json:"-"`
}

// cpuConfigJSON contains the JSON metadata for the struct [CPUConfig]
type cpuConfigJSON struct {
	Vcpu        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CPUConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cpuConfigJSON) RawJSON() string {
	return r.raw
}

// CPU configuration details.
type CPUConfigParam struct {
	// virtual CPUs
	Vcpu param.Field[int64] `json:"vcpu,required"`
}

func (r CPUConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Memory configuration details.
type MemoryConfig struct {
	// memory size
	Size int64            `json:"size,required"`
	JSON memoryConfigJSON `json:"-"`
}

// memoryConfigJSON contains the JSON metadata for the struct [MemoryConfig]
type memoryConfigJSON struct {
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryConfigJSON) RawJSON() string {
	return r.raw
}

// Memory configuration details.
type MemoryConfigParam struct {
	// memory size
	Size param.Field[int64] `json:"size,required"`
}

func (r MemoryConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OSImage struct {
	CreatedAt   time.Time   `json:"created_at,required" format:"date-time"`
	DisplayName string      `json:"display_name,required"`
	Name        string      `json:"name,required"`
	JSON        osImageJSON `json:"-"`
}

// osImageJSON contains the JSON metadata for the struct [OSImage]
type osImageJSON struct {
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OSImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r osImageJSON) RawJSON() string {
	return r.raw
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
	ID           string `json:"id,required"`
	BootVolumeID string `json:"boot_volume_id,required"`
	// CPU configuration details.
	CPUConfig     CPUConfig `json:"cpu_config,required"`
	CreatedAt     time.Time `json:"created_at,required" format:"date-time"`
	DataVolumeIDs []string  `json:"data_volume_ids,required"`
	// Memory configuration details.
	MemoryConfig MemoryConfig          `json:"memory_config,required"`
	Name         string                `json:"name,required"`
	PrivateIP    string                `json:"private_ip,required,nullable"`
	PublicIP     string                `json:"public_ip,required,nullable"`
	Region       shared.RegionName     `json:"region,required"`
	Status       shared.ResourceStatus `json:"status,required"`
	SubnetID     string                `json:"subnet_id,required"`
	UpdatedAt    time.Time             `json:"updated_at,required" format:"date-time"`
	VPCID        string                `json:"vpc_id,required"`
	VPCName      string                `json:"vpc_name,required"`
	JSON         vmJSON                `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [VM]
type vmJSON struct {
	ID            apijson.Field
	BootVolumeID  apijson.Field
	CPUConfig     apijson.Field
	CreatedAt     apijson.Field
	DataVolumeIDs apijson.Field
	MemoryConfig  apijson.Field
	Name          apijson.Field
	PrivateIP     apijson.Field
	PublicIP      apijson.Field
	Region        apijson.Field
	Status        apijson.Field
	SubnetID      apijson.Field
	UpdatedAt     apijson.Field
	VPCID         apijson.Field
	VPCName       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VM) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmJSON) RawJSON() string {
	return r.raw
}

type VMList struct {
	Items []VM       `json:"items,required"`
	JSON  vmListJSON `json:"-"`
}

// vmListJSON contains the JSON metadata for the struct [VMList]
type vmListJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VMList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vmListJSON) RawJSON() string {
	return r.raw
}

type VMNewParams struct {
	// Boot volume create request.
	BootVolume param.Field[VMNewParamsBootVolume] `json:"boot_volume,required"`
	// CPU configuration details.
	CPUConfig param.Field[CPUConfigParam] `json:"cpu_config,required"`
	// Memory configuration details.
	MemoryConfig    param.Field[MemoryConfigParam] `json:"memory_config,required"`
	Name            param.Field[string]            `json:"name,required"`
	OSImageName     param.Field[string]            `json:"os_image_name,required"`
	PublicIPEnabled param.Field[bool]              `json:"public_ip_enabled,required"`
	Region          param.Field[shared.RegionName] `json:"region,required"`
	// SSH key details.
	SSHKey      param.Field[SSHKeyParam]             `json:"ssh_key,required"`
	SubnetID    param.Field[string]                  `json:"subnet_id,required"`
	DataVolumes param.Field[[]VMNewParamsDataVolume] `json:"data_volumes"`
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

// VM data volume create request.
type VMNewParamsDataVolume struct {
	Name param.Field[string] `json:"name,required"`
	Size param.Field[int64]  `json:"size,required"`
}

func (r VMNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMUpdateParams struct {
	// CPU configuration details.
	CPUConfig param.Field[CPUConfigParam] `json:"cpu_config"`
	// Memory configuration details.
	MemoryConfig    param.Field[MemoryConfigParam] `json:"memory_config"`
	Name            param.Field[string]            `json:"name"`
	PublicIPEnabled param.Field[bool]              `json:"public_ip_enabled"`
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
