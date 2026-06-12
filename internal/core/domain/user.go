package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

const (
	MinUsernameLen = 3
	MaxUsernameLen = 60
	MinFullNameLen = 3
	MaxFullNameLen = 60
)

func (u *User) Validate() error {
	if err := core_validation.ValidateIsIntInBounds(len(u.Username), MinUsernameLen, MaxUsernameLen); err != nil {
		return fmt.Errorf("invalid username: %w: %w", err, core_errors.ErrInvalidArgument)
	}
	if err := core_validation.ValidateIsIntInBounds(len(u.FullName), MinFullNameLen, MaxFullNameLen); err != nil {
		return fmt.Errorf("invalid fullname: %w: %w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}
