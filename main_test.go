package main

import (
	"context"
	"strings"
	"testing"
	"time"

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
		Args:   []string{"binary", "--version"},
		Stdout: &strings.Builder{},
		Stderr: &strings.Builder{},
	}
	err := runServer(context.Background(), conf)

	assert.Equal(t, "roaw version: testVersion (testCommit - testDateStr)\n", conf.Stdout.(*strings.Builder).String())
	assert.Equal(t, "", conf.Stderr.(*strings.Builder).String())
	assert.NoError(t, err)
}

func TestRunServer_Healthcheck(t *testing.T) {
	conf := configs.Config{
		Getenv: func(s string) string { return "" },
		Args:   []string{"binary"},
		Stdout: &strings.Builder{},
		Stderr: &strings.Builder{},
	}
	ctx, cancel := context.WithCancel(context.Background())
	var err error
	go func() {
		err = runServer(ctx, conf)
	}()

	<-time.After(2 * time.Second)

	assert.Equal(t, "", conf.Stdout.(*strings.Builder).String())
	assert.Equal(t, "", conf.Stderr.(*strings.Builder).String())
	assert.NoError(t, err)
	cancel()
}
