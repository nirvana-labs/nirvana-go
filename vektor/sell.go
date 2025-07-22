// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// SellService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSellService] method instead.
type SellService struct {
	Options []option.RequestOption
	Quotes  SellQuoteService
	Sell    SellSellService
}

// NewSellService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSellService(opts ...option.RequestOption) (r SellService) {
	r = SellService{}
	r.Options = opts
	r.Quotes = NewSellQuoteService(opts...)
	r.Sell = NewSellSellService(opts...)
	return
}
