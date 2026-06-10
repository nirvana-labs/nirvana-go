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

// VolumeCostService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeCostService] method instead.
type VolumeCostService struct {
	Options []option.RequestOption
}

// NewVolumeCostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVolumeCostService(opts ...option.RequestOption) (r VolumeCostService) {
	r = VolumeCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed Volume.
func (r *VolumeCostService) New(ctx context.Context, body VolumeCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/compute/volumes/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed Volume update plus a diff against
// the current state.
func (r *VolumeCostService) Update(ctx context.Context, volumeID string, body VolumeCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/compute/volumes/%s/cost", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type VolumeCostNewParams struct {
	// Name of the Volume.
	Name string `json:"name" api:"required"`
	// Project ID the Volume belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Size of the Volume in GB.
	Size int64 `json:"size" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type VolumeType `json:"type,omitzero" api:"required"`
	// ID of the VM the Volume is attached to.
	VMID param.Opt[string] `json:"vm_id,omitzero"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VolumeCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeCostUpdateParams struct {
	// Name of the Volume.
	Name param.Opt[string] `json:"name,omitzero"`
	// Size of the Volume in GB.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VolumeCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
