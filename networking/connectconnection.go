// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ConnectConnectionService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectConnectionService] method instead.
type ConnectConnectionService struct {
	Options []option.RequestOption
}

// NewConnectConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectConnectionService(opts ...option.RequestOption) (r ConnectConnectionService) {
	r = ConnectConnectionService{}
	r.Options = opts
	return
}

// Create a Connect Connection
func (r *ConnectConnectionService) New(ctx context.Context, body ConnectConnectionNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/networking/connect/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update Connect Connection details
func (r *ConnectConnectionService) Update(ctx context.Context, connectionID string, body ConnectConnectionUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/connect/connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all Connect Connections
func (r *ConnectConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectConnectionList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/networking/connect/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Connect Connection
func (r *ConnectConnectionService) Delete(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/connect/connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Connect Connection details
func (r *ConnectConnectionService) Get(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *ConnectConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/connect/connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ConnectConnectionNewParams struct {
	// Connect Connection speed in Mbps
	//
	// Any of 50, 200, 500, 1000, 2000.
	BandwidthMbps int64 `json:"bandwidth_mbps,omitzero,required"`
	// CIDRs for the Connect Connection
	CIDRs []string `json:"cidrs,omitzero,required"`
	// Name of the Connect Connection
	Name string `json:"name,required"`
	// Provider CIDRs
	ProviderCIDRs []string `json:"provider_cidrs,omitzero,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,omitzero,required"`
	// AWS provider configuration
	AWS ConnectConnectionAWSConfigRequestParam `json:"aws,omitzero"`
	// Tags to attach to the Connect Connection
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ConnectConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectConnectionUpdateParams struct {
	// Name of the Connect Connection.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Connect Connection
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ConnectConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
