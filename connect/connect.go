// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connect

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// ConnectService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectService] method instead.
type ConnectService struct {
	Options []option.RequestOption
	Flux    FluxService
}

// NewConnectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConnectService(opts ...option.RequestOption) (r ConnectService) {
	r = ConnectService{}
	r.Options = opts
	r.Flux = NewFluxService(opts...)
	return
}
