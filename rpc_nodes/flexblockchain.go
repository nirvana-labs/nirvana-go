// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// FlexBlockchainService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFlexBlockchainService] method instead.
type FlexBlockchainService struct {
	Options []option.RequestOption
}

// NewFlexBlockchainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFlexBlockchainService(opts ...option.RequestOption) (r FlexBlockchainService) {
	r = FlexBlockchainService{}
	r.Options = opts
	return
}

// List all Flex Blockchains
func (r *FlexBlockchainService) List(ctx context.Context, opts ...option.RequestOption) (res *RPCNodesFlexBlockchainList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rpc_nodes/flex/blockchains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
