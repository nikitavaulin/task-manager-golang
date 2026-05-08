package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
	ErrNoRepeatRule    = errors.New("repeat rule is empty")
	ErrForbidden       = errors.New("forbidden")
)
