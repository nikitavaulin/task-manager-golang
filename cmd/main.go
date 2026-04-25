package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

const webDirPath = "web"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	router := core_http_server.NewRouter()
	router.RegisterFileServer("/", webDirPath)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewHTTPServerConfig(),
	)

	httpServer.RegisterRouters(router)

	if err := httpServer.Run(ctx); err != nil {
		fmt.Printf("HTTP server run error: %v\n", err)
		os.Exit(1)
	}
}
