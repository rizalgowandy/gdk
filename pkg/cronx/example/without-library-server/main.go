package main

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/converter"
	"github.com/peractio/gdk/pkg/cronx"
	"github.com/peractio/gdk/pkg/cronx/interceptor"
	"github.com/peractio/gdk/pkg/cronx/page"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
)

type SendEmail struct{}

func (s SendEmail) Run(ctx context.Context) error {
	logx.INF(ctx, nil, "send email is running")
	return nil
}

type PayBill struct{}

func (p PayBill) Run(ctx context.Context) error {
	logx.INF(ctx, nil, "pay bill is running")
	return nil
}

type AlwaysError struct{}

func (a AlwaysError) Run(ctx context.Context) error {
	err := errorx.E("some super long error message that come from executing the process")
	logx.ERR(ctx, err, "always error is running")
	return err
}

type EveryJob struct{}

func (EveryJob) Run(ctx context.Context) error {
	logx.INF(ctx, nil, "every job is running")
	return nil
}

type Subscription struct{}

func (Subscription) Run(ctx context.Context) error {
	md, ok := cronx.GetJobMetadata(ctx)
	if !ok {
		return errorx.E("cannot job metadata")
	}

	logx.INF(ctx, md, "Subscription is running")
	return nil
}

func main() {
	// ===========================
	// With Custom Configuration
	// ===========================
	// Create cron middleware.
	// The order is important.
	// The first one will be executed first.
	cronMiddleware := cronx.Chain(
		interceptor.RequestID,
		interceptor.Recover(),
		interceptor.Logger(),
		interceptor.DefaultWorkerPool(),
	)

	// Create a cron with custom config and middleware.
	cronx.New(cronx.Config{
		Address: "", // Disable built-in server frontend.
		Location: func() *time.Location { // Change timezone to Jakarta.
			jakarta, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				secondsEastOfUTC := int((7 * time.Hour).Seconds())
				jakarta = time.FixedZone("WIB", secondsEastOfUTC)
			}
			return jakarta
		}(),
	}, cronMiddleware)
	// Since we want to create custom HTTP server.
	// Do not forget to shutdown the cron gracefully manually here.
	defer cronx.Stop()

	// Register jobs.
	RegisterJobs()

	// ===========================
	// Start Main Server
	// ===========================
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())

	// Register routes because we want to create custom HTTP server.
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, cronx.GetInfo())
	})
	e.GET("/jobs", func(c echo.Context) error {
		index, _ := page.GetStatusTemplate()
		return index.Execute(c.Response().Writer, cronx.GetStatusData())
	})
	e.GET("/api/jobs", func(c echo.Context) error {
		return c.JSON(http.StatusOK, cronx.GetStatusJSON())
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func RegisterJobs() {
	ctx := logx.NewContext()

	// Struct name will become the name for the current job.
	if err := cronx.Schedule("@every 5s", SendEmail{}); err != nil {
		// create log and send alert we fail to register job.
		logx.ERR(ctx, err, "register send email must success")
	}

	// Create some jobs with the same struct.
	// Duplication is okay.
	for i := 0; i < 3; i++ {
		spec := "@every " + converter.String(i+1) + "m"
		if err := cronx.Schedule(spec, PayBill{}); err != nil {
			logx.ERR(ctx, err, "register pay bill must success")
		}
	}

	// Create some jobs with broken spec.
	for i := 0; i < 3; i++ {
		spec := "broken spec " + converter.String(i+1)
		if err := cronx.Schedule(spec, PayBill{}); err != nil {
			logx.ERR(ctx, err, "register pay bill must success")
		}
	}

	// Create a job with run that will always be error.
	if err := cronx.Schedule("@every 30s", AlwaysError{}); err != nil {
		logx.ERR(ctx, err, "register always error must success")
	}

	// Create a custom job with missing name.
	if err := cronx.Schedule("0 */1 * * *", cronx.Func(func(context.Context) error {
		logx.INF(ctx, nil, "nameless job is running")
		return nil
	})); err != nil {
		logx.ERR(ctx, err, "register nameless job must success")
	}

	// Create a job with v1 specification that includes seconds.
	if err := cronx.Schedule("0 0 1 * * *", Subscription{}); err != nil {
		logx.ERR(ctx, err, "register subscription must success")
	}

	// Create a job with multiple schedules
	if err := cronx.Schedules("0 0 4 * * *#0 0 7 * * *#0 0 11 * * *", "#", Subscription{}); err != nil {
		logx.ERR(ctx, err, "register subscription must success")
	}

	const (
		everyInterval    = 20
		jobIDToBeRemoved = 2
	)

	// Create a job that run every 20 sec.
	cronx.Every(everyInterval*time.Second, EveryJob{})

	// Remove a job.
	cronx.Remove(jobIDToBeRemoved)

	// Get all current registered job.
	logx.INF(ctx, cronx.GetEntries(), "current jobs")
}
