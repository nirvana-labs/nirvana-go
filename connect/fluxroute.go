// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connect

import (
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
