package mcpserver

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/hassaku63/how-to-reading-technical-paper/internal/version"
)

// RunStdio starts an MCP server over stdio using the official Go SDK.
// No tools, prompts, or resources are registered yet.
func RunStdio(ctx context.Context) error {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "paper-reading-mcp",
		Version: version.Version,
	}, nil)

	// Run blocks until the client disconnects or the context is canceled.
	// StdioTransport uses the current process's stdin/stdout.
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		// The SDK returns context.Canceled when ctx is canceled; log for clarity.
		log.Printf("MCP server stopped: %v", err)
		return err
	}
	return nil
}
