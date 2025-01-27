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

// ComputeVolumeService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeVolumeService] method instead.
type ComputeVolumeService struct {
	Options []option.RequestOption
}

// NewComputeVolumeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeVolumeService(opts ...option.RequestOption) (r *ComputeVolumeService) {
	r = &ComputeVolumeService{}
	r.Options = opts
	return
}

// Create a Volume. Only data volumes can be created.
func (r *ComputeVolumeService) New(ctx context.Context, body ComputeVolumeNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Volume. Boot or data volumes can be updated.
func (r *ComputeVolumeService) Update(ctx context.Context, volumeID string, body ComputeVolumeUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all volumes
func (r *ComputeVolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *ComputeVolumeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Volume. Boot or data volumes can be deleted.
func (r *ComputeVolumeService) Delete(ctx context.Context, volumeID string, body ComputeVolumeDeleteParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get a Volume.
func (r *ComputeVolumeService) Get(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	// Volume kind.
	Kind VolumeKind `json:"kind,required"`
	Size int64      `json:"size,required"`
	// Storage type.
	Type      StorageType `json:"type,required"`
	UpdatedAt string      `json:"updated_at,required"`
	JSON      volumeJSON  `json:"-"`
}

// volumeJSON contains the JSON metadata for the struct [Volume]
type volumeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Kind        apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Volume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeJSON) RawJSON() string {
	return r.raw
}

// Volume kind.
type VolumeKind string

const (
	VolumeKindBoot VolumeKind = "boot"
	VolumeKindData VolumeKind = "data"
)

func (r VolumeKind) IsKnown() bool {
	switch r {
	case VolumeKindBoot, VolumeKindData:
		return true
	}
	return false
}

type ComputeVolumeListResponse struct {
	Items []Volume                      `json:"items,required"`
	JSON  computeVolumeListResponseJSON `json:"-"`
}

// computeVolumeListResponseJSON contains the JSON metadata for the struct
// [ComputeVolumeListResponse]
type computeVolumeListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeVolumeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeVolumeListResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeVolumeNewParams struct {
	Size param.Field[int64]  `json:"size,required"`
	VMID param.Field[string] `json:"vm_id,required"`
	// Storage type.
	Type param.Field[StorageType] `json:"type"`
}

func (r ComputeVolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputeVolumeUpdateParams struct {
	Size param.Field[int64]  `json:"size,required"`
	VMID param.Field[string] `json:"vm_id,required"`
}

func (r ComputeVolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputeVolumeDeleteParams struct {
	VMID param.Field[string] `json:"vm_id,required"`
}

func (r ComputeVolumeDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
