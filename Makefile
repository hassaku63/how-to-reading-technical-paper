# Simple Go Makefile

GO ?= go
PKG_SERVER := ./cmd/server
PKG_DEBUG_CLIENT := ./cmd/debug-client
BIN_DIR := bin
SERVER_BIN := $(BIN_DIR)/server
DEBUG_CLIENT_BIN := $(BIN_DIR)/debug-client

.PHONY: all build build-server build-debug-client run-server run-debug-client fmt clean test-debug-client

all: build

build: build-server build-debug-client

build-server:
	@mkdir -p $(BIN_DIR)
	$(GO) build -trimpath -ldflags "-s -w" -o $(SERVER_BIN) $(PKG_SERVER)

build-debug-client:
	@mkdir -p $(BIN_DIR)
	$(GO) build -trimpath -ldflags "-s -w" -o $(DEBUG_CLIENT_BIN) $(PKG_DEBUG_CLIENT)

run-server:
	$(GO) run $(PKG_SERVER)

run-debug-client:
	$(GO) run $(PKG_DEBUG_CLIENT)

# Test debug client against MCP server using stdio transport
test-debug-client:
	@echo "Testing MCP debug client against server using stdio..."
	@echo "Usage: make test-debug-client"
	$(DEBUG_CLIENT_BIN) -server $(SERVER_BIN)

# Test debug client with verbose output
test-debug-client-verbose:
	@echo "Testing MCP debug client with verbose output..."
	$(DEBUG_CLIENT_BIN) -server $(SERVER_BIN) -verbose

# Show debug client help
debug-client-help:
	@echo "Showing debug client help..."
	$(DEBUG_CLIENT_BIN) -help

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN_DIR)

