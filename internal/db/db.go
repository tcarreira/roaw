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
	postgresSchema = `
DROP TABLE IF EXISTS roaw_user;

CREATE TABLE IF NOT EXISTS roaw_user (
	id VARCHAR(255),
	name VARCHAR(255),
	email VARCHAR(255),
	provider VARCHAR(255),
	provider_id VARCHAR(255),
	access_token VARCHAR(255),
	refresh_token VARCHAR(255),
	avatar_url VARCHAR(255),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

ALTER TABLE ONLY roaw_user
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX users_emails_unique
    ON roaw_user USING btree (email);

`
)

func CreateSchema(db *sqlx.DB) error {
	var err error

	switch db.DriverName() {
	case "sqlite":
		_, err = db.Exec(sqliteSchema)
	case "postgres":
		_, err = db.Exec(postgresSchema)
	default:
		err = fmt.Errorf("DB Driver not supported: %s", db.DriverName())
	}
	if err != nil {
		slog.Error(fmt.Sprintf("error creating schema: %s", err))
	}
	return err
}
