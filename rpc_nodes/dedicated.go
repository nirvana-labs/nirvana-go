// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// DedicatedService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDedicatedService] method instead.
type DedicatedService struct {
	Options     []option.RequestOption
	Blockchains DedicatedBlockchainService
}

// NewDedicatedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDedicatedService(opts ...option.RequestOption) (r DedicatedService) {
	r = DedicatedService{}
	r.Options = opts
	r.Blockchains = NewDedicatedBlockchainService(opts...)
	return
}

// List all RPC Node Dedicated you created
func (r *DedicatedService) List(ctx context.Context, opts ...option.RequestOption) (res *RPCNodesDedicatedList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rpc_nodes/dedicated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about an RPC Node Dedicated
func (r *DedicatedService) Get(ctx context.Context, nodeID string, opts ...option.RequestOption) (res *RPCNodesDedicated, err error) {
	opts = append(r.Options[:], opts...)
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return
	}
	path := fmt.Sprintf("v1/rpc_nodes/dedicated/%s", nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RPCNodesDedicated struct {
	ID         string    `json:"id,required"`
	Blockchain string    `json:"blockchain,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	Endpoint   string    `json:"endpoint,required"`
	Name       string    `json:"name,required"`
	Network    string    `json:"network,required"`
	UpdatedAt  time.Time `json:"updated_at,required" format:"date-time"`
	UserID     string    `json:"user_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Blockchain  respjson.Field
		CreatedAt   respjson.Field
		Endpoint    respjson.Field
		Name        respjson.Field
		Network     respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RPCNodesDedicated) RawJSON() string { return r.JSON.raw }
func (r *RPCNodesDedicated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Blockchain supported by a dedicated node.
type RPCNodesDedicatedBlockchain struct {
	Blockchain string `json:"blockchain,required"`
	Network    string `json:"network,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain  respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RPCNodesDedicatedBlockchain) RawJSON() string { return r.JSON.raw }
func (r *RPCNodesDedicatedBlockchain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RPCNodesDedicatedBlockchainList struct {
	Items []RPCNodesDedicatedBlockchain `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RPCNodesDedicatedBlockchainList) RawJSON() string { return r.JSON.raw }
func (r *RPCNodesDedicatedBlockchainList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RPCNodesDedicatedList struct {
	Items []RPCNodesDedicated `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RPCNodesDedicatedList) RawJSON() string { return r.JSON.raw }
func (r *RPCNodesDedicatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
