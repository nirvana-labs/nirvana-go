// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vms

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// OSImageService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOSImageService] method instead.
type OSImageService struct {
	Options []option.RequestOption
}

// NewOSImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOSImageService(opts ...option.RequestOption) (r *OSImageService) {
	r = &OSImageService{}
	r.Options = opts
	return
}

// List all OS Images
func (r *OSImageService) List(ctx context.Context, opts ...option.RequestOption) (res *[]OSImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "compute/vms/os_images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
