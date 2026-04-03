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

// ClusterPersistentVolumeClaimService contains methods and other services that
// help with interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterPersistentVolumeClaimService] method instead.
type ClusterPersistentVolumeClaimService struct {
	Options []option.RequestOption
}

// NewClusterPersistentVolumeClaimService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClusterPersistentVolumeClaimService(opts ...option.RequestOption) (r ClusterPersistentVolumeClaimService) {
	r = ClusterPersistentVolumeClaimService{}
	r.Options = opts
	return
}

// List all persistent volume claims in an NKS cluster
func (r *ClusterPersistentVolumeClaimService) List(ctx context.Context, clusterID string, query ClusterPersistentVolumeClaimListParams, opts ...option.RequestOption) (res *pagination.Cursor[PersistentVolumeClaim], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/persistent_volume_claims", clusterID)
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

// List all persistent volume claims in an NKS cluster
func (r *ClusterPersistentVolumeClaimService) ListAutoPaging(ctx context.Context, clusterID string, query ClusterPersistentVolumeClaimListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PersistentVolumeClaim] {
	return pagination.NewCursorAutoPager(r.List(ctx, clusterID, query, opts...))
}

// Get details about an NKS persistent volume claim
func (r *ClusterPersistentVolumeClaimService) Get(ctx context.Context, clusterID string, persistentVolumeClaimID string, opts ...option.RequestOption) (res *PersistentVolumeClaim, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	if persistentVolumeClaimID == "" {
		err = errors.New("missing required persistent_volume_claim_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/persistent_volume_claims/%s", clusterID, persistentVolumeClaimID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// NKS persistent volume claim details.
type PersistentVolumeClaim struct {
	// Unique identifier for the persistent volume claim.
	ID string `json:"id" api:"required"`
	// Cluster this persistent volume claim belongs to.
	ClusterID string `json:"cluster_id" api:"required"`
	// When the persistent volume claim was first discovered.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the persistent volume claim.
	Name string `json:"name" api:"required"`
	// Size of the persistent volume claim in GB.
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
	// When the persistent volume claim was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClusterID   respjson.Field
		CreatedAt   respjson.Field
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
func (r PersistentVolumeClaim) RawJSON() string { return r.JSON.raw }
func (r *PersistentVolumeClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersistentVolumeClaimList struct {
	Items []PersistentVolumeClaim `json:"items" api:"required"`
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
func (r PersistentVolumeClaimList) RawJSON() string { return r.JSON.raw }
func (r *PersistentVolumeClaimList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterPersistentVolumeClaimListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterPersistentVolumeClaimListParams]'s query parameters
// as `url.Values`.
func (r ClusterPersistentVolumeClaimListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
