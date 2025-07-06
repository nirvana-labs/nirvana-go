// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// DedicatedService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDedicatedService] method instead.
type DedicatedService struct {
	Options     []option.RequestOption
	Blockchains DedicatedBlockchainService
}

// NewDedicatedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDedicatedService(opts ...option.RequestOption) (r DedicatedService) {
	r = DedicatedService{}
	r.Options = opts
	r.Blockchains = NewDedicatedBlockchainService(opts...)
	return
}
