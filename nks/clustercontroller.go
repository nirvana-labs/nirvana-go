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

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ClusterControllerService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterControllerService] method instead.
type ClusterControllerService struct {
	Options []option.RequestOption
	Volumes ClusterControllerVolumeService
}

// NewClusterControllerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterControllerService(opts ...option.RequestOption) (r ClusterControllerService) {
	r = ClusterControllerService{}
	r.Options = opts
	r.Volumes = NewClusterControllerVolumeService(opts...)
	return
}

// List all controllers in an NKS cluster
func (r *ClusterControllerService) List(ctx context.Context, clusterID string, query ClusterControllerListParams, opts ...option.RequestOption) (res *pagination.Cursor[NKSController], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/controllers", clusterID)
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

// List all controllers in an NKS cluster
func (r *ClusterControllerService) ListAutoPaging(ctx context.Context, clusterID string, query ClusterControllerListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NKSController] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, query, opts...))
}

// Get details about an NKS controller
func (r *ClusterControllerService) Get(ctx context.Context, clusterID string, controllerID string, opts ...option.RequestOption) (res *NKSController, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if controllerID == "" {
		err = errors.New("missing required controller_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/controllers/%s", clusterID, controllerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS controller details.
type NKSController struct {
	// Unique identifier for the controller.
	ID string `json:"id" api:"required"`
	// CPU configuration.
	CPUConfig NKSNodePoolCPUConfigResponse `json:"cpu_config" api:"required"`
	// When the controller was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Memory configuration.
	MemoryConfig NKSNodePoolMemoryConfigResponse `json:"memory_config" api:"required"`
	// Name of the controller.
	Name string `json:"name" api:"required"`
	// Private IP address of the controller.
	PrivateIP string `json:"private_ip" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the controller was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CPUConfig    respjson.Field
		CreatedAt    respjson.Field
		MemoryConfig respjson.Field
		Name         respjson.Field
		PrivateIP    respjson.Field
		Status       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NKSController) RawJSON() string { return r.JSON.raw }
func (r *NKSController) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NKSControllerList struct {
	Items []NKSController `json:"items" api:"required"`
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
func (r NKSControllerList) RawJSON() string { return r.JSON.raw }
func (r *NKSControllerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterControllerListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterControllerListParams]'s query parameters as
// `url.Values`.
func (r ClusterControllerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
