# Re:Earth CMS MCP Server

A Model Context Protocol (MCP) server for Re:Earth CMS API, built with Go.

## Project Structure

```console
.
├── main.go                      # Entry point, tool registration
├── internal                     # Internal application code
├── .github/workflows/           # GitHub Actions workflows
└── .goreleaser.yaml             # GoReleaser configuration
```

## How to Build

### Requirements

- Go 1.25+

### Build

```bash
go build .
```

### Build with Version

```bash
go build -ldflags "-X github.com/KeisukeYamashita/reearth-cms-mcp/internal/version.Version=v1.0.0" .
```

## How to Run

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `REEARTH_CMS_TOKEN` | Yes | API token for authentication |
| `REEARTH_CMS_WORKSPACE_ID` | Yes | Workspace ID to operate on |
| `REEARTH_CMS_BASE_URL` | No | Base URL (default: `https://api.cms.reearth.io/api`) |

## Available Tools

| Tool | Description |
|------|-------------|
| `comment_to_asset` | Add a comment to an asset |
| `comment_to_item` | Add a comment to an item |
| `create_item` | Create a new item |
| `delete_item` | Delete an item |
| `get_asset` | Get a specific asset by ID |
| `get_item` | Get a specific item by ID |
| `get_model` | Get a specific model by ID |
| `list_items` | List items in a model |
| `list_models` | List all models in a project |
| `update_item` | Update an existing item |

## Best Practices Summary

### Development Guidelines

- **Write idiomatic Go code** following standard conventions and patterns
- **Use interface-driven development** with explicit dependency injection
- **Prefer composition over inheritance** with small, purpose-specific interfaces
- **Write short, focused functions** with single responsibility
- **Handle errors explicitly** using wrapped errors for traceability
- **Avoid global state** and use constructor functions for dependency injection
- **Sort functions, methods, tools and types alphabetically** for easier navigation

### Testing Strategy

- **Write comprehensive unit tests** using table-driven patterns
- **Mock external interfaces** cleanly using generated or handwritten mocks
- **Separate fast unit tests** from slower integration and E2E tests
- **Ensure test coverage** for every exported function with behavioral checks
- **Use parallel execution** where appropriate to speed up test runs

### Observability and Monitoring

- **Use OpenTelemetry** for distributed tracing, metrics, and structured logging
- **Propagate context** across all service boundaries (HTTP, gRPC, DB, external APIs)
- **Record important attributes** like request parameters, user ID, and error messages
- **Monitor key metrics** including request latency, throughput, error rate, and resource usage
- **Use structured logging** with JSON format for better observability tool integration

### Security and Resilience

- **Apply input validation** rigorously, especially on external inputs
- **Use secure defaults** for JWT, cookies, and configuration settings
- **Implement retries, exponential backoff, and timeouts** on all external calls
- **Use circuit breakers and rate limiting** for service protection
- **Isolate sensitive operations** with clear permission boundaries

### Performance and Concurrency

- **Use goroutines safely** with proper synchronization mechanisms
- **Implement goroutine cancellation** using context propagation
- **Minimize allocations** and profile before optimizing
- **Use benchmarks** to track performance regressions
- **Guard shared state** with channels or sync primitives
