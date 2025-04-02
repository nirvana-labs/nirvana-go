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
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	return
}

// Create a new API key
func (r *APIKeyService) New(ctx context.Context, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List all API keys you created
func (r *APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *APIKeyList, err error) {
	opts = append(r.Options[:], opts...)
	path := "api_keys"
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
	path := fmt.Sprintf("api_keys/%s", apiKeyID)
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
	path := fmt.Sprintf("api_keys/%s", apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// API key response.
type APIKey struct {
	// API key ID.
	ID string `json:"id,required"`
	// Time on which the API key was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Time after which the API key is not valid.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// API key name.
	Name string `json:"name,required"`
	// Status of the API key.
	Status APIKeyStatus `json:"status,required"`
	// User ID that owns the API key.
	UserID string `json:"user_id,required"`
	// API key.
	Key string `json:"key"`
	// Time before which the API key is not valid.
	NotBefore time.Time  `json:"not_before" format:"date-time"`
	JSON      apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	UserID      apijson.Field
	Key         apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

// Status of the API key.
type APIKeyStatus string

const (
	APIKeyStatusActive   APIKeyStatus = "active"
	APIKeyStatusInactive APIKeyStatus = "inactive"
	APIKeyStatusExpired  APIKeyStatus = "expired"
)

func (r APIKeyStatus) IsKnown() bool {
	switch r {
	case APIKeyStatusActive, APIKeyStatusInactive, APIKeyStatusExpired:
		return true
	}
	return false
}

type APIKeyList struct {
	Items []APIKey       `json:"items,required"`
	JSON  apiKeyListJSON `json:"-"`
}

// apiKeyListJSON contains the JSON metadata for the struct [APIKeyList]
type apiKeyListJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListJSON) RawJSON() string {
	return r.raw
}
