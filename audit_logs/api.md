# AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLog">AuditLog</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogActor">AuditLogActor</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogList">AuditLogList</a>
- <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogType">AuditLogType</a>

Methods:

- <code title="get /v1/audit_logs">client.AuditLogs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogListParams">AuditLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/packages/pagination#Cursor">Cursor</a>[<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLog">AuditLog</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/audit_logs/{audit_log_id}">client.AuditLogs.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditLogID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs">audit_logs</a>.<a href="https://pkg.go.dev/github.com/nirvana-labs/nirvana-go/audit_logs#AuditLog">AuditLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
