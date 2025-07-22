// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// WrapService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWrapService] method instead.
type WrapService struct {
	Options []option.RequestOption
	Wrap    WrapWrapService
	Unwrap  WrapUnwrapService
}

// NewWrapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWrapService(opts ...option.RequestOption) (r WrapService) {
	r = WrapService{}
	r.Options = opts
	r.Wrap = NewWrapWrapService(opts...)
	r.Unwrap = NewWrapUnwrapService(opts...)
	return
}
