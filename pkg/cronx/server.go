package cronx

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/cronx/pages"
)

// NewServer creates a new http server.
// - /			=> current server status.
// - /jobs		=> current jobs as frontend html.
// - /api/jobs	=> current jobs as json.
func NewServer(config Config, commandCtrl *CommandController) {
	if config.Location == nil {
		config.Location = defaultConfig.Location
	}

	// Create a server.
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	// Create controller.
	ctrl := &Controller{
		CommandController: commandCtrl,
		CreatedTime:       time.Now().In(config.Location),
		Location:          config.Location,
	}

	// Register routes.
	e.GET("/", ctrl.HealthCheck)
	e.GET("/jobs", ctrl.Jobs)
	e.GET("/api/jobs", ctrl.APIJobs)

	// Overcome issue with socket-master respawning 2nd app,
	// We will keep trying to run the server.
	// If the current address is busy,
	// sleep then try again until the address has become available.
	for {
		if err := e.Start(config.Address); err != nil {
			time.Sleep(SleepDuration)
		}
	}
}

// Controller is http server controller.
type Controller struct {
	// CommandController controls all the underlying job.
	CommandController *CommandController
	// CreatedTime describes when the command controller created.
	CreatedTime time.Time
	// Location describes the timezone current cron is running.
	// By default the timezone will be the same timezone as the server.
	Location *time.Location
}

// HealthCheck returns server status.
func (c *Controller) HealthCheck(context echo.Context) error {
	currentTime := time.Now().In(c.Location)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusText(http.StatusOK),
		"data": map[string]interface{}{
			"location":     c.Location.String(),
			"created_time": c.CreatedTime.String(),
			"current_time": currentTime.String(),
			"up_time":      currentTime.Sub(c.CreatedTime).String(),
		},
	})
}

// Jobs return job status as frontend template.
func (c *Controller) Jobs(context echo.Context) error {
	index, _ := pages.GetStatusTemplate()
	return index.Execute(context.Response().Writer, c.CommandController.GetStatusData())
}

// APIJobs returns job status as json.
func (c *Controller) APIJobs(context echo.Context) error {
	return context.JSON(http.StatusOK, c.CommandController.GetStatusJSON())
}
