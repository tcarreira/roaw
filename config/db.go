package config

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/jmoiron/sqlx"
)

func NewDB(dbType, connStr string) *sqlx.DB {
	resultChan := make(chan *sqlx.DB)
	go func() {
		if db, err := sqlx.Connect(dbType, connStr); err == nil {
			resultChan <- db
		}
		time.Sleep(10 * time.Millisecond)
		for {
			if db, err := sqlx.Connect(dbType, connStr); err == nil {
				resultChan <- db
			}
			slog.Warn(fmt.Sprintf("Could not connect to database (%s:%s). Retrying...", dbType, connStr))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	select {
	case db := <-resultChan:
		return db
	case <-time.After(10 * time.Second):
		panic("Timeout: Could not create database connection")
	}
}
