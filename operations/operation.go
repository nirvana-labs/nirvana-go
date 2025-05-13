// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package operations

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
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
func (r *OperationService) List(ctx context.Context, opts ...option.RequestOption) (res *OperationList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/operations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about a specific operation
func (r *OperationService) Get(ctx context.Context, operationID string, opts ...option.RequestOption) (res *Operation, err error) {
	opts = append(r.Options[:], opts...)
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
	// Unique identifier for the operation.
	ID string `json:"id,required"`
	// When the operation was created.
	CreatedAt string `json:"created_at,required"`
	// Kind of operation.
	//
	// Any of "vm", "volume", "vpc", "firewall_rule".
	Kind OperationKind `json:"kind,required"`
	// ID of the resource that the operation is acting on.
	ResourceID string `json:"resource_id,required"`
	// Status of the operation.
	//
	// Any of "pending", "running", "done", "failed", "unknown".
	Status OperationStatus `json:"status,required"`
	// Type of operation.
	//
	// Any of "create", "update", "delete".
	Type OperationType `json:"type,required"`
	// When the operation was updated.
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
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

// Kind of operation.
type OperationKind string

const (
	OperationKindVM           OperationKind = "vm"
	OperationKindVolume       OperationKind = "volume"
	OperationKindVPC          OperationKind = "vpc"
	OperationKindFirewallRule OperationKind = "firewall_rule"
)

type OperationList struct {
	Items []Operation `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationList) RawJSON() string { return r.JSON.raw }
func (r *OperationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the operation.
type OperationStatus string

const (
	OperationStatusPending OperationStatus = "pending"
	OperationStatusRunning OperationStatus = "running"
	OperationStatusDone    OperationStatus = "done"
	OperationStatusFailed  OperationStatus = "failed"
	OperationStatusUnknown OperationStatus = "unknown"
)

// Type of operation.
type OperationType string

const (
	OperationTypeCreate OperationType = "create"
	OperationTypeUpdate OperationType = "update"
	OperationTypeDelete OperationType = "delete"
)
