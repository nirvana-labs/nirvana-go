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
	path := "v1/compute/vms"
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
	path := fmt.Sprintf("v1/compute/vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all VMs
func (r *VMService) List(ctx context.Context, opts ...option.RequestOption) (res *VMList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/compute/vms"
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
	path := fmt.Sprintf("v1/compute/vms/%s", vmID)
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
	path := fmt.Sprintf("v1/compute/vms/%s", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// CPU configuration for the VM.
type CPUConfig struct {
	// Number of virtual CPUs.
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

// CPU configuration for the VM.
type CPUConfigParam struct {
	// Number of virtual CPUs.
	Vcpu param.Field[int64] `json:"vcpu,required"`
}

func (r CPUConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Memory configuration for the VM.
type MemoryConfig struct {
	// Size of the memory in GB.
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

// Memory configuration for the VM.
type MemoryConfigParam struct {
	// Size of the memory in GB.
	Size param.Field[int64] `json:"size,required"`
}

func (r MemoryConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OS image details.
type OSImage struct {
	// When the OS image was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name of the OS image.
	DisplayName string `json:"display_name,required"`
	// Name of the OS image.
	Name string      `json:"name,required"`
	JSON osImageJSON `json:"-"`
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

// Public SSH key configuration for the VM.
type SSHKeyParam struct {
	// Public key to and and use to access the VM.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r SSHKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// VM details.
type VM struct {
	// Unique identifier for the VM.
	ID string `json:"id,required"`
	// ID of the boot volume for the VM.
	BootVolumeID string `json:"boot_volume_id,required"`
	// CPU configuration for the VM.
	CPUConfig CPUConfig `json:"cpu_config,required"`
	// When the VM was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// IDs of the data volumes for the VM.
	DataVolumeIDs []string `json:"data_volume_ids,required"`
	// Memory configuration for the VM.
	MemoryConfig MemoryConfig `json:"memory_config,required"`
	// Name of the VM.
	Name string `json:"name,required"`
	// Private IP of the VM.
	PrivateIP string `json:"private_ip,required,nullable"`
	// Public IP of the VM.
	PublicIP string `json:"public_ip,required,nullable"`
	// Region the resource is in.
	Region shared.RegionName `json:"region,required"`
	// Status of the resource.
	Status shared.ResourceStatus `json:"status,required"`
	// ID of the subnet for the VM.
	SubnetID string `json:"subnet_id,required"`
	// When the VM was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the VPC for the VM.
	VPCID string `json:"vpc_id,required"`
	// Name of the VPC for the VM.
	VPCName string `json:"vpc_name,required"`
	JSON    vmJSON `json:"-"`
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
	// Boot volume for the VM.
	BootVolume param.Field[VMNewParamsBootVolume] `json:"boot_volume,required"`
	// CPU configuration for the VM.
	CPUConfig param.Field[CPUConfigParam] `json:"cpu_config,required"`
	// Memory configuration for the VM.
	MemoryConfig param.Field[MemoryConfigParam] `json:"memory_config,required"`
	// Name of the VM.
	Name param.Field[string] `json:"name,required"`
	// Name of the OS image to use for the VM.
	OSImageName param.Field[string] `json:"os_image_name,required"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled param.Field[bool] `json:"public_ip_enabled,required"`
	// Region the resource is in.
	Region param.Field[shared.RegionName] `json:"region,required"`
	// Public SSH key configuration for the VM.
	SSHKey param.Field[SSHKeyParam] `json:"ssh_key,required"`
	// ID of the subnet to use for the VM.
	SubnetID param.Field[string] `json:"subnet_id,required"`
	// Data volumes for the VM.
	DataVolumes param.Field[[]VMNewParamsDataVolume] `json:"data_volumes"`
}

func (r VMNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boot volume for the VM.
type VMNewParamsBootVolume struct {
	// Size of the volume in GB.
	Size param.Field[int64] `json:"size,required"`
}

func (r VMNewParamsBootVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// VM data volume create request.
type VMNewParamsDataVolume struct {
	// Name of the volume.
	Name param.Field[string] `json:"name,required"`
	// Size of the volume in GB.
	Size param.Field[int64] `json:"size,required"`
}

func (r VMNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMUpdateParams struct {
	// CPU configuration for the VM.
	CPUConfig param.Field[CPUConfigParam] `json:"cpu_config"`
	// Memory configuration for the VM.
	MemoryConfig param.Field[MemoryConfigParam] `json:"memory_config"`
	// Name of the VM.
	Name param.Field[string] `json:"name"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled param.Field[bool] `json:"public_ip_enabled"`
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
