// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

// Region the resource is in.
type RegionName string

const (
	RegionNameUsSea1 RegionName = "us-sea-1"
	RegionNameUsSva1 RegionName = "us-sva-1"
	RegionNameUsChi1 RegionName = "us-chi-1"
	RegionNameUsWdc1 RegionName = "us-wdc-1"
	RegionNameEuLon1 RegionName = "eu-lon-1"
	RegionNameEuAms1 RegionName = "eu-ams-1"
	RegionNameEuFrk1 RegionName = "eu-frk-1"
	RegionNameApSin1 RegionName = "ap-sin-1"
	RegionNameApSeo1 RegionName = "ap-seo-1"
	RegionNameApTyo1 RegionName = "ap-tyo-1"
)

func (r RegionName) IsKnown() bool {
	switch r {
	case RegionNameUsSea1, RegionNameUsSva1, RegionNameUsChi1, RegionNameUsWdc1, RegionNameEuLon1, RegionNameEuAms1, RegionNameEuFrk1, RegionNameApSin1, RegionNameApSeo1, RegionNameApTyo1:
		return true
	}
	return false
}

// Status of the resource.
type ResourceStatus string

const (
	ResourceStatusPending  ResourceStatus = "pending"
	ResourceStatusCreating ResourceStatus = "creating"
	ResourceStatusUpdating ResourceStatus = "updating"
	ResourceStatusReady    ResourceStatus = "ready"
	ResourceStatusDeleting ResourceStatus = "deleting"
	ResourceStatusDeleted  ResourceStatus = "deleted"
	ResourceStatusError    ResourceStatus = "error"
)

func (r ResourceStatus) IsKnown() bool {
	switch r {
	case ResourceStatusPending, ResourceStatusCreating, ResourceStatusUpdating, ResourceStatusReady, ResourceStatusDeleting, ResourceStatusDeleted, ResourceStatusError:
		return true
	}
	return false
}
