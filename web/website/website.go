package website

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/config"
)

func RegisterRoutes(e *echo.Echo, path string, embedFS fs.FS) {
	HandleGroupWithConfigs(e, path, embedFS, config.GetConfigs())
}

func HandleGroupWithConfigs(e *echo.Echo, path string, embedFS fs.FS, conf *config.Config) {
	g := e.Group(path)

	g.StaticFS("", embedFS)

	g.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html.j2", BuildState(c))
	})
	g.GET("/index.html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html.j2", BuildState(c))
	})
	g.GET("/version", func(c echo.Context) error {
		return c.Render(http.StatusOK, "version.html.j2", map[string]string{"version": conf.GetVersionString()})
	})

	g.GET("/auth/strava", AuthCallback)
	g.GET("/auth/logout", AuthLogout)
}
