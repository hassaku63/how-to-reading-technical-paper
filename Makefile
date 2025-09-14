# Simple Go Makefile

GO ?= go
PKG_SERVER := ./cmd/server
PKG_CLIENT := ./cmd/client
BIN_DIR := bin
SERVER_BIN := $(BIN_DIR)/server
CLIENT_BIN := $(BIN_DIR)/client

.PHONY: all build build-server build-client run-server run-client fmt clean test-client

all: build

build: build-server build-client

build-server:
	@mkdir -p $(BIN_DIR)
	$(GO) build -trimpath -ldflags "-s -w" -o $(SERVER_BIN) $(PKG_SERVER)

build-client:
	@mkdir -p $(BIN_DIR)
	$(GO) build -trimpath -ldflags "-s -w" -o $(CLIENT_BIN) $(PKG_CLIENT)

run-server:
	$(GO) run $(PKG_SERVER)

run-client:
	$(GO) run $(PKG_CLIENT)

# Test client against a running MCP server
test-client:
	@echo "Testing MCP client against server..."
	@echo "Make sure the server is running first: make run-server"
	@echo "Then run: make test-client URL=http://localhost:8080"
	@if [ -z "$(URL)" ]; then \
		echo "Usage: make test-client URL=http://localhost:8080"; \
		exit 1; \
	fi
	$(CLIENT_BIN) $(URL)

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN_DIR)

