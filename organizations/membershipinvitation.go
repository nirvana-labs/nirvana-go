// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// MembershipInvitationService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipInvitationService] method instead.
type MembershipInvitationService struct {
	Options []option.RequestOption
}

// NewMembershipInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMembershipInvitationService(opts ...option.RequestOption) (r MembershipInvitationService) {
	r = MembershipInvitationService{}
	r.Options = opts
	return
}

// Invite an email address to join an organization with a role. The membership is
// created when the invitee accepts. Owners may invite with any role; admins may
// invite with admin, billing or member.
func (r *MembershipInvitationService) New(ctx context.Context, organizationID string, body MembershipInvitationNewParams, opts ...option.RequestOption) (res *OrganizationInvitation, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/invitations", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the invitations to an organization that are still pending: not yet
// accepted, revoked or expired. Only owners and admins may list invitations.
func (r *MembershipInvitationService) List(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationInvitationList, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/invitations", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke a pending invitation so it can no longer be accepted. Owners may revoke
// any invitation; admins may revoke invitations with the admin, billing or member
// role.
func (r *MembershipInvitationService) Delete(ctx context.Context, organizationID string, invitationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return err
	}
	if invitationID == "" {
		err = errors.New("missing required invitation_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/invitations/%s", organizationID, invitationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Send a pending invitation's email again. Owners may resend any invitation;
// admins may resend invitations with the admin, billing or member role.
func (r *MembershipInvitationService) Resend(ctx context.Context, organizationID string, invitationID string, opts ...option.RequestOption) (res *OrganizationInvitation, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if invitationID == "" {
		err = errors.New("missing required invitation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/memberships/invitations/%s/resend", organizationID, invitationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Organization invitation details.
type OrganizationInvitation struct {
	// Invitation ID.
	ID string `json:"id" api:"required"`
	// When the invitation was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address the invitation was sent to.
	Email string `json:"email" api:"required"`
	// When the invitation expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// ID of the user who sent the invitation, or null when the sender is not known.
	InviterUserID string `json:"inviter_user_id" api:"required"`
	// Organization ID.
	OrganizationID string `json:"organization_id" api:"required"`
	// Role of the user in the organization.
	//
	// Any of "owner", "admin", "billing", "member".
	Role OrganizationInvitationRole `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Email          respjson.Field
		ExpiresAt      respjson.Field
		InviterUserID  respjson.Field
		OrganizationID respjson.Field
		Role           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationInvitation) RawJSON() string { return r.JSON.raw }
func (r *OrganizationInvitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role of the user in the organization.
type OrganizationInvitationRole string

const (
	OrganizationInvitationRoleOwner   OrganizationInvitationRole = "owner"
	OrganizationInvitationRoleAdmin   OrganizationInvitationRole = "admin"
	OrganizationInvitationRoleBilling OrganizationInvitationRole = "billing"
	OrganizationInvitationRoleMember  OrganizationInvitationRole = "member"
)

type OrganizationInvitationList struct {
	Items []OrganizationInvitation `json:"items" api:"required"`
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
func (r OrganizationInvitationList) RawJSON() string { return r.JSON.raw }
func (r *OrganizationInvitationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipInvitationNewParams struct {
	// Email address to invite. Inviting an address that already belongs to a member of
	// the organization fails.
	Email string `json:"email" api:"required"`
	// Role of the user in the organization.
	//
	// Any of "owner", "admin", "billing", "member".
	Role MembershipInvitationNewParamsRole `json:"role,omitzero" api:"required"`
	paramObj
}

func (r MembershipInvitationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MembershipInvitationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MembershipInvitationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role of the user in the organization.
type MembershipInvitationNewParamsRole string

const (
	MembershipInvitationNewParamsRoleOwner   MembershipInvitationNewParamsRole = "owner"
	MembershipInvitationNewParamsRoleAdmin   MembershipInvitationNewParamsRole = "admin"
	MembershipInvitationNewParamsRoleBilling MembershipInvitationNewParamsRole = "billing"
	MembershipInvitationNewParamsRoleMember  MembershipInvitationNewParamsRole = "member"
)
