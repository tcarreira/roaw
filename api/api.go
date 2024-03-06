package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/configs"
	"github.com/tcarreira/roaw/internal/db"
)

func RegisterHealthcheck(e *echo.Echo, conf configs.Config, method, path string) {
	e.Add(method, path, func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})
}

func RegisterDBMigrate(e *echo.Echo, conf configs.Config, method, path string) {
	e.Add(method, path, func(c echo.Context) error {
		return db.CreateSchema(conf.Db)
	})
}
