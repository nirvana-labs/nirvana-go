// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Cost quote returned by POST /cost. current_summary and updated_summary hold the
// org billing summary now and with this resource; omitted when the caller cannot
// view billing.
type CostQuote struct {
	// Currency the quote is denominated in. Always "USD" in v1.
	Currency string `json:"currency" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyTotal string `json:"monthly_total" api:"required" format:"decimal"`
	// Timestamp the quote was priced at.
	PricedAt time.Time `json:"priced_at" api:"required" format:"date-time"`
	// Priced rows, one per usage dimension emitted by the resource.
	UsageDimensions []CostQuoteUsageDimension `json:"usage_dimensions" api:"required"`
	// Forward-looking billing summary for an organization. All costs are run-rate
	// projections from the organization's current active usage ("≈ $X/mo at current
	// usage").
	CurrentSummary OrganizationBillingSummary `json:"current_summary"`
	// Forward-looking billing summary for an organization. All costs are run-rate
	// projections from the organization's current active usage ("≈ $X/mo at current
	// usage").
	UpdatedSummary OrganizationBillingSummary `json:"updated_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency        respjson.Field
		MonthlyTotal    respjson.Field
		PricedAt        respjson.Field
		UsageDimensions respjson.Field
		CurrentSummary  respjson.Field
		UpdatedSummary  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuote) RawJSON() string { return r.JSON.raw }
func (r *CostQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Priced row for a single usage dimension emitted by a resource.
type CostQuoteUsageDimension struct {
	// Usage dimension being priced (e.g. compute_vcpu, storage_abs_gb).
	Dimension string `json:"dimension" api:"required"`
	// User-facing label for the dimension (e.g. "vCPU (hours)").
	DimensionDisplayName string `json:"dimension_display_name" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmount string `json:"monthly_amount" api:"required" format:"decimal"`
	// Quantity of the dimension being priced.
	Quantity int64 `json:"quantity" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	UnitPrice string `json:"unit_price" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension            respjson.Field
		DimensionDisplayName respjson.Field
		MonthlyAmount        respjson.Field
		Quantity             respjson.Field
		UnitPrice            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUsageDimension) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUsageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost quote returned by PATCH /:id/cost: current-state quote, post-update quote,
// and signed diff. current_summary and updated_summary omitted when the caller
// cannot view billing.
type CostQuoteUpdate struct {
	// Quote for the proposed (post-update) resource state.
	After CostQuoteUpdateAfter `json:"after" api:"required"`
	// Quote for the proposed (post-update) resource state.
	Before CostQuoteUpdateBefore `json:"before" api:"required"`
	// Currency the quote is denominated in. Always "USD" in v1.
	Currency string `json:"currency" api:"required"`
	// Per-dimension and total deltas: after minus before.
	Diff CostQuoteUpdateDiff `json:"diff" api:"required"`
	// Timestamp the quote was priced at.
	PricedAt time.Time `json:"priced_at" api:"required" format:"date-time"`
	// Forward-looking billing summary for an organization. All costs are run-rate
	// projections from the organization's current active usage ("≈ $X/mo at current
	// usage").
	CurrentSummary OrganizationBillingSummary `json:"current_summary"`
	// Forward-looking billing summary for an organization. All costs are run-rate
	// projections from the organization's current active usage ("≈ $X/mo at current
	// usage").
	UpdatedSummary OrganizationBillingSummary `json:"updated_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After          respjson.Field
		Before         respjson.Field
		Currency       respjson.Field
		Diff           respjson.Field
		PricedAt       respjson.Field
		CurrentSummary respjson.Field
		UpdatedSummary respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdate) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote for the proposed (post-update) resource state.
type CostQuoteUpdateAfter struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyTotal string `json:"monthly_total" api:"required" format:"decimal"`
	// Priced rows, one per usage dimension emitted by the resource.
	UsageDimensions []CostQuoteUpdateAfterUsageDimension `json:"usage_dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyTotal    respjson.Field
		UsageDimensions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateAfter) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Priced row for a single usage dimension emitted by a resource.
type CostQuoteUpdateAfterUsageDimension struct {
	// Usage dimension being priced (e.g. compute_vcpu, storage_abs_gb).
	Dimension string `json:"dimension" api:"required"`
	// User-facing label for the dimension (e.g. "vCPU (hours)").
	DimensionDisplayName string `json:"dimension_display_name" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmount string `json:"monthly_amount" api:"required" format:"decimal"`
	// Quantity of the dimension being priced.
	Quantity int64 `json:"quantity" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	UnitPrice string `json:"unit_price" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension            respjson.Field
		DimensionDisplayName respjson.Field
		MonthlyAmount        respjson.Field
		Quantity             respjson.Field
		UnitPrice            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateAfterUsageDimension) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateAfterUsageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote for the proposed (post-update) resource state.
type CostQuoteUpdateBefore struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyTotal string `json:"monthly_total" api:"required" format:"decimal"`
	// Priced rows, one per usage dimension emitted by the resource.
	UsageDimensions []CostQuoteUpdateBeforeUsageDimension `json:"usage_dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyTotal    respjson.Field
		UsageDimensions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateBefore) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Priced row for a single usage dimension emitted by a resource.
type CostQuoteUpdateBeforeUsageDimension struct {
	// Usage dimension being priced (e.g. compute_vcpu, storage_abs_gb).
	Dimension string `json:"dimension" api:"required"`
	// User-facing label for the dimension (e.g. "vCPU (hours)").
	DimensionDisplayName string `json:"dimension_display_name" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmount string `json:"monthly_amount" api:"required" format:"decimal"`
	// Quantity of the dimension being priced.
	Quantity int64 `json:"quantity" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	UnitPrice string `json:"unit_price" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension            respjson.Field
		DimensionDisplayName respjson.Field
		MonthlyAmount        respjson.Field
		Quantity             respjson.Field
		UnitPrice            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateBeforeUsageDimension) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateBeforeUsageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-dimension and total deltas: after minus before.
type CostQuoteUpdateDiff struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyTotalDelta string `json:"monthly_total_delta" api:"required" format:"decimal"`
	// Per-dimension diff entries. Includes every dimension touched by the update.
	UsageDimensions []CostQuoteUpdateDiffUsageDimension `json:"usage_dimensions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyTotalDelta respjson.Field
		UsageDimensions   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateDiff) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateDiff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-dimension diff entry. Both before and after are always present.
type CostQuoteUpdateDiffUsageDimension struct {
	// Priced row after the update. Always present.
	After CostQuoteUpdateDiffUsageDimensionAfter `json:"after" api:"required"`
	// Priced row after the update. Always present.
	Before CostQuoteUpdateDiffUsageDimensionBefore `json:"before" api:"required"`
	// Usage dimension being priced (e.g. compute_vcpu, storage_abs_gb).
	Dimension string `json:"dimension" api:"required"`
	// User-facing label for the dimension (e.g. "vCPU (hours)").
	DimensionDisplayName string `json:"dimension_display_name" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmountDelta string `json:"monthly_amount_delta" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After                respjson.Field
		Before               respjson.Field
		Dimension            respjson.Field
		DimensionDisplayName respjson.Field
		MonthlyAmountDelta   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateDiffUsageDimension) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateDiffUsageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Priced row after the update. Always present.
type CostQuoteUpdateDiffUsageDimensionAfter struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmount string `json:"monthly_amount" api:"required" format:"decimal"`
	Quantity      int64  `json:"quantity" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	UnitPrice string `json:"unit_price" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyAmount respjson.Field
		Quantity      respjson.Field
		UnitPrice     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateDiffUsageDimensionAfter) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateDiffUsageDimensionAfter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Priced row after the update. Always present.
type CostQuoteUpdateDiffUsageDimensionBefore struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyAmount string `json:"monthly_amount" api:"required" format:"decimal"`
	Quantity      int64  `json:"quantity" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	UnitPrice string `json:"unit_price" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonthlyAmount respjson.Field
		Quantity      respjson.Field
		UnitPrice     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CostQuoteUpdateDiffUsageDimensionBefore) RawJSON() string { return r.JSON.raw }
func (r *CostQuoteUpdateDiffUsageDimensionBefore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Forward-looking billing summary for an organization. All costs are run-rate
// projections from the organization's current active usage ("≈ $X/mo at current
// usage").
type OrganizationBillingSummary struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	DailyCost string `json:"daily_cost" api:"required" format:"decimal"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	EffectiveBalance string `json:"effective_balance" api:"required" format:"decimal"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	MonthlyCost string `json:"monthly_cost" api:"required" format:"decimal"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	RechargeThresholdFraction string `json:"recharge_threshold_fraction" api:"required" format:"decimal"`
	// Projected date the balance reaches the recharge threshold at the current
	// run-rate. Null when there is no active usage (never charges).
	EstimatedNextChargeAt time.Time `json:"estimated_next_charge_at" api:"nullable" format:"date-time"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	RunwayMonths string `json:"runway_months" api:"nullable" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DailyCost                 respjson.Field
		EffectiveBalance          respjson.Field
		MonthlyCost               respjson.Field
		RechargeThresholdFraction respjson.Field
		EstimatedNextChargeAt     respjson.Field
		RunwayMonths              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationBillingSummary) RawJSON() string { return r.JSON.raw }
func (r *OrganizationBillingSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination response details.
type Pagination struct {
	NextCursor     string `json:"next_cursor" api:"required"`
	PreviousCursor string `json:"previous_cursor" api:"required"`
	TotalCount     int64  `json:"total_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor     respjson.Field
		PreviousCursor respjson.Field
		TotalCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region the resource is in.
type RegionName string

const (
	RegionNameUsSva2 RegionName = "us-sva-2"
)

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

// IP filter rules.
type SourceIPRuleParam struct {
	// List of IPv4 CIDR addresses to allow.
	Allowed []string `json:"allowed,omitzero"`
	// List of IPv4 CIDR addresses to deny.
	Blocked []string `json:"blocked,omitzero"`
	paramObj
}

func (r SourceIPRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceIPRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceIPRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IP filter rules.
type SourceIPRuleResponse struct {
	// List of IPv4 CIDR addresses to allow.
	Allowed []string `json:"allowed" api:"required"`
	// List of IPv4 CIDR addresses to deny.
	Blocked []string `json:"blocked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Blocked     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceIPRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceIPRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
