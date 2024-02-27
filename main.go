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
	"github.com/tcarreira/roaw/configs"
	"github.com/tcarreira/roaw/web/website"
)

var (
	//go:embed assets/* web/templates/*
	embedFS embed.FS

	version = "dev"
	commit  = ""
	dateStr = ""
)

type FlagError struct {
	Msg string
}

func (e *FlagError) Error() string { return e.Msg }

func newEchoServer(confs configs.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: confs.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection: "1; mode=block",
		XFrameOptions: "SAMEORIGIN",
		HSTSMaxAge:    3600,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore(confs.SessionSecret())))

	return e
}

func runServer(conf configs.Config) error {
	e := newEchoServer(conf)
	api.RegisterRoutes(e, conf, "/api")

	e.Renderer = website.NewRenderer(embedFS)
	website.RegisterRoutes(e, conf, "", embedFS)

	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagVersion := flagSet.Bool("version", false, "Print version information and quit")
	if err := flagSet.Parse(conf.Args[1:]); err != nil {
		return &FlagError{err.Error()}
	}
	if *flagVersion {
		fmt.Println(conf.GetVersionString())
		return nil
	}

	port := conf.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return e.Start(":" + port)

}

func main() {
	fmt.Println("Starting Roaw - Run Once A Week")

	conf := configs.NewConfig(
		configs.Version{
			Version: version,
			Commit:  commit,
			DateStr: dateStr,
		},
		os.Args,
		os.Getenv,
		os.Stdout,
		os.Stderr,
	)

	if err := runServer(conf); err != nil {
		if _, ok := err.(*FlagError); !ok {
			fmt.Printf(":: %v\n", err)
		}
	}
}
