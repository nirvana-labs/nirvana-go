// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
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
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
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

// Get the organization's total usage cost per UTC day over a date range (max 90
// days), summing open and closed resources. One entry per day, oldest first.
// Defaults to the last 30 days.
func (r *BillingService) Cost(ctx context.Context, organizationID string, query BillingCostParams, opts ...option.RequestOption) (res *OrganizationDailyCost, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/cost", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the organization's billing history: prepaid credits, top-ups, and manual
// adjustments, newest first. Paginated with an opaque cursor.
func (r *BillingService) History(ctx context.Context, organizationID string, query BillingHistoryParams, opts ...option.RequestOption) (res *BillingHistoryEntryList, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/history", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the organization's billing summary: effective balance, monthly and daily
// run-rate cost, runway, and the projected next-recharge date. Costs are run-rate
// projections.
func (r *BillingService) Summary(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *shared.OrganizationBillingSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/summary", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Charge the card on file and credit the prepaid balance. A unique Idempotency-Key
// header is required; reuse it across retries so a timed-out top-up is not charged
// twice.
func (r *BillingService) TopUp(ctx context.Context, organizationID string, params BillingTopUpParams, opts ...option.RequestOption) (res *shared.OrganizationBillingSummary, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/topup", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A single billing history line item: a prepaid credit or a manual adjustment.
type BillingHistoryEntry struct {
	// Unique identifier for the entry.
	ID string `json:"id" api:"required"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	Amount string `json:"amount" api:"required" format:"decimal"`
	// When the entry was recorded.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// Kind of entry.
	//
	// Any of "grant", "adjustment".
	Type BillingHistoryEntryType `json:"type" api:"required"`
	// Human-readable note describing the entry, when available.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Type        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingHistoryEntry) RawJSON() string { return r.JSON.raw }
func (r *BillingHistoryEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryEntryList struct {
	Items []BillingHistoryEntry `json:"items" api:"required"`
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
func (r BillingHistoryEntryList) RawJSON() string { return r.JSON.raw }
func (r *BillingHistoryEntryList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of entry.
type BillingHistoryEntryType string

const (
	BillingHistoryEntryTypeGrant      BillingHistoryEntryType = "grant"
	BillingHistoryEntryTypeAdjustment BillingHistoryEntryType = "adjustment"
)

// Total usage cost for a single UTC day.
type DailyCostPoint struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	Cost string `json:"cost" api:"required" format:"decimal"`
	// UTC calendar day (YYYY-MM-DD).
	Date string `json:"date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cost        respjson.Field
		Date        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DailyCostPoint) RawJSON() string { return r.JSON.raw }
func (r *DailyCostPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Daily usage cost over a date range: one entry per UTC day (zero on idle days),
// summing open and closed resources. Suitable for a daily cost bar chart.
type OrganizationDailyCost struct {
	// ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// One entry per UTC day in the range, oldest first.
	Days []DailyCostPoint `json:"days" api:"required"`
	// Inclusive start of the range, as a UTC calendar day (YYYY-MM-DD).
	From string `json:"from" api:"required"`
	// Inclusive end of the range, as a UTC calendar day (YYYY-MM-DD).
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Days        respjson.Field
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationDailyCost) RawJSON() string { return r.JSON.raw }
func (r *OrganizationDailyCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingCostParams struct {
	// Inclusive start day, YYYY-MM-DD (UTC). Defaults to 30 days before to.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// Inclusive end day, YYYY-MM-DD (UTC). Defaults to today.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [BillingCostParams]'s query parameters as `url.Values`.
func (r BillingCostParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillingHistoryParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BillingHistoryParams]'s query parameters as `url.Values`.
func (r BillingHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillingTopUpParams struct {
	// Amount to charge and credit, in USD. Must be greater than 0, at most two decimal
	// places, and at most 10000.
	Amount         string `json:"amount" api:"required" format:"decimal"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	paramObj
}

func (r BillingTopUpParams) MarshalJSON() (data []byte, err error) {
	type shadow BillingTopUpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingTopUpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
