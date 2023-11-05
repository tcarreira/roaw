package website

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// AuthCallback handles callback from provider
func AuthCallback(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["current_user_id"] = "hardcoded-user-id"
	sess.AddFlash("You have been logged in", "success")
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(302, "/")
}

// AuthLogout will close and clear the login session
func AuthLogout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	delete(sess.Values, "current_user_id")
	sess.AddFlash("You have been logged out", "success")

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(302, "/")
}
