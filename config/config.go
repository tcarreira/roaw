package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

var gConfig *Config

type Version struct {
	Version string
	Commit  string
	DateStr string
}

type Config struct {
	Version Version
	Db      *sqlx.DB
}

func SetupGlobalConfig(c Config) {
	gConfig = &c
}

func defaultGlobalConfig() Config {
	db, err := sqlx.Connect("sqlite", ":memory")
	if err != nil {
		panic("cannot create a sqlite :memory")
	}
	return Config{
		Version: Version{"test", "", ""},
		Db:      db,
	}
}

func GetConfigs() *Config {
	if gConfig == nil {
		SetupGlobalConfig(defaultGlobalConfig())
	}
	return gConfig
}

func (c *Config) GetVersionString() string {
	if c.Version.Commit == "" && c.Version.DateStr == "" {
		return fmt.Sprintf("roaw version: %s", c.Version.Version)
	}
	return fmt.Sprintf("roaw version: %s (%s - %s)", c.Version.Version, c.Version.Commit, c.Version.DateStr)
}

func WebsiteTemplatesPath() string {
	p := os.Getenv("ROAW_TEMPLATES_PATH")
	if p != "" {
		return p
	}

	wd, _ := os.Getwd()
	tryPaths := []string{
		"/app/templates",
		filepath.Join(wd, "web", "templates"),
		filepath.Join(wd, "..", "web", "templates"),
	}

	for _, p := range tryPaths {
		if stat, err := os.Stat(p); err == nil && stat.IsDir() {
			return p
		}
	}

	panic("Could not find website html templates")
}
