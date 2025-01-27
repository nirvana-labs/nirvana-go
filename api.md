# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#ResourceStatus">ResourceStatus</a>

# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>

Methods:

- <code title="get /operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#CPUParam">CPUParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#RamParam">RamParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#SSHKeyParam">SSHKeyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#CPU">CPU</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Ram">Ram</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VMList">VMList</a>

Methods:

- <code title="post /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMNewParams">ComputeVMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMUpdateParams">ComputeVMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VMList">VMList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVMOSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#OSImage">OSImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#StorageType">StorageType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeListResponse">ComputeVolumeListResponse</a>

Methods:

- <code title="post /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeNewParams">ComputeVolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeUpdateParams">ComputeVolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeListResponse">ComputeVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeDeleteParams">ComputeVolumeDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#ComputeVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Networking

## VPCs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Subnet">Subnet</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VPC">VPC</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VPCList">VPCList</a>

Methods:

- <code title="post /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingVPCService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingVPCNewParams">NetworkingVPCNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingVPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VPCList">VPCList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingVPCService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingVPCService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#VPC">VPC</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FirewallRules

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRuleEndpointParam">FirewallRuleEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRule">FirewallRule</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRuleEndpoint">FirewallRuleEndpoint</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRuleList">FirewallRuleList</a>

Methods:

- <code title="post /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleNewParams">NetworkingFirewallRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleUpdateParams">NetworkingFirewallRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRuleList">FirewallRuleList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#NetworkingFirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go">nirvana</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
