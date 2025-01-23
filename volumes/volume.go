// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volumes

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
func NewVolumeService(opts ...option.RequestOption) (r *VolumeService) {
	r = &VolumeService{}
	r.Options = opts
	return
}

// Create a Volume. Only data volumes can be created.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "volumes"
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
	path := fmt.Sprintf("volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a Volume. Boot or data volumes can be deleted.
func (r *VolumeService) Delete(ctx context.Context, volumeID string, body VolumeDeleteParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
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

// Volume details.
type Volume struct {
	ID   string `json:"id,required"`
	Size int64  `json:"size,required"`
	// Storage type.
	Type StorageType `json:"type,required"`
	JSON volumeJSON  `json:"-"`
}

// volumeJSON contains the JSON metadata for the struct [Volume]
type volumeJSON struct {
	ID          apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Volume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeJSON) RawJSON() string {
	return r.raw
}

type VolumeNewParams struct {
	Size param.Field[int64]  `json:"size,required"`
	VMID param.Field[string] `json:"vm_id,required"`
	// Storage type.
	Type param.Field[StorageType] `json:"type"`
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeUpdateParams struct {
	Size param.Field[int64]  `json:"size,required"`
	VMID param.Field[string] `json:"vm_id,required"`
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeDeleteParams struct {
	VMID param.Field[string] `json:"vm_id,required"`
}

func (r VolumeDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
