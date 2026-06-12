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

	CategoryID int64
	UserID     int64
}

func NewTaskUninitialized(
	title string,
	date string,
	comment string,
	repeat string,
	categoryId int64,
	userId int64,
) *Task {
	return NewTask(
		UninitializedID,
		title,
		date,
		comment,
		repeat,
		categoryId,
		userId,
	)
}

func NewTask(
	id int64,
	title string,
	date string,
	comment string,
	repeat string,
	categoryId int64,
	userId int64,
) *Task {
	if len(date) == 0 {
		date = time.Now().Format(DateLayout)
	}
	return &Task{
		ID:         id,
		Title:      title,
		Date:       date,
		Comment:    comment,
		Repeat:     repeat,
		CategoryID: categoryId,
		UserID:     userId,
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
