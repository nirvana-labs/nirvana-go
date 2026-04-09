// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// ClusterKubeconfigService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterKubeconfigService] method instead.
type ClusterKubeconfigService struct {
	Options []option.RequestOption
}

// NewClusterKubeconfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterKubeconfigService(opts ...option.RequestOption) (r ClusterKubeconfigService) {
	r = ClusterKubeconfigService{}
	r.Options = opts
	return
}

// Get the kubeconfig for an NKS cluster
func (r *ClusterKubeconfigService) Get(ctx context.Context, clusterID string, opts ...option.RequestOption) (res *Kubeconfig, err error) {
	opts = slices.Concat(r.Options, opts)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/nks/clusters/%s/kubeconfig", clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Kubeconfig for an NKS Cluster.
type Kubeconfig struct {
	// Kubeconfig content for the Cluster.
	Kubeconfig string `json:"kubeconfig" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kubeconfig  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Kubeconfig) RawJSON() string { return r.JSON.raw }
func (r *Kubeconfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
