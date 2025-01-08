// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/nirvana-labs/nirvana-go/internal/apijson"
)

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
	OperationKindVPC          OperationKind = "VPC"
	OperationKindFirewallRule OperationKind = "FIREWALL_RULE"
)

func (r OperationKind) IsKnown() bool {
	switch r {
	case OperationKindVM, OperationKindVPC, OperationKindFirewallRule:
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
)

func (r OperationStatus) IsKnown() bool {
	switch r {
	case OperationStatusPending, OperationStatusRunning, OperationStatusDone, OperationStatusFailed:
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

type RegionName string

const (
	RegionNameAmsterdam     RegionName = "amsterdam"
	RegionNameChicago       RegionName = "chicago"
	RegionNameFrankfurt     RegionName = "frankfurt"
	RegionNameHongkong      RegionName = "hongkong"
	RegionNameLondon        RegionName = "london"
	RegionNameMumbai        RegionName = "mumbai"
	RegionNameSaopaulo      RegionName = "saopaulo"
	RegionNameSeattle       RegionName = "seattle"
	RegionNameSiliconvalley RegionName = "siliconvalley"
	RegionNameSingapore     RegionName = "singapore"
	RegionNameStockholm     RegionName = "stockholm"
	RegionNameSydney        RegionName = "sydney"
	RegionNameTokyo         RegionName = "tokyo"
	RegionNameWashingtondc  RegionName = "washingtondc"
)

func (r RegionName) IsKnown() bool {
	switch r {
	case RegionNameAmsterdam, RegionNameChicago, RegionNameFrankfurt, RegionNameHongkong, RegionNameLondon, RegionNameMumbai, RegionNameSaopaulo, RegionNameSeattle, RegionNameSiliconvalley, RegionNameSingapore, RegionNameStockholm, RegionNameSydney, RegionNameTokyo, RegionNameWashingtondc:
		return true
	}
	return false
}

type ResourceStatus string

const (
	ResourceStatusPending  ResourceStatus = "PENDING"
	ResourceStatusCreating ResourceStatus = "CREATING"
	ResourceStatusUpdating ResourceStatus = "UPDATING"
	ResourceStatusReady    ResourceStatus = "READY"
	ResourceStatusDeleting ResourceStatus = "DELETING"
	ResourceStatusDeleted  ResourceStatus = "DELETED"
	ResourceStatusFailed   ResourceStatus = "FAILED"
)

func (r ResourceStatus) IsKnown() bool {
	switch r {
	case ResourceStatusPending, ResourceStatusCreating, ResourceStatusUpdating, ResourceStatusReady, ResourceStatusDeleting, ResourceStatusDeleted, ResourceStatusFailed:
		return true
	}
	return false
}
