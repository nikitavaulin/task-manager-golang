package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	tools_envparser "github.com/nikitavaulin/task-manager-golang/internal/core/tools/env_parser"
	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
	repeat_service "github.com/nikitavaulin/task-manager-golang/internal/features/repeat_task/service"
	repeat_task_transport_http "github.com/nikitavaulin/task-manager-golang/internal/features/repeat_task/transport"
	task_repository "github.com/nikitavaulin/task-manager-golang/internal/features/task/repository"
	task_service "github.com/nikitavaulin/task-manager-golang/internal/features/task/service"
	task_transport_http "github.com/nikitavaulin/task-manager-golang/internal/features/task/transport"
	"github.com/nikitavaulin/task-manager-golang/pkg/db"
)

const (
	webDirPath    = "web"
	dbFileDefault = "scheduler.db"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbFile := tools_envparser.GetEnvVarOrDefault("TODO_DBFILE", dbFileDefault)
	dbConn, err := db.Init(dbFile)
	if err != nil {
		fmt.Printf("failed to connect to db: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	repeatTaskService := repeat_service.NewRepeatTaskService()
	repeatTaskTransport := repeat_task_transport_http.NewRepeatTaskHTTPTransportHandler(repeatTaskService)

	taskRepository := task_repository.NewTaskRepository(dbConn.DB)
	taskService := task_service.NewTaskService(taskRepository, repeatTaskService)
	taskTransport := task_transport_http.NewTaskHTTPTransportHandler(taskService)

	router := core_http_server.NewRouter()
	router.RegisterFileServer("/", webDirPath)
	router.RegisterRoutes(repeatTaskTransport.Routes()...)
	router.RegisterRoutes(taskTransport.Routes()...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewHTTPServerConfig(),
	)

	httpServer.RegisterRouters(router)

	if err := httpServer.Run(ctx); err != nil {
		fmt.Printf("HTTP server run error: %v\n", err)
		os.Exit(1)
	}
}
