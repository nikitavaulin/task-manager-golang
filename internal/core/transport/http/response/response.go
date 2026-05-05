package core_http_response

import (
	"encoding/json"
	"errors"
	"net/http"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

type HTTPResponseHandler struct {
	rw http.ResponseWriter
}

func NewHTTPResponseHandler(rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		rw: rw,
	}
}

func (h *HTTPResponseHandler) TextResponse(responseBody string, statusCode int) {
	h.rw.WriteHeader(statusCode)
	h.rw.Write([]byte(responseBody))
}

func (h *HTTPResponseHandler) JSONResponse(responseBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)
	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		// log
	}
}

func (h *HTTPResponseHandler) NoContentResponse(statusCode int) {
	h.rw.WriteHeader(statusCode)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var statusCode int
	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict

	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest

	default:
		statusCode = http.StatusInternalServerError
	}
	h.errorResponse(err, msg, statusCode)

}

func (h *HTTPResponseHandler) errorResponse(err error, msg string, statusCode int) {
	// log
	responseBody := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}
	h.JSONResponse(responseBody, statusCode)
}
