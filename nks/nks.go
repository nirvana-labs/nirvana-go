// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// NKSService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNKSService] method instead.
type NKSService struct {
	Options  []option.RequestOption
	Clusters ClusterService
}

// NewNKSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNKSService(opts ...option.RequestOption) (r NKSService) {
	r = NKSService{}
	r.Options = opts
	r.Clusters = NewClusterService(opts...)
	return
}
