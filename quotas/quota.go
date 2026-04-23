// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quotas

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
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
func (r *QuotaService) List(ctx context.Context, opts ...option.RequestOption) (res *QuotaList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	// Any of "us-sva-1", "us-sva-2", "us-chi-1".
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
	// Quota resource detail.
	MemoryGB QuotaResourceDetail `json:"memory_gb" api:"required"`
	// Quota resource detail.
	Vcpu QuotaResourceDetail `json:"vcpu" api:"required"`
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
	// Quota resource detail.
	ConnectConnections QuotaResourceDetail `json:"connect_connections" api:"required"`
	// Quota resource detail.
	PublicIPs QuotaResourceDetail `json:"public_ips" api:"required"`
	// Quota resource detail.
	VPCs QuotaResourceDetail `json:"vpcs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectConnections respjson.Field
		PublicIPs          respjson.Field
		VPCs               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaNetworking) RawJSON() string { return r.JSON.raw }
func (r *QuotaNetworking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NKS quota.
type QuotaNKS struct {
	// Quota resource detail.
	Clusters QuotaResourceDetail `json:"clusters" api:"required"`
	// Quota resource detail.
	NodePoolMemoryGB QuotaResourceDetail `json:"node_pool_memory_gb" api:"required"`
	// Quota resource detail.
	NodePoolVcpu QuotaResourceDetail `json:"node_pool_vcpu" api:"required"`
	// Quota resource detail.
	PublicIPs QuotaResourceDetail `json:"public_ips" api:"required"`
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

// Quota resource detail.
type QuotaResourceDetail struct {
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
func (r QuotaResourceDetail) RawJSON() string { return r.JSON.raw }
func (r *QuotaResourceDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage quota.
type QuotaStorage struct {
	// Quota resource detail.
	ABS QuotaResourceDetail `json:"abs" api:"required"`
	// Quota resource detail.
	LocalNvme QuotaResourceDetail `json:"local_nvme" api:"required"`
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

type QuotaGetParamsRegion string

const (
	QuotaGetParamsRegionUsSva1 QuotaGetParamsRegion = "us-sva-1"
	QuotaGetParamsRegionUsSva2 QuotaGetParamsRegion = "us-sva-2"
	QuotaGetParamsRegionUsChi1 QuotaGetParamsRegion = "us-chi-1"
)
