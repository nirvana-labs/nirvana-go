// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"github.com/nirvana-labs/nirvana-go/v3/option"
)

// StorageService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options []option.RequestOption
	Buckets BucketService
	Keys    KeyService
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	r.Buckets = NewBucketService(opts...)
	r.Keys = NewKeyService(opts...)
	return
}
