package task_category_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type IDResponse struct {
	ID string `json:"id"`
}

func (h *TaskCategoryTransportHTTP) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	var category domain.TaskCategory
	if err := core_http_request.Decode(r, &category); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode category")
		return
	}

	categoryID, err := h.categoryService.CreateCategory(category)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create category")
		return
	}

	responseHandler.JSONResponse(IDResponse{fmt.Sprint(categoryID)}, http.StatusCreated)
}
