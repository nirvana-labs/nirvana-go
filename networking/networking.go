// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// NetworkingService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkingService] method instead.
type NetworkingService struct {
	Options       []option.RequestOption
	VPCs          VPCService
	FirewallRules FirewallRuleService
}

// NewNetworkingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkingService(opts ...option.RequestOption) (r NetworkingService) {
	r = NetworkingService{}
	r.Options = opts
	r.VPCs = NewVPCService(opts...)
	r.FirewallRules = NewFirewallRuleService(opts...)
	return
}
