package domain

import (
	"fmt"
	"time"

	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

type Task struct {
	ID      int64
	Title   string
	Date    string
	Comment string
	Repeat  string
}

func NewTaskUninitialized(
	title string,
	date string,
	comment string,
	repeat string,
) *Task {
	return NewTask(
		UninitializedID,
		title,
		date,
		comment,
		repeat,
	)
}

func NewTask(
	id int64,
	title string,
	date string,
	comment string,
	repeat string,
) *Task {
	if len(date) == 0 {
		date = time.Now().Format(DateLayout)
	}
	return &Task{
		ID:      id,
		Title:   title,
		Date:    date,
		Comment: comment,
		Repeat:  repeat,
	}
}

func (t *Task) Validate() error {
	if len(t.Title) == 0 {
		return fmt.Errorf("title cannot be empty")
	}
	if _, err := time.Parse(DateLayout, t.Date); err != nil {
		return fmt.Errorf("incorrect date format: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}
