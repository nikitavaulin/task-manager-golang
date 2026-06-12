package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func Decode(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf(
			"failed to decode json: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}
