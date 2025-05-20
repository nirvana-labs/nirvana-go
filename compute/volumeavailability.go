// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
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
func (r *VolumeAvailabilityService) New(ctx context.Context, body VolumeAvailabilityNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/compute/volumes/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Check Volume Update Availability
func (r *VolumeAvailabilityService) Update(ctx context.Context, volumeID string, body VolumeAvailabilityUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s/availability", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

type VolumeAvailabilityNewParams struct {
	// Volume data volume create request.
	VolumeCreateRequest VolumeCreateRequestParam
	paramObj
}

func (r VolumeAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.VolumeCreateRequest)
}
func (r *VolumeAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return r.VolumeCreateRequest.UnmarshalJSON(data)
}

type VolumeAvailabilityUpdateParams struct {
	// Volume update request.
	VolumeUpdateRequest VolumeUpdateRequestParam
	paramObj
}

func (r VolumeAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.VolumeUpdateRequest)
}
func (r *VolumeAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return r.VolumeUpdateRequest.UnmarshalJSON(data)
}
