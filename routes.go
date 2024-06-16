package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/api"
	"github.com/tcarreira/roaw/api/users"
	"github.com/tcarreira/roaw/configs"
	"github.com/tcarreira/roaw/web/website"
)

func registerAllRoutes(e *echo.Echo, conf configs.Config) {
	api.RegisterHealthcheck(e, conf, http.MethodGet, "/api/healthcheck")
	api.RegisterDBMigrate(e, conf, http.MethodPost, "/api/admin/db/migrate")
	users.RegisterHandler(e, conf, "/api/users")

	e.Renderer = website.NewRenderer(embedFS)
	website.RegisterRoutes(e, conf, "", embedFS)
}
