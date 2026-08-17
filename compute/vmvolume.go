// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nirvana-labs/nirvana-go/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
)

// VMVolumeService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMVolumeService] method instead.
type VMVolumeService struct {
	Options []option.RequestOption
}

// NewVMVolumeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVMVolumeService(opts ...option.RequestOption) (r VMVolumeService) {
	r = VMVolumeService{}
	r.Options = opts
	return
}

// List VM's Volumes
func (r *VMVolumeService) List(ctx context.Context, vmID string, query VMVolumeListParams, opts ...option.RequestOption) (res *pagination.Cursor[Volume], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if vmID == "" {
		err = errors.New("missing required vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/compute/vms/%s/volumes", vmID)
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

// List VM's Volumes
func (r *VMVolumeService) ListAutoPaging(ctx context.Context, vmID string, query VMVolumeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Volume] {
	return pagination.NewCursorAutoPager(r.List(ctx, vmID, query, opts...))
}

type VMVolumeListParams struct {
	// Pagination cursor returned by a previous request. Only valid for the same VM,
	// filters and sort order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the Volume name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at, name, status, size
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by Volume kind
	//
	// Any of "boot", "data".
	Kind VMVolumeListParamsKind `query:"kind,omitzero" json:"-"`
	// Filter by Volume status
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "error".
	Status VMVolumeListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by tags. Repeat the parameter to require several tags; a Volume must
	// carry all of them.
	Tags []string `query:"tags,omitzero" json:"-"`
	// Filter by storage type
	//
	// Any of "nvme", "abs".
	Type VMVolumeListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VMVolumeListParams]'s query parameters as `url.Values`.
func (r VMVolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by Volume kind
type VMVolumeListParamsKind string

const (
	VMVolumeListParamsKindBoot VMVolumeListParamsKind = "boot"
	VMVolumeListParamsKindData VMVolumeListParamsKind = "data"
)

// Filter by Volume status
type VMVolumeListParamsStatus string

const (
	VMVolumeListParamsStatusPending  VMVolumeListParamsStatus = "pending"
	VMVolumeListParamsStatusCreating VMVolumeListParamsStatus = "creating"
	VMVolumeListParamsStatusUpdating VMVolumeListParamsStatus = "updating"
	VMVolumeListParamsStatusReady    VMVolumeListParamsStatus = "ready"
	VMVolumeListParamsStatusDeleting VMVolumeListParamsStatus = "deleting"
	VMVolumeListParamsStatusError    VMVolumeListParamsStatus = "error"
)

// Filter by storage type
type VMVolumeListParamsType string

const (
	VMVolumeListParamsTypeNvme VMVolumeListParamsType = "nvme"
	VMVolumeListParamsTypeABS  VMVolumeListParamsType = "abs"
)
