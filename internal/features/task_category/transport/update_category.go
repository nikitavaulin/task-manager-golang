package task_category_transport_http

import (
	"net/http"

	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

type UpdateCategoryNameRequestDTO struct {
	CategoryName string `json:"category_name"`
}

func (h *TaskCategoryTransportHTTP) UpdateCategoryName(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	categoryID, err := core_http_request.GetPathValueID(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category ID")
		return
	}

	var dto UpdateCategoryNameRequestDTO
	if err := core_http_request.Decode(r, &dto); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode update category name request")
		return
	}

	if err := h.categoryService.UpdateCategoryName(categoryID, dto.CategoryName); err != nil {
		responseHandler.ErrorResponse(err, "failed to update category name")
		return
	}

	responseHandler.NoContentResponse(http.StatusOK)
}
