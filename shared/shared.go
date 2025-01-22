// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

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
