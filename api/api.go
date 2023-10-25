package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw2023/api/users"
)

func RegisterRoutes(e *echo.Echo, path string) {
	g := e.Group(path)
	g.GET("/healthcheck", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})
	users.RegisterHandler(e, path+"/users")
}
