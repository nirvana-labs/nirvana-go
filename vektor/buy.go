// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// BuyService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuyService] method instead.
type BuyService struct {
	Options []option.RequestOption
	Quotes  BuyQuoteService
	Buy     BuyBuyService
}

// NewBuyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuyService(opts ...option.RequestOption) (r BuyService) {
	r = BuyService{}
	r.Options = opts
	r.Quotes = NewBuyQuoteService(opts...)
	r.Buy = NewBuyBuyService(opts...)
	return
}
