package task_category_repository

import "database/sql"

type TaskCategoryRepository struct {
	db *sql.DB
}

func NewTaskCategoryRepository(db *sql.DB) *TaskCategoryRepository {
	return &TaskCategoryRepository{
		db: db,
	}
}
