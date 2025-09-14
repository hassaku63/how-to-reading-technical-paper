package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func printJSON(title string, data interface{}) {
	fmt.Printf("\n=== %s ===\n", title)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	// Define command line flags
	var (
		serverCmd = flag.String("server", "", "Server command to execute (required)")
		verbose   = flag.Bool("verbose", false, "Enable verbose output")
		timeout   = flag.Duration("timeout", 0, "Connection timeout (0 = no timeout)")
		help      = flag.Bool("help", false, "Show help message")
	)

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "MCP Debug Client - Test MCP servers using stdio transport\n\n")
		fmt.Fprintf(os.Stderr, "This is a debug-only client for testing MCP server implementations.\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -server ./bin/server\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -server ./bin/server -verbose\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -server 'go run ./cmd/server' -timeout 30s\n", os.Args[0])
	}

	flag.Parse()

	// Show help if requested
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Validate required flags
	if *serverCmd == "" {
		fmt.Fprintf(os.Stderr, "Error: -server flag is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("MCP Debug Client Starting...")
	if *verbose {
		fmt.Printf("Server command: %s\n", *serverCmd)
		fmt.Printf("Verbose mode: enabled\n")
		if *timeout > 0 {
			fmt.Printf("Timeout: %v\n", *timeout)
		}
	}

	// Create context with optional timeout
	ctx := context.Background()
	if *timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	// Create a new client using the official Go SDK
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "mcp-debug-client",
		Version: "1.0.0",
	}, nil)

	// Connect to the server using CommandTransport (stdio)
	// Parse the server command to handle shell commands properly
	var cmd *exec.Cmd
	if *verbose {
		fmt.Printf("Executing command: %s\n", *serverCmd)
	}

	// For simple commands, use exec.Command directly
	// For complex shell commands, use shell execution
	cmd = exec.Command("sh", "-c", *serverCmd)

	transport := &mcp.CommandTransport{
		Command: cmd,
	}

	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer session.Close()

	fmt.Println("Connected to MCP server successfully")
	if *verbose {
		fmt.Println("Connection established via stdio transport")
	}

	// Step 1: List all primitives
	fmt.Println("\n1. Listing primitives...")

	// List Prompts
	if *verbose {
		fmt.Println("\n1.1. Listing Prompts...")
	}
	prompts, err := session.ListPrompts(ctx, &mcp.ListPromptsParams{})
	if err != nil {
		log.Printf("Warning: Failed to list prompts: %v", err)
	} else {
		printJSON("Prompts", prompts)
		if *verbose {
			fmt.Printf("Found %d prompts\n", len(prompts.Prompts))
		}
	}

	// List Resources
	if *verbose {
		fmt.Println("\n1.2. Listing Resources...")
	}
	resources, err := session.ListResources(ctx, &mcp.ListResourcesParams{})
	if err != nil {
		log.Printf("Warning: Failed to list resources: %v", err)
	} else {
		printJSON("Resources", resources)
		if *verbose {
			fmt.Printf("Found %d resources\n", len(resources.Resources))
		}
	}

	// List Tools
	if *verbose {
		fmt.Println("\n1.3. Listing Tools...")
	}
	tools, err := session.ListTools(ctx, &mcp.ListToolsParams{})
	if err != nil {
		log.Printf("Warning: Failed to list tools: %v", err)
	} else {
		printJSON("Tools", tools)
		if *verbose {
			fmt.Printf("Found %d tools\n", len(tools.Tools))
		}
	}

	fmt.Println("\n2. Shutdown complete")
	if *verbose {
		fmt.Println("All operations completed successfully")
	}
	fmt.Println("MCP Debug Client finished successfully")
}
