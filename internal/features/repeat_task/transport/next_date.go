package repeat_task_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_http_request "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/task-manager-golang/internal/core/transport/http/response"
)

func (h *RepeatTaskHTTPTransportHandler) GetNextDate(rw http.ResponseWriter, r *http.Request) {
	responseHandler := core_http_response.NewHTTPResponseHandler(rw)

	nowStr, start, repeat, err := getNowDateRepeatQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get query params")
		return
	}

	now, err := getDateNowParam(nowStr)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'now' date query param")
		return
	}

	nextDate, err := h.repeatTaskService.NextDate(now, start, repeat)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to calc next date")
		return
	}

	responseHandler.TextResponse(nextDate, http.StatusOK)
}

func getDateNowParam(nowStr *string) (time.Time, error) {
	if nowStr == nil {
		return time.Now(), nil
	}
	now, err := parseDate(*nowStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}
	return now, nil
}

func parseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse(domain.DateLayout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return date, nil
}

func getNowDateRepeatQueryParams(r *http.Request) (*string, string, string, error) {
	now := core_http_request.GetStringQueryParam("now", r)

	date := core_http_request.GetStringQueryParam("date", r)
	if date == nil {
		return nil, "", "", fmt.Errorf("date query param is empty: %w", core_errors.ErrInvalidArgument)
	}

	repeat := core_http_request.GetStringQueryParam("repeat", r)
	if repeat == nil {
		return nil, "", "", fmt.Errorf("repeat query param is empty: %w", core_errors.ErrInvalidArgument)
	}

	return now, *date, *repeat, nil
}
