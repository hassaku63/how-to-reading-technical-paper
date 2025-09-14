package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	srv "github.com/hassaku63/how-to-reading-technical-paper/internal/mcpserver"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := srv.RunStdio(ctx); err != nil {
		log.Printf("server exited with error: %v", err)
		os.Exit(1)
	}
}
