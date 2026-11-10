# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#Operation">Operation</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationChanges">OperationChanges</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationDetails">OperationDetails</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationFieldDiff">OperationFieldDiff</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationKind">OperationKind</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationList">OperationList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationStatus">OperationStatus</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationType">OperationType</a>

Methods:

- <code title="get /v1/operations">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationListParams">OperationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#Operation">Operation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/{operation_id}">client.Operations.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#OperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations">operations</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/v3/operations#Operation">Operation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
