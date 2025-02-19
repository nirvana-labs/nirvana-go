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
	path := "compute/volumes"
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
	path := fmt.Sprintf("compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all volumes
func (r *VolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *VolumeList, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/volumes"
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
	path := fmt.Sprintf("compute/volumes/%s", volumeID)
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
	Name string     `json:"name,required"`
	Size int64      `json:"size,required"`
	// Storage type.
	Type      StorageType `json:"type,required"`
	UpdatedAt string      `json:"updated_at,required"`
	VMID      string      `json:"vm_id,required"`
	JSON      volumeJSON  `json:"-"`
}

// volumeJSON contains the JSON metadata for the struct [Volume]
type volumeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	VMID        apijson.Field
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

type VolumeList struct {
	Items []Volume       `json:"items,required"`
	JSON  volumeListJSON `json:"-"`
}

// volumeListJSON contains the JSON metadata for the struct [VolumeList]
type volumeListJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r volumeListJSON) RawJSON() string {
	return r.raw
}

type VolumeNewParams struct {
	Size param.Field[int64]  `json:"size,required"`
	VMID param.Field[string] `json:"vm_id,required"`
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeUpdateParams struct {
	Size param.Field[int64] `json:"size,required"`
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
