// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// BridgeService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeService] method instead.
type BridgeService struct {
	Options []option.RequestOption
	Quotes  BridgeQuoteService
}

// NewBridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBridgeService(opts ...option.RequestOption) (r BridgeService) {
	r = BridgeService{}
	r.Options = opts
	r.Quotes = NewBridgeQuoteService(opts...)
	return
}
