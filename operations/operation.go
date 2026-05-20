// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package operations

import (
	"context"
	"encoding/json"
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
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Operation details.
type Operation struct {
	// Unique identifier for the Operation.
	ID string `json:"id" api:"required"`
	// When the Operation was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Kind of Operation.
	//
	// Any of "vm", "volume", "vpc", "firewall_rule", "nks_cluster", "nks_node_pool".
	Kind OperationKind `json:"kind" api:"required"`
	// Project ID the Operation belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// ID of the resource that the Operation is acting on.
	ResourceID string `json:"resource_id" api:"required"`
	// Status of the Operation.
	//
	// Any of "pending", "running", "done", "failed", "unknown".
	Status OperationStatus `json:"status" api:"required"`
	// Type of Operation.
	//
	// Any of "create", "update", "delete", "restart".
	Type OperationType `json:"type" api:"required"`
	// When the Operation was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Structured details about what an operation is changing.
	Details OperationDetails `json:"details"`
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
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Operation) RawJSON() string { return r.JSON.raw }
func (r *Operation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured details about what an operation is changing.
type OperationDetails struct {
	// Map of changed field names to their from/to diffs. Keys depend on the parent
	// operation's kind+type.
	Changes map[string]OperationDetailsChange `json:"changes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationDetails) RawJSON() string { return r.JSON.raw }
func (r *OperationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single field's before/after pair on an operation. Values are scalars (string,
// number, boolean) or string arrays.
type OperationDetailsChange struct {
	// Previous value.
	From OperationDetailsChangeFromUnion `json:"from" api:"required"`
	// New value.
	To OperationDetailsChangeToUnion `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationDetailsChange) RawJSON() string { return r.JSON.raw }
func (r *OperationDetailsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OperationDetailsChangeFromUnion contains all possible properties and values from
// [string], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OperationDetailsChangeFromUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u OperationDetailsChangeFromUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeFromUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeFromUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeFromUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OperationDetailsChangeFromUnion) RawJSON() string { return u.JSON.raw }

func (r *OperationDetailsChangeFromUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OperationDetailsChangeToUnion contains all possible properties and values from
// [string], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfStringArray]
type OperationDetailsChangeToUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u OperationDetailsChangeToUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeToUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeToUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OperationDetailsChangeToUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OperationDetailsChangeToUnion) RawJSON() string { return u.JSON.raw }

func (r *OperationDetailsChangeToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of Operation.
type OperationKind string

const (
	OperationKindVM           OperationKind = "vm"
	OperationKindVolume       OperationKind = "volume"
	OperationKindVPC          OperationKind = "vpc"
	OperationKindFirewallRule OperationKind = "firewall_rule"
	OperationKindNKSCluster   OperationKind = "nks_cluster"
	OperationKindNKSNodePool  OperationKind = "nks_node_pool"
)

type OperationList struct {
	Items []Operation `json:"items" api:"required"`
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
	ProjectID string `query:"project_id" api:"required" json:"-"`
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
