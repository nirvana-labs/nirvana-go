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
)

// VolumeAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeAvailabilityService] method instead.
type VolumeAvailabilityService struct {
	Options []option.RequestOption
}

// NewVolumeAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVolumeAvailabilityService(opts ...option.RequestOption) (r VolumeAvailabilityService) {
	r = VolumeAvailabilityService{}
	r.Options = opts
	return
}

// Check Volume Create Availability
func (r *VolumeAvailabilityService) New(ctx context.Context, body VolumeAvailabilityNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "v1/compute/volumes/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check Volume Update Availability
func (r *VolumeAvailabilityService) Update(ctx context.Context, volumeID string, body VolumeAvailabilityUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s/availability", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type VolumeAvailabilityNewParams struct {
	// Name of the volume.
	Name string `json:"name,required"`
	// Size of the volume in GB.
	Size int64 `json:"size,required"`
	// ID of the VM the volume is attached to.
	VMID string `json:"vm_id,required"`
	paramObj
}

func (r VolumeAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeAvailabilityUpdateParams struct {
	// Name of the volume.
	Name param.Opt[string] `json:"name,omitzero"`
	// Size of the volume in GB.
	Size param.Opt[int64] `json:"size,omitzero"`
	paramObj
}

func (r VolumeAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAvailabilityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
