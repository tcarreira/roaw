package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tcarreira/roaw/api"
	"github.com/tcarreira/roaw/config"
	"github.com/tcarreira/roaw/web/website"
)

var (
	//go:embed assets/* web/templates/*
	embedFS embed.FS

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
		// Db: config.NewDB("postgres", "postgres://roawuser:roawpass@localhost:5432/roaw?sslmode=disable"),
		Db: config.NewDB("sqlite", "db.sqlite"),
	})

	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection: "1; mode=block",
		XFrameOptions: "SAMEORIGIN",
		HSTSMaxAge:    3600,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore(config.SessionSecret())))

	api.RegisterRoutes(e, "/api")

	e.Renderer = website.NewRenderer(embedFS)
	website.RegisterRoutes(e, "", embedFS)

	flagVersion := flag.Bool("version", false, "Print version information and quit")
	flag.Parse()
	if *flagVersion {
		fmt.Println(config.GetConfigs().GetVersionString())
		os.Exit(0)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
