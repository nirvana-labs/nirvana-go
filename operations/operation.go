// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package operations

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

// OperationService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	Options []option.RequestOption
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r OperationService) {
	r = OperationService{}
	r.Options = opts
	return
}

// List all operations
func (r *OperationService) List(ctx context.Context, query OperationListParams, opts ...option.RequestOption) (res *pagination.Cursor[Operation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/operations"
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

// List all operations
func (r *OperationService) ListAutoPaging(ctx context.Context, query OperationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Operation] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get details about a specific operation
func (r *OperationService) Get(ctx context.Context, operationID string, opts ...option.RequestOption) (res *Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Operation details.
type Operation struct {
	// Unique identifier for the Operation.
	ID string `json:"id,required"`
	// When the Operation was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of Operation.
	//
	// Any of "vm", "volume", "vpc", "firewall_rule".
	Kind OperationKind `json:"kind,required"`
	// Project ID the Operation belongs to.
	ProjectID string `json:"project_id,required"`
	// ID of the resource that the Operation is acting on.
	ResourceID string `json:"resource_id,required"`
	// Status of the Operation.
	//
	// Any of "pending", "running", "done", "failed", "unknown".
	Status OperationStatus `json:"status,required"`
	// Type of Operation.
	//
	// Any of "create", "update", "delete", "restart".
	Type OperationType `json:"type,required"`
	// When the Operation was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		ProjectID   respjson.Field
		ResourceID  respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Operation) RawJSON() string { return r.JSON.raw }
func (r *Operation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of Operation.
type OperationKind string

const (
	OperationKindVM           OperationKind = "vm"
	OperationKindVolume       OperationKind = "volume"
	OperationKindVPC          OperationKind = "vpc"
	OperationKindFirewallRule OperationKind = "firewall_rule"
)

type OperationList struct {
	Items []Operation `json:"items,required"`
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
func (r OperationList) RawJSON() string { return r.JSON.raw }
func (r *OperationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Operation.
type OperationStatus string

const (
	OperationStatusPending OperationStatus = "pending"
	OperationStatusRunning OperationStatus = "running"
	OperationStatusDone    OperationStatus = "done"
	OperationStatusFailed  OperationStatus = "failed"
	OperationStatusUnknown OperationStatus = "unknown"
)

// Type of Operation.
type OperationType string

const (
	OperationTypeCreate  OperationType = "create"
	OperationTypeUpdate  OperationType = "update"
	OperationTypeDelete  OperationType = "delete"
	OperationTypeRestart OperationType = "restart"
)

type OperationListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id,required" json:"-"`
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationListParams]'s query parameters as `url.Values`.
func (r OperationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
