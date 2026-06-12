package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(60) NOT NULL UNIQUE,
		full_name VARCHAR(60) DEFAULT 'Пользователь',
		password_hash TEXT NOT NULL
	);

	CREATE TABLE task_categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category_name VARCHAR(60) NOT NULL,
		user_id INTEGER NOT NULL,

		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(256) NOT NULL,
		comment TEXT,
		date  CHAR(8) NOT NULL,
		repeat VARCHAR(128),
		category_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,

		FOREIGN KEY (category_id) REFERENCES task_categories(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE INDEX task_date ON tasks(date);
	CREATE INDEX user_username ON users(username);
`

type DBConnection struct {
	*sql.DB
}

func Init(dbFile string) (*DBConnection, error) {
	_, err := os.Stat(dbFile)

	var isNeedInstall bool
	if err != nil {
		isNeedInstall = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if isNeedInstall {
		if _, err := db.Exec(schema); err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to apply schema db: %w", err)
		}
	}

	return &DBConnection{DB: db}, nil
}
