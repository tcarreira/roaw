package users

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tcarreira/roaw/configs"
	"github.com/tcarreira/roaw/internal/db"
	"github.com/tcarreira/roaw/pkg/types"
)

func RegisterHandler(e *echo.Echo, conf configs.Config, path string) {
	g := e.Group(path)
	g.GET("", UsersListHandler(conf))
	g.POST("", UsersCreateHandler(conf))
	g.GET("/:uuid", UsersReadHandler)
	g.PUT("/:uuid", UsersUpdateHandler)
	g.DELETE("/:uuid", UsersDeleteHandler)
}

func UsersListHandler(conf configs.Config) func(c echo.Context) error {
	return func(c echo.Context) error {
		u, err := db.ListAllUsers(conf.Db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}

func UsersCreateHandler(conf configs.Config) func(c echo.Context) error {
	return func(c echo.Context) error {
		now := time.Now()
		u := types.User{}
		c.Bind(&u)
		u.ID = uuid.NewString()
		u.CreatedAt = now
		u.UpdatedAt = now

		err := db.UserCreate(conf.Db, &u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, u)
	}
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
