// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// ExecutionService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecutionService] method instead.
type ExecutionService struct {
	Options []option.RequestOption
	Steps   ExecutionStepService
}

// NewExecutionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExecutionService(opts ...option.RequestOption) (r ExecutionService) {
	r = ExecutionService{}
	r.Options = opts
	r.Steps = NewExecutionStepService(opts...)
	return
}

// Get a list of executions
func (r *ExecutionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vektor/executions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get an execution
func (r *ExecutionService) Get(ctx context.Context, executionID string, opts ...option.RequestOption) (res *Execution, err error) {
	opts = slices.Concat(r.Options, opts)
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vektor/executions/%s", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
