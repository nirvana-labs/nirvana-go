// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// LendService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendService] method instead.
type LendService struct {
	Options       []option.RequestOption
	Markets       LendMarketService
	Positions     LendPositionService
	Lend          LendLendService
	Withdraw      LendWithdrawService
	SetCollateral LendSetCollateralService
}

// NewLendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLendService(opts ...option.RequestOption) (r LendService) {
	r = LendService{}
	r.Options = opts
	r.Markets = NewLendMarketService(opts...)
	r.Positions = NewLendPositionService(opts...)
	r.Lend = NewLendLendService(opts...)
	r.Withdraw = NewLendWithdrawService(opts...)
	r.SetCollateral = NewLendSetCollateralService(opts...)
	return
}
