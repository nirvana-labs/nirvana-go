// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// VMAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMAvailabilityService] method instead.
type VMAvailabilityService struct {
	Options []option.RequestOption
}

// NewVMAvailabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVMAvailabilityService(opts ...option.RequestOption) (r VMAvailabilityService) {
	r = VMAvailabilityService{}
	r.Options = opts
	return
}

// Check VM Create Availability
func (r *VMAvailabilityService) New(ctx context.Context, body VMAvailabilityNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "v1/compute/vms/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check VM Update Availability
func (r *VMAvailabilityService) Update(ctx context.Context, vmID string, body VMAvailabilityUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/vms/%s/availability", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type VMAvailabilityNewParams struct {
	// Boot volume for the VM.
	BootVolume VMAvailabilityNewParamsBootVolume `json:"boot_volume,omitzero,required"`
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
	DataVolumes []VMAvailabilityNewParamsDataVolume `json:"data_volumes,omitzero"`
	paramObj
}

func (r VMAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VMAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume for the VM.
//
// The property Size is required.
type VMAvailabilityNewParamsBootVolume struct {
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r VMAvailabilityNewParamsBootVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMAvailabilityNewParamsBootVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMAvailabilityNewParamsBootVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VM data volume create request.
//
// The properties Name, Size are required.
type VMAvailabilityNewParamsDataVolume struct {
	// Name of the volume.
	Name string `json:"name,required"`
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r VMAvailabilityNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMAvailabilityNewParamsDataVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMAvailabilityNewParamsDataVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMAvailabilityUpdateParams struct {
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

func (r VMAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VMAvailabilityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
