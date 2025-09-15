package mcpserver

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/hassaku63/how-to-reading-technical-paper/internal/prompts"
	"github.com/hassaku63/how-to-reading-technical-paper/internal/resources"
	"github.com/hassaku63/how-to-reading-technical-paper/internal/tools/slash"
	"github.com/hassaku63/how-to-reading-technical-paper/internal/version"
)

// RunStdio starts an MCP server over stdio using the official Go SDK.
// No tools, prompts, or resources are registered yet.
func RunStdio(ctx context.Context) error {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "paper-reading-mcp",
		Version: version.Version,
	}, nil)

	// Logging middleware
	// server.AddReceivingMiddleware(withLogging)

	// Register resources from embedded assets.
	resources.RegisterTemplatesScreeningChecklist(server)
	resources.RegisterTemplatesReadingNotes(server)
	resources.RegisterPatternsCommonStructures(server)
	resources.RegisterCriteriaPaperEvaluation(server)
	resources.RegisterMethodologySurveyGuide(server)
	resources.RegisterMethodologyImplementationGuide(server)
	resources.RegisterActionableSurveySteps(server)
	resources.RegisterActionableImplementationSteps(server)

	// Register prompts
	prompts.RegisterSurveyReadingFlow(server)
	prompts.RegisterImplementationReadingFlow(server)
	prompts.RegisterOutputGeneration(server)
	prompts.RegisterCriticalAnalysis(server)
	prompts.RegisterComparisonMatrix(server)

	// Register tools
	slash.RegisterExportClaudeCommandsTool(server)
	slash.RegisterPreviewClaudeCommandsTool(server)

	// Run blocks until the client disconnects or the c ontext is canceled.
	// StdioTransport uses the current process's stdin/stdout.
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		// The SDK returns context.Canceled when ctx is canceled; log for clarity.
		log.Printf("MCP server stopped: %v", err)
		return err
	}
	return nil
}
