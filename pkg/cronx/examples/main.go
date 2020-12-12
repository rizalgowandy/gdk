package main

import (
	"runtime/debug"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peractio/gdk/pkg/converter"
	"github.com/peractio/gdk/pkg/cronx"
	"github.com/peractio/gdk/pkg/stack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type sendEmail struct{}

func (e sendEmail) Run() {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "sendEmail").
		Msg("every 5 sec send reminder emails")
}

type payBill struct{}

func (p payBill) Run() {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "payBill").
		Msg("every 1 min pay bill")
}

func main() {
	// Create a cron controller with default config.
	cronx.Default()

	// Create a cron with custom config.
	cronx.New(cronx.Config{
		Address:  ":8000", // Determines if we want the library to serve the frontend.
		PoolSize: 1000,    // Determines how many jobs can be run at a time.
		PanicRecover: func(j *cronx.Job) {
			if err := recover(); err != nil {
				log.WithLevel(zerolog.PanicLevel).
					Interface("err", err).
					Interface("stack", stack.ToArr(stack.Trim(debug.Stack()))).
					Interface("job", j).
					Msg("recovered")
			}
		}, // Inject panic middleware with custom logger and alert.
	})

	// Register a new cron job.
	// Struct name will become the name for the current job.
	if err := cronx.Schedule("@every 5s", sendEmail{}); err != nil {
		// create log and send alert we fail to register job.
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register sendEmail must success")
	}
	for i := 0; i < 15; i++ {
		spec := "@every " + converter.ToStr(i+1) + "m"
		if err := cronx.Schedule(spec, payBill{}); err != nil {
			log.WithLevel(zerolog.ErrorLevel).
				Err(err).
				Msg("register payBill must success")
		}
	}
	for i := 0; i < 3; i++ {
		spec := "broken spec " + converter.ToStr(i+1)
		if err := cronx.Schedule(spec, payBill{}); err != nil {
			log.WithLevel(zerolog.ErrorLevel).
				Err(err).
				Msg("register payBill must success")
		}
	}

	// Get all current registered job.
	log.WithLevel(zerolog.InfoLevel).
		Interface("entries", cronx.GetEntries()).
		Msg("current jobs")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
