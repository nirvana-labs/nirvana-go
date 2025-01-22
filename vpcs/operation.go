// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vpcs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
)

// OperationService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	Options []option.RequestOption
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r *OperationService) {
	r = &OperationService{}
	r.Options = opts
	return
}

// Get details of a VPC operation
func (r *OperationService) Get(ctx context.Context, operationID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("vpcs/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
