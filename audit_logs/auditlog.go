// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/nirvana-labs/nirvana-go/shared"
)

// AuditLogService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.Options = opts
	return
}

// List Audit Log entries for an organization
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.Cursor[AuditLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/audit_logs"
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

// List Audit Log entries for an organization
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AuditLog] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Get an Audit Log entry
func (r *AuditLogService) Get(ctx context.Context, auditLogID string, opts ...option.RequestOption) (res *AuditLog, err error) {
	opts = slices.Concat(r.Options, opts)
	if auditLogID == "" {
		err = errors.New("missing required audit_log_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/audit_logs/%s", auditLogID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Audit log entry.
type AuditLog struct {
	// Unique identifier for the audit log entry.
	ID string `json:"id" api:"required"`
	// The action that was performed.
	Action string `json:"action" api:"required"`
	// The entity that performed the action.
	Actor AuditLogActor `json:"actor" api:"required"`
	// Client IP address.
	ClientIP string `json:"client_ip" api:"required"`
	// When the action occurred.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// HTTP method of the request.
	Method string `json:"method" api:"required"`
	// Request path.
	Path string `json:"path" api:"required"`
	// HTTP status code of the response.
	StatusCode int64 `json:"status_code" api:"required"`
	// User agent string.
	UserAgent string `json:"user_agent" api:"required"`
	// The target resource of the action.
	Target AuditLogTarget `json:"target" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Actor       respjson.Field
		ClientIP    respjson.Field
		CreatedAt   respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		StatusCode  respjson.Field
		UserAgent   respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLog) RawJSON() string { return r.JSON.raw }
func (r *AuditLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The target resource of the action.
type AuditLogTarget struct {
	// Unique identifier for the target resource.
	ID string `json:"id" api:"required"`
	// Type of the target resource.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogTarget) RawJSON() string { return r.JSON.raw }
func (r *AuditLogTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The entity that performed the action.
type AuditLogActor struct {
	// Unique identifier for the actor.
	ID string `json:"id" api:"required"`
	// Display name of the actor.
	Name string `json:"name" api:"required"`
	// Type of actor.
	//
	// Any of "user", "api_key".
	Type AuditLogType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogActor) RawJSON() string { return r.JSON.raw }
func (r *AuditLogActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogList struct {
	Items []AuditLog `json:"items" api:"required"`
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
func (r AuditLogList) RawJSON() string { return r.JSON.raw }
func (r *AuditLogList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of actor.
type AuditLogType string

const (
	AuditLogTypeUser   AuditLogType = "user"
	AuditLogTypeAPIKey AuditLogType = "api_key"
)

type AuditLogListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
