package website

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw2023/config"
)

func RegisterRoutes(e *echo.Echo, path string) {
	HandleGroupWithConfigs(e, path, config.GetConfigs())
}

func HandleGroupWithConfigs(e *echo.Echo, path string, conf *config.Config) {
	g := e.Group(path)
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, from roaw!")
	})
	g.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, conf.GetVersionString()+"\n")
	})
}
