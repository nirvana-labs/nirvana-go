// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/v3/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v3/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/shared"
)

// BucketCostService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketCostService] method instead.
type BucketCostService struct {
	Options []option.RequestOption
}

// NewBucketCostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketCostService(opts ...option.RequestOption) (r BucketCostService) {
	r = BucketCostService{}
	r.Options = opts
	return
}

// Return a priced monthly cost quote for the proposed Bucket storing size_gb of
// data. A Bucket is charged per GB it stores, prorated by the hour.
func (r *BucketCostService) New(ctx context.Context, body BucketCostNewParams, opts ...option.RequestOption) (res *shared.CostQuote, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/storage/buckets/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BucketCostNewParams struct {
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
	// Amount of data to price the Bucket at storing, in GB (1 GB = 10^9 bytes). A
	// Bucket is charged for what it stores, so a quote without it is empty.
	SizeGB param.Opt[int64] `json:"size_gb,omitzero"`
	// Tags to attach to the Bucket.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r BucketCostNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketCostNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketCostNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
