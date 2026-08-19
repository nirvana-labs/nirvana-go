// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// VMMetricDescriptorService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMMetricDescriptorService] method instead.
type VMMetricDescriptorService struct {
	Options []option.RequestOption
}

// NewVMMetricDescriptorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVMMetricDescriptorService(opts ...option.RequestOption) (r VMMetricDescriptorService) {
	r = VMMetricDescriptorService{}
	r.Options = opts
	return
}

// Describe every metric a VM reports: its name, unit, the range its values fall
// in, and whether a value can be null. Read this instead of holding a list of
// metric names, so a metric published later is picked up without a client change.
func (r *VMMetricDescriptorService) List(ctx context.Context, opts ...option.RequestOption) (res *VMMetricDescriptorList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/compute/metric_descriptors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Describes one metric a VM reports: how to name it, what its values mean, and the
// range they fall in.
type VMMetricDescriptor struct {
	// What the metric measures.
	Description string `json:"description" api:"required"`
	// Highest value the metric reports, or null when it has no ceiling of its own.
	MaxValue float64 `json:"max_value" api:"required"`
	// Fully-qualified name of the metric, and the only spelling the metric query
	// parameter accepts.
	Metric string `json:"metric" api:"required"`
	// Lowest value the metric reports.
	MinValue float64 `json:"min_value" api:"required"`
	// Whether a point's value can be null. A null means the VM reported no observation
	// for that period, which is what a stopped VM looks like.
	Nullable bool `json:"nullable" api:"required"`
	// Unit the values are expressed in.
	//
	// Any of "ratio", "bytes", "cores", "bytes_per_second", "operations_per_second".
	Unit shared.VMMetricUnit `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		MaxValue    respjson.Field
		Metric      respjson.Field
		MinValue    respjson.Field
		Nullable    respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMMetricDescriptor) RawJSON() string { return r.JSON.raw }
func (r *VMMetricDescriptor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMMetricDescriptorList struct {
	Items []VMMetricDescriptor `json:"items" api:"required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMMetricDescriptorList) RawJSON() string { return r.JSON.raw }
func (r *VMMetricDescriptorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
