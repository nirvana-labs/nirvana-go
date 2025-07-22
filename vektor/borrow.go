// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// BorrowService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBorrowService] method instead.
type BorrowService struct {
	Options   []option.RequestOption
	Markets   BorrowMarketService
	Positions BorrowPositionService
	Accounts  BorrowAccountService
}

// NewBorrowService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBorrowService(opts ...option.RequestOption) (r BorrowService) {
	r = BorrowService{}
	r.Options = opts
	r.Markets = NewBorrowMarketService(opts...)
	r.Positions = NewBorrowPositionService(opts...)
	r.Accounts = NewBorrowAccountService(opts...)
	return
}
