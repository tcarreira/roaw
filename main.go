package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tcarreira/roaw/api"
	"github.com/tcarreira/roaw/config"
	"github.com/tcarreira/roaw/web/website"
)

var (
	version = "dev"
	commit  = ""
	dateStr = ""
)

func main() {
	fmt.Println("Starting Roaw - Run Once A Week")

	config.SetupGlobalConfig(config.Config{
		Version: config.Version{
			Version: version,
			Commit:  commit,
			DateStr: dateStr,
		},
		Db: config.NewDB("sqlite", ":memory:"),
	})

	flagVersion := flag.Bool("version", false, "Print version information and quit")
	flag.Parse()
	if *flagVersion {
		fmt.Println(config.GetConfigs().GetVersionString())
		os.Exit(0)
	}

	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	website.RegisterRoutes(e, "")
	api.RegisterRoutes(e, "/api")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
