// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connect

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// FluxService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFluxService] method instead.
type FluxService struct {
	Options   []option.RequestOption
	Providers FluxProviderService
}

// NewFluxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFluxService(opts ...option.RequestOption) (r FluxService) {
	r = FluxService{}
	r.Options = opts
	r.Providers = NewFluxProviderService(opts...)
	return
}

// Create a Connect Flux
func (r *FluxService) New(ctx context.Context, body FluxNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/flux"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update Connect Flux details
func (r *FluxService) Update(ctx context.Context, fluxID string, body FluxUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if fluxID == "" {
		err = errors.New("missing required flux_id parameter")
		return
	}
	path := fmt.Sprintf("v1/connect/flux/%s", fluxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Connect Flux
func (r *FluxService) List(ctx context.Context, opts ...option.RequestOption) (res *FluxList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/flux"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Connect Flux
func (r *FluxService) Delete(ctx context.Context, fluxID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if fluxID == "" {
		err = errors.New("missing required flux_id parameter")
		return
	}
	path := fmt.Sprintf("v1/connect/flux/%s", fluxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Connect Flux details
func (r *FluxService) Get(ctx context.Context, fluxID string, opts ...option.RequestOption) (res *Flux, err error) {
	opts = slices.Concat(r.Options, opts)
	if fluxID == "" {
		err = errors.New("missing required flux_id parameter")
		return
	}
	path := fmt.Sprintf("v1/connect/flux/%s", fluxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Connect flux details.
type Flux struct {
	// Unique identifier for the Connect Flux
	ID string `json:"id,required"`
	// ASN
	ASN int64 `json:"asn,required"`
	// AWS provider configuration
	AWS FluxProviderAWSConfig `json:"aws,required"`
	// Connect Flux speed in Mbps
	//
	// Any of 50, 200, 500, 1000, 2000.
	BandwidthMbps int64 `json:"bandwidth_mbps,required"`
	// CIDRs for the Connect Flux
	CIDRs []string `json:"cidrs,required"`
	// When the Connect Flux was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the Connect Flux
	Name string `json:"name,required"`
	// Provider ASN
	ProviderASN int64 `json:"provider_asn,required"`
	// Provider CIDRs for the Connect Flux
	ProviderCIDRs []string `json:"provider_cidrs,required"`
	// Provider Router IP
	ProviderRouterIP string `json:"provider_router_ip,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Router IP
	RouterIP string `json:"router_ip,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// When the Connect Flux was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ASN              respjson.Field
		AWS              respjson.Field
		BandwidthMbps    respjson.Field
		CIDRs            respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		ProviderASN      respjson.Field
		ProviderCIDRs    respjson.Field
		ProviderRouterIP respjson.Field
		Region           respjson.Field
		RouterIP         respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Flux) RawJSON() string { return r.JSON.raw }
func (r *Flux) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FluxList struct {
	Items []Flux `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FluxList) RawJSON() string { return r.JSON.raw }
func (r *FluxList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider supported for Connect Flux.
type FluxProvider struct {
	// Provider name.
	Name string `json:"name,required"`
	// Provider region name.
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FluxProvider) RawJSON() string { return r.JSON.raw }
func (r *FluxProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS provider configuration
type FluxProviderAWSConfig struct {
	// AWS region where the connection is established
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FluxProviderAWSConfig) RawJSON() string { return r.JSON.raw }
func (r *FluxProviderAWSConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS provider configuration
//
// The properties AccountNumber, Region are required.
type FluxProviderAWSConfigRequestParam struct {
	// AWS account number
	AccountNumber string `json:"account_number,required"`
	// AWS region where the connection will be established
	Region string `json:"region,required"`
	paramObj
}

func (r FluxProviderAWSConfigRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FluxProviderAWSConfigRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FluxProviderAWSConfigRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FluxProviderList struct {
	Items []FluxProvider `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FluxProviderList) RawJSON() string { return r.JSON.raw }
func (r *FluxProviderList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FluxNewParams struct {
	// Connect Flux speed in Mbps
	//
	// Any of 50, 200, 500, 1000, 2000.
	BandwidthMbps int64 `json:"bandwidth_mbps,omitzero,required"`
	// CIDRs for the Connect Flux
	CIDRs []string `json:"cidrs,omitzero,required"`
	// Name of the Connect Flux
	Name string `json:"name,required"`
	// Provider CIDRs
	ProviderCIDRs []string `json:"provider_cidrs,omitzero,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// AWS provider configuration
	AWS FluxProviderAWSConfigRequestParam `json:"aws,omitzero"`
	paramObj
}

func (r FluxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FluxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FluxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FluxUpdateParams struct {
	// Name of the Connect Flux.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r FluxUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FluxUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FluxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
