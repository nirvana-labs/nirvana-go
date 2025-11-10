// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

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
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/pagination"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// FirewallRuleService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFirewallRuleService] method instead.
type FirewallRuleService struct {
	Options []option.RequestOption
}

// NewFirewallRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallRuleService(opts ...option.RequestOption) (r FirewallRuleService) {
	r = FirewallRuleService{}
	r.Options = opts
	return
}

// Create a firewall rule
func (r *FirewallRuleService) New(ctx context.Context, vpcID string, body FirewallRuleNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/firewall_rules", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a firewall rule
func (r *FirewallRuleService) Update(ctx context.Context, vpcID string, firewallRuleID string, body FirewallRuleUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all firewall rules
func (r *FirewallRuleService) List(ctx context.Context, vpcID string, query FirewallRuleListParams, opts ...option.RequestOption) (res *pagination.Cursor[FirewallRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/firewall_rules", vpcID)
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

// List all firewall rules
func (r *FirewallRuleService) ListAutoPaging(ctx context.Context, vpcID string, query FirewallRuleListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[FirewallRule] {
	return pagination.NewCursorAutoPager(r.List(ctx, vpcID, query, opts...))
}

// Delete a firewall rule
func (r *FirewallRuleService) Delete(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a firewall rule
func (r *FirewallRuleService) Get(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("v1/networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Firewall rule details.
type FirewallRule struct {
	// Unique identifier for the Firewall Rule.
	ID string `json:"id,required"`
	// When the Firewall Rule was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC.
	DestinationAddress string `json:"destination_address,required"`
	// Destination ports of the Firewall Rule.
	DestinationPorts []string `json:"destination_ports,required"`
	// Name of the Firewall Rule.
	Name string `json:"name,required"`
	// Protocol of the Firewall Rule.
	//
	// Any of "tcp", "udp".
	Protocol FirewallRuleProtocol `json:"protocol,required"`
	// Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask
	// of 0.
	SourceAddress string `json:"source_address,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Tags to attach to the Firewall Rule.
	Tags []string `json:"tags,required"`
	// When the Firewall Rule was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the VPC the Firewall Rule belongs to.
	VPCID string `json:"vpc_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		DestinationAddress respjson.Field
		DestinationPorts   respjson.Field
		Name               respjson.Field
		Protocol           respjson.Field
		SourceAddress      respjson.Field
		Status             respjson.Field
		Tags               respjson.Field
		UpdatedAt          respjson.Field
		VPCID              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FirewallRule) RawJSON() string { return r.JSON.raw }
func (r *FirewallRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol of the Firewall Rule.
type FirewallRuleProtocol string

const (
	FirewallRuleProtocolTcp FirewallRuleProtocol = "tcp"
	FirewallRuleProtocolUdp FirewallRuleProtocol = "udp"
)

type FirewallRuleList struct {
	Items []FirewallRule `json:"items,required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FirewallRuleList) RawJSON() string { return r.JSON.raw }
func (r *FirewallRuleList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleNewParams struct {
	// Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC. Must be
	// in network-aligned/canonical form.
	DestinationAddress string `json:"destination_address,required"`
	// Destination ports of the Firewall Rule.
	DestinationPorts []string `json:"destination_ports,omitzero,required"`
	// Name of the Firewall Rule.
	Name string `json:"name,required"`
	// Protocol of the Firewall Rule.
	//
	// Any of "tcp", "udp".
	Protocol FirewallRuleNewParamsProtocol `json:"protocol,omitzero,required"`
	// Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask
	// of 0. Must be in network-aligned/canonical form.
	SourceAddress string `json:"source_address,required"`
	// Tags to attach to the Firewall Rule.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FirewallRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FirewallRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol of the Firewall Rule.
type FirewallRuleNewParamsProtocol string

const (
	FirewallRuleNewParamsProtocolTcp FirewallRuleNewParamsProtocol = "tcp"
	FirewallRuleNewParamsProtocolUdp FirewallRuleNewParamsProtocol = "udp"
)

type FirewallRuleUpdateParams struct {
	// Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC. Must be
	// in network-aligned/canonical form.
	DestinationAddress param.Opt[string] `json:"destination_address,omitzero"`
	// Name of the Firewall Rule.
	Name param.Opt[string] `json:"name,omitzero"`
	// Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask
	// of 0. Must be in network-aligned/canonical form.
	SourceAddress param.Opt[string] `json:"source_address,omitzero"`
	// Destination ports of the Firewall Rule.
	DestinationPorts []string `json:"destination_ports,omitzero"`
	// Protocol of the Firewall Rule.
	//
	// Any of "tcp", "udp".
	Protocol FirewallRuleUpdateParamsProtocol `json:"protocol,omitzero"`
	// Tags to attach to the Firewall Rule.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FirewallRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FirewallRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol of the Firewall Rule.
type FirewallRuleUpdateParamsProtocol string

const (
	FirewallRuleUpdateParamsProtocolTcp FirewallRuleUpdateParamsProtocol = "tcp"
	FirewallRuleUpdateParamsProtocolUdp FirewallRuleUpdateParamsProtocol = "udp"
)

type FirewallRuleListParams struct {
	// Pagination cursor returned by a previous request
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FirewallRuleListParams]'s query parameters as `url.Values`.
func (r FirewallRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
