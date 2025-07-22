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

// RegistryVenueService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryVenueService] method instead.
type RegistryVenueService struct {
	Options []option.RequestOption
}

// NewRegistryVenueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryVenueService(opts ...option.RequestOption) (r RegistryVenueService) {
	r = RegistryVenueService{}
	r.Options = opts
	return
}

// List supported venues
func (r *RegistryVenueService) List(ctx context.Context, body RegistryVenueListParams, opts ...option.RequestOption) (res *RegistryVenueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vektor/registry/venues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegistryVenueListResponse struct {
	// A list of venues
	Items []Venue `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryVenueListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegistryVenueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegistryVenueListParams struct {
	// A venue ID, represented as a TypeID with `venue` prefix
	ID param.Opt[string] `json:"id,omitzero"`
	// A venue symbol
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	paramObj
}

func (r RegistryVenueListParams) MarshalJSON() (data []byte, err error) {
	type shadow RegistryVenueListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegistryVenueListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
