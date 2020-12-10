# Cronx

Cronx is a wrapper for _robfig/cron_. It includes a live monitoring of current schedule and state of active jobs that can be outputted as JSON or HTML template.

### Quick Start
Create _**main.go**_.

```go
package main

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/peractio/gdk/pkg/cronx"
	"github.com/peractio/gdk/pkg/stack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// In order to create a job you need to create a struct that has Run() method.
type sendEmail struct{}

func (e sendEmail) Run() {
	time.Sleep(time.Second * 3)
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "sendEmail").
		Msg("every 5 sec send reminder emails")
}

func main() {
	// Create a cron controller with default config that:
	// - runs on port :8998
	// - has a max running jobs limit 1000
	// - with built in panic recovery
	cronx.Start()
	
	// Register a new cron job.
	// Struct name will become the name for the current job.
	if err := cronx.Schedule("@every 5s", sendEmail{}); err != nil {
		// create log and send alert we fail to register job.
		log.WithLevel(zerolog.ErrorLevel).
			Err(err).
			Msg("register sendEmail must success")
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8080"))
}
```

Start server
```shell
$ go run main.go
```

Browse to
- http://localhost:8998/jobs/html => see the html page.
![](https://raw.githubusercontent.com/peractio/gdk/main/pkg/cronx/screenshots/status_page.png?token=ABUOECPMR742WX5EQZAPTVC72OLD2)
- http://localhost:8998/jobs => see the json response.
```json
{
  "data": [
    {
      "id": 1,
      "job": {
        "name": "sendEmail",
        "status": "RUNNING",
        "latency": "3.000299794s"
      },
      "next_run": "2020-12-11T22:36:35+07:00",
      "prev_run": "2020-12-11T22:36:30+07:00"
    }
  ]
}
```

### Custom Configuration
```go
// Create a cron with custom config.
cronx.StartWithConfig(cronx.Config{
    Address:  ":8000", // Determines if we want the library to serve the frontend.
    PoolSize: 1,       // Determines how many jobs can be run at a time.
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
```

### Schedule Specification Format
Please refer to this [link](https://pkg.go.dev/github.com/robfig/cron?readme=expanded#section-readme/).

### FAQS
#### Q: I don't want the use the built-in server, but my own!
#### A: Yes, you can. This library is very modular.
```go
// Create a custom config and leave the address as empty string.
cronx.StartWithConfig(cronx.Config{
    Address:  "",
})

// GetStatusData will return the []cronx.StatusData.
// You can use this data like any other Golang array.
// You can print it, or even serves it using your own router.
res := cronx.GetStatusData() 

// An example using gin as the router.
r := gin.Default()
r.GET("/custom-path", func(c *gin.Context) {
    c.JSON(http.StatusOK, map[string]interface{}{
    	"data": res,
	})
})
```