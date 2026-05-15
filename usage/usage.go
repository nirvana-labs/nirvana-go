// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// UsageService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.Options = opts
	return
}

// List per-resource usage records for the current organization. Each item is one
// resource with its nested dimension history (active and closed segments).
func (r *UsageService) List(ctx context.Context, query UsageListParams, opts ...option.RequestOption) (res *pagination.Cursor[Usage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/usage"
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

// List per-resource usage records for the current organization. Each item is one
// resource with its nested dimension history (active and closed segments).
func (r *UsageService) ListAutoPaging(ctx context.Context, query UsageListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Usage] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get the usage record for a single resource (metadata plus dimension history) for
// the current organization.
func (r *UsageService) Get(ctx context.Context, resourceID string, opts ...option.RequestOption) (res *Usage, err error) {
	opts = slices.Concat(r.Options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/usage/%s", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Usage record for a single metered resource.
type Usage struct {
	Dimensions []UsageDimension `json:"dimensions" api:"required"`
	EndedAt    string           `json:"ended_at" api:"required"`
	ProjectID  string           `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region     shared.RegionName `json:"region" api:"required"`
	ResourceID string            `json:"resource_id" api:"required"`
	// Kind of metered resource a ledger row belongs to.
	//
	// Any of "vm", "volume", "vpc", "connect_connection", "nks_cluster",
	// "nks_node_pool", "nks_load_balancer".
	ResourceType UsageResourceType `json:"resource_type" api:"required"`
	StartedAt    string            `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimensions   respjson.Field
		EndedAt      respjson.Field
		ProjectID    respjson.Field
		Region       respjson.Field
		ResourceID   respjson.Field
		ResourceType respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Usage) RawJSON() string { return r.JSON.raw }
func (r *Usage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level dimension entry; bundle heads expand via children.
type UsageDimension struct {
	ID        string               `json:"id" api:"required"`
	Dimension string               `json:"dimension" api:"required"`
	EndedAt   string               `json:"ended_at" api:"required"`
	Quantity  int64                `json:"quantity" api:"required"`
	StartedAt string               `json:"started_at" api:"required"`
	Children  []UsageDimensionLeaf `json:"children"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Dimension   respjson.Field
		EndedAt     respjson.Field
		Quantity    respjson.Field
		StartedAt   respjson.Field
		Children    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageDimension) RawJSON() string { return r.JSON.raw }
func (r *UsageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One ledger segment for a (resource, dimension) pair.
type UsageDimensionLeaf struct {
	ID        string `json:"id" api:"required"`
	Dimension string `json:"dimension" api:"required"`
	EndedAt   string `json:"ended_at" api:"required"`
	Quantity  int64  `json:"quantity" api:"required"`
	StartedAt string `json:"started_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Dimension   respjson.Field
		EndedAt     respjson.Field
		Quantity    respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageDimensionLeaf) RawJSON() string { return r.JSON.raw }
func (r *UsageDimensionLeaf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageList struct {
	Items []Usage `json:"items" api:"required"`
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
func (r UsageList) RawJSON() string { return r.JSON.raw }
func (r *UsageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of metered resource a ledger row belongs to.
type UsageResourceType string

const (
	UsageResourceTypeVM                UsageResourceType = "vm"
	UsageResourceTypeVolume            UsageResourceType = "volume"
	UsageResourceTypeVPC               UsageResourceType = "vpc"
	UsageResourceTypeConnectConnection UsageResourceType = "connect_connection"
	UsageResourceTypeNKSCluster        UsageResourceType = "nks_cluster"
	UsageResourceTypeNKSNodePool       UsageResourceType = "nks_node_pool"
	UsageResourceTypeNKSLoadBalancer   UsageResourceType = "nks_load_balancer"
)

type UsageListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageListParams]'s query parameters as `url.Values`.
func (r UsageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
