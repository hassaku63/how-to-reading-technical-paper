# MCP Server Skeleton Development

This repository includes a minimal MCP Server skeleton in Go. It currently registers no tools, prompts, or resources; it provides a place to wire the official Model Context Protocol Go SDK.

## Layout

- `cmd/server/main.go`: entrypoint, signal handling, invokes server run.
- `internal/mcpserver/server.go`: skeleton server over stdio; TODO to integrate the Go SDK.
- `internal/version/version.go`: central version constant.

## Building and Running

1. Ensure Go is installed (as specified in `go.mod`).
2. Build the server:
   - `go build ./cmd/server`
3. Run the server (stdio transport expected):
   - `./server`

Note: The server is a skeleton and will only emit a banner line, then wait for termination. It does not yet speak the MCP protocol until the SDK is wired.

## Next Steps (SDK wiring)

- Add dependency: `github.com/modelcontextprotocol/go-sdk`
- Create an SDK server instance, register features, and serve over stdio.
- Gradually register Tools/Prompts/Resources from `docs/mcp-specs/*`.

