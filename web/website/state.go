package website

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type State struct {
	IsLoggedIn bool
}

func BuildState(c echo.Context) State {
	return State{
		IsLoggedIn: isLoggedIn(c),
	}
}

func isLoggedIn(c echo.Context) bool {
	sess, err := session.Get("session", c)
	fmt.Println("############", err, sess.Values)
	if err != nil {
		return false
	}
	userID := sess.Values["current_user_id"]
	return userID != nil && userID != ""
}
