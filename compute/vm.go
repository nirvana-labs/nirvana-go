// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"encoding/json"
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

// VMService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMService] method instead.
type VMService struct {
	Options      []option.RequestOption
	Availability VMAvailabilityService
	Volumes      VMVolumeService
	OSImages     VMOSImageService
}

// NewVMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMService(opts ...option.RequestOption) (r VMService) {
	r = VMService{}
	r.Options = opts
	r.Availability = NewVMAvailabilityService(opts...)
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

// Restart a VM
func (r *VMService) Restart(ctx context.Context, vmID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/vms/%s/restart", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// CPU configuration for the VM.
type CPUConfig struct {
	// Number of virtual CPUs.
	Vcpu int64 `json:"vcpu,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Vcpu        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CPUConfig) RawJSON() string { return r.JSON.raw }
func (r *CPUConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CPUConfig to a CPUConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CPUConfigParam.Overrides()
func (r CPUConfig) ToParam() CPUConfigParam {
	return param.Override[CPUConfigParam](json.RawMessage(r.RawJSON()))
}

// CPU configuration for the VM.
//
// The property Vcpu is required.
type CPUConfigParam struct {
	// Number of virtual CPUs.
	Vcpu int64 `json:"vcpu,required"`
	paramObj
}

func (r CPUConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CPUConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CPUConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Memory configuration for the VM.
type MemoryConfig struct {
	// Size of the memory in GB.
	Size int64 `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryConfig) RawJSON() string { return r.JSON.raw }
func (r *MemoryConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MemoryConfig to a MemoryConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MemoryConfigParam.Overrides()
func (r MemoryConfig) ToParam() MemoryConfigParam {
	return param.Override[MemoryConfigParam](json.RawMessage(r.RawJSON()))
}

// Memory configuration for the VM.
//
// The property Size is required.
type MemoryConfigParam struct {
	// Size of the memory in GB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r MemoryConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow MemoryConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OS Image details.
type OSImage struct {
	// When the OS Image was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name of the OS Image.
	DisplayName string `json:"display_name,required"`
	// Name of the OS Image.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OSImage) RawJSON() string { return r.JSON.raw }
func (r *OSImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public SSH key configuration for the VM.
//
// The property PublicKey is required.
type SSHKeyParam struct {
	// Public key to and and use to access the VM.
	PublicKey string `json:"public_key,required"`
	paramObj
}

func (r SSHKeyParam) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSHKeyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VM details.
type VM struct {
	// Unique identifier for the VM.
	ID string `json:"id,required"`
	// ID of the boot volume attached to the VM.
	BootVolumeID string `json:"boot_volume_id,required"`
	// CPU configuration for the VM.
	CPUConfig CPUConfig `json:"cpu_config,required"`
	// When the VM was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// IDs of the data volumes attached to the VM.
	DataVolumeIDs []string `json:"data_volume_ids,required"`
	// Memory configuration for the VM.
	MemoryConfig MemoryConfig `json:"memory_config,required"`
	// Name of the VM.
	Name string `json:"name,required"`
	// Private IP of the VM.
	PrivateIP string `json:"private_ip,required"`
	// Public IP of the VM.
	PublicIP string `json:"public_ip,required"`
	// Whether the public IP is enabled for the VM.
	PublicIPEnabled bool `json:"public_ip_enabled,required"`
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
	// ID of the subnet the VM is in.
	SubnetID string `json:"subnet_id,required"`
	// When the VM was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the VPC the VM is in.
	VPCID string `json:"vpc_id,required"`
	// Name of the VPC the VM is in.
	VPCName string `json:"vpc_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BootVolumeID    respjson.Field
		CPUConfig       respjson.Field
		CreatedAt       respjson.Field
		DataVolumeIDs   respjson.Field
		MemoryConfig    respjson.Field
		Name            respjson.Field
		PrivateIP       respjson.Field
		PublicIP        respjson.Field
		PublicIPEnabled respjson.Field
		Region          respjson.Field
		Status          respjson.Field
		SubnetID        respjson.Field
		UpdatedAt       respjson.Field
		VPCID           respjson.Field
		VPCName         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VM) RawJSON() string { return r.JSON.raw }
func (r *VM) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMList struct {
	Items []VM `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMList) RawJSON() string { return r.JSON.raw }
func (r *VMList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMNewParams struct {
	// Boot volume for the VM.
	BootVolume VMNewParamsBootVolume `json:"boot_volume,omitzero,required"`
	// CPU configuration for the VM.
	CPUConfig CPUConfigParam `json:"cpu_config,omitzero,required"`
	// Memory configuration for the VM.
	MemoryConfig MemoryConfigParam `json:"memory_config,omitzero,required"`
	// Name of the VM.
	Name string `json:"name,required"`
	// Name of the OS Image to use for the VM.
	OSImageName string `json:"os_image_name,required"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled bool `json:"public_ip_enabled,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// Public SSH key configuration for the VM.
	SSHKey SSHKeyParam `json:"ssh_key,omitzero,required"`
	// ID of the subnet to use for the VM.
	SubnetID string `json:"subnet_id,required"`
	// Data volumes for the VM.
	DataVolumes []VMNewParamsDataVolume `json:"data_volumes,omitzero"`
	paramObj
}

func (r VMNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VMNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume for the VM.
//
// The property Size is required.
type VMNewParamsBootVolume struct {
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r VMNewParamsBootVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMNewParamsBootVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMNewParamsBootVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VM data volume create request.
//
// The properties Name, Size are required.
type VMNewParamsDataVolume struct {
	// Name of the volume.
	Name string `json:"name,required"`
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r VMNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMNewParamsDataVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMNewParamsDataVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMUpdateParams struct {
	// Name of the VM.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled param.Opt[bool] `json:"public_ip_enabled,omitzero"`
	// CPU configuration for the VM.
	CPUConfig CPUConfigParam `json:"cpu_config,omitzero"`
	// Memory configuration for the VM.
	MemoryConfig MemoryConfigParam `json:"memory_config,omitzero"`
	paramObj
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VMUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
