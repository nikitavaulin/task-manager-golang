package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
	core_validation "github.com/nikitavaulin/task-manager-golang/internal/core/tools/validation"
)

type TaskCategory struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`

	UserID int64 `json:"user_id"`
}

const (
	MinCategoryNameLen = 3
	MaxCategoryNameLen = 60
)

func (c *TaskCategory) Validate() error {
	if err := c.ValidateCategoryName(c.CategoryName); err != nil {
		return fmt.Errorf("invalid category name: %w: %w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}

func (c *TaskCategory) ValidateCategoryName(categoryName string) error {
	if err := core_validation.ValidateIsIntInBounds(len(categoryName), MinCategoryNameLen, MaxCategoryNameLen); err != nil {
		return fmt.Errorf("invalid category name length: %w: %w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}
