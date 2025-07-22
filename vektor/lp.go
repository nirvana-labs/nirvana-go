// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// LPService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLPService] method instead.
type LPService struct {
	Options       []option.RequestOption
	Pools         LPPoolService
	Positions     LPPositionService
	DepositQuote  LPDepositQuoteService
	WithdrawQuote LPWithdrawQuoteService
}

// NewLPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLPService(opts ...option.RequestOption) (r LPService) {
	r = LPService{}
	r.Options = opts
	r.Pools = NewLPPoolService(opts...)
	r.Positions = NewLPPositionService(opts...)
	r.DepositQuote = NewLPDepositQuoteService(opts...)
	r.WithdrawQuote = NewLPWithdrawQuoteService(opts...)
	return
}
