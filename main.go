package main

import (
	"flag"
	"fmt"
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	version = "dev"
	commit  = ""
	dateStr = ""
)

func main() {
	fmt.Println("Starting Roaw - Run Once A Week")

	verInfo := fmt.Sprintf("roaw version: %s (%s - %s)\n", version, commit, dateStr)
	flagVersion := flag.Bool("version", false, "Print version information and quit")
	flag.Parse()
	if *flagVersion {
		fmt.Println(verInfo)
		os.Exit(0)
	}

	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, from roaw!")
	})
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, verInfo)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
