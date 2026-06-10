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
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ConnectConnectionCostService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectConnectionCostService] method instead.
type ConnectConnectionCostService struct {
	Options []option.RequestOption
}

// NewConnectConnectionCostService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectConnectionCostService(opts ...option.RequestOption) (r ConnectConnectionCostService) {
	r = ConnectConnectionCostService{}
	r.Options = opts
	return
}

// Return a priced cost quote for the proposed Connect Connection.
func (r *ConnectConnectionCostService) New(ctx context.Context, body ConnectConnectionCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/networking/connect/connections/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a priced cost quote for the proposed Connect Connection update plus a
// diff against the current state.
func (r *ConnectConnectionCostService) Update(ctx context.Context, connectionID string, body ConnectConnectionCostUpdateParams, opts ...option.RequestOption) (res *shared.CostQuoteUpdate, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/networking/connect/connections/%s/cost", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ConnectConnectionCostNewParams struct {
	// Connect Connection speed in Mbps
	//
	// Any of 50, 200, 500, 1000, 2000.
	BandwidthMbps int64 `json:"bandwidth_mbps,omitzero" api:"required"`
	// CIDRs for the Connect Connection. Must be in network-aligned/canonical form.
	CIDRs []string `json:"cidrs,omitzero" api:"required"`
	// Name of the Connect Connection
	Name string `json:"name" api:"required"`
	// Project ID the Connect Connection belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Provider CIDRs. Must be in network-aligned/canonical form.
	ProviderCIDRs []string `json:"provider_cidrs,omitzero" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// AWS provider configuration
	AWS ConnectConnectionAWSConfigRequestParam `json:"aws,omitzero"`
	// Tags to attach to the Connect Connection
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ConnectConnectionCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectConnectionCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectConnectionCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectConnectionCostUpdateParams struct {
	// Name of the Connect Connection.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Connect Connection
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ConnectConnectionCostUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectConnectionCostUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectConnectionCostUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
