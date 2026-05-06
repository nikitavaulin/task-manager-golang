package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(256) NOT NULL,
    comment TEXT,
    date  CHAR(8) NOT NULL,
    repeat VARCHAR(128)
);

CREATE INDEX scheduler_date ON scheduler(date);
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
			return nil, fmt.Errorf("failed to apply schema db: %w", err)
		}
	}

	return &DBConnection{DB: db}, nil
}
