# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#ResourceStatus">ResourceStatus</a>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/user">user</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/user#User">User</a>

Methods:

- <code title="get /v1/user">client.User.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/user#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/user">user</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/user#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyList">APIKeyList</a>

Methods:

- <code title="post /v1/api_keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyNewParams">APIKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/api_keys/{api_key_id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyUpdateParams">APIKeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api_keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyList">APIKeyList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api_keys/{api_key_id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/api_keys/{api_key_id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationKind">OperationKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationList">OperationList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationStatus">OperationStatus</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationType">OperationType</a>

Methods:

- <code title="get /v1/operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationList">OperationList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#CPUConfigParam">CPUConfigParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#MemoryConfigParam">MemoryConfigParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#SSHKeyParam">SSHKeyParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#CPUConfig">CPUConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#MemoryConfig">MemoryConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMList">VMList</a>

Methods:

- <code title="post /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMNewParams">VMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMUpdateParams">VMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMList">VMList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/compute/vms/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityNewParams">VMAvailabilityNewParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityUpdateParams">VMAvailabilityUpdateParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Volumes

Methods:

- <code title="get /v1/compute/vms/{vm_id}/volumes">client.Compute.VMs.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeList">VolumeList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /v1/compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMOSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#StorageType">StorageType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeList">VolumeList</a>

Methods:

- <code title="post /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeUpdateParams">VolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeList">VolumeList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/compute/volumes/availability">client.Compute.Volumes.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeAvailabilityNewParams">VolumeAvailabilityNewParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/volumes/{volume_id}/availability">client.Compute.Volumes.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeAvailabilityUpdateParams">VolumeAvailabilityUpdateParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Networking

## VPCs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#Subnet">Subnet</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPC">VPC</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCList">VPCList</a>

Methods:

- <code title="post /v1/networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCNewParams">VPCNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCUpdateParams">VPCUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCList">VPCList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/vpcs/{vpc_id}">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPC">VPC</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/networking/vpcs/availability">client.Networking.VPCs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCAvailabilityNewParams">VPCAvailabilityNewParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/networking/vpcs/{vpc_id}/availability">client.Networking.VPCs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCAvailabilityUpdateParams">VPCAvailabilityUpdateParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FirewallRules

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleList">FirewallRuleList</a>

Methods:

- <code title="post /v1/networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleNewParams">FirewallRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleUpdateParams">FirewallRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleList">FirewallRuleList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RPCNodes

## Flex

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlex">RPCNodesFlex</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlexBlockchain">RPCNodesFlexBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlexBlockchainList">RPCNodesFlexBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlexList">RPCNodesFlexList</a>

Methods:

- <code title="get /v1/rpc_nodes/flex">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlexList">RPCNodesFlexList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlex">RPCNodesFlex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/flex/blockchains">client.RPCNodes.Flex.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesFlexBlockchainList">RPCNodesFlexBlockchainList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Dedicated

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicated">RPCNodesDedicated</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicatedBlockchain">RPCNodesDedicatedBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicatedBlockchainList">RPCNodesDedicatedBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicatedList">RPCNodesDedicatedList</a>

Methods:

- <code title="get /v1/rpc_nodes/dedicated">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicatedList">RPCNodesDedicatedList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/dedicated/{node_id}">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicated">RPCNodesDedicated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/dedicated/blockchains">client.RPCNodes.Dedicated.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#RPCNodesDedicatedBlockchainList">RPCNodesDedicatedBlockchainList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Connect

## Flux

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect">connect</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#ConnectFlux">ConnectFlux</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect">connect</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#ConnectFluxList">ConnectFluxList</a>

Methods:

- <code title="get /v1/connect/flux">client.Connect.Flux.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#FluxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect">connect</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#ConnectFluxList">ConnectFluxList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/connect/flux/{flux_id}">client.Connect.Flux.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#FluxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fluxID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect">connect</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/connect#ConnectFlux">ConnectFlux</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
