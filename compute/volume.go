// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
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

// VolumeService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	return
}

// Create a Volume. Only data volumes can be created.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/compute/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Volume. Boot or data volumes can be updated.
func (r *VolumeService) Update(ctx context.Context, volumeID string, body VolumeUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all volumes
func (r *VolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *VolumeList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/compute/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Volume. Boot or data volumes can be deleted.
func (r *VolumeService) Delete(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a Volume.
func (r *VolumeService) Get(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Storage type the volume is using.
type StorageType string

const (
	StorageTypeNvme StorageType = "nvme"
)

// Volume details.
type Volume struct {
	// Unique identifier for the volume.
	ID string `json:"id,required"`
	// When the volume was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Volume kind.
	//
	// Any of "boot", "data".
	Kind VolumeKind `json:"kind,required"`
	// Name of the volume.
	Name string `json:"name,required"`
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Storage type the volume is using.
	//
	// Any of "nvme".
	Type StorageType `json:"type,required"`
	// When the volume was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the VM the volume is attached to.
	VMID string `json:"vm_id,required"`
	// Name of the VM the volume is attached to.
	VMName string `json:"vm_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		VMID        respjson.Field
		VMName      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume kind.
type VolumeKind string

const (
	VolumeKindBoot VolumeKind = "boot"
	VolumeKindData VolumeKind = "data"
)

type VolumeList struct {
	Items []Volume `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeList) RawJSON() string { return r.JSON.raw }
func (r *VolumeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeNewParams struct {
	// Name of the volume.
	Name string `json:"name,required"`
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	// ID of the VM the volume is attached to.
	VMID string `json:"vm_id,required"`
	paramObj
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeUpdateParams struct {
	// Name of the volume.
	Name param.Opt[string] `json:"name,omitzero"`
	// Size of the volume in GB.
	Size param.Opt[int64] `json:"size,omitzero"`
	paramObj
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
