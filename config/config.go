package config

import (
	"fmt"

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
