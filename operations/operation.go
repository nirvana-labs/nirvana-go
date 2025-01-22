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
func NewOperationService(opts ...option.RequestOption) (r *OperationService) {
	r = &OperationService{}
	r.Options = opts
	return
}

// Get details about a specific operation
func (r *OperationService) Get(ctx context.Context, operationID string, opts ...option.RequestOption) (res *Operation, err error) {
	opts = append(r.Options[:], opts...)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Operation details.
type Operation struct {
	ID         string          `json:"id,required"`
	Kind       OperationKind   `json:"kind,required"`
	ResourceID string          `json:"resource_id,required"`
	Status     OperationStatus `json:"status,required"`
	Type       OperationType   `json:"type,required"`
	JSON       operationJSON   `json:"-"`
}

// operationJSON contains the JSON metadata for the struct [Operation]
type operationJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	ResourceID  apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Operation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationJSON) RawJSON() string {
	return r.raw
}

type OperationKind string

const (
	OperationKindVM           OperationKind = "VM"
	OperationKindVolume       OperationKind = "VOLUME"
	OperationKindVPC          OperationKind = "VPC"
	OperationKindFirewallRule OperationKind = "FIREWALL_RULE"
)

func (r OperationKind) IsKnown() bool {
	switch r {
	case OperationKindVM, OperationKindVolume, OperationKindVPC, OperationKindFirewallRule:
		return true
	}
	return false
}

type OperationStatus string

const (
	OperationStatusPending OperationStatus = "PENDING"
	OperationStatusRunning OperationStatus = "RUNNING"
	OperationStatusDone    OperationStatus = "DONE"
	OperationStatusFailed  OperationStatus = "FAILED"
	OperationStatusUnknown OperationStatus = "UNKNOWN"
)

func (r OperationStatus) IsKnown() bool {
	switch r {
	case OperationStatusPending, OperationStatusRunning, OperationStatusDone, OperationStatusFailed, OperationStatusUnknown:
		return true
	}
	return false
}

type OperationType string

const (
	OperationTypeCreate OperationType = "CREATE"
	OperationTypeUpdate OperationType = "UPDATE"
	OperationTypeDelete OperationType = "DELETE"
)

func (r OperationType) IsKnown() bool {
	switch r {
	case OperationTypeCreate, OperationTypeUpdate, OperationTypeDelete:
		return true
	}
	return false
}
