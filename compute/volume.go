// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
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
	Options      []option.RequestOption
	Availability VolumeAvailabilityService
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	r.Availability = NewVolumeAvailabilityService(opts...)
	return
}

// Create a Volume. Only data volumes can be created.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/compute/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Volume. Boot or data volumes can be updated.
func (r *VolumeService) Update(ctx context.Context, volumeID string, body VolumeUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all volumes
func (r *VolumeService) List(ctx context.Context, query VolumeListParams, opts ...option.RequestOption) (res *pagination.Cursor[Volume], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/compute/volumes"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all volumes
func (r *VolumeService) ListAutoPaging(ctx context.Context, query VolumeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Volume] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a Volume. Boot or data volumes can be deleted.
func (r *VolumeService) Delete(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Attach a volume to a VM
func (r *VolumeService) Attach(ctx context.Context, volumeID string, body VolumeAttachParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s/attach", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach a volume from a VM
func (r *VolumeService) Detach(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s/detach", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get a Volume.
func (r *VolumeService) Get(ctx context.Context, volumeID string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("v1/compute/volumes/%s", volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Volume details.
type Volume struct {
	// Unique identifier for the Volume.
	ID string `json:"id,required"`
	// When the Volume was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Volume kind.
	//
	// Any of "boot", "data".
	Kind VolumeKind `json:"kind,required"`
	// Name of the Volume.
	Name string `json:"name,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "us-wdc-1", "eu-frk-1",
	// "ap-sin-1", "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Size of the Volume in GB.
	Size int64 `json:"size,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type VolumeType `json:"type,required"`
	// When the Volume was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the VM the Volume is attached to.
	VMID string `json:"vm_id,required"`
	// Name of the VM the Volume is attached to.
	VMName string `json:"vm_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Region      respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
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
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeList) RawJSON() string { return r.JSON.raw }
func (r *VolumeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the Volume.
type VolumeType string

const (
	VolumeTypeNvme VolumeType = "nvme"
	VolumeTypeABS  VolumeType = "abs"
)

type VolumeNewParams struct {
	// Name of the Volume.
	Name string `json:"name,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "us-wdc-1", "eu-frk-1",
	// "ap-sin-1", "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// Size of the Volume in GB.
	Size int64 `json:"size,required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type VolumeType `json:"type,omitzero,required"`
	// ID of the VM the Volume is attached to.
	VMID param.Opt[string] `json:"vm_id,omitzero"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
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
	// Name of the Volume.
	Name param.Opt[string] `json:"name,omitzero"`
	// Size of the Volume in GB.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Tags to attach to the Volume.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VolumeListParams]'s query parameters as `url.Values`.
func (r VolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VolumeAttachParams struct {
	// ID of the VM to attach the Volume to.
	VMID string `json:"vm_id,required"`
	paramObj
}

func (r VolumeAttachParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
