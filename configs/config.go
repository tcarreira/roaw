package configs

import (
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Version struct {
	Version string
	Commit  string
	DateStr string
}

type Config struct {
	Version    Version
	Args       []string
	Getenv     func(string) string
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	Db         *sqlx.DB
	ListenPort uint
}

func NewConfig(
	version Version,
	args []string,
	getenv func(string) string,
	stdout,
	stderr io.Writer,
) Config {
	return Config{
		Version: version,
		Args:    args,
		Getenv:  getenv,
		Stdout:  stdout,
		Stderr:  stderr,
		Db:      NewDB("sqlite", "db.sqlite"),
	}
}

func (c *Config) GetVersionString() string {
	if c.Version.Commit == "" && c.Version.DateStr == "" {
		return fmt.Sprintf("roaw version: %s", c.Version.Version)
	}
	return fmt.Sprintf("roaw version: %s (%s - %s)", c.Version.Version, c.Version.Commit, c.Version.DateStr)
}

func (c *Config) SessionSecret() []byte {
	secret := c.Getenv("ROAW_SESSION_SECRET")
	if secret == "" {
		u := uuid.New()
		return u[:]
	}
	return []byte(secret)
}
