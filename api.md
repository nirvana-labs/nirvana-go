# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#ResourceStatus">ResourceStatus</a>

# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationListResponse">OperationListResponse</a>

Methods:

- <code title="get /operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationListResponse">OperationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#SSHKeyParam">SSHKeyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMList">VMList</a>

Methods:

- <code title="post /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMNewParams">VMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMUpdateParams">VMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMList">VMList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMOSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeListResponse">VolumeListResponse</a>

Methods:

- <code title="post /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeUpdateParams">VolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeListResponse">VolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Networking

## VPCs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#Subnet">Subnet</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPC">VPC</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCList">VPCList</a>

Methods:

- <code title="post /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCNewParams">VPCNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCList">VPCList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPC">VPC</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FirewallRules

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleEndpointParam">FirewallRuleEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleEndpoint">FirewallRuleEndpoint</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleList">FirewallRuleList</a>

Methods:

- <code title="post /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleNewParams">FirewallRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleUpdateParams">FirewallRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleList">FirewallRuleList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
