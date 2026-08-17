// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// VMMetricService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMMetricService] method instead.
type VMMetricService struct {
	Options []option.RequestOption
}

// NewVMMetricService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVMMetricService(opts ...option.RequestOption) (r VMMetricService) {
	r = VMMetricService{}
	r.Options = opts
	return
}

// Read a VM's resource metrics over an interval. Every series covers the same
// periods, so they line up index for index, and a period the VM reported no
// observation for carries a null value.
func (r *VMMetricService) List(ctx context.Context, vmID string, query VMMetricListParams, opts ...option.RequestOption) (res *VMMetrics, err error) {
	opts = slices.Concat(r.Options, opts)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/compute/vms/%s/metrics", vmID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// One period's value.
type VMMetricPoint struct {
	// End of the period the value covers, so a point timestamped 00:05 over
	// five-minute periods describes 00:00 through 00:05.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Value over the period, in the series' unit. Null means the VM reported no
	// observation for this period, which is what a stopped VM looks like.
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMMetricPoint) RawJSON() string { return r.JSON.raw }
func (r *VMMetricPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One metric's values over the interval served.
type VMMetricSeries struct {
	// Fully-qualified name of the metric.
	Metric string `json:"metric" api:"required"`
	// Values over the interval, oldest first. Every series in a response covers the
	// same periods, so they line up index for index.
	Points []VMMetricPoint `json:"points" api:"required"`
	// Unit the values are expressed in.
	//
	// Any of "ratio", "bytes", "cores".
	Unit shared.VMMetricUnit `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Points      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMMetricSeries) RawJSON() string { return r.JSON.raw }
func (r *VMMetricSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A VM's metrics over an interval: one series per metric, on a shared grid of
// periods.
type VMMetrics struct {
	// How the samples inside one period were folded into a single value.
	//
	// Any of "mean", "max", "min".
	Aggregation VMMetricsAggregation `json:"aggregation" api:"required"`
	// End of the interval served, exclusive.
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// One series per requested metric, in the order they were requested.
	Metrics []VMMetricSeries `json:"metrics" api:"required"`
	// Width of one period, and so the spacing between consecutive points.
	//
	// Any of "5m", "15m", "1h", "6h", "24h".
	Period VMMetricsPeriod `json:"period" api:"required"`
	// Start of the interval served, inclusive. It can be later than the requested
	// start when the request asked for more than the available history.
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// ID of the VM the series belong to.
	VMID string `json:"vm_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregation respjson.Field
		EndTime     respjson.Field
		Metrics     respjson.Field
		Period      respjson.Field
		StartTime   respjson.Field
		VMID        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMMetrics) RawJSON() string { return r.JSON.raw }
func (r *VMMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the samples inside one period were folded into a single value.
type VMMetricsAggregation string

const (
	VMMetricsAggregationMean VMMetricsAggregation = "mean"
	VMMetricsAggregationMax  VMMetricsAggregation = "max"
	VMMetricsAggregationMin  VMMetricsAggregation = "min"
)

// Width of one period, and so the spacing between consecutive points.
type VMMetricsPeriod string

const (
	VMMetricsPeriod5m  VMMetricsPeriod = "5m"
	VMMetricsPeriod15m VMMetricsPeriod = "15m"
	VMMetricsPeriod1h  VMMetricsPeriod = "1h"
	VMMetricsPeriod6h  VMMetricsPeriod = "6h"
	VMMetricsPeriod24h VMMetricsPeriod = "24h"
)

type VMMetricListParams struct {
	// End of the interval, exclusive, as an RFC 3339 timestamp. Defaults to now.
	EndTime param.Opt[time.Time] `query:"end_time,omitzero" format:"date-time" json:"-"`
	// Start of the interval, inclusive, as an RFC 3339 timestamp. Defaults to an hour
	// before end_time. A start older than the 30 days of history kept is served from
	// where that history begins.
	StartTime param.Opt[time.Time] `query:"start_time,omitzero" format:"date-time" json:"-"`
	// How the samples inside one period are folded into a single value.
	//
	// Any of "mean", "max", "min".
	Aggregation VMMetricListParamsAggregation `query:"aggregation,omitzero" json:"-"`
	// Metric to return. Repeat the parameter for several; every metric is returned
	// when it is left out.
	//
	// Any of "compute.nirvanalabs.io/vm/cpu/used_cores",
	// "compute.nirvanalabs.io/vm/cpu/utilization",
	// "compute.nirvanalabs.io/vm/memory/used_bytes",
	// "compute.nirvanalabs.io/vm/memory/utilization".
	Metric []string `query:"metric,omitzero" json:"-"`
	// Width of one period, and so the spacing between points. An interval holding more
	// than 1440 periods is rejected; the error names a period that fits.
	//
	// Any of "5m", "15m", "1h", "6h", "24h".
	Period VMMetricListParamsPeriod `query:"period,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VMMetricListParams]'s query parameters as `url.Values`.
func (r VMMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// How the samples inside one period are folded into a single value.
type VMMetricListParamsAggregation string

const (
	VMMetricListParamsAggregationMean VMMetricListParamsAggregation = "mean"
	VMMetricListParamsAggregationMax  VMMetricListParamsAggregation = "max"
	VMMetricListParamsAggregationMin  VMMetricListParamsAggregation = "min"
)

// Width of one period, and so the spacing between points. An interval holding more
// than 1440 periods is rejected; the error names a period that fits.
type VMMetricListParamsPeriod string

const (
	VMMetricListParamsPeriod5m  VMMetricListParamsPeriod = "5m"
	VMMetricListParamsPeriod15m VMMetricListParamsPeriod = "15m"
	VMMetricListParamsPeriod1h  VMMetricListParamsPeriod = "1h"
	VMMetricListParamsPeriod6h  VMMetricListParamsPeriod = "6h"
	VMMetricListParamsPeriod24h VMMetricListParamsPeriod = "24h"
)
