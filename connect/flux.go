// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connect

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
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
	Options []option.RequestOption
}

// NewFluxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFluxService(opts ...option.RequestOption) (r FluxService) {
	r = FluxService{}
	r.Options = opts
	return
}

// List all Connect Flux
func (r *FluxService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectFluxList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/connect/flux"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about a Connect Flux
func (r *FluxService) Get(ctx context.Context, fluxID string, opts ...option.RequestOption) (res *ConnectFlux, err error) {
	opts = append(r.Options[:], opts...)
	if fluxID == "" {
		err = errors.New("missing required flux_id parameter")
		return
	}
	path := fmt.Sprintf("v1/connect/flux/%s", fluxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Connect flux details.
type ConnectFlux struct {
	// Unique identifier for the connect flux
	ID string `json:"id,required"`
	// Connect flux speed in Mbps
	//
	// Any of 50, 100, 200, 500, 1000, 2000, 5000.
	BandwidthMbps int64 `json:"bandwidth_mbps,required"`
	// CIDRs
	Cidrs []string `json:"cidrs,required"`
	// When the connect flux was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the connect flux
	Name string `json:"name,required"`
	// Provider CIDRs
	ProviderCidrs []string `json:"provider_cidrs,required"`
	// Provider name
	ProviderName string `json:"provider_name,required"`
	// Provider region
	ProviderRegion string `json:"provider_region,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-lon-1", "eu-ams-1",
	// "eu-frk-1", "ap-sin-1", "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// When the connect flux was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		BandwidthMbps  respjson.Field
		Cidrs          respjson.Field
		CreatedAt      respjson.Field
		Name           respjson.Field
		ProviderCidrs  respjson.Field
		ProviderName   respjson.Field
		ProviderRegion respjson.Field
		Region         respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectFlux) RawJSON() string { return r.JSON.raw }
func (r *ConnectFlux) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectFluxList struct {
	Items []ConnectFlux `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectFluxList) RawJSON() string { return r.JSON.raw }
func (r *ConnectFluxList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
