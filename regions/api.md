# Regions

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#Region">Region</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#RegionAvailability">RegionAvailability</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#RegionList">RegionList</a>

Methods:

- <code title="get /v1/regions">client.Regions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#RegionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#RegionListParams">RegionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#Region">Region</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/regions/{name}">client.Regions.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#RegionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions">regions</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/regions#Region">Region</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
