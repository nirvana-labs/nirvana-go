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
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// OrganizationService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	return
}

// Create a new organization
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing organization
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List organizations
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.Cursor[Organization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/organizations"
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

// List organizations
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Organization] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get details about an Organization
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("v1/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Organization response.
type Organization struct {
	// Organization ID.
	ID string `json:"id,required"`
	// When the Organization was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Current user's membership details.
	Membership OrganizationMembership `json:"membership,required"`
	// Organization name.
	Name string `json:"name,required"`
	// Services that the Organization has access to.
	Services OrganizationServices `json:"services,required"`
	// When the Organization was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Authentication provider organization ID.
	AuthID string `json:"auth_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Membership  respjson.Field
		Name        respjson.Field
		Services    respjson.Field
		UpdatedAt   respjson.Field
		AuthID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Services that the Organization has access to.
type OrganizationServices struct {
	Cloud bool `json:"cloud"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cloud       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationServices) RawJSON() string { return r.JSON.raw }
func (r *OrganizationServices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationList struct {
	Items []Organization `json:"items,required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationList) RawJSON() string { return r.JSON.raw }
func (r *OrganizationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current user's membership details.
type OrganizationMembership struct {
	// Membership ID.
	ID string `json:"id,required"`
	// Role of the user in the organization.
	//
	// Any of "owner", "member".
	Role OrganizationMembershipRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	OrganizationMembershipRoleOwner  OrganizationMembershipRole = "owner"
	OrganizationMembershipRoleMember OrganizationMembershipRole = "member"
)

type OrganizationNewParams struct {
	// Organization name.
	Name string `json:"name,required"`
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateParams struct {
	// Organization name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
