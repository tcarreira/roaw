package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo, path string) {
	g := e.Group(path)
	g.GET("", UsersListHandler)
	g.POST("", UsersCreateHandler)
	g.GET("/:uuid", UsersReadHandler)
	g.PUT("/:uuid", UsersUpdateHandler)
	g.DELETE("/:uuid", UsersDeleteHandler)
}

func UsersListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]any{{"id": "666", "username": "tbd"}})
}

func UsersCreateHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]any{"id": "666", "username": "tbd"})
}

func UsersReadHandler(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid != "666" {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]any{"id": "666", "username": "tbd"})
}

func UsersUpdateHandler(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid != "666" {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]any{"id": "666", "username": "tbd"})
}

func UsersDeleteHandler(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid == "666" {
		return c.NoContent(http.StatusNoContent)
	}
	return c.NoContent(http.StatusNotFound)
}
