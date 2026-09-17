# Compute

## VMs

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#SSHKeyRequestParam">SSHKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#CPUConfig">CPUConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#MemoryConfig">MemoryConfig</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#OSImage">OSImage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VM">VM</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMList">VMList</a>

Methods:

- <code title="post /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMNewParams">VMNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMUpdateParams">VMUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMListParams">VMListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VM">VM</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/vms/{vm_id}">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VM">VM</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/compute/vms/{vm_id}/restart">client.Compute.VMs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMService.Restart">Restart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/compute/vms/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMAvailabilityNewParams">VMAvailabilityNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /v1/compute/vms/{vm_id}/availability">client.Compute.VMs.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMAvailabilityUpdateParams">VMAvailabilityUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Cost

Methods:

- <code title="post /v1/compute/vms/cost">client.Compute.VMs.Cost.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMCostService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMCostNewParams">VMCostNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared#CostQuote">CostQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/vms/{vm_id}/cost">client.Compute.VMs.Cost.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMCostService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMCostUpdateParams">VMCostUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared#CostQuoteUpdate">CostQuoteUpdate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Volumes

Methods:

- <code title="get /v1/compute/vms/{vm_id}/volumes">client.Compute.VMs.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMVolumeListParams">VMVolumeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#Volume">Volume</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### OSImages

Methods:

- <code title="get /v1/compute/vms/os_images">client.Compute.VMs.OSImages.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMOSImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMOSImageListParams">VMOSImageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#OSImage">OSImage</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Metrics

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricPoint">VMMetricPoint</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricSeries">VMMetricSeries</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetrics">VMMetrics</a>

Methods:

- <code title="get /v1/compute/vms/{vm_id}/metrics">client.Compute.VMs.Metrics.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricListParams">VMMetricListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetrics">VMMetrics</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### MetricDescriptors

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricDescriptor">VMMetricDescriptor</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricDescriptorList">VMMetricDescriptorList</a>

Methods:

- <code title="get /v1/compute/metric_descriptors">client.Compute.VMs.MetricDescriptors.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricDescriptorService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VMMetricDescriptorList">VMMetricDescriptorList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeType">VolumeType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#Volume">Volume</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeKind">VolumeKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeList">VolumeList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeType">VolumeType</a>

Methods:

- <code title="post /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeNewParams">VolumeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeUpdateParams">VolumeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/volumes">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeListParams">VolumeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#Volume">Volume</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/compute/volumes/{volume_id}/attach">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeAttachParams">VolumeAttachParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/compute/volumes/{volume_id}/detach">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.Detach">Detach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/compute/volumes/{volume_id}">client.Compute.Volumes.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Availability

Methods:

- <code title="post /v1/compute/volumes/availability">client.Compute.Volumes.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeAvailabilityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeAvailabilityNewParams">VolumeAvailabilityNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /v1/compute/volumes/{volume_id}/availability">client.Compute.Volumes.Availability.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeAvailabilityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeAvailabilityUpdateParams">VolumeAvailabilityUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Cost

Methods:

- <code title="post /v1/compute/volumes/cost">client.Compute.Volumes.Cost.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeCostService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeCostNewParams">VolumeCostNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared#CostQuote">CostQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/compute/volumes/{volume_id}/cost">client.Compute.Volumes.Cost.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeCostService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute">compute</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/compute#VolumeCostUpdateParams">VolumeCostUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/shared#CostQuoteUpdate">CostQuoteUpdate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
