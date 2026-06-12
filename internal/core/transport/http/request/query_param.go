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
		return nil, fmt.Errorf("failed to convert str to int: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return &number, nil
}

func GetIDQueryParamMust(key string, r *http.Request) (int64, error) {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return 0, nil
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return id, nil
}

func GetStringQueryParam(key string, r *http.Request) *string {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return nil
	}
	return &value
}
