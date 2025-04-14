// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// UserService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Get details about an authenticated user
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type User struct {
	// Unique identifier for the user.
	ID string `json:"id,required"`
	// Email address of the user.
	Email string `json:"email,required"`
	// Services that the user has access to.
	Services UserServices `json:"services,required"`
	JSON     userJSON     `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

// Services that the user has access to.
type UserServices struct {
	Cloud bool             `json:"cloud"`
	JSON  userServicesJSON `json:"-"`
}

// userServicesJSON contains the JSON metadata for the struct [UserServices]
type userServicesJSON struct {
	Cloud       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userServicesJSON) RawJSON() string {
	return r.raw
}
