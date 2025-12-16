// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regions

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

// RegionService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r RegionService) {
	r = RegionService{}
	r.Options = opts
	return
}

// List all regions
func (r *RegionService) List(ctx context.Context, query RegionListParams, opts ...option.RequestOption) (res *pagination.Cursor[Region], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/regions"
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

// List all regions
func (r *RegionService) ListAutoPaging(ctx context.Context, query RegionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Region] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a region by name
func (r *RegionService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *Region, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("v1/regions/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Region response with product availability.
type Region struct {
	// Availability status of the region.
	//
	// Any of "live", "preview", "maintenance", "sunset".
	Availability RegionAvailability `json:"availability,required"`
	// Compute products available in this region.
	Compute RegionCompute `json:"compute,required"`
	// Name of the region.
	Name string `json:"name,required"`
	// Networking products available in this region.
	Networking RegionNetworking `json:"networking,required"`
	// Storage products available in this region.
	Storage RegionStorage `json:"storage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability respjson.Field
		Compute      respjson.Field
		Name         respjson.Field
		Networking   respjson.Field
		Storage      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Region) RawJSON() string { return r.JSON.raw }
func (r *Region) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute products available in this region.
type RegionCompute struct {
	// VMs indicates if Virtual Machines are available.
	VMs bool `json:"vms,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VMs         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCompute) RawJSON() string { return r.JSON.raw }
func (r *RegionCompute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Networking products available in this region.
type RegionNetworking struct {
	// Connect indicates if Nirvana Connect is available.
	Connect bool `json:"connect,required"`
	// VPCs indicates if Virtual Private Clouds are available.
	VPCs bool `json:"vpcs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connect     respjson.Field
		VPCs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionNetworking) RawJSON() string { return r.JSON.raw }
func (r *RegionNetworking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage products available in this region.
type RegionStorage struct {
	// ABS indicates if Accelerated Block Storage is available.
	ABS bool `json:"abs,required"`
	// LocalNVMe indicates if locally-attached NVMe storage is available.
	LocalNvme bool `json:"local_nvme,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ABS         respjson.Field
		LocalNvme   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionStorage) RawJSON() string { return r.JSON.raw }
func (r *RegionStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Availability status of the region.
type RegionAvailability string

const (
	RegionAvailabilityLive        RegionAvailability = "live"
	RegionAvailabilityPreview     RegionAvailability = "preview"
	RegionAvailabilityMaintenance RegionAvailability = "maintenance"
	RegionAvailabilitySunset      RegionAvailability = "sunset"
)

type RegionList struct {
	Items []Region `json:"items,required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionList) RawJSON() string { return r.JSON.raw }
func (r *RegionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegionListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegionListParams]'s query parameters as `url.Values`.
func (r RegionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
