package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/api/users"
	"github.com/tcarreira/roaw/config"
	"github.com/tcarreira/roaw/internal/db"
)

func RegisterRoutes(e *echo.Echo, path string) {
	users.RegisterHandler(e, path+"/users")

	g := e.Group(path)
	g.GET("/healthcheck", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})
	g.POST("/admin/db/migrate", func(c echo.Context) error {
		return db.CreateSchema(config.GetConfigs().Db)
	})

}
