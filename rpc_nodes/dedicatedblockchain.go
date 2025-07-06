// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rpc_nodes

import (
	"context"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
)

// DedicatedBlockchainService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDedicatedBlockchainService] method instead.
type DedicatedBlockchainService struct {
	Options []option.RequestOption
}

// NewDedicatedBlockchainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDedicatedBlockchainService(opts ...option.RequestOption) (r DedicatedBlockchainService) {
	r = DedicatedBlockchainService{}
	r.Options = opts
	return
}

// List all Dedicated Blockchains
func (r *DedicatedBlockchainService) List(ctx context.Context, opts ...option.RequestOption) (res *RPCNodesDedicatedBlockchainList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rpc_nodes/dedicated/blockchains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
