package main

import (
	"errors"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/converter"
	"github.com/peractio/gdk/pkg/cronx"
	"github.com/peractio/gdk/pkg/stack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type sendEmail struct{}

func (e sendEmail) Run() error {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "sendEmail").
		Msg("every 5 sec send reminder emails")
	return nil
}

type payBill struct{}

func (p payBill) Run() error {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "payBill").
		Msg("every 1 min pay bill")
	return nil
}

type alwaysError struct{}

func (a alwaysError) Run() error {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "alwaysError").
		Msg("every 30 sec error")
	time.Sleep(5 * time.Second)
	return errors.New("some super long error message that come from executing the process")
}

type everyJob struct{}

func (everyJob) Run() error {
	return nil
}

type subscription struct{}

func (subscription) Run() error {
	return nil
}

func main() {
	// Create a cron controller with default config.
	cronx.Default()

	// Create a cron with custom config.
	cronx.New(cronx.Config{
		Address:  ":8000", // Determines if we want the library to serve the frontend.
		PoolSize: 1000,    // Determines how many jobs can be run at a time.
		PanicRecover: func(j *cronx.Job) { // Add panic middleware.
			if err := recover(); err != nil {
				log.WithLevel(zerolog.PanicLevel).
					Interface("err", err).
					Interface("stack", stack.ToArr(stack.Trim(debug.Stack()))).
					Interface("job", j).
					Msg("recovered")
			}
		},
		Location: func() *time.Location { // Change timezone to Jakarta.
			jakarta, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				secondsEastOfUTC := int((7 * time.Hour).Seconds())
				jakarta = time.FixedZone("WIB", secondsEastOfUTC)
			}
			return jakarta
		}(),
	})

	// Register a new cron job.
	// Struct name will become the name for the current job.
	if err := cronx.Schedule("@every 5s", sendEmail{}); err != nil {
		// create log and send alert we fail to register job.
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register sendEmail must success")
	}
	// Example of registering job with the same struct.
	// Duplication is okay.
	for i := 0; i < 3; i++ {
		spec := "@every " + converter.ToStr(i+1) + "m"
		if err := cronx.Schedule(spec, payBill{}); err != nil {
			log.WithLevel(zerolog.ErrorLevel).
				Err(err).
				Msg("register payBill must success")
		}
	}
	// Example jobs with broken spec.
	for i := 0; i < 3; i++ {
		spec := "broken spec " + converter.ToStr(i+1)
		if err := cronx.Schedule(spec, payBill{}); err != nil {
			log.WithLevel(zerolog.ErrorLevel).
				Err(err).
				Msg("register payBill must success")
		}
	}
	// Example of job with run that will always be error.
	if err := cronx.Schedule("@every 30s", alwaysError{}); err != nil {
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register alwaysError must success")
	}
	// Custom job with missing name.
	if err := cronx.Schedule("0 */1 * * *", cronx.Func(func() error {
		log.WithLevel(zerolog.InfoLevel).
			Str("job", "nameless job").
			Msg("every 1h will be run")
		return nil
	})); err != nil {
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register alwaysError must success")
	}
	// Job with v1 specification that includes seconds.
	if err := cronx.Schedule("0 0 4 * * *", subscription{}); err != nil {
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register alwaysError must success")
	}
	// Create a job that run every 20 sec.
	cronx.Every(20*time.Second, everyJob{})

	// Remove a job.
	cronx.Remove(2)

	// Get all current registered job.
	log.WithLevel(zerolog.InfoLevel).
		Interface("entries", cronx.GetEntries()).
		Msg("current jobs")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
