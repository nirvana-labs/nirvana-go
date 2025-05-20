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

// VMAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMAvailabilityService] method instead.
type VMAvailabilityService struct {
	Options []option.RequestOption
}

// NewVMAvailabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVMAvailabilityService(opts ...option.RequestOption) (r VMAvailabilityService) {
	r = VMAvailabilityService{}
	r.Options = opts
	return
}

// Check VM Create Availability
func (r *VMAvailabilityService) New(ctx context.Context, body VMAvailabilityNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/compute/vms/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Check VM Update Availability
func (r *VMAvailabilityService) Update(ctx context.Context, vmID string, body VMAvailabilityUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/vms/%s/availability", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

type VMAvailabilityNewParams struct {
	// VM create request.
	VMCreateRequest VMCreateRequestParam
	paramObj
}

func (r VMAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.VMCreateRequest)
}
func (r *VMAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return r.VMCreateRequest.UnmarshalJSON(data)
}

type VMAvailabilityUpdateParams struct {
	// VM update request.
	VMUpdateRequest VMUpdateRequestParam
	paramObj
}

func (r VMAvailabilityUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.VMUpdateRequest)
}
func (r *VMAvailabilityUpdateParams) UnmarshalJSON(data []byte) error {
	return r.VMUpdateRequest.UnmarshalJSON(data)
}
