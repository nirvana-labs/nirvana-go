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
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// AddressService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r AddressService) {
	r = AddressService{}
	r.Options = opts
	return
}

// Create the address for an organization
func (r *AddressService) New(ctx context.Context, organizationID string, body AddressNewParams, opts ...option.RequestOption) (res *OrganizationAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/address", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update the address for an organization
func (r *AddressService) Update(ctx context.Context, organizationID string, body AddressUpdateParams, opts ...option.RequestOption) (res *OrganizationAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/address", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get the address for an organization
func (r *AddressService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationAddress, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/organizations/%s/address", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Organization address details.
type OrganizationAddress struct {
	// Address ID.
	ID string `json:"id" api:"required"`
	// City or locality.
	City string `json:"city" api:"required"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"required"`
	// When the address was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// First line of the street address.
	Line1 string `json:"line1" api:"required"`
	// Second line of the street address. Null when not provided.
	Line2 string `json:"line2" api:"required"`
	// Organization ID the address belongs to.
	OrganizationID string `json:"organization_id" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State, province, or region. Null when not provided.
	State string `json:"state" api:"required"`
	// Tax identification number. Null when not provided.
	TaxID string `json:"tax_id" api:"required"`
	// Type of the tax identification number. Null when not provided.
	TaxIDType string `json:"tax_id_type" api:"required"`
	// When the address was updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		City           respjson.Field
		Country        respjson.Field
		CreatedAt      respjson.Field
		Line1          respjson.Field
		Line2          respjson.Field
		OrganizationID respjson.Field
		PostalCode     respjson.Field
		State          respjson.Field
		TaxID          respjson.Field
		TaxIDType      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationAddress) RawJSON() string { return r.JSON.raw }
func (r *OrganizationAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressNewParams struct {
	// City or locality.
	City string `json:"city" api:"required"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country string `json:"country" api:"required"`
	// First line of the street address.
	Line1 string `json:"line1" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// Second line of the street address (suite, unit, building).
	Line2 param.Opt[string] `json:"line2,omitzero"`
	// State, province, or region. Required by some tax jurisdictions (e.g. US, CA).
	State param.Opt[string] `json:"state,omitzero"`
	// Tax identification number (e.g. VAT, EIN, ABN). Optional.
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Type of the tax identification number (e.g. eu_vat, us_ein, gb_vat, au_abn).
	// Optional.
	TaxIDType param.Opt[string] `json:"tax_id_type,omitzero"`
	paramObj
}

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressUpdateParams struct {
	// Second line of the street address (suite, unit, building). Omit to leave
	// unchanged, send null to clear, or send a value to set it.
	Line2 param.Opt[string] `json:"line2,omitzero"`
	// State, province, or region. Omit to leave unchanged, send null to clear, or send
	// a value to set it.
	State param.Opt[string] `json:"state,omitzero"`
	// Tax identification number (e.g. VAT, EIN, ABN). Omit to leave unchanged, send
	// null to clear, or send a value to set it.
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Type of the tax identification number (e.g. eu_vat, us_ein, gb_vat, au_abn).
	// Omit to leave unchanged, send null to clear, or send a value to set it.
	TaxIDType param.Opt[string] `json:"tax_id_type,omitzero"`
	// City or locality.
	City param.Opt[string] `json:"city,omitzero"`
	// Two-letter ISO 3166-1 alpha-2 country code.
	Country param.Opt[string] `json:"country,omitzero"`
	// First line of the street address.
	Line1 param.Opt[string] `json:"line1,omitzero"`
	// Postal or ZIP code.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	paramObj
}

func (r AddressUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
