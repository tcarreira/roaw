package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tcarreira/roaw/configs"
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

func runServer(ctx context.Context, conf configs.Config) error {
	// Handle command line flags
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagVersion := flagSet.Bool("version", false, "Print version information and quit")
	if err := flagSet.Parse(conf.Args[1:]); err != nil {
		return &FlagError{err.Error()}
	}
	if *flagVersion {
		fmt.Fprintln(conf.Stdout, conf.GetVersionString())
		return nil
	}

	// Setup routes
	e := newEchoServer(conf)
	registerAllRoutes(e, conf)

	// Start http server

	port := conf.Getenv("PORT")
	if port == "" {
		port = "0"
	}

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	<-ctx.Done()
	newCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := e.Shutdown(newCtx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := runServer(ctx, conf); err != nil {
		if _, ok := err.(*FlagError); !ok {
			fmt.Printf(":: %v\n", err)
		}
	}
}
