package slash

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	appassets "github.com/hassaku63/how-to-reading-technical-paper/assets"
	"github.com/hassaku63/how-to-reading-technical-paper/internal/version"
)

// PromptMeta holds metadata for a registered prompt
type PromptMeta struct {
	Prompt    *mcp.Prompt
	AssetPath string
}

// Internal registry of prompts with their metadata
var promptRegistry = make(map[string]*PromptMeta)

// RegisterPromptMeta registers a prompt with its metadata for export
func RegisterPromptMeta(name string, prompt *mcp.Prompt, assetPath string) {
	promptRegistry[name] = &PromptMeta{
		Prompt:    prompt,
		AssetPath: assetPath,
	}
}

// ExportClaudeCommandsArgs represents the input parameters for export_claude_commands tool
type ExportClaudeCommandsArgs struct {
	TargetDir  string `json:"target_dir,omitempty" jsonschema:"title:Target Directory,description:出力先ディレクトリ。省略時は .claude/commands/"`
	NamePrefix string `json:"name_prefix,omitempty" jsonschema:"title:Name Prefix,description:コマンド名に付与する接頭辞（デフォルト: rp-）,default:rp-"`
	DryRun     bool   `json:"dry_run,omitempty" jsonschema:"title:Dry Run,description:trueのときファイル保存せず生成内容のみを返す"`
}

// ExportClaudeCommandsResult represents the output of export_claude_commands tool
type ExportClaudeCommandsResult struct {
	Status       string   `json:"status"`
	Written      bool     `json:"written"`
	Paths        []string `json:"paths,omitempty"`
	CommandCount int      `json:"command_count"`
	Version      string   `json:"version"`
	Hash         string   `json:"hash"`
	Preview      *string  `json:"preview,omitempty"`
}

// PreviewClaudeCommandsArgs represents the input parameters for preview_claude_commands tool
type PreviewClaudeCommandsArgs struct {
	NamePrefix string `json:"name_prefix,omitempty" jsonschema:"title:Name Prefix,description:コマンド名に付与する接頭辞（デフォルト: rp-）,default:rp-"`
}

// PreviewClaudeCommandsResult represents the output of preview_claude_commands tool
type PreviewClaudeCommandsResult struct {
	Status       string                 `json:"status"`
	CommandCount int                    `json:"command_count"`
	Version      string                 `json:"version"`
	Files        []PreviewFileContent   `json:"files"`
}

// PreviewFileContent represents a file in the preview result
type PreviewFileContent struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// placeholderRe matches {arg} patterns to convert to ${N}
var placeholderRe = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)

// extractDefaultFromDescription extracts a default value from an argument description if present.
// Supports patterns like:
// - "デフォルト: 値" / "デフォルト：値" / "デフォルトは 値"
// - "既定: 値" / "既定：値" / "既定は 値"
// - "default: value"
// The value ends before sentence delimiters or closing parentheses.
func extractDefaultFromDescription(description string) (string, bool) {
    if description == "" {
        return "", false
    }

    patterns := []*regexp.Regexp{
        regexp.MustCompile(`(?i)(?:デフォルト|既定)[\s]*[:：][\s]*([^\n。．、,，)）]+)`),
        regexp.MustCompile(`(?i)(?:デフォルト|既定)は[\s]*([^\n。．、,，)）]+)`),
        regexp.MustCompile(`(?i)default[\s]*[:：][\s]*([^\n。．、,，)）]+)`),
    }

    for _, re := range patterns {
        m := re.FindStringSubmatch(description)
        if len(m) >= 2 {
            v := strings.TrimSpace(m[1])
            v = strings.Trim(v, `"'「」『』“”`)
            if v != "" {
                if strings.HasPrefix(v, "例") || strings.HasPrefix(strings.ToLower(v), "e.g") {
                    continue
                }
                return v, true
            }
        }
    }
    return "", false
}

// convertPlaceholders converts {arg} to ${N} based on argument order
func convertPlaceholders(template string, args []*mcp.PromptArgument) string {
	// Create argument position mapping
	argToPos := make(map[string]int)
	pos := 1

	// Required arguments first
	for _, arg := range args {
		if arg.Required {
			argToPos[arg.Name] = pos
			pos++
		}
	}

	// Optional arguments after required ones
	for _, arg := range args {
		if !arg.Required {
			argToPos[arg.Name] = pos
			pos++
		}
	}

	// Replace placeholders
	return placeholderRe.ReplaceAllStringFunc(template, func(match string) string {
		argName := strings.Trim(match, "{}")
		if pos, exists := argToPos[argName]; exists {
			// For optional arguments, add default if available
			for _, arg := range args {
				if arg.Name == argName && !arg.Required {
					if defVal, ok := extractDefaultFromDescription(arg.Description); ok {
						// avoid unbalanced braces in default text
						safe := strings.ReplaceAll(defVal, "}", ")")
						return fmt.Sprintf("${%d:-%s}", pos, safe)
					}
					return fmt.Sprintf("${%d}", pos)
				}
			}
			return fmt.Sprintf("${%d}", pos)
		}
		return match // Leave unchanged if not found
	})
}

// generateArgumentHint creates argument-hint string for frontmatter
func generateArgumentHint(args []*mcp.PromptArgument) string {
	var hints []string

	// Required arguments first
	for _, arg := range args {
		if arg.Required {
			if arg.Title != "" {
				hints = append(hints, fmt.Sprintf("<%s>", arg.Title))
			} else {
				hints = append(hints, fmt.Sprintf("<%s>", arg.Name))
			}
		}
	}

	// Optional arguments after required ones
	for _, arg := range args {
		if !arg.Required {
			if arg.Title != "" {
				hints = append(hints, fmt.Sprintf("[%s]", arg.Title))
			} else {
				hints = append(hints, fmt.Sprintf("[%s]", arg.Name))
			}
		}
	}

	return strings.Join(hints, " ")
}

// generateArgumentComments creates detailed argument documentation for frontmatter
func generateArgumentComments(args []*mcp.PromptArgument) string {
	if len(args) == 0 {
		return ""
	}

	var comments []string
	comments = append(comments, "# Arguments:")

	for _, arg := range args {
		status := "optional"
		if arg.Required {
			status = "required"
		}

		var parts []string
		parts = append(parts, fmt.Sprintf("#   %s (%s)", arg.Name, status))

		if arg.Title != "" {
			parts = append(parts, arg.Title)
		}

		if arg.Description != "" {
			parts = append(parts, "- " + arg.Description)
		}

		comments = append(comments, strings.Join(parts, ": "))
	}

	return strings.Join(comments, "\n")
}

// generateMarkdownContent creates markdown content for a prompt
func generateMarkdownContent(meta *PromptMeta, namePrefix string) (string, error) {
	// Load template content
	data, err := fs.ReadFile(appassets.FS, meta.AssetPath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", meta.AssetPath, err)
	}

	templateContent := string(data)
	convertedContent := convertPlaceholders(templateContent, meta.Prompt.Arguments)

	// Generate argument comments
	argComments := generateArgumentComments(meta.Prompt.Arguments)
	var commentsSection string
	if argComments != "" {
		commentsSection = "\n" + argComments
	}

	// Generate frontmatter
	frontmatter := fmt.Sprintf(`---
description: %s
argument-hint: %s
model: claude-3-5-sonnet-20241022
# Generated by paper-reading-mcp v%s
# Source prompt: %s%s
---

%s`,
		meta.Prompt.Description,
		generateArgumentHint(meta.Prompt.Arguments),
		version.Version,
		meta.Prompt.Name,
		commentsSection,
		convertedContent)

	return frontmatter, nil
}

// generateFileName creates filename for a prompt
func generateFileName(promptName, namePrefix string) string {
	prefix := namePrefix
	if prefix == "" {
		prefix = "rp-"
	}
	return prefix + promptName + ".md"
}

// calculateHash creates a SHA-256 hash of all generated content
func calculateHash(contents map[string]string) string {
	h := sha256.New()
	for path, content := range contents {
		h.Write([]byte(path))
		h.Write([]byte(content))
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}

// ExportClaudeCommands implements the export_claude_commands tool
func ExportClaudeCommands(ctx context.Context, req *mcp.CallToolRequest, args ExportClaudeCommandsArgs) (*mcp.CallToolResult, ExportClaudeCommandsResult, error) {
	targetDir := args.TargetDir
	if targetDir == "" {
		targetDir = ".claude/commands"
	}

	namePrefix := args.NamePrefix
	if namePrefix == "" {
		namePrefix = "rp-"
	}

	// Generate all content
	contents := make(map[string]string)
	var paths []string

	for _, meta := range promptRegistry {
		fileName := generateFileName(meta.Prompt.Name, namePrefix)
		filePath := filepath.Join(targetDir, fileName)

		content, err := generateMarkdownContent(meta, namePrefix)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Failed to generate content: %v", err)},
				},
			}, ExportClaudeCommandsResult{Status: "error"}, nil
		}

		contents[filePath] = content
		paths = append(paths, filePath)
	}

	hash := calculateHash(contents)

	result := ExportClaudeCommandsResult{
		Status:       "ok",
		Written:      !args.DryRun,
		Paths:        paths,
		CommandCount: len(contents),
		Version:      version.Version,
		Hash:         hash,
	}

	if args.DryRun {
		// Return preview instead of writing files
		preview := "Generated files:\n"
		for path, content := range contents {
			preview += fmt.Sprintf("\n=== %s ===\n%s\n", path, content)
		}
		result.Preview = &preview

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Dry run completed. See result for preview."},
			},
		}, result, nil
	}

	// Create target directory if it doesn't exist
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Failed to create directory %s: %v", targetDir, err)},
			},
		}, ExportClaudeCommandsResult{Status: "error"}, nil
	}

	// Write files
	for path, content := range contents {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Failed to write file %s: %v", path, err)},
				},
			}, ExportClaudeCommandsResult{Status: "error"}, nil
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Successfully exported %d commands to %s", len(contents), targetDir)},
		},
	}, result, nil
}

// PreviewClaudeCommands implements the preview_claude_commands tool
func PreviewClaudeCommands(ctx context.Context, req *mcp.CallToolRequest, args PreviewClaudeCommandsArgs) (*mcp.CallToolResult, PreviewClaudeCommandsResult, error) {
	namePrefix := args.NamePrefix
	if namePrefix == "" {
		namePrefix = "rp-"
	}

	var files []PreviewFileContent
	targetDir := ".claude/commands"

	for _, meta := range promptRegistry {
		fileName := generateFileName(meta.Prompt.Name, namePrefix)
		filePath := filepath.Join(targetDir, fileName)

		content, err := generateMarkdownContent(meta, namePrefix)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Failed to generate content: %v", err)},
				},
			}, PreviewClaudeCommandsResult{Status: "error"}, nil
		}

		files = append(files, PreviewFileContent{
			Path:    filePath,
			Content: content,
		})
	}

	result := PreviewClaudeCommandsResult{
		Status:       "ok",
		CommandCount: len(files),
		Version:      version.Version,
		Files:        files,
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Preview generated for %d commands", len(files))},
		},
	}, result, nil
}

// RegisterExportClaudeCommandsTool registers the export_claude_commands tool using the new MCP API
func RegisterExportClaudeCommandsTool(s *mcp.Server) {
	tool := &mcp.Tool{
		Name:        "export_claude_commands",
		Description: "MCP Prompts を Claude Code 用の Slash Command ファイルとしてエクスポートし、.claude/commands/ に保存する。",
	}

	mcp.AddTool(s, tool, ExportClaudeCommands)
}

// RegisterPreviewClaudeCommandsTool registers the preview_claude_commands tool using the new MCP API
func RegisterPreviewClaudeCommandsTool(s *mcp.Server) {
	tool := &mcp.Tool{
		Name:        "preview_claude_commands",
		Description: "Claude Code 用 Slash Command ファイルの生成結果をプレビューし、保存せずに返す（CI や人手確認向け）。",
	}

	mcp.AddTool(s, tool, PreviewClaudeCommands)
}