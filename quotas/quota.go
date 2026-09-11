// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quotas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// QuotaService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options []option.RequestOption
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r QuotaService) {
	r = QuotaService{}
	r.Options = opts
	return
}

// List quota usage and limits for the current organization across all regions
func (r *QuotaService) List(ctx context.Context, query QuotaListParams, opts ...option.RequestOption) (res *pagination.Cursor[Quota], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/quotas"
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

// List quota usage and limits for the current organization across all regions
func (r *QuotaService) ListAutoPaging(ctx context.Context, query QuotaListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Quota] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get quota usage and limits for the current organization in a single region
func (r *QuotaService) Get(ctx context.Context, region QuotaGetParamsRegion, opts ...option.RequestOption) (res *Quota, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/quotas/%v", region)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Quota response.
type Quota struct {
	// Compute quota.
	Compute QuotaCompute `json:"compute" api:"required"`
	// Networking quota.
	Networking QuotaNetworking `json:"networking" api:"required"`
	// NKS quota.
	NKS QuotaNKS `json:"nks" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region" api:"required"`
	// Storage quota.
	Storage QuotaStorage `json:"storage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compute     respjson.Field
		Networking  respjson.Field
		NKS         respjson.Field
		Region      respjson.Field
		Storage     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quota) RawJSON() string { return r.JSON.raw }
func (r *Quota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute quota.
type QuotaCompute struct {
	// Quota dimension detail.
	MemoryGB QuotaDimensionDetail `json:"memory_gb" api:"required"`
	// Quota dimension detail.
	Vcpu QuotaDimensionDetail `json:"vcpu" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryGB    respjson.Field
		Vcpu        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaCompute) RawJSON() string { return r.JSON.raw }
func (r *QuotaCompute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quota dimension detail.
type QuotaDimensionDetail struct {
	Limit     int64 `json:"limit" api:"required"`
	Remaining int64 `json:"remaining" api:"required"`
	Used      int64 `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Remaining   respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaDimensionDetail) RawJSON() string { return r.JSON.raw }
func (r *QuotaDimensionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaList struct {
	Items []Quota `json:"items" api:"required"`
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
func (r QuotaList) RawJSON() string { return r.JSON.raw }
func (r *QuotaList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networking quota.
type QuotaNetworking struct {
	// Quota dimension detail.
	PublicIPs QuotaDimensionDetail `json:"public_ips" api:"required"`
	// Quota dimension detail.
	VPCs QuotaDimensionDetail `json:"vpcs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublicIPs   respjson.Field
		VPCs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaNetworking) RawJSON() string { return r.JSON.raw }
func (r *QuotaNetworking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NKS quota.
type QuotaNKS struct {
	// Quota dimension detail.
	Clusters QuotaDimensionDetail `json:"clusters" api:"required"`
	// Quota dimension detail.
	NodePoolMemoryGB QuotaDimensionDetail `json:"node_pool_memory_gb" api:"required"`
	// Quota dimension detail.
	NodePoolVcpu QuotaDimensionDetail `json:"node_pool_vcpu" api:"required"`
	// Quota dimension detail.
	PublicIPs QuotaDimensionDetail `json:"public_ips" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clusters         respjson.Field
		NodePoolMemoryGB respjson.Field
		NodePoolVcpu     respjson.Field
		PublicIPs        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaNKS) RawJSON() string { return r.JSON.raw }
func (r *QuotaNKS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage quota.
type QuotaStorage struct {
	// Quota dimension detail.
	ABS QuotaDimensionDetail `json:"abs" api:"required"`
	// Quota dimension detail.
	LocalNvme QuotaDimensionDetail `json:"local_nvme" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ABS         respjson.Field
		LocalNvme   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaStorage) RawJSON() string { return r.JSON.raw }
func (r *QuotaStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaListParams struct {
	// Pagination cursor returned by a previous request. Only valid for the same
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by region
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: region
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QuotaListParams]'s query parameters as `url.Values`.
func (r QuotaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuotaGetParamsRegion string

const (
	QuotaGetParamsRegionUsSva2 QuotaGetParamsRegion = "us-sva-2"
)
