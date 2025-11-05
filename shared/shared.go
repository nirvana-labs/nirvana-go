// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Pagination response details.
type Pagination struct {
	NextCursor     string `json:"next_cursor,required"`
	PreviousCursor string `json:"previous_cursor,required"`
	TotalCount     int64  `json:"total_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor     respjson.Field
		PreviousCursor respjson.Field
		TotalCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region the resource is in.
type RegionName string

const (
	RegionNameUsSea1 RegionName = "us-sea-1"
	RegionNameUsSva1 RegionName = "us-sva-1"
	RegionNameUsChi1 RegionName = "us-chi-1"
	RegionNameUsWdc1 RegionName = "us-wdc-1"
	RegionNameEuFrk1 RegionName = "eu-frk-1"
	RegionNameApSin1 RegionName = "ap-sin-1"
	RegionNameApSeo1 RegionName = "ap-seo-1"
	RegionNameApTyo1 RegionName = "ap-tyo-1"
)

// Status of the resource.
type ResourceStatus string

const (
	ResourceStatusPending  ResourceStatus = "pending"
	ResourceStatusCreating ResourceStatus = "creating"
	ResourceStatusUpdating ResourceStatus = "updating"
	ResourceStatusReady    ResourceStatus = "ready"
	ResourceStatusDeleting ResourceStatus = "deleting"
	ResourceStatusDeleted  ResourceStatus = "deleted"
	ResourceStatusError    ResourceStatus = "error"
)
