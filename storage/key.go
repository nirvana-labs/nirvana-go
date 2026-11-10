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

	"github.com/nirvana-labs/nirvana-go/v3/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v3/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v3/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v3/shared"
)

// KeyService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r KeyService) {
	r = KeyService{}
	r.Options = opts
	return
}

// Create an S3 access key for a project's Buckets in a region. The response is the
// only one that includes the secret access key.
func (r *KeyService) New(ctx context.Context, body KeyNewParams, opts ...option.RequestOption) (res *ObjectStorageKeyWithSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/storage/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Rename an Object Storage Key. Its access key and secret do not change.
func (r *KeyService) Update(ctx context.Context, keyID string, body KeyUpdateParams, opts ...option.RequestOption) (res *ObjectStorageKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/storage/keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all Object Storage Keys. Secret access keys are not included.
func (r *KeyService) List(ctx context.Context, query KeyListParams, opts ...option.RequestOption) (res *pagination.Cursor[ObjectStorageKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/storage/keys"
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

// List all Object Storage Keys. Secret access keys are not included.
func (r *KeyService) ListAutoPaging(ctx context.Context, query KeyListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ObjectStorageKey] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an Object Storage Key. It stops working immediately.
func (r *KeyService) Delete(ctx context.Context, keyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/storage/keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get details about an Object Storage Key. The secret access key is not included.
func (r *KeyService) Get(ctx context.Context, keyID string, opts ...option.RequestOption) (res *ObjectStorageKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/storage/keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Object Storage Key details. The secret access key is only returned when the key
// is created.
type ObjectStorageKey struct {
	// Unique identifier for the Object Storage Key.
	ID string `json:"id" api:"required"`
	// S3 access key ID.
	AccessKey string `json:"access_key" api:"required"`
	// When the Object Storage Key was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the Object Storage Key.
	Name string `json:"name" api:"required"`
	// Project ID the Object Storage Key belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error", "suspended".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the Object Storage Key was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccessKey   respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectStorageKey) RawJSON() string { return r.JSON.raw }
func (r *ObjectStorageKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectStorageKeyList struct {
	Items []ObjectStorageKey `json:"items" api:"required"`
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
func (r ObjectStorageKeyList) RawJSON() string { return r.JSON.raw }
func (r *ObjectStorageKeyList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object Storage Key details, including the secret access key. Returned only when
// the key is created.
type ObjectStorageKeyWithSecret struct {
	// Unique identifier for the Object Storage Key.
	ID string `json:"id" api:"required"`
	// S3 access key ID.
	AccessKey string `json:"access_key" api:"required"`
	// When the Object Storage Key was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of the Object Storage Key.
	Name string `json:"name" api:"required"`
	// Project ID the Object Storage Key belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// S3 secret access key. It is only returned here and cannot be retrieved again:
	// store it securely, or delete the key and create a new one if it is lost.
	Secret string `json:"secret" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error", "suspended".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// When the Object Storage Key was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccessKey   respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		Secret      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectStorageKeyWithSecret) RawJSON() string { return r.JSON.raw }
func (r *ObjectStorageKeyWithSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyNewParams struct {
	// Name of the Object Storage Key.
	Name string `json:"name" api:"required"`
	// Project ID the Object Storage Key belongs to. The key can access the project's
	// Buckets in its region.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	paramObj
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyUpdateParams struct {
	// Name of the Object Storage Key.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r KeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow KeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id" api:"required" json:"-"`
	// Pagination cursor returned by a previous request. Only valid for the same
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the Object Storage Key name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by region
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at, name
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [KeyListParams]'s query parameters as `url.Values`.
func (r KeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
