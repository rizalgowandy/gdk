package cronx

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/cronx/page"
)

// SleepDuration defines the duration to sleep the server if the defined address is busy.
const SleepDuration = time.Second * 10

// NewServer creates a new http server.
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

	// Overcome issue with socket-master respawning 2nd app,
	// We will keep trying to run the server.
	// If the current address is busy,
	// sleep then try again until the address has become available.
	for {
		if err := e.Start(commandCtrl.Address); err != nil {
			time.Sleep(SleepDuration)
		}
	}
}

// ServerController is http server controller.
type ServerController struct {
	// CommandController controls all the underlying job.
	CommandController *CommandController
}

// HealthCheck returns server status.
func (c *ServerController) HealthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, c.CommandController.Info())
}

// Jobs return job status as frontend template.
func (c *ServerController) Jobs(context echo.Context) error {
	index, _ := page.GetStatusTemplate()
	return index.Execute(context.Response().Writer, c.CommandController.StatusData())
}

// APIJobs returns job status as json.
func (c *ServerController) APIJobs(context echo.Context) error {
	return context.JSON(http.StatusOK, c.CommandController.StatusJSON())
}
