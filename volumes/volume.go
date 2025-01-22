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

// Create a Volume
func (r *VolumeService) New(ctx context.Context, vmID string, body VolumeNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("vms/%s/volumes", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a volume
func (r *VolumeService) Delete(ctx context.Context, vmID string, volumeID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("vms/%s/volumes/%s", vmID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	Size param.Field[int64] `json:"size,required"`
	// Storage type.
	Type param.Field[StorageType] `json:"type"`
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
