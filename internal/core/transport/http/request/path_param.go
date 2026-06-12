package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func GetPathValueID(r *http.Request, key string) (int64, error) {
	value := r.PathValue(key)
	if len(value) == 0 {
		return 0, fmt.Errorf("ID path value by key %s is empty: %w", key, core_errors.ErrInvalidArgument)
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID path value by key %s is empty: %w: %w", key, err, core_errors.ErrInvalidArgument)
	}

	return id, nil
}
