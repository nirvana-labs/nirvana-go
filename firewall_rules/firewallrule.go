// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rules

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/param"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/operations"
	"github.com/nirvana-labs/nirvana-go/option"
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
func NewFirewallRuleService(opts ...option.RequestOption) (r *FirewallRuleService) {
	r = &FirewallRuleService{}
	r.Options = opts
	return
}

// Create a firewall rule
func (r *FirewallRuleService) New(ctx context.Context, vpcID string, body FirewallRuleNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s/firewall_rules", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a firewall rule
func (r *FirewallRuleService) Update(ctx context.Context, vpcID string, firewallRuleID string, body FirewallRuleUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all firewall rules
func (r *FirewallRuleService) List(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *FirewallRuleList, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s/firewall_rules", vpcID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a firewall rule
func (r *FirewallRuleService) Delete(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a firewall rule
func (r *FirewallRuleService) Get(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = append(r.Options[:], opts...)
	if vpcID == "" {
		err = errors.New("missing required vpc_id parameter")
		return
	}
	if firewallRuleID == "" {
		err = errors.New("missing required firewall_rule_id parameter")
		return
	}
	path := fmt.Sprintf("networking/vpcs/%s/firewall_rules/%s", vpcID, firewallRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Firewall rule details.
type FirewallRule struct {
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	// Firewall rule endpoint.
	Destination FirewallRuleEndpoint `json:"destination,required"`
	Name        string               `json:"name,required"`
	Protocol    string               `json:"protocol,required"`
	// Firewall rule endpoint.
	Source    FirewallRuleEndpoint  `json:"source,required"`
	Status    shared.ResourceStatus `json:"status,required"`
	UpdatedAt string                `json:"updated_at,required"`
	VPCID     string                `json:"vpc_id,required"`
	JSON      firewallRuleJSON      `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Destination apijson.Field
	Name        apijson.Field
	Protocol    apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	VPCID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

// Firewall rule endpoint.
type FirewallRuleEndpoint struct {
	Address string                   `json:"address"`
	Ports   []string                 `json:"ports"`
	JSON    firewallRuleEndpointJSON `json:"-"`
}

// firewallRuleEndpointJSON contains the JSON metadata for the struct
// [FirewallRuleEndpoint]
type firewallRuleEndpointJSON struct {
	Address     apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEndpointJSON) RawJSON() string {
	return r.raw
}

// Firewall rule endpoint.
type FirewallRuleEndpointParam struct {
	Address param.Field[string]   `json:"address"`
	Ports   param.Field[[]string] `json:"ports"`
}

func (r FirewallRuleEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallRuleList struct {
	Items []FirewallRule       `json:"items,required"`
	JSON  firewallRuleListJSON `json:"-"`
}

// firewallRuleListJSON contains the JSON metadata for the struct
// [FirewallRuleList]
type firewallRuleListJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleListJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleNewParams struct {
	// Firewall rule endpoint.
	Destination param.Field[FirewallRuleEndpointParam] `json:"destination,required"`
	Name        param.Field[string]                    `json:"name,required"`
	// Supported protocols.
	Protocol param.Field[string] `json:"protocol,required"`
	// Firewall rule endpoint.
	Source param.Field[FirewallRuleEndpointParam] `json:"source,required"`
}

func (r FirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallRuleUpdateParams struct {
	// Firewall rule endpoint.
	Destination param.Field[FirewallRuleEndpointParam] `json:"destination,required"`
	Name        param.Field[string]                    `json:"name,required"`
	// Supported protocols.
	Protocol param.Field[FirewallRuleUpdateParamsProtocol] `json:"protocol,required"`
	// Firewall rule endpoint.
	Source param.Field[FirewallRuleEndpointParam] `json:"source,required"`
}

func (r FirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Supported protocols.
type FirewallRuleUpdateParamsProtocol string

const (
	FirewallRuleUpdateParamsProtocolTcp FirewallRuleUpdateParamsProtocol = "tcp"
	FirewallRuleUpdateParamsProtocolUdp FirewallRuleUpdateParamsProtocol = "udp"
)

func (r FirewallRuleUpdateParamsProtocol) IsKnown() bool {
	switch r {
	case FirewallRuleUpdateParamsProtocolTcp, FirewallRuleUpdateParamsProtocolUdp:
		return true
	}
	return false
}
