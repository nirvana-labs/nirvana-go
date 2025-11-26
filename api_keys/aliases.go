// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_keys

import (
	"github.com/nirvana-labs/nirvana-go/internal/apierror"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Pagination response details.
//
// This is an alias to an internal type.
type Pagination = shared.Pagination

// Region the resource is in.
//
// This is an alias to an internal type.
type RegionName = shared.RegionName

// Equals "us-sea-1"
const RegionNameUsSea1 = shared.RegionNameUsSea1

// Equals "us-sva-1"
const RegionNameUsSva1 = shared.RegionNameUsSva1

// Equals "us-sva-2"
const RegionNameUsSva2 = shared.RegionNameUsSva2

// Equals "us-chi-1"
const RegionNameUsChi1 = shared.RegionNameUsChi1

// Equals "us-wdc-1"
const RegionNameUsWdc1 = shared.RegionNameUsWdc1

// Equals "eu-frk-1"
const RegionNameEuFrk1 = shared.RegionNameEuFrk1

// Equals "ap-sin-1"
const RegionNameApSin1 = shared.RegionNameApSin1

// Equals "ap-seo-1"
const RegionNameApSeo1 = shared.RegionNameApSeo1

// Equals "ap-tyo-1"
const RegionNameApTyo1 = shared.RegionNameApTyo1

// Status of the resource.
//
// This is an alias to an internal type.
type ResourceStatus = shared.ResourceStatus

// Equals "pending"
const ResourceStatusPending = shared.ResourceStatusPending

// Equals "creating"
const ResourceStatusCreating = shared.ResourceStatusCreating

// Equals "updating"
const ResourceStatusUpdating = shared.ResourceStatusUpdating

// Equals "ready"
const ResourceStatusReady = shared.ResourceStatusReady

// Equals "deleting"
const ResourceStatusDeleting = shared.ResourceStatusDeleting

// Equals "deleted"
const ResourceStatusDeleted = shared.ResourceStatusDeleted

// Equals "error"
const ResourceStatusError = shared.ResourceStatusError
