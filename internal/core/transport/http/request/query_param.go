package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func GetIntQueryParam(key string, r *http.Request) (*int, error) {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return nil, nil
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert int to str: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return &number, nil
}

func GetStringQueryParam(key string, r *http.Request) *string {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return nil
	}
	return &value
}
