// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// RegistryErrorService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryErrorService] method instead.
type RegistryErrorService struct {
	Options []option.RequestOption
}

// NewRegistryErrorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryErrorService(opts ...option.RequestOption) (r RegistryErrorService) {
	r = RegistryErrorService{}
	r.Options = opts
	return
}

// A list with one example of each error type
func (r *RegistryErrorService) List(ctx context.Context, body RegistryErrorListParams, opts ...option.RequestOption) (res *RegistryErrorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/registry/errors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryErrorListResponse struct {
	// A list of errors
	Items []RegistryErrorListResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryErrorListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryErrorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type RegistryErrorListResponseItem struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryErrorListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *RegistryErrorListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryErrorListParams struct {
	Errors []string `json:"errors,omitzero"`
	paramObj
}

func (r RegistryErrorListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryErrorListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryErrorListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
