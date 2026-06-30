// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// BillingService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingService] method instead.
type BillingService struct {
	Options []option.RequestOption
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r BillingService) {
	r = BillingService{}
	r.Options = opts
	return
}

// Get the organization's billing summary: effective balance, monthly and daily
// run-rate cost, runway, and the projected next-recharge date. Costs are run-rate
// projections.
func (r *BillingService) Summary(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationBillingSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/summary", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
