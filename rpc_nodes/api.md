# RPCNodes

## Flex

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchain">FlexBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainList">FlexBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexList">FlexList</a>

Methods:

- <code title="post /v1/rpc_nodes/flex">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexNewParams">FlexNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexUpdateParams">FlexUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/flex">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexListParams">FlexListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/rpc_nodes/flex/{node_id}">client.RPCNodes.Flex.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Flex">Flex</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/flex/blockchains">client.RPCNodes.Flex.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchainListParams">FlexBlockchainListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#FlexBlockchain">FlexBlockchain</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Dedicated

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchain">DedicatedBlockchain</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainList">DedicatedBlockchainList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedList">DedicatedList</a>

Methods:

- <code title="get /v1/rpc_nodes/dedicated">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedListParams">DedicatedListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/rpc_nodes/dedicated/{node_id}">client.RPCNodes.Dedicated.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#Dedicated">Dedicated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Blockchains

Methods:

- <code title="get /v1/rpc_nodes/dedicated/blockchains">client.RPCNodes.Dedicated.Blockchains.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchainListParams">DedicatedBlockchainListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes">rpc_nodes</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/rpc_nodes#DedicatedBlockchain">DedicatedBlockchain</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
