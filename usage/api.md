# Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#Usage">Usage</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageDimension">UsageDimension</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageDimensionLeaf">UsageDimensionLeaf</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageList">UsageList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageResourceType">UsageResourceType</a>

Methods:

- <code title="get /v1/usage">client.Usage.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageListParams">UsageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#Usage">Usage</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/usage/{resource_id}">client.Usage.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, resourceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage">usage</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v2/usage#Usage">Usage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
