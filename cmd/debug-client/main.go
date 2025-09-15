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

func printJSON(title string, data interface{}) error {
	log.Printf("\n=== %s ===\n", title)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal %s: %w", title, err)
	}
	fmt.Println(string(jsonData))
	return nil
}

type ListTarget string

func (lt ListTarget) String() string {
	return string(lt)
}

const (
	LIST_TARGET_ALL       ListTarget = "all"
	LIST_TARGET_PROMPTS   ListTarget = "prompts"
	LIST_TARGET_RESOURCES ListTarget = "resources"
	LIST_TARGET_TOOLS     ListTarget = "tools"
)

func validateListTarget(listTarget string) error {
	lt := ListTarget(listTarget)
	if lt != LIST_TARGET_ALL && lt != LIST_TARGET_PROMPTS && lt != LIST_TARGET_RESOURCES && lt != LIST_TARGET_TOOLS {
		return fmt.Errorf("invalid list target: %s", listTarget)
	}
	return nil
}

func main() {
	// Define command line flags
	var (
		serverCmd  = flag.String("server", "", "Server command to execute (required)")
		verbose    = flag.Bool("verbose", false, "Enable verbose output")
		timeout    = flag.Duration("timeout", 0, "Connection timeout (0 = no timeout)")
		help       = flag.Bool("help", false, "Show help message")
		listTarget = flag.String("list", "all", "What to list: all, prompts, resources, tools (default: all)")
	)

	flag.Parse()

	// Show help if requested
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Validate required flags
	if *serverCmd == "" {
		log.Fatalf("Error: -server flag is required\n\n")
	}

	if err := validateListTarget(*listTarget); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("MCP Debug Client Starting...")
	if *verbose {
		log.Printf("Server command: %s\n", *serverCmd)
		log.Printf("Verbose mode: enabled\n")
		log.Printf("List target: %s\n", *listTarget)
		if *timeout > 0 {
			log.Printf("Timeout: %v\n", *timeout)
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
		log.Printf("Executing command: %s\n", *serverCmd)
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

	log.Println("Connected to MCP server successfully")
	if *verbose {
		log.Println("Connection established via stdio transport")
	}

	// Step 1: List requested primitives
	log.Printf("1. Listing %s...\n", *listTarget)

	// List Prompts
	target := ListTarget(*listTarget)
	shouldListPrompts := target == LIST_TARGET_ALL || target == LIST_TARGET_PROMPTS
	if shouldListPrompts {
		if *verbose {
			log.Println("1.1. Listing Prompts...")
		}
		prompts, err := session.ListPrompts(ctx, &mcp.ListPromptsParams{})
		if err != nil {
			log.Printf("Warning: Failed to list prompts: %v", err)
		} else {
			if err := printJSON("Prompts", prompts); err != nil {
				log.Fatalf("Failed to print prompts: %v", err)
			}
			if *verbose {
				log.Printf("Found %d prompts\n", len(prompts.Prompts))
			}
		}
	}

	// List Resources
	shouldListResources := target == LIST_TARGET_ALL || target == LIST_TARGET_RESOURCES
	if shouldListResources {
		if *verbose {
			log.Println("1.2. Listing Resources...")
		}
		resources, err := session.ListResources(ctx, &mcp.ListResourcesParams{})
		if err != nil {
			log.Printf("Warning: Failed to list resources: %v", err)
		} else {
			if err := printJSON("Resources", resources); err != nil {
				log.Fatalf("Failed to print resources: %v", err)
			}
			if *verbose {
				log.Printf("Found %d resources\n", len(resources.Resources))
			}
		}
	}

	// List Tools
	shouldListTools := target == LIST_TARGET_ALL || target == LIST_TARGET_TOOLS
	if shouldListTools {
		if *verbose {
			log.Println("1.3. Listing Tools...")
		}
		tools, err := session.ListTools(ctx, &mcp.ListToolsParams{})
		if err != nil {
			log.Printf("Warning: Failed to list tools: %v", err)
		} else {
			if err := printJSON("Tools", tools); err != nil {
				log.Fatalf("Failed to print tools: %v", err)
			}
			if *verbose {
				log.Printf("Found %d tools\n", len(tools.Tools))
			}
		}
	}

	log.Println("2. Shutdown complete")
	if *verbose {
		log.Println("All operations completed successfully")
	}
	log.Println("MCP Debug Client finished successfully")
}
