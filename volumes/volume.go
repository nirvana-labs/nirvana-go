// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package volumes

import (
	"github.com/nirvana-labs/nirvana-go/internal/apijson"
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
