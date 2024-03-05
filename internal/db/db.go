package db

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

var (
	sqliteSchema = `
CREATE TABLE IF NOT EXISTS roaw_user (
	id TEXT PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE,
	provider TEXT,
	provider_id TEXT,
	access_token TEXT,
	refresh_token TEXT,
	avatar_url TEXT,
	created_at DATETIME,
	updated_at DATETIME
);
`
)

func CreateSchema(db *sqlx.DB) error {
	var err error

	switch db.DriverName() {
	case "sqlite":
		_, err = db.Exec(sqliteSchema)
	default:
		err = fmt.Errorf("DB Driver not supported: %s", db.DriverName())
	}
	if err != nil {
		slog.Error(fmt.Sprintf("error creating schema: %s", err))
	}
	return err
}
