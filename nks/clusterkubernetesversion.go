// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nks

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// ClusterKubernetesVersionService contains methods and other services that help
// with interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterKubernetesVersionService] method instead.
type ClusterKubernetesVersionService struct {
	Options []option.RequestOption
}

// NewClusterKubernetesVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClusterKubernetesVersionService(opts ...option.RequestOption) (r ClusterKubernetesVersionService) {
	r = ClusterKubernetesVersionService{}
	r.Options = opts
	return
}

// List all supported Kubernetes versions for NKS clusters
func (r *ClusterKubernetesVersionService) List(ctx context.Context, query ClusterKubernetesVersionListParams, opts ...option.RequestOption) (res *pagination.Cursor[ClusterKubernetesVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/nks/kubernetes_versions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all supported Kubernetes versions for NKS clusters
func (r *ClusterKubernetesVersionService) ListAutoPaging(ctx context.Context, query ClusterKubernetesVersionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ClusterKubernetesVersionListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// NKS Kubernetes version details.
type ClusterKubernetesVersionListResponse struct {
	// When the Kubernetes version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the Kubernetes version.
	DisplayName string `json:"display_name" api:"required"`
	// Name of the Kubernetes version.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterKubernetesVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterKubernetesVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClusterKubernetesVersionListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClusterKubernetesVersionListParams]'s query parameters as
// `url.Values`.
func (r ClusterKubernetesVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
