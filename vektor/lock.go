// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// LockService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLockService] method instead.
type LockService struct {
	Options   []option.RequestOption
	Markets   LockMarketService
	Positions LockPositionService
}

// NewLockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLockService(opts ...option.RequestOption) (r LockService) {
	r = LockService{}
	r.Options = opts
	r.Markets = NewLockMarketService(opts...)
	r.Positions = NewLockPositionService(opts...)
	return
}
