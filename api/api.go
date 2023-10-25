package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, group string) {
	g := e.Group(group)
	g.GET("/healthcheck", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})
}
