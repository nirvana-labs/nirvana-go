// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// SecurityService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityService] method instead.
type SecurityService struct {
	Options []option.RequestOption
}

// NewSecurityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityService(opts ...option.RequestOption) (r SecurityService) {
	r = SecurityService{}
	r.Options = opts
	return
}

// Update the current user's security settings
func (r *SecurityService) Update(ctx context.Context, body SecurityUpdateParams, opts ...option.RequestOption) (res *UserSecurity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/user/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get the current user's security settings
func (r *SecurityService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserSecurity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/user/security"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// User security settings response.
type UserSecurity struct {
	// IP filter rules.
	SourceIPRule UserSecuritySourceIPRule `json:"source_ip_rule,required"`
	// When the user security settings were created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// When the user security settings were updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceIPRule respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserSecurity) RawJSON() string { return r.JSON.raw }
func (r *UserSecurity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IP filter rules.
type UserSecuritySourceIPRule struct {
	// List of IPv4 CIDR addresses to allow.
	Allowed []string `json:"allowed,required"`
	// List of IPv4 CIDR addresses to deny.
	Blocked []string `json:"blocked,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Blocked     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserSecuritySourceIPRule) RawJSON() string { return r.JSON.raw }
func (r *UserSecuritySourceIPRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityUpdateParams struct {
	// IP filter rules.
	SourceIPRule SecurityUpdateParamsSourceIPRule `json:"source_ip_rule,omitzero"`
	paramObj
}

func (r SecurityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IP filter rules.
type SecurityUpdateParamsSourceIPRule struct {
	// List of IPv4 CIDR addresses to allow.
	Allowed []string `json:"allowed,omitzero"`
	// List of IPv4 CIDR addresses to deny.
	Blocked []string `json:"blocked,omitzero"`
	paramObj
}

func (r SecurityUpdateParamsSourceIPRule) MarshalJSON() (data []byte, err error) {
	type shadow SecurityUpdateParamsSourceIPRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityUpdateParamsSourceIPRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
