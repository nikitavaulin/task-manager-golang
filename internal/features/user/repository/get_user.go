package user_repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikitavaulin/task-manager-golang/internal/core/domain"
	core_errors "github.com/nikitavaulin/task-manager-golang/internal/core/errors"
)

func (r *UserRepository) GetUserByID(userID int64) (domain.User, error) {
	query := `
		SELECT * FROM users
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, userID)

	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with ID=%d: %v: %w", userID, err, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("failed to get user from db: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (domain.User, error) {
	query := `
		SELECT * FROM users
		WHERE username = $1;
	`

	row := r.db.QueryRow(query, username)

	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with username=%s: %v: %w", username, err, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("failed to get user from db: %w", err)
	}

	return user, nil
}
