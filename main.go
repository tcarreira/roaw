package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	version = "dev"
	commit  = ""
	dateStr = ""
)

func main() {
	fmt.Println("Starting Roaw - Run Once A Week")
	vStr := fmt.Sprintf("%s (%s - %s)\n", version, commit, dateStr)

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Print(vStr)
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, from roaw!")
	})
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, vStr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
