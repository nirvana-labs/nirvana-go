// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/v2/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v2/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/packages/param"
	"github.com/nirvana-labs/nirvana-go/v2/shared"
)

// BucketAvailabilityService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketAvailabilityService] method instead.
type BucketAvailabilityService struct {
	Options []option.RequestOption
}

// NewBucketAvailabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketAvailabilityService(opts ...option.RequestOption) (r BucketAvailabilityService) {
	r = BucketAvailabilityService{}
	r.Options = opts
	return
}

// Check if a Bucket can be created
func (r *BucketAvailabilityService) New(ctx context.Context, body BucketAvailabilityNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/storage/buckets/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type BucketAvailabilityNewParams struct {
	// Name of the Bucket, used to address it over S3. It must be 3 to 63 characters of
	// lowercase letters, digits, dots and hyphens, start and end with a letter or
	// digit, not contain two adjacent dots and not be formatted as an IP address.
	// "admin" and "ops-api" are reserved. The name cannot be changed once the Bucket
	// is created.
	Name string `json:"name" api:"required"`
	// Project ID the Bucket belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Tags to attach to the Bucket.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BucketAvailabilityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketAvailabilityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketAvailabilityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
