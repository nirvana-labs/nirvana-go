// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/operations"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// BucketService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
type BucketService struct {
	Options      []option.RequestOption
	Availability BucketAvailabilityService
	Cost         BucketCostService
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r BucketService) {
	r = BucketService{}
	r.Options = opts
	r.Availability = NewBucketAvailabilityService(opts...)
	r.Cost = NewBucketCostService(opts...)
	return
}

// Create a Bucket. The Bucket is created asynchronously: follow the returned
// operation until it is done.
func (r *BucketService) New(ctx context.Context, body BucketNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/storage/buckets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a Bucket's tags. A Bucket's name is its S3 name and cannot be changed.
func (r *BucketService) Update(ctx context.Context, bucketID string, body BucketUpdateParams, opts ...option.RequestOption) (res *Bucket, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketID == "" {
		err = errors.New("missing required bucket_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/storage/buckets/%s", bucketID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all Buckets
func (r *BucketService) List(ctx context.Context, query BucketListParams, opts ...option.RequestOption) (res *pagination.Cursor[Bucket], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/storage/buckets"
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

// List all Buckets
func (r *BucketService) ListAutoPaging(ctx context.Context, query BucketListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Bucket] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a Bucket. The Bucket is deleted asynchronously: follow the returned
// operation until it is done. A Bucket that still holds objects is only deleted
// with force.
func (r *BucketService) Delete(ctx context.Context, bucketID string, body BucketDeleteParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketID == "" {
		err = errors.New("missing required bucket_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/storage/buckets/%s", bucketID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Get details about a Bucket
func (r *BucketService) Get(ctx context.Context, bucketID string, opts ...option.RequestOption) (res *Bucket, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketID == "" {
		err = errors.New("missing required bucket_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/storage/buckets/%s", bucketID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Bucket details.
type Bucket struct {
	// Unique identifier for the Bucket.
	ID string `json:"id" api:"required"`
	// When the Bucket was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// S3 endpoint of the Bucket's region. Point S3 clients at it, with an Object
	// Storage Key of the same project and region.
	Endpoint string `json:"endpoint" api:"required"`
	// Name of the Bucket, used to address it over S3.
	Name string `json:"name" api:"required"`
	// Number of objects in the Bucket. Updated periodically, so it can lag behind
	// recent writes.
	ObjectCount int64 `json:"object_count" api:"required"`
	// Project ID the Bucket belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Total size of the objects in the Bucket, in bytes. Updated periodically, so it
	// can lag behind recent writes.
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Tags attached to the Bucket.
	Tags []string `json:"tags" api:"required"`
	// When the Bucket was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Path-style URL of the Bucket.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Endpoint    respjson.Field
		Name        respjson.Field
		ObjectCount respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		SizeBytes   respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Bucket) RawJSON() string { return r.JSON.raw }
func (r *Bucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketList struct {
	Items []Bucket `json:"items" api:"required"`
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
func (r BucketList) RawJSON() string { return r.JSON.raw }
func (r *BucketList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketNewParams struct {
	// Name of the Bucket, used to address it over S3. It must be 3 to 63 characters of
	// lowercase letters, digits, dots and hyphens, start and end with a letter or
	// digit, not contain two adjacent dots and not be formatted as an IP address.
	// "admin" and "ops-api" are reserved. The name cannot be changed once the Bucket
	// is created.
	Name string `json:"name" api:"required"`
	// Project ID the Bucket belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Tags to attach to the Bucket.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BucketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketUpdateParams struct {
	// Tags to attach to the Bucket.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BucketUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id" api:"required" json:"-"`
	// Pagination cursor returned by a previous request. Only valid for the same
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the Bucket name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by region
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at, name, status, size_bytes, object_count
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by Bucket status
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "error".
	Status BucketListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by tags. Repeat the parameter to require several tags; a Bucket must
	// carry all of them.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BucketListParams]'s query parameters as `url.Values`.
func (r BucketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by Bucket status
type BucketListParamsStatus string

const (
	BucketListParamsStatusPending  BucketListParamsStatus = "pending"
	BucketListParamsStatusCreating BucketListParamsStatus = "creating"
	BucketListParamsStatusUpdating BucketListParamsStatus = "updating"
	BucketListParamsStatusReady    BucketListParamsStatus = "ready"
	BucketListParamsStatusDeleting BucketListParamsStatus = "deleting"
	BucketListParamsStatusError    BucketListParamsStatus = "error"
)

type BucketDeleteParams struct {
	// Permanently delete every object in the Bucket, including all versions, before
	// deleting it. Without it, deleting a Bucket that holds objects fails.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BucketDeleteParams]'s query parameters as `url.Values`.
func (r BucketDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
