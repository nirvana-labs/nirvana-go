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

// NetworkingFirewallRuleService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkingFirewallRuleService] method instead.
type NetworkingFirewallRuleService struct {
	Options []option.RequestOption
}

// NewNetworkingFirewallRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkingFirewallRuleService(opts ...option.RequestOption) (r *NetworkingFirewallRuleService) {
	r = &NetworkingFirewallRuleService{}
	r.Options = opts
	return
}

// Create a firewall rule
func (r *NetworkingFirewallRuleService) New(ctx context.Context, vpcID string, body NetworkingFirewallRuleNewParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
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
func (r *NetworkingFirewallRuleService) Update(ctx context.Context, vpcID string, firewallRuleID string, body NetworkingFirewallRuleUpdateParams, opts ...option.RequestOption) (res *operations.Operation, err error) {
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
func (r *NetworkingFirewallRuleService) List(ctx context.Context, vpcID string, opts ...option.RequestOption) (res *FirewallRuleList, err error) {
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
func (r *NetworkingFirewallRuleService) Delete(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *operations.Operation, err error) {
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
func (r *NetworkingFirewallRuleService) Get(ctx context.Context, vpcID string, firewallRuleID string, opts ...option.RequestOption) (res *FirewallRule, err error) {
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

type NetworkingFirewallRuleNewParams struct {
	// Firewall rule endpoint.
	Destination param.Field[FirewallRuleEndpointParam] `json:"destination,required"`
	Name        param.Field[string]                    `json:"name,required"`
	// Supported protocols.
	Protocol param.Field[string] `json:"protocol,required"`
	// Firewall rule endpoint.
	Source param.Field[FirewallRuleEndpointParam] `json:"source,required"`
}

func (r NetworkingFirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkingFirewallRuleUpdateParams struct {
	// Firewall rule endpoint.
	Destination param.Field[FirewallRuleEndpointParam] `json:"destination,required"`
	Name        param.Field[string]                    `json:"name,required"`
	// Supported protocols.
	Protocol param.Field[NetworkingFirewallRuleUpdateParamsProtocol] `json:"protocol,required"`
	// Firewall rule endpoint.
	Source param.Field[FirewallRuleEndpointParam] `json:"source,required"`
}

func (r NetworkingFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Supported protocols.
type NetworkingFirewallRuleUpdateParamsProtocol string

const (
	NetworkingFirewallRuleUpdateParamsProtocolTcp NetworkingFirewallRuleUpdateParamsProtocol = "tcp"
	NetworkingFirewallRuleUpdateParamsProtocolUdp NetworkingFirewallRuleUpdateParamsProtocol = "udp"
)

func (r NetworkingFirewallRuleUpdateParamsProtocol) IsKnown() bool {
	switch r {
	case NetworkingFirewallRuleUpdateParamsProtocolTcp, NetworkingFirewallRuleUpdateParamsProtocolUdp:
		return true
	}
	return false
}
