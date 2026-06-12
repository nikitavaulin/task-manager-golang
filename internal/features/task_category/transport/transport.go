package task_category_transport_http

import (
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/service"
	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

type TaskCategoryTransportHTTP struct {
	categoryService service.TaskCategoryService
}

func NewTaskCategoryTransportHTTP(categoryService service.TaskCategoryService) *TaskCategoryTransportHTTP {
	return &TaskCategoryTransportHTTP{
		categoryService: categoryService,
	}
}

func (h *TaskCategoryTransportHTTP) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/task-categories",
			Handler: h.CreateCategory,
			Auth:    true,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/task-categories",
			Handler: h.GetCategories,
			Auth:    true,
		},
		{
			Method:  http.MethodGet,
			Path:    "/api/task-categories/{id}",
			Handler: h.GetCategory,
			Auth:    true,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/api/task-categories/{id}",
			Handler: h.UpdateCategoryName,
			Auth:    true,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/api/task-categories/{id}",
			Handler: h.DeleteCategory,
			Auth:    true,
		},
	}
}
