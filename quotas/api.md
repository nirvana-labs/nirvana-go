# Quotas

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#Quota">Quota</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaCompute">QuotaCompute</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaDimensionDetail">QuotaDimensionDetail</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaList">QuotaList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaNetworking">QuotaNetworking</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaNKS">QuotaNKS</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaStorage">QuotaStorage</a>

Methods:

- <code title="get /v1/quotas">client.Quotas.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaListParams">QuotaListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#Quota">Quota</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/quotas/{region}">client.Quotas.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, region <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#QuotaGetParamsRegion">QuotaGetParamsRegion</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas">quotas</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/quotas#Quota">Quota</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
