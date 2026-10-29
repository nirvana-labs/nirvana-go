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

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// MembershipService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
type MembershipService struct {
	Options     []option.RequestOption
	Invitations MembershipInvitationService
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r MembershipService) {
	r = MembershipService{}
	r.Options = opts
	r.Invitations = NewMembershipInvitationService(opts...)
	return
}

// Change the role of a member of an organization. Owners may assign any role.
// Admins may assign admin, billing or member, and cannot change the role of an
// owner. The last owner cannot be demoted, and members cannot change their own
// role except for an owner stepping down while another owner remains.
func (r *MembershipService) Update(ctx context.Context, organizationID string, membershipID string, body MembershipUpdateParams, opts ...option.RequestOption) (res *OrganizationMembership, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/%s", organizationID, membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all memberships for an organization
func (r *MembershipService) List(ctx context.Context, organizationID string, query MembershipListParams, opts ...option.RequestOption) (res *pagination.Cursor[OrganizationMembership], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships", organizationID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all memberships for an organization
func (r *MembershipService) ListAutoPaging(ctx context.Context, organizationID string, query MembershipListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[OrganizationMembership] {
	return pagination.NewCursorAutoPager(r.List(ctx, organizationID, query, opts...))
}

// Remove a member from an organization. Admins cannot remove an owner, and the
// last owner cannot be removed.
func (r *MembershipService) Delete(ctx context.Context, organizationID string, membershipID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/%s", organizationID, membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get details about an organization membership
func (r *MembershipService) Get(ctx context.Context, organizationID string, membershipID string, opts ...option.RequestOption) (res *OrganizationMembership, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/%s", organizationID, membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the authenticated user's own membership in an organization, including their
// role. Available to every member of the organization, whatever their role.
func (r *MembershipService) Me(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationMembership, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/me", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Organization membership details.
type OrganizationMembership struct {
	// Membership ID.
	ID string `json:"id" api:"required"`
	// When the membership was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address of the user.
	Email string `json:"email" api:"required"`
	// First name of the user.
	FirstName string `json:"first_name" api:"required"`
	// Last name of the user.
	LastName string `json:"last_name" api:"required"`
	// Organization ID.
	OrganizationID string `json:"organization_id" api:"required"`
	// Role of the user in the organization.
	//
	// Any of "owner", "admin", "billing", "member".
	Role OrganizationMembershipRole `json:"role" api:"required"`
	// When the membership was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// User ID.
	UserID string `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Email          respjson.Field
		FirstName      respjson.Field
		LastName       respjson.Field
		OrganizationID respjson.Field
		Role           respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationMembership) RawJSON() string { return r.JSON.raw }
func (r *OrganizationMembership) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role of the user in the organization.
type OrganizationMembershipRole string

const (
	OrganizationMembershipRoleOwner   OrganizationMembershipRole = "owner"
	OrganizationMembershipRoleAdmin   OrganizationMembershipRole = "admin"
	OrganizationMembershipRoleBilling OrganizationMembershipRole = "billing"
	OrganizationMembershipRoleMember  OrganizationMembershipRole = "member"
)

type OrganizationMembershipList struct {
	Items []OrganizationMembership `json:"items" api:"required"`
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
func (r OrganizationMembershipList) RawJSON() string { return r.JSON.raw }
func (r *OrganizationMembershipList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUpdateParams struct {
	// Role of the user in the organization.
	//
	// Any of "owner", "admin", "billing", "member".
	Role MembershipUpdateParamsRole `json:"role,omitzero" api:"required"`
	paramObj
}

func (r MembershipUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MembershipUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MembershipUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role of the user in the organization.
type MembershipUpdateParamsRole string

const (
	MembershipUpdateParamsRoleOwner   MembershipUpdateParamsRole = "owner"
	MembershipUpdateParamsRoleAdmin   MembershipUpdateParamsRole = "admin"
	MembershipUpdateParamsRoleBilling MembershipUpdateParamsRole = "billing"
	MembershipUpdateParamsRoleMember  MembershipUpdateParamsRole = "member"
)

type MembershipListParams struct {
	// Pagination cursor returned by a previous request. Only valid for the same
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at, role
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by the member's user ID
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Filter by membership role
	//
	// Any of "owner", "admin", "billing", "member".
	Role MembershipListParamsRole `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MembershipListParams]'s query parameters as `url.Values`.
func (r MembershipListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by membership role
type MembershipListParamsRole string

const (
	MembershipListParamsRoleOwner   MembershipListParamsRole = "owner"
	MembershipListParamsRoleAdmin   MembershipListParamsRole = "admin"
	MembershipListParamsRoleBilling MembershipListParamsRole = "billing"
	MembershipListParamsRoleMember  MembershipListParamsRole = "member"
)
