package assets

import "embed"

// Embedded assets for MCP resources and prompts.
//
//go:embed resources/** prompts/**
var FS embed.FS
