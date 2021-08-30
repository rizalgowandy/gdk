package main

import (
	"context"
	"time"

	"github.com/rizalgowandy/gdk/pkg/converter"
	"github.com/rizalgowandy/gdk/pkg/cronx"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/logx"
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
	// With Default Configuration
	// ===========================
	// Create a cron controller with default config.
	// - running on port :8998
	// - location is time.Local
	// - without any middleware
	cronx.Default()

	// Register jobs.
	RegisterJobs()

	// Start HTTP server.
	cronx.Serve()
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
