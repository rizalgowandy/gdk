package cronx

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peractio/gdk/pkg/cronx/pages"
	"github.com/robfig/cron/v3"
)

// SleepDuration defines the duration to sleep the server if the defined address is busy.
const SleepDuration = time.Second * 10

// CommandController controls all the underlying job.
type CommandController struct {
	// Commander holds all the underlying cron jobs.
	Commander *cron.Cron

	// WorkerPool determine the limit of the number of jobs allowed to run concurrently.
	WorkerPool chan struct{}

	// PanicRecover is deferred function that will be executed before executing each job.
	PanicRecover func(j *Job)

	// Address determines the address will we serve the json and frontend status.
	// Empty string meaning we won't serve the current job status.
	Address string

	// UnregisteredJobs describes the list of jobs that have been failed to be registered.
	UnregisteredJobs []*Job

	// CreatedTime describes when the command controller created.
	CreatedTime time.Time
}

// Default starts all the underlying cron jobs.
// If address is not empty, create a server with routes:
// - /			=> current server status.
// - /jobs 		=> current jobs as json.
// - /jobs/html => current jobs as frontend html.
func (c *CommandController) Start() {
	// Start the commander.
	if c.Commander == nil {
		c.Commander = cron.New()
	}
	c.Commander.Start()

	// Check if client want to start a server to serve json and frontend.
	if c.Address == "" {
		return
	}

	go func() {
		// Create a server.
		e := echo.New()
		e.HideBanner = true
		e.HidePort = true
		e.Use(middleware.Recover())
		e.Use(middleware.RemoveTrailingSlash())

		// Register routes.
		e.GET("/", func(context echo.Context) error {
			return context.JSON(http.StatusOK, map[string]interface{}{
				"status": http.StatusText(http.StatusOK),
				"data": map[string]interface{}{
					"current_time": time.Now().String(),
					"created_time": c.CreatedTime.String(),
					"up_time":      time.Since(c.CreatedTime).String(),
				},
			})
		})
		e.GET("/jobs", func(context echo.Context) error {
			return context.JSON(http.StatusOK, GetStatusJSON())
		})
		index, _ := pages.GetStatusTemplate()
		e.GET("jobs/html", func(context echo.Context) error {
			return index.Execute(context.Response().Writer, GetStatusData())
		})

		// Overcome issue with socket-master respawning 2nd app,
		// We will keep trying to run the server,
		// if the current address is busy,
		// sleep then try again until the address has become available.
		for {
			if err := e.Start(c.Address); err != nil {
				time.Sleep(SleepDuration)
			}
		}
	}()
}

// NewCommandController create a command controller with a specific config.
func NewCommandController(config Config) *CommandController {
	return &CommandController{
		Commander:    cron.New(),
		WorkerPool:   make(chan struct{}, config.PoolSize),
		PanicRecover: config.PanicRecover,
		Address:      config.Address,
		CreatedTime:  time.Now(),
	}
}
