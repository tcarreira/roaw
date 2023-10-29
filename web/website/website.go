package website

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/config"
)

func RegisterRoutes(e *echo.Echo, path string) {
	HandleGroupWithConfigs(e, path, config.GetConfigs())
}

func HandleGroupWithConfigs(e *echo.Echo, path string, conf *config.Config) {
	g := e.Group(path)

	g.Static("/assets", "assets")

	g.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html.j2", nil)
	})
	g.GET("/index.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html.j2", nil)
	})
	g.GET("/version", func(c echo.Context) error {
		return c.Render(http.StatusOK, "version.html.j2", map[string]string{"version": conf.GetVersionString()})
	})

}
