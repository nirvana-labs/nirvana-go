// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// FlexService contains methods and other services that help with interacting with
// the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFlexService] method instead.
type FlexService struct {
	Options     []option.RequestOption
	Blockchains FlexBlockchainService
}

// NewFlexService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFlexService(opts ...option.RequestOption) (r FlexService) {
	r = FlexService{}
	r.Options = opts
	r.Blockchains = NewFlexBlockchainService(opts...)
	return
}

// List all RPC Node Flex you created
func (r *FlexService) List(ctx context.Context, opts ...option.RequestOption) (res *FlexList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/rpc_nodes/flex"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details about an RPC Node Flex
func (r *FlexService) Get(ctx context.Context, nodeID string, opts ...option.RequestOption) (res *Flex, err error) {
	opts = slices.Concat(r.Options, opts)
	if nodeID == "" {
		err = errors.New("missing required node_id parameter")
		return
	}
	path := fmt.Sprintf("v1/rpc_nodes/flex/%s", nodeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// RPC Node Flex details.
type Flex struct {
	// Unique identifier for the RPC Node Flex.
	ID string `json:"id,required"`
	// Blockchain type.
	Blockchain string `json:"blockchain,required"`
	// When the RPC Node Flex was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// RPC endpoint URL.
	Endpoint string `json:"endpoint,required"`
	// Name of the RPC Node Flex.
	Name string `json:"name,required"`
	// Network type (e.g., mainnet, testnet).
	Network string `json:"network,required"`
	// When the RPC Node Flex was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Blockchain  respjson.Field
		CreatedAt   respjson.Field
		Endpoint    respjson.Field
		Name        respjson.Field
		Network     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Flex) RawJSON() string { return r.JSON.raw }
func (r *Flex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Blockchain supported by the Flex RPC Node.
type FlexBlockchain struct {
	// Blockchain type.
	Blockchain string `json:"blockchain,required"`
	// Network type (e.g., mainnet, testnet).
	Network string `json:"network,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain  respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlexBlockchain) RawJSON() string { return r.JSON.raw }
func (r *FlexBlockchain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlexBlockchainList struct {
	Items []FlexBlockchain `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlexBlockchainList) RawJSON() string { return r.JSON.raw }
func (r *FlexBlockchainList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlexList struct {
	Items []Flex `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlexList) RawJSON() string { return r.JSON.raw }
func (r *FlexList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
