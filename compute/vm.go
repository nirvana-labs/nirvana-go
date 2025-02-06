// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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
	OSImages *VMOSImageService
}

// NewVMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMService(opts ...option.RequestOption) (r *VMService) {
	r = &VMService{}
	r.Options = opts
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

type OSImage struct {
	CreatedAt   string      `json:"created_at,required"`
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
	ID           string `json:"id,required"`
	BootVolumeID string `json:"boot_volume_id,required"`
	// CPU details.
	CPUConfig     CPU      `json:"cpu_config,required"`
	CreatedAt     string   `json:"created_at,required"`
	DataVolumeIDs []string `json:"data_volume_ids,required"`
	// RAM details.
	MemConfig Ram                   `json:"mem_config,required"`
	Name      string                `json:"name,required"`
	PrivateIP string                `json:"private_ip,required"`
	PublicIP  string                `json:"public_ip,required"`
	Region    shared.RegionName     `json:"region,required"`
	Status    shared.ResourceStatus `json:"status,required"`
	UpdatedAt string                `json:"updated_at,required"`
	VPCID     string                `json:"vpc_id,required"`
	JSON      vmJSON                `json:"-"`
}

// vmJSON contains the JSON metadata for the struct [VM]
type vmJSON struct {
	ID            apijson.Field
	BootVolumeID  apijson.Field
	CPUConfig     apijson.Field
	CreatedAt     apijson.Field
	DataVolumeIDs apijson.Field
	MemConfig     apijson.Field
	Name          apijson.Field
	PrivateIP     apijson.Field
	PublicIP      apijson.Field
	Region        apijson.Field
	Status        apijson.Field
	UpdatedAt     apijson.Field
	VPCID         apijson.Field
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
	// CPU details.
	CPU             param.Field[CPUParam] `json:"cpu,required"`
	Name            param.Field[string]   `json:"name,required"`
	OSImageName     param.Field[string]   `json:"os_image_name,required"`
	PublicIPEnabled param.Field[bool]     `json:"public_ip_enabled,required"`
	// RAM details.
	Ram    param.Field[RamParam]          `json:"ram,required"`
	Region param.Field[shared.RegionName] `json:"region,required"`
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

// VM data volume create request.
type VMNewParamsDataVolume struct {
	Size param.Field[int64] `json:"size,required"`
}

func (r VMNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VMUpdateParams struct {
	// CPU details.
	CPU param.Field[CPUParam] `json:"cpu"`
	// RAM details.
	Ram param.Field[RamParam] `json:"ram"`
}

func (r VMUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
