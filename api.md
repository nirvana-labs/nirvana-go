# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#ResourceStatus">ResourceStatus</a>

# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>

Methods:

- <code title="get /operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#CPUParam">CPUParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#RamParam">RamParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#SSHKeyParam">SSHKeyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#CPU">CPU</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#Ram">Ram</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMList">VMList</a>

Methods:

- <code title="post /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMNewParams">ComputeVMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMUpdateParams">ComputeVMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMList">VMList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#ComputeVMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#OSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#OSImage">OSImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#StorageType">StorageType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeListResponse">ComputeVolumeListResponse</a>

Methods:

- <code title="post /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeNewParams">ComputeVolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeUpdateParams">ComputeVolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeListResponse">ComputeVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeDeleteParams">ComputeVolumeDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#ComputeVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Networking

## VPCs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#Subnet">Subnet</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPC">VPC</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCList">VPCList</a>

Methods:

- <code title="post /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#NetworkingVPCService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#NetworkingVPCNewParams">NetworkingVPCNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#NetworkingVPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCList">VPCList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#NetworkingVPCService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#NetworkingVPCService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPC">VPC</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FirewallRules

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRuleEndpointParam">FirewallRuleEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRule">FirewallRule</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRuleEndpoint">FirewallRuleEndpoint</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRuleList">FirewallRuleList</a>

Methods:

- <code title="post /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleNewParams">NetworkingFirewallRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleUpdateParams">NetworkingFirewallRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRuleList">FirewallRuleList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#NetworkingFirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules">firewall_rules</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rules#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
