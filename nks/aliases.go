// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

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

// Cost quote returned by POST /cost.
//
// This is an alias to an internal type.
type CostQuote = shared.CostQuote

// Priced row for a single usage dimension emitted by a resource.
//
// This is an alias to an internal type.
type CostQuoteUsageDimension = shared.CostQuoteUsageDimension

// Cost quote returned by PATCH /:id/cost: the current-state quote, the post-update
// quote, and the signed diff.
//
// This is an alias to an internal type.
type CostQuoteUpdate = shared.CostQuoteUpdate

// Quote for the proposed (post-update) resource state.
//
// This is an alias to an internal type.
type CostQuoteUpdateAfter = shared.CostQuoteUpdateAfter

// Priced row for a single usage dimension emitted by a resource.
//
// This is an alias to an internal type.
type CostQuoteUpdateAfterUsageDimension = shared.CostQuoteUpdateAfterUsageDimension

// Quote for the proposed (post-update) resource state.
//
// This is an alias to an internal type.
type CostQuoteUpdateBefore = shared.CostQuoteUpdateBefore

// Priced row for a single usage dimension emitted by a resource.
//
// This is an alias to an internal type.
type CostQuoteUpdateBeforeUsageDimension = shared.CostQuoteUpdateBeforeUsageDimension

// Per-dimension and total deltas: after minus before.
//
// This is an alias to an internal type.
type CostQuoteUpdateDiff = shared.CostQuoteUpdateDiff

// Per-dimension diff entry. Both before and after are always present.
//
// This is an alias to an internal type.
type CostQuoteUpdateDiffUsageDimension = shared.CostQuoteUpdateDiffUsageDimension

// Priced row after the update. Always present.
//
// This is an alias to an internal type.
type CostQuoteUpdateDiffUsageDimensionAfter = shared.CostQuoteUpdateDiffUsageDimensionAfter

// Priced row after the update. Always present.
//
// This is an alias to an internal type.
type CostQuoteUpdateDiffUsageDimensionBefore = shared.CostQuoteUpdateDiffUsageDimensionBefore

// Pagination response details.
//
// This is an alias to an internal type.
type Pagination = shared.Pagination

// Region the resource is in.
//
// This is an alias to an internal type.
type RegionName = shared.RegionName

// Equals "us-sva-2"
const RegionNameUsSva2 = shared.RegionNameUsSva2

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

// IP filter rules.
//
// This is an alias to an internal type.
type SourceIPRuleParam = shared.SourceIPRuleParam

// IP filter rules.
//
// This is an alias to an internal type.
type SourceIPRuleResponse = shared.SourceIPRuleResponse
