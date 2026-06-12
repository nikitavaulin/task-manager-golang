package user_repository

import (
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
)

func (r *UserRepository) CreateUser(user domain.User) (int64, error) {
	query := `
		INSERT INTO users
			(username, full_name, password_hash)
		VALUES
			($1, $2, $3);
	`

	result, err := r.db.Exec(
		query,
		user.Username,
		user.FullName,
		user.PasswordHash,
	)
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to create user in db: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.UninitializedID, fmt.Errorf("failed to get user ID: %w", err)
	}

	return id, nil
}
