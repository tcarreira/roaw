package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tcarreira/roaw/configs"
)

func TestRunServer_VersionFlag(t *testing.T) {
	conf := configs.Config{
		Version: configs.Version{
			Version: "testVersion",
			Commit:  "testCommit",
			DateStr: "testDateStr",
		},
		Getenv: func(s string) string { return "" },
		Args:   []string{"bin", "--version"},
		Stdout: &strings.Builder{},
		Stderr: &strings.Builder{},
	}
	err := runServer(conf)

	assert.Equal(t, "roaw version: testVersion (testCommit - testDateStr)\n", conf.Stdout.(*strings.Builder).String())
	assert.Equal(t, "", conf.Stderr.(*strings.Builder).String())
	assert.NoError(t, err)
}
