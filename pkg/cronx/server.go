package cronx

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/cronx/page"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/tags"
)

const TimeoutDuration = time.Second * 10

// NewServer creates an HTTP server.
// - /			=> current server status.
// - /jobs		=> current jobs as frontend html.
// - /api/jobs	=> current jobs as json.
func NewServer(commandCtrl *CommandController) {
	if commandCtrl.Location == nil {
		commandCtrl.Location = defaultConfig.Location
	}

	// Create server.
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())

	// Create server controller.
	ctrl := &ServerController{CommandController: commandCtrl}

	// Register routes.
	e.GET("/", ctrl.HealthCheck)
	e.GET("/jobs", ctrl.Jobs)
	e.GET("/api/jobs", ctrl.APIJobs)

	// Start server.s
	go func() {
		if err := e.Start(commandCtrl.Address); err != nil && err != http.ErrServerClosed {
			logx.FTL(logx.NewContext(), err, "shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a certain timeout.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Stop cron jobs.
	ctx := commandController.Commander.Stop()
	ctx = logx.ContextWithRequestID(ctx)
	select {
	case <-ctx.Done():
		logx.INF(ctx, nil, "cron has been shutdown")
	case <-time.After(TimeoutDuration):
		logx.WRN(
			ctx,
			errorx.E("timeout", errorx.Fields{tags.Duration: TimeoutDuration.String()}),
			"cron shutdown failure",
		)
	}

	ctx, cancel := context.WithTimeout(ctx, TimeoutDuration)
	defer cancel()

	// Shutdown server.
	if err := e.Shutdown(ctx); err != nil && err != context.Canceled {
		logx.FTL(ctx, err, "server shutdown failure")
	} else {
		logx.INF(ctx, nil, "server has been shutdown")
	}
}

// ServerController is http server controller.
type ServerController struct {
	// CommandController controls all the underlying job.
	CommandController *CommandController
}

// HealthCheck returns server status.
func (c *ServerController) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, c.CommandController.Info())
}

// Jobs return job status as frontend template.
func (c *ServerController) Jobs(ctx echo.Context) error {
	index, _ := page.GetStatusTemplate()
	return index.Execute(ctx.Response().Writer, c.CommandController.StatusData())
}

// APIJobs returns job status as json.
func (c *ServerController) APIJobs(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, c.CommandController.StatusJSON())
}
