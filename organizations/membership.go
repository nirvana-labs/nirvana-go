// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// MembershipService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
type MembershipService struct {
	Options []option.RequestOption
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r MembershipService) {
	r = MembershipService{}
	r.Options = opts
	return
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

type MembershipListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MembershipListParams]'s query parameters as `url.Values`.
func (r MembershipListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
