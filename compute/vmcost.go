// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// VMCostService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMCostService] method instead.
type VMCostService struct {
	Options []option.RequestOption
}

// NewVMCostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMCostService(opts ...option.RequestOption) (r VMCostService) {
	r = VMCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed VM.
func (r *VMCostService) New(ctx context.Context, body VMCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/compute/vms/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed VM update plus a diff against the
// current state.
func (r *VMCostService) Update(ctx context.Context, vmID string, body VMCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/compute/vms/%s/cost", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type VMCostNewParams struct {
	// Boot volume for the VM.
	BootVolume VMCostNewParamsBootVolume `json:"boot_volume,omitzero" api:"required"`
	// Instance type name.
	InstanceType string `json:"instance_type" api:"required"`
	// Name of the VM.
	Name string `json:"name" api:"required"`
	// Name of the OS Image to use for the VM.
	OSImageName string `json:"os_image_name" api:"required"`
	// Project ID to create the VM in.
	ProjectID string `json:"project_id" api:"required"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled bool `json:"public_ip_enabled" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Public SSH key configuration for the VM.
	SSHKey SSHKeyRequestParam `json:"ssh_key,omitzero" api:"required"`
	// ID of the subnet to use for the VM.
	SubnetID string `json:"subnet_id" api:"required"`
	// Data volumes for the VM.
	DataVolumes []VMCostNewParamsDataVolume `json:"data_volumes,omitzero"`
	// Tags to attach to the VM.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VMCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VMCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume for the VM.
//
// The properties Size, Type are required.
type VMCostNewParamsBootVolume struct {
	// Size of the Volume in GB.
	Size int64 `json:"size" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type VolumeType `json:"type,omitzero" api:"required"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VMCostNewParamsBootVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMCostNewParamsBootVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMCostNewParamsBootVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VM data volume create request.
//
// The properties Name, Size, Type are required.
type VMCostNewParamsDataVolume struct {
	// Name of the Volume.
	Name string `json:"name" api:"required"`
	// Size of the Volume in GB.
	Size int64 `json:"size" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type VolumeType `json:"type,omitzero" api:"required"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VMCostNewParamsDataVolume) MarshalJSON() (data []byte, err error) {
	type shadow VMCostNewParamsDataVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMCostNewParamsDataVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMCostUpdateParams struct {
	// Instance type name.
	InstanceType param.Opt[string] `json:"instance_type,omitzero"`
	// Name of the VM.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether to enable public IP for the VM.
	PublicIPEnabled param.Opt[bool] `json:"public_ip_enabled,omitzero"`
	// Tags to attach to the VM.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VMCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VMCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
