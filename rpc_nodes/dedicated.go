// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

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

// DedicatedService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDedicatedService] method instead.
type DedicatedService struct {
	Options     []option.RequestOption
	Blockchains DedicatedBlockchainService
}

// NewDedicatedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDedicatedService(opts ...option.RequestOption) (r DedicatedService) {
	r = DedicatedService{}
	r.Options = opts
	r.Blockchains = NewDedicatedBlockchainService(opts...)
	return
}

// List all RPC Node Dedicated you created
func (r *DedicatedService) List(ctx context.Context, query DedicatedListParams, opts ...option.RequestOption) (res *pagination.Cursor[Dedicated], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/rpc_nodes/dedicated"
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

// List all RPC Node Dedicated you created
func (r *DedicatedService) ListAutoPaging(ctx context.Context, query DedicatedListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Dedicated] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get details about an RPC Node Dedicated
func (r *DedicatedService) Get(ctx context.Context, nodeID string, opts ...option.RequestOption) (res *Dedicated, err error) {
	opts = slices.Concat(r.Options, opts)
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return
	}
	path := fmt.Sprintf("v1/rpc_nodes/dedicated/%s", nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// RPC Node Dedicated details.
type Dedicated struct {
	// Unique identifier for the RPC Node Dedicated.
	ID string `json:"id,required"`
	// Blockchain type.
	Blockchain string `json:"blockchain,required"`
	// When the RPC Node Dedicated was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// RPC endpoint URL.
	Endpoint string `json:"endpoint,required"`
	// Name of the RPC Node Dedicated.
	Name string `json:"name,required"`
	// Network type (e.g., mainnet, testnet).
	Network string `json:"network,required"`
	// Project identifier associated with the RPC Node Dedicated.
	ProjectID string `json:"project_id,required"`
	// Tags to attach to the RPC Node Dedicated.
	Tags []string `json:"tags,required"`
	// When the RPC Node Dedicated was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Blockchain  respjson.Field
		CreatedAt   respjson.Field
		Endpoint    respjson.Field
		Name        respjson.Field
		Network     respjson.Field
		ProjectID   respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Dedicated) RawJSON() string { return r.JSON.raw }
func (r *Dedicated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Blockchain supported by the RPC Node Dedicated.
type DedicatedBlockchain struct {
	// Blockchain type.
	Blockchain string `json:"blockchain,required"`
	// Network type (e.g., mainnet, testnet).
	Network string `json:"network,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain  respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DedicatedBlockchain) RawJSON() string { return r.JSON.raw }
func (r *DedicatedBlockchain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DedicatedBlockchainList struct {
	Items []DedicatedBlockchain `json:"items,required"`
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
func (r DedicatedBlockchainList) RawJSON() string { return r.JSON.raw }
func (r *DedicatedBlockchainList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DedicatedList struct {
	Items []Dedicated `json:"items,required"`
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
func (r DedicatedList) RawJSON() string { return r.JSON.raw }
func (r *DedicatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DedicatedListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id,required" json:"-"`
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DedicatedListParams]'s query parameters as `url.Values`.
func (r DedicatedListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
