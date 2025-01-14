# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#OperationKind">OperationKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#OperationStatus">OperationStatus</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#OperationType">OperationType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#ResourceStatus">ResourceStatus</a>

# VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#CPUParam">CPUParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#RamParam">RamParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#RamUnit">RamUnit</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#SSHKeyParam">SSHKeyParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#StorageParam">StorageParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#StorageUnit">StorageUnit</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#CPU">CPU</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#Ram">Ram</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#RamUnit">RamUnit</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#Storage">Storage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#StorageUnit">StorageUnit</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMList">VMList</a>

Methods:

- <code title="post /vms">client.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMNewParams">VMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /vms/{vm_id}">client.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMUpdateParams">VMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vms">client.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMListParams">VMListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMList">VMList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /vms/{vm_id}">client.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vms/{vm_id}">client.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms">vms</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Operations

Methods:

- <code title="get /vms/operations/{operation_id}">client.VMs.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vms#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VPCs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#Subnet">Subnet</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPC">VPC</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCList">VPCList</a>

Methods:

- <code title="post /vpcs">client.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCNewParams">VPCNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vpcs">client.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCListParams">VPCListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCList">VPCList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /vpcs/{vpc_id}">client.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vpcs/{vpc_id}">client.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPCService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs">vpcs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#VPC">VPC</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Operations

Methods:

- <code title="get /vpcs/operations/{operation_id}">client.VPCs.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vpcs#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FirewallRule

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleEndpointParam">FirewallRuleEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRule">FirewallRule</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleEndpoint">FirewallRuleEndpoint</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleList">FirewallRuleList</a>

Methods:

- <code title="post /vpcs/{vpc_id}/firewall_rules">client.FirewallRule.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleNewParams">FirewallRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.FirewallRule.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleUpdateParams">FirewallRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vpcs/{vpc_id}/firewall_rules">client.FirewallRule.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleList">FirewallRuleList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.FirewallRule.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.FirewallRule.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule">firewall_rule</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Operations

Methods:

- <code title="get /vpcs/{vpc_id}/firewall_rules/operations/{operation_id}">client.FirewallRule.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/firewall_rule#OperationService.Operations">Operations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Volumes

Methods:

- <code title="post /vms/{vm_id}/volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes">volumes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/volumes#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
