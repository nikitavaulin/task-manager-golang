package task_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

type TaskHTTPTransportHandler struct {
	taskService TaskService
}

type TaskService interface {
	CreateTask(task domain.Task) (int64, error)
	UpdateTask(task domain.Task) error
	GetTasks(limit int, search *string) ([]domain.Task, error)
	GetTask(taskID int64) (domain.Task, error)
	DeleteTask(taskID int64) error
	SetTaskDone(taskID int64) error
}

func NewTaskHTTPTransportHandler(taskService TaskService) *TaskHTTPTransportHandler {
	return &TaskHTTPTransportHandler{
		taskService: taskService,
	}
}

func (h *TaskHTTPTransportHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/task",
			Handler: h.CreateTask,
			Auth:    true,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/tasks",
			Handler: h.GetTasks,
			Auth:    true,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/task",
			Handler: h.GetTask,
			Auth:    true,
		},
		{
			Method:  http.MethodPut,
			Path:    "/api/task",
			Handler: h.UpdateTask,
			Auth:    true,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/task",
			Handler: h.DeleteTask,
			Auth:    true,
		},
		{
			Method:  http.MethodPost,
			Path:    "/api/task/done",
			Handler: h.SetTaskDone,
			Auth:    true,
		},
	}
}
