# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationKind">OperationKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationList">OperationList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationStatus">OperationStatus</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationType">OperationType</a>

Methods:

- <code title="get /v1/operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationListParams">OperationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
