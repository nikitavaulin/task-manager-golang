package task_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
)

func getTaskID(r *http.Request) (int64, error) {
	id, err := core_http_request.GetIntQueryParam("id", r)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to get ID query param")
	}
	if id == nil {
		return domain.UninitializedID, fmt.Errorf("ID cannot be empty: %w", core_errors.ErrInvalidArgument)
	}
	return int64(*id), err
}
