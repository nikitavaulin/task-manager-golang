package repeat_task_transport_http

import (
	"net/http"
	"time"

	core_http_server "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/server"
)

type RepeatTaskHTTPTransportHandler struct {
	repeatTaskService RepeatTaskService
}

type RepeatTaskService interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
}

func NewRepeatTaskHTTPTransportHandler(service RepeatTaskService) *RepeatTaskHTTPTransportHandler {
	return &RepeatTaskHTTPTransportHandler{
		repeatTaskService: service,
	}
}

func (h *RepeatTaskHTTPTransportHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/api/nextdate",
			Handler: h.GetNextDate,
		},
	}
}
