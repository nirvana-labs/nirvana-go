// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// BillingRechargePolicyService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingRechargePolicyService] method instead.
type BillingRechargePolicyService struct {
	Options []option.RequestOption
}

// NewBillingRechargePolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillingRechargePolicyService(opts ...option.RequestOption) (r BillingRechargePolicyService) {
	r = BillingRechargePolicyService{}
	r.Options = opts
	return
}

// Update the organization's recharge mode: manual for self-serve top-ups, or
// automatic to charge the card on file at the recharge threshold (fixed and
// proportional required).
func (r *BillingRechargePolicyService) Update(ctx context.Context, organizationID string, body BillingRechargePolicyUpdateParams, opts ...option.RequestOption) (res *OrganizationRechargePolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/recharge_policy", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get the organization's recharge configuration: the top-up mode and the fixed and
// proportional threshold components.
func (r *BillingRechargePolicyService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationRechargePolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/billing/recharge_policy", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// PolicyArgs carries the threshold parameters. Required when policy is
// "automatic"; must be omitted when policy is "manual".
type AutomaticPolicyArgs struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	Fixed string `json:"fixed" api:"required" format:"decimal"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	RunwayDays string `json:"runway_days" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fixed       respjson.Field
		RunwayDays  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomaticPolicyArgs) RawJSON() string { return r.JSON.raw }
func (r *AutomaticPolicyArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutomaticPolicyArgs to a AutomaticPolicyArgsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutomaticPolicyArgsParam.Overrides()
func (r AutomaticPolicyArgs) ToParam() AutomaticPolicyArgsParam {
	return param.Override[AutomaticPolicyArgsParam](json.RawMessage(r.RawJSON()))
}

// PolicyArgs carries the threshold parameters. Required when policy is
// "automatic"; must be omitted when policy is "manual".
//
// The properties Fixed, RunwayDays are required.
type AutomaticPolicyArgsParam struct {
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	Fixed string `json:"fixed" api:"required" format:"decimal"`
	// Arbitrary-precision decimal serialized as a string (e.g. "58.40").
	RunwayDays string `json:"runway_days" api:"required" format:"decimal"`
	paramObj
}

func (r AutomaticPolicyArgsParam) MarshalJSON() (data []byte, err error) {
	type shadow AutomaticPolicyArgsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutomaticPolicyArgsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An organization's current recharge policy. policy_args is null for a manual
// policy.
type OrganizationRechargePolicy struct {
	// Policy is the top-up mode.
	//
	// Any of "manual", "automatic".
	Policy RechargePolicyMode `json:"policy"`
	// PolicyArgs carries the threshold parameters. Required when policy is
	// "automatic"; must be omitted when policy is "manual".
	PolicyArgs AutomaticPolicyArgs `json:"policy_args" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Policy      respjson.Field
		PolicyArgs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationRechargePolicy) RawJSON() string { return r.JSON.raw }
func (r *OrganizationRechargePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Policy is the top-up mode.
type RechargePolicyMode string

const (
	RechargePolicyModeManual    RechargePolicyMode = "manual"
	RechargePolicyModeAutomatic RechargePolicyMode = "automatic"
)

type BillingRechargePolicyUpdateParams struct {
	// Policy is the top-up mode.
	//
	// Any of "manual", "automatic".
	Policy RechargePolicyMode `json:"policy,omitzero" api:"required"`
	// PolicyArgs carries the threshold parameters. Required when policy is
	// "automatic"; must be omitted when policy is "manual".
	PolicyArgs AutomaticPolicyArgsParam `json:"policy_args,omitzero"`
	paramObj
}

func (r BillingRechargePolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BillingRechargePolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingRechargePolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
