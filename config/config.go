package config

import "fmt"

var gConfig *Config

type Config struct {
	Version string
	Commit  string
	DateStr string
}

func SetupGlobalConfig(c Config) {
	gConfig = &c
}

func defaultGlobalConfig() Config {
	return Config{
		Version: "test",
	}
}

func GetConfigs() *Config {
	if gConfig == nil {
		SetupGlobalConfig(defaultGlobalConfig())
	}
	return gConfig
}

func (c *Config) GetVersionString() string {
	if c.Commit == "" && c.DateStr == "" {
		return fmt.Sprintf("roaw version: %s", c.Version)
	}
	return fmt.Sprintf("roaw version: %s (%s - %s)", c.Version, c.Commit, c.DateStr)
}
