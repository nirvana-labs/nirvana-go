# Shared Params Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#RegionName">RegionName</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/shared#Pagination">Pagination</a>
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
- <code title="get /v1/api_keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKeyListParams">APIKeyListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys">api_keys</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/api_keys#APIKey">APIKey</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
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

- <code title="get /v1/operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationListParams">OperationListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectList">ProjectList</a>

Methods:

- <code title="post /v1/projects">client.Projects.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/projects/{project_id}">client.Projects.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/projects">client.Projects.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#Project">Project</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/projects/{project_id}">client.Projects.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/projects/{project_id}">client.Projects.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects">projects</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/projects#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#CPUConfigRequestParam">CPUConfigRequestParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#MemoryConfigRequestParam">MemoryConfigRequestParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#SSHKeyRequestParam">SSHKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#CPUConfig">CPUConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#MemoryConfig">MemoryConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMList">VMList</a>

Methods:

- <code title="post /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMNewParams">VMNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMUpdateParams">VMUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMListParams">VMListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/compute/vms/{vm_id}/restart">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMService.Restart">Restart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/compute/vms/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityNewParams">VMAvailabilityNewParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMAvailabilityUpdateParams">VMAvailabilityUpdateParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Volumes

Methods:

- <code title="get /v1/compute/vms/{vm_id}/volumes">client.Compute.VMs.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMVolumeListParams">VMVolumeListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /v1/compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMOSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VMOSImageListParams">VMOSImageListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#OSImage">OSImage</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeList">VolumeList</a>

Methods:

- <code title="post /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeUpdateParams">VolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#VolumeListParams">VolumeListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/compute#Volume">Volume</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
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
- <code title="get /v1/networking/vpcs">client.Networking.VPCs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPCListParams">VPCListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#VPC">VPC</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
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
- <code title="get /v1/networking/vpcs/{vpc_id}/firewall_rules">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleListParams">FirewallRuleListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/vpcs/{vpc_id}/firewall_rules/{firewall_rule_id}">client.Networking.FirewallRules.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vpcID <a href="https://pkg.go.dev/builtin#string">string</a>, firewallRuleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#FirewallRule">FirewallRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connect

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionAWSConfigRequestParam">ConnectConnectionAWSConfigRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnection">ConnectConnection</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionAWSConfig">ConnectConnectionAWSConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionList">ConnectConnectionList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectRoute">ConnectRoute</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectRouteList">ConnectRouteList</a>

### Connections

Methods:

- <code title="post /v1/networking/connect/connections">client.Networking.Connect.Connections.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionNewParams">ConnectConnectionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/networking/connect/connections/{connection_id}">client.Networking.Connect.Connections.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionUpdateParams">ConnectConnectionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/connect/connections">client.Networking.Connect.Connections.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionListParams">ConnectConnectionListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnection">ConnectConnection</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/networking/connect/connections/{connection_id}">client.Networking.Connect.Connections.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/networking/connect/connections/{connection_id}">client.Networking.Connect.Connections.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectConnection">ConnectConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Routes

Methods:

- <code title="get /v1/networking/connect/routes">client.Networking.Connect.Routes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectRouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectRouteListParams">ConnectRouteListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking">networking</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/networking#ConnectRoute">ConnectRoute</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RPCNodes

## Flex

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchain">FlexBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainList">FlexBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexList">FlexList</a>

Methods:

- <code title="post /v1/rpc_nodes/flex">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexNewParams">FlexNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexUpdateParams">FlexUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/flex">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexListParams">FlexListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/flex/blockchains">client.RPCNodes.Flex.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainListParams">FlexBlockchainListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchain">FlexBlockchain</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Dedicated

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchain">DedicatedBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainList">DedicatedBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedList">DedicatedList</a>

Methods:

- <code title="get /v1/rpc_nodes/dedicated">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedListParams">DedicatedListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/dedicated/{node_id}">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/dedicated/blockchains">client.RPCNodes.Dedicated.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainListParams">DedicatedBlockchainListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchain">DedicatedBlockchain</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vektor

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AssetID">AssetID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AssetIDOrAddressEVMOrAssetSymbol">AssetIDOrAddressEVMOrAssetSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AssetSymbol">AssetSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockNumber">BlockNumber</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainID">BlockchainID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainIDOrBlockchainSymbol">BlockchainIDOrBlockchainSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainName">BlockchainName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainSymbol">BlockchainSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ChainType">ChainType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Decimal">Decimal</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#EVMChainDataParam">EVMChainDataParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#HexString">HexString</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendBorrowMarketID">LendBorrowMarketID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NetworkName">NetworkName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NFTParam">NFTParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NFTCollectionParam">NFTCollectionParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Timestamp">Timestamp</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#TimestampOrBlockNumberUnionParam">TimestampOrBlockNumberUnionParam</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VenueIDOrVenueSymbol">VenueIDOrVenueSymbol</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AddressEVM">AddressEVM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#APY">APY</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Asset">Asset</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AssetID">AssetID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#AssetSymbol">AssetSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Balance">Balance</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockNumber">BlockNumber</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Blockchain">Blockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainID">BlockchainID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainName">BlockchainName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BlockchainSymbol">BlockchainSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Blockstamp">Blockstamp</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccount">BorrowAccount</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarket">BorrowMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPosition">BorrowPosition</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BridgeQuote">BridgeQuote</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyQuote">BuyQuote</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ChainType">ChainType</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Decimal">Decimal</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ErrorListOutput">ErrorListOutput</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#EVMChainData">EVMChainData</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionError">ExecutionError</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionEVMTransactionEIP1559Payload">ExecutionEVMTransactionEIP1559Payload</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionEVMTransactionState">ExecutionEVMTransactionState</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionID">ExecutionID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionState">ExecutionState</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepID">ExecutionStepID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#HexString">HexString</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#IncentivizeMarket">IncentivizeMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendBorrowAPYs">LendBorrowAPYs</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendBorrowMarketID">LendBorrowMarketID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarket">LendMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPosition">LendPosition</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockMarket">LockMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockPosition">LockPosition</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPool">LPPool</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolSolidlyAttributes">LPPoolSolidlyAttributes</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolUniswapV3Attributes">LPPoolUniswapV3Attributes</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPosition">LPPosition</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionAttributes">LPPositionAttributes</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPQuote">LPQuote</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NetworkName">NetworkName</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NFT">NFT</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#NFTCollection">NFTCollection</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#OffChainHistoricalRange">OffChainHistoricalRange</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#OnChainHistoricalRange">OnChainHistoricalRange</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfo0x">QuoteInfo0x</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfo0xFill">QuoteInfo0xFill</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfo0xRoute">QuoteInfo0xRoute</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfo0xToken">QuoteInfo0xToken</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfoCurve">QuoteInfoCurve</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfoUniswapV2">QuoteInfoUniswapV2</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#QuoteInfoUniswapV3">QuoteInfoUniswapV3</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLendBorrowMarket">RegistryLendBorrowMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLPPool">RegistryLPPool</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellQuote">SellQuote</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Timestamp">Timestamp</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VektorError">VektorError</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VektorErrorList">VektorErrorList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Venue">Venue</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VenueID">VenueID</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VenueSymbol">VenueSymbol</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteMarket">VoteMarket</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteReward">VoteReward</a>

## Registry

### Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryAssetListResponse">RegistryAssetListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/assets">client.Vektor.Registry.Assets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryAssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryAssetListParams">RegistryAssetListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryAssetListResponse">RegistryAssetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBlockchainListResponse">RegistryBlockchainListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/blockchains">client.Vektor.Registry.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBlockchainListParams">RegistryBlockchainListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBlockchainListResponse">RegistryBlockchainListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Venues

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryVenueListResponse">RegistryVenueListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/venues">client.Vektor.Registry.Venues.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryVenueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryVenueListParams">RegistryVenueListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryVenueListResponse">RegistryVenueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Errors

Methods:

- <code title="post /v1/vektor/registry/errors">client.Vektor.Registry.Errors.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryErrorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryErrorListParams">RegistryErrorListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ErrorListOutput">ErrorListOutput</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### LendMarkets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLendMarketListResponse">RegistryLendMarketListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/lend/markets">client.Vektor.Registry.LendMarkets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLendMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLendMarketListParams">RegistryLendMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLendMarketListResponse">RegistryLendMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### BorrowMarkets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBorrowMarketListResponse">RegistryBorrowMarketListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/borrow/markets">client.Vektor.Registry.BorrowMarkets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBorrowMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBorrowMarketListParams">RegistryBorrowMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryBorrowMarketListResponse">RegistryBorrowMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### LPPools

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLPPoolListResponse">RegistryLPPoolListResponse</a>

Methods:

- <code title="post /v1/vektor/registry/lp/pools">client.Vektor.Registry.LPPools.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLPPoolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLPPoolListParams">RegistryLPPoolListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#RegistryLPPoolListResponse">RegistryLPPoolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Balances

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListResponse">BalanceListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListHistoricalResponse">BalanceListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/balances">client.Vektor.Balances.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListParams">BalanceListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListResponse">BalanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/balances/historical">client.Vektor.Balances.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListHistoricalParams">BalanceListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BalanceListHistoricalResponse">BalanceListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Prices

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListResponse">PriceListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListHistoricalResponse">PriceListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/prices">client.Vektor.Prices.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListParams">PriceListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListResponse">PriceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/prices/historical">client.Vektor.Prices.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListHistoricalParams">PriceListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#PriceListHistoricalResponse">PriceListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lend

### Markets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListResponse">LendMarketListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListHistoricalResponse">LendMarketListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/lend/markets">client.Vektor.Lend.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListParams">LendMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListResponse">LendMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/lend/markets/historical">client.Vektor.Lend.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListHistoricalParams">LendMarketListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendMarketListHistoricalResponse">LendMarketListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListResponse">LendPositionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListHistoricalResponse">LendPositionListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/lend/positions">client.Vektor.Lend.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListParams">LendPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListResponse">LendPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/lend/positions/historical">client.Vektor.Lend.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListHistoricalParams">LendPositionListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendPositionListHistoricalResponse">LendPositionListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lend

Methods:

- <code title="post /v1/vektor/lend/lend">client.Vektor.Lend.Lend.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendLendService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendLendNewParams">LendLendNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Withdraw

Methods:

- <code title="post /v1/vektor/lend/withdraw">client.Vektor.Lend.Withdraw.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendWithdrawService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendWithdrawNewParams">LendWithdrawNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### SetCollateral

Methods:

- <code title="post /v1/vektor/lend/set_collateral">client.Vektor.Lend.SetCollateral.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendSetCollateralService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LendSetCollateralNewParams">LendSetCollateralNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Borrow

### Markets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListResponse">BorrowMarketListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListHistoricalResponse">BorrowMarketListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/borrow/markets">client.Vektor.Borrow.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListParams">BorrowMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListResponse">BorrowMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/borrow/markets/historical">client.Vektor.Borrow.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListHistoricalParams">BorrowMarketListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowMarketListHistoricalResponse">BorrowMarketListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListResponse">BorrowPositionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListHistoricalResponse">BorrowPositionListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/borrow/positions">client.Vektor.Borrow.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListParams">BorrowPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListResponse">BorrowPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/borrow/positions/historical">client.Vektor.Borrow.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListHistoricalParams">BorrowPositionListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowPositionListHistoricalResponse">BorrowPositionListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListResponse">BorrowAccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListHistoricalResponse">BorrowAccountListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/borrow/accounts">client.Vektor.Borrow.Accounts.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListParams">BorrowAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListResponse">BorrowAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/borrow/accounts/historical">client.Vektor.Borrow.Accounts.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListHistoricalParams">BorrowAccountListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowAccountListHistoricalResponse">BorrowAccountListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Borrow

Methods:

- <code title="post /v1/vektor/borrow/borrow">client.Vektor.Borrow.Borrow.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowBorrowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowBorrowNewParams">BorrowBorrowNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Repay

Methods:

- <code title="post /v1/vektor/borrow/repay">client.Vektor.Borrow.Repay.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowRepayService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BorrowRepayNewParams">BorrowRepayNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LP

### Pools

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListResponse">LPPoolListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListHistoricalResponse">LPPoolListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/lp/pools">client.Vektor.LP.Pools.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListParams">LPPoolListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListResponse">LPPoolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/lp/pools/historical">client.Vektor.LP.Pools.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListHistoricalParams">LPPoolListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPoolListHistoricalResponse">LPPoolListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListResponse">LPPositionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListHistoricalResponse">LPPositionListHistoricalResponse</a>

Methods:

- <code title="post /v1/vektor/lp/positions">client.Vektor.LP.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListParams">LPPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListResponse">LPPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/lp/positions/historical">client.Vektor.LP.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionService.ListHistorical">ListHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListHistoricalParams">LPPositionListHistoricalParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPPositionListHistoricalResponse">LPPositionListHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### DepositQuote

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPDepositQuoteNewResponse">LPDepositQuoteNewResponse</a>

Methods:

- <code title="post /v1/vektor/lp/deposit_quote">client.Vektor.LP.DepositQuote.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPDepositQuoteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPDepositQuoteNewParams">LPDepositQuoteNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPDepositQuoteNewResponse">LPDepositQuoteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### WithdrawQuote

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPWithdrawQuoteNewResponse">LPWithdrawQuoteNewResponse</a>

Methods:

- <code title="post /v1/vektor/lp/withdraw_quote">client.Vektor.LP.WithdrawQuote.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPWithdrawQuoteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPWithdrawQuoteNewParams">LPWithdrawQuoteNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LPWithdrawQuoteNewResponse">LPWithdrawQuoteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Buy

### Quotes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyQuoteListResponse">BuyQuoteListResponse</a>

Methods:

- <code title="post /v1/vektor/buy/quotes">client.Vektor.Buy.Quotes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyQuoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyQuoteListParams">BuyQuoteListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyQuoteListResponse">BuyQuoteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Buy

Methods:

- <code title="post /v1/vektor/buy/buy">client.Vektor.Buy.Buy.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyBuyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BuyBuyNewParams">BuyBuyNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sell

### Quotes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellQuoteListResponse">SellQuoteListResponse</a>

Methods:

- <code title="post /v1/vektor/sell/quotes">client.Vektor.Sell.Quotes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellQuoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellQuoteListParams">SellQuoteListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellQuoteListResponse">SellQuoteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Sell

Methods:

- <code title="post /v1/vektor/sell/sell">client.Vektor.Sell.Sell.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellSellService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#SellSellNewParams">SellSellNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Move

Methods:

- <code title="post /v1/vektor/move/move">client.Vektor.Move.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#MoveService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#MoveNewParams">MoveNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Wrap

### Wrap

Methods:

- <code title="post /v1/vektor/wrap/wrap">client.Vektor.Wrap.Wrap.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#WrapWrapService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#WrapWrapNewParams">WrapWrapNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Unwrap

Methods:

- <code title="post /v1/vektor/wrap/unwrap">client.Vektor.Wrap.Unwrap.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#WrapUnwrapService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#WrapUnwrapNewParams">WrapUnwrapNewParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bridge

### Quotes

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BridgeQuoteListResponse">BridgeQuoteListResponse</a>

Methods:

- <code title="post /v1/vektor/bridge/quotes">client.Vektor.Bridge.Quotes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BridgeQuoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BridgeQuoteListParams">BridgeQuoteListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#BridgeQuoteListResponse">BridgeQuoteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Lock

### Markets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockMarketListResponse">LockMarketListResponse</a>

Methods:

- <code title="post /v1/vektor/lock/markets">client.Vektor.Lock.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockMarketListParams">LockMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockMarketListResponse">LockMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockPositionListResponse">LockPositionListResponse</a>

Methods:

- <code title="post /v1/vektor/lock/positions">client.Vektor.Lock.Positions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockPositionListParams">LockPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#LockPositionListResponse">LockPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Vote

### Markets

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteMarketListResponse">VoteMarketListResponse</a>

Methods:

- <code title="post /v1/vektor/vote/markets">client.Vektor.Vote.Markets.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteMarketService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteMarketListParams">VoteMarketListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteMarketListResponse">VoteMarketListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rewards

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteRewardListResponse">VoteRewardListResponse</a>

Methods:

- <code title="post /v1/vektor/vote/rewards">client.Vektor.Vote.Rewards.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteRewardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteRewardListParams">VoteRewardListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#VoteRewardListResponse">VoteRewardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Incentivize

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#IncentivizeListResponse">IncentivizeListResponse</a>

Methods:

- <code title="post /v1/vektor/incentivize/markets">client.Vektor.Incentivize.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#IncentivizeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#IncentivizeListParams">IncentivizeListParams</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#IncentivizeListResponse">IncentivizeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Executions

Methods:

- <code title="get /v1/vektor/executions">client.Vektor.Executions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vektor/executions/{execution_id}">client.Vektor.Executions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#Execution">Execution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Steps

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepGetResponse">ExecutionStepGetResponse</a>

Methods:

- <code title="get /v1/vektor/executions/{execution_id}/steps/{step_id}">client.Vektor.Executions.Steps.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepGetResponse">ExecutionStepGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vektor/executions/{execution_id}/steps/{step_id}/sign">client.Vektor.Executions.Steps.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepService.Sign">Sign</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor">vektor</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/vektor#ExecutionStepSignParams">ExecutionStepSignParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
