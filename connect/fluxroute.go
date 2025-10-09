// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connect

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// FluxRouteService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFluxRouteService] method instead.
type FluxRouteService struct {
	Options []option.RequestOption
}

// NewFluxRouteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFluxRouteService(opts ...option.RequestOption) (r FluxRouteService) {
	r = FluxRouteService{}
	r.Options = opts
	return
}

// List all supported routes with regions for Connect Flux.
func (r *FluxRouteService) List(ctx context.Context, opts ...option.RequestOption) (res *FluxRouteList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/connect/flux/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
