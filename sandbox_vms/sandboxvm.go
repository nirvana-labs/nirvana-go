// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sandbox_vms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nirvana-labs/nirvana-go/v3/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/v3/internal/apiquery"
	"github.com/nirvana-labs/nirvana-go/v3/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/v3/operations"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/v3/packages/param"
	"github.com/nirvana-labs/nirvana-go/v3/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/v3/shared"
)

// SandboxVMService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxVMService] method instead.
type SandboxVMService struct {
	Options []option.RequestOption
	Cost    CostService
}

// NewSandboxVMService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxVMService(opts ...option.RequestOption) (r SandboxVMService) {
	r = SandboxVMService{}
	r.Options = opts
	r.Cost = NewCostService(opts...)
	return
}

// Create a Sandbox VM to run sandboxes on. A Sandbox VM can be paused and resumed,
// and is billed for compute only while it is running.
func (r *SandboxVMService) New(ctx context.Context, body SandboxVMNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sandbox_vms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a Sandbox VM's name and tags
func (r *SandboxVMService) Update(ctx context.Context, sandboxVMID string, body SandboxVMUpdateParams, opts ...option.RequestOption) (res *SandboxVM, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the Sandbox VMs in a project
func (r *SandboxVMService) List(ctx context.Context, query SandboxVMListParams, opts ...option.RequestOption) (res *pagination.Cursor[SandboxVM], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/sandbox_vms"
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

// List the Sandbox VMs in a project
func (r *SandboxVMService) ListAutoPaging(ctx context.Context, query SandboxVMListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[SandboxVM] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a Sandbox VM, with its sandboxes and its disk. A Sandbox VM can be
// deleted when it is ready or failed, whether it is running, paused or powered
// off.
func (r *SandboxVMService) Delete(ctx context.Context, sandboxVMID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a Sandbox VM
func (r *SandboxVMService) Get(ctx context.Context, sandboxVMID string, opts ...option.RequestOption) (res *SandboxVM, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Pause a Sandbox VM. Its sandboxes stop, keeping their files and installed
// packages but not their running processes or memory, and compute billing stops.
// After an hour paused the Sandbox VM is powered off. A pause sent while another
// change runs waits for it, and pausing a paused Sandbox VM changes nothing.
func (r *SandboxVMService) Pause(ctx context.Context, sandboxVMID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s/pause", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Restart a Sandbox VM's VM. Its sandboxes keep their files and installed packages
// but not their running processes or memory. A restart sent while another change
// runs waits for it. A billing block returns 402.
func (r *SandboxVMService) Restart(ctx context.Context, sandboxVMID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s/restart", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Resume a paused or powered-off Sandbox VM, restarting its sandboxes. A resume
// sent while another change runs waits for it, and resuming a running Sandbox VM
// changes nothing. Resuming can fail if the organization's sandbox quota was used
// up while the Sandbox VM was paused. A billing block returns 402.
func (r *SandboxVMService) Resume(ctx context.Context, sandboxVMID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxVMID == "" {
		err = errors.New("missing required sandbox_vm_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sandbox_vms/%s/resume", sandboxVMID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A VM sandboxes run on, in one region. It can be paused and resumed.
type SandboxVM struct {
	// Unique identifier for the Sandbox VM.
	ID string `json:"id" api:"required"`
	// When the Sandbox VM was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Hostname the Sandbox VM's sandbox API is served on over HTTPS. It stays the same
	// for the life of the Sandbox VM. Null until assigned.
	EndpointHostname string `json:"endpoint_hostname" api:"required"`
	// Version of the image the Sandbox VM runs. Null until it is running.
	ImageVersion string `json:"image_version" api:"required"`
	// Instance type of the Sandbox VM.
	InstanceType string `json:"instance_type" api:"required"`
	// Name of the Sandbox VM.
	Name string `json:"name" api:"required"`
	// Version of OpenSandbox the Sandbox VM runs. Null until it is running.
	OpensandboxVersion string `json:"opensandbox_version" api:"required"`
	// When the Sandbox VM was paused. Null while it is running.
	PausedAt time.Time `json:"paused_at" api:"required" format:"date-time"`
	// Whether the Sandbox VM's sandboxes can run.
	//
	// Any of "running", "paused", "powered_off".
	PowerState SandboxVMPowerState `json:"power_state" api:"required"`
	// When the paused Sandbox VM is, or was, powered off. Null while it is running.
	PowersOffAt time.Time `json:"powers_off_at" api:"required" format:"date-time"`
	// ID of the project the Sandbox VM is in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region" api:"required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error", "suspended".
	Status shared.ResourceStatus `json:"status" api:"required"`
	// Tags attached to the Sandbox VM.
	Tags []string `json:"tags" api:"required"`
	// When the Sandbox VM was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the VPC the Sandbox VM is in. Null until it is placed in one.
	VPCID string `json:"vpc_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		EndpointHostname   respjson.Field
		ImageVersion       respjson.Field
		InstanceType       respjson.Field
		Name               respjson.Field
		OpensandboxVersion respjson.Field
		PausedAt           respjson.Field
		PowerState         respjson.Field
		PowersOffAt        respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		Status             respjson.Field
		Tags               respjson.Field
		UpdatedAt          respjson.Field
		VPCID              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxVM) RawJSON() string { return r.JSON.raw }
func (r *SandboxVM) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Sandbox VM's sandboxes can run.
type SandboxVMPowerState string

const (
	SandboxVMPowerStateRunning    SandboxVMPowerState = "running"
	SandboxVMPowerStatePaused     SandboxVMPowerState = "paused"
	SandboxVMPowerStatePoweredOff SandboxVMPowerState = "powered_off"
)

type SandboxVMList struct {
	Items []SandboxVM `json:"items" api:"required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxVMList) RawJSON() string { return r.JSON.raw }
func (r *SandboxVMList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxVMNewParams struct {
	// Name of the Sandbox VM. Unique within its project.
	Name string `json:"name" api:"required"`
	// ID of the project to create the Sandbox VM in.
	ProjectID string `json:"project_id" api:"required"`
	// Region the resource is in.
	//
	// Any of "us-sva-2", "us-wdc-1".
	Region shared.RegionName `json:"region,omitzero" api:"required"`
	// Instance type of the Sandbox VM. Defaults to n1-highcpu-2 (2 vCPU, 4 GiB), the
	// only size available.
	InstanceType param.Opt[string] `json:"instance_type,omitzero"`
	// ID of the VPC to place the Sandbox VM in. It must be a ready VPC in the Sandbox
	// VM's project and region. Omit it to use the project's oldest ready VPC in the
	// region, or a new VPC when there is none.
	VPCID param.Opt[string] `json:"vpc_id,omitzero"`
	// Tags to attach to the Sandbox VM.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SandboxVMNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SandboxVMNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxVMNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxVMUpdateParams struct {
	// Name of the Sandbox VM. Unique within its project.
	Name param.Opt[string] `json:"name,omitzero"`
	// Tags to attach to the Sandbox VM.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SandboxVMUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SandboxVMUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxVMUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxVMListParams struct {
	// Project ID of resources to request
	ProjectID string `query:"project_id" api:"required" json:"-"`
	// Pagination cursor returned by a previous request. Only valid for the same sort
	// order.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a case-insensitive substring of the Sandbox VM name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by region
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Comma-separated sort terms in precedence order, each field:asc or field:desc.
	// Fields: created_at, updated_at
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter by tags. Repeat the parameter to require several tags; a Sandbox VM must
	// carry all of them.
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxVMListParams]'s query parameters as `url.Values`.
func (r SandboxVMListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
