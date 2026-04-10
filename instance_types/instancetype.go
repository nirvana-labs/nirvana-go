// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package instance_types

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

// InstanceTypeService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceTypeService] method instead.
type InstanceTypeService struct {
	Options []option.RequestOption
}

// NewInstanceTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceTypeService(opts ...option.RequestOption) (r InstanceTypeService) {
	r = InstanceTypeService{}
	r.Options = opts
	return
}

// List instance types
func (r *InstanceTypeService) List(ctx context.Context, query InstanceTypeListParams, opts ...option.RequestOption) (res *pagination.Cursor[InstanceType], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/instance_types"
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

// List instance types
func (r *InstanceTypeService) ListAutoPaging(ctx context.Context, query InstanceTypeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[InstanceType] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get an instance type by region and name
func (r *InstanceTypeService) Get(ctx context.Context, region InstanceTypeGetParamsRegion, name string, opts ...option.RequestOption) (res *InstanceType, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instance_types/%v/%s", region, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Instance type.
type InstanceType struct {
	Chipset string `json:"chipset" api:"required"`
	// When the Instance Type was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Family    string    `json:"family" api:"required"`
	MemoryGB  int64     `json:"memory_gb" api:"required"`
	Name      string    `json:"name" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-sva-2", "us-chi-1", "ap-sin-1".
	Region shared.RegionName `json:"region" api:"required"`
	Series string            `json:"series" api:"required"`
	// When the Instance Type was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	Vcpu      int64     `json:"vcpu" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chipset     respjson.Field
		CreatedAt   respjson.Field
		Family      respjson.Field
		MemoryGB    respjson.Field
		Name        respjson.Field
		Region      respjson.Field
		Series      respjson.Field
		UpdatedAt   respjson.Field
		Vcpu        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceType) RawJSON() string { return r.JSON.raw }
func (r *InstanceType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceTypeList struct {
	Items []InstanceType `json:"items" api:"required"`
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
func (r InstanceTypeList) RawJSON() string { return r.JSON.raw }
func (r *InstanceTypeList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceTypeListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceTypeListParams]'s query parameters as `url.Values`.
func (r InstanceTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstanceTypeGetParamsRegion string

const (
	InstanceTypeGetParamsRegionUsSea1 InstanceTypeGetParamsRegion = "us-sea-1"
	InstanceTypeGetParamsRegionUsSva1 InstanceTypeGetParamsRegion = "us-sva-1"
	InstanceTypeGetParamsRegionUsSva2 InstanceTypeGetParamsRegion = "us-sva-2"
	InstanceTypeGetParamsRegionUsChi1 InstanceTypeGetParamsRegion = "us-chi-1"
	InstanceTypeGetParamsRegionApSin1 InstanceTypeGetParamsRegion = "ap-sin-1"
)
