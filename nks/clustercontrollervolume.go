// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ClusterControllerVolumeService contains methods and other services that help
// with interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterControllerVolumeService] method instead.
type ClusterControllerVolumeService struct {
	Options []option.RequestOption
}

// NewClusterControllerVolumeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterControllerVolumeService(opts ...option.RequestOption) (r ClusterControllerVolumeService) {
	r = ClusterControllerVolumeService{}
	r.Options = opts
	return
}

// List all volumes attached to an NKS controller
func (r *ClusterControllerVolumeService) List(ctx context.Context, clusterID string, controllerID string, query ClusterControllerVolumeListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSControllerVolume], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if controllerID == "" {
		err = errors.New("missing required controller_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/controllers/%s/volumes", clusterID, controllerID)
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

// List all volumes attached to an NKS controller
func (r *ClusterControllerVolumeService) ListAutoPaging(ctx context.Context, clusterID string, controllerID string, query ClusterControllerVolumeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSControllerVolume] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, controllerID, query, opts...))
}

// Get details about a volume attached to an NKS controller
func (r *ClusterControllerVolumeService) Get(ctx context.Context, clusterID string, controllerID string, volumeID string, opts ...option.RequestOption) (res *NKSControllerVolume, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if controllerID == "" {
		err = errors.New("missing required controller_id parameter")
		return nil, err
	}
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/controllers/%s/volumes/%s", clusterID, controllerID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS controller volume details.
type NKSControllerVolume struct {
	// Unique identifier for the volume.
	ID string `json:"id" api:"required"`
	// When the volume was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Volume kind.
	//
	// Any of "boot", "data".
	Kind compute.VolumeKind `json:"kind" api:"required"`
	// Name of the volume.
	Name string `json:"name" api:"required"`
	// Size of the volume in GB.
	Size int64 `json:"size" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Type of the Volume.
	//
	// Any of "nvme", "abs".
	Type compute.VolumeType `json:"type" api:"required"`
	// When the volume was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSControllerVolume) RawJSON() string { return r.JSON.raw }
func (r *NKSControllerVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSControllerVolumeList struct {
	Items []NKSControllerVolume `json:"items" api:"required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSControllerVolumeList) RawJSON() string { return r.JSON.raw }
func (r *NKSControllerVolumeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterControllerVolumeListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterControllerVolumeListParams]'s query parameters as
// `url.Values`.
func (r ClusterControllerVolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
