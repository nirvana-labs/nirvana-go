// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_keys

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// APIKeyService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	return
}

// Create a new API key
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an API key's name
func (r *APIKeyService) Update(ctx context.Context, apiKeyID string, body APIKeyUpdateParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all API keys you created
func (r *APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *APIKeyList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an API key
func (r *APIKeyService) Delete(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get details about an API key
func (r *APIKeyService) Get(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/api_keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// API key response.
type APIKey struct {
	// API key ID.
	ID string `json:"id,required"`
	// When the API key was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// When the API key expires and is no longer valid.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// API key name.
	Name string `json:"name,required"`
	// Status of the API key.
	//
	// Any of "active", "inactive", "expired".
	Status APIKeyStatus `json:"status,required"`
	// When the API key was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID that owns the API key.
	UserID string `json:"user_id,required"`
	// API key. Only returned on creation.
	Key string `json:"key"`
	// When the API key starts to be valid.
	StartsAt time.Time `json:"starts_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		Key         respjson.Field
		StartsAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the API key.
type APIKeyStatus string

const (
	APIKeyStatusActive   APIKeyStatus = "active"
	APIKeyStatusInactive APIKeyStatus = "inactive"
	APIKeyStatusExpired  APIKeyStatus = "expired"
)

type APIKeyList struct {
	Items []APIKey `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyList) RawJSON() string { return r.JSON.raw }
func (r *APIKeyList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	// When the API key expires and is no longer valid.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// API key name.
	Name string `json:"name,required"`
	// When the API key starts to be valid.
	StartsAt param.Opt[time.Time] `json:"starts_at,omitzero" format:"date-time"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyUpdateParams struct {
	// API key name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r APIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
