# Cronx
Cronx is a wrapper for _robfig/cron_.
It includes a live monitoring of current schedule and state of active jobs that can be outputted as JSON or HTML template.

## Available Status
* **Down** => Job fails to be registered.
* **Up** => Job has just been created.
* **Running** => Job is currently running.
* **Idle** => Job is waiting for next execution time.
* **Error** => Job fails on the last run.

## Quick Start
Create a _**main.go**_ file.
```go
package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/cronx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// In order to create a job you need to create a struct that has Run() method.
type sendEmail struct{}

func (s sendEmail) Run(ctx context.Context) error {
	log.WithLevel(zerolog.InfoLevel).
		Str("job", "sendEmail").
		Msg("every 5 sec send reminder emails")
	return nil
}

func main() {
	// Create a cron controller with default config that:
	// - running on port :8998
	// - location is time.Local
	// - without any middleware
	cronx.Default()
	defer cronx.Stop()

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
Get dependencies
```shell
$ go mod vendor -v
```

Start server
```shell
$ go run main.go
```

Browse to
- http://localhost:8998 => see server health status.
- http://localhost:8998/jobs => see the html page.
- http://localhost:8998/api/jobs => see the json response.
```json
{
  "data": [
    {
      "id": 1,
      "job": {
        "name": "sendEmail",
        "status": "RUNNING",
        "latency": "3.000299794s",
        "error": ""
      },
      "next_run": "2020-12-11T22:36:35+07:00",
      "prev_run": "2020-12-11T22:36:30+07:00"
    }
  ]
}
```

## Custom Configuration
```go
// Create a cron with custom config.
cronx.New(cronx.Config{
    Address:  ":8998", // Determines if we want the library to serve the frontend.
    Location: func() *time.Location { // Change timezone to Jakarta.
        jakarta, err := time.LoadLocation("Asia/Jakarta")
        if err != nil {
            secondsEastOfUTC := int((7 * time.Hour).Seconds())
            jakarta = time.FixedZone("WIB", secondsEastOfUTC)
        }
        return jakarta
    }(),
})
```

## Interceptor / Middleware
Interceptor or commonly known as middleware is an operation that commonly executed before any of other operation. 
This library has the capability to add multiple middlewares that will be executed before or after the real job.
It means you can log the running job, send telemetry, or protect the application from going down because of panic by adding middlewares.
The idea of a middleware is to be declared once, and be executed on all registered jobs.
Hence, reduce the code duplication on each job implementation.

### Adding Interceptor / Middleware
```go
// Create cron middleware.
// The order is important.
// The first one will be executed first.
middleware := cronx.Chain(
    interceptor.Recover(), // Auto recover from panic.
    interceptor.Logger(), // Log start and finish process.
    interceptor.DefaultWorkerPool(), // Limit concurrent running job.
)

cronx.Default(middleware)
```

### Custom Interceptor / Middleware
```go
// Sleep is a middleware that sleep a few second after job has been executed.
func Sleep() cronx.Interceptor {
	return func(ctx context.Context, job *cronx.Job, handler cronx.Handler) error {
		err := handler(ctx, job)
		time.Sleep(10 * time.Second)
		return err
	}
}
```

## Schedule Specification Format

### Schedule
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Optional   | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

### Predefined schedules
Entry                  | Description                                | Equivalent
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *

### Intervals
```
@every <duration>
```
For example, "@every 1h30m10s" would indicate a schedule that activates after 1 hour, 30 minutes, 10 seconds, and then every interval after that.

Please refer to this [link](https://pkg.go.dev/github.com/robfig/cron?readme=expanded#section-readme/) for more detail.

## FAQ

### What are the available commands?
Here the list of commonly used commands.
```go
// Schedule sets a job to run at specific time.
// Example:
//  @every 5m
//  0 */10 * * * * => every 10m
Schedule(spec string, job JobItf) error

// Schedules sets a job to run multiple times at specific time.
// Symbol */,-? should never be used as separator character.
// These symbols are reserved for cron specification.
//
// Example:
//  Spec		: "0 0 1 * * *#0 0 2 * * *#0 0 3 * * *
//  Separator	: "#"
//  This input schedules the job to run 3 times.
Schedules(spec, separator string, job JobItf) error

// Every executes the given job at a fixed interval.
// The interval provided is the time between the job ending and the job being run again.
// The time that the job takes to run is not included in the interval.
// Minimal time is 1 sec.
Every(duration time.Duration, job JobItf)
```
Go to `cronx/cronx.go` to see the list of available commands.

### What are the available interceptors?
Go to `cronx/interceptors` to see the available interceptors.

### Can I use my own router without starting the built-in router?
Yes, you can. This library is very modular.
```go
// Create a custom config and leave the address as empty string.
// Empty string meaning the library won't start the built-in server.
cronx.New(cronx.Config{
    Address:  "",
})

// GetStatusData will return the []cronx.StatusData.
// You can use this data like any other Golang data structure.
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

### Can I still get the built-in template if I use my own router?
Yes, you can.
```go
// GetStatusTemplate will return the built-in status page template.
index, _ := page.GetStatusTemplate()

// An example using echo as the router.
e := echo.New()
index, _ := page.GetStatusTemplate()
e.GET("jobs/html", func(context echo.Context) error {
    // Serve the template to the writer and pass the current status data.
    return index.Execute(context.Response().Writer, cronx.GetStatusData())
})
```

### Server is located in the US, but my user is in Jakarta, can I change the cron timezone?
Yes, you can.
By default, the cron timezone will follow the server location timezone using `time.Local`.
If you placed the server in the US, it will use the US timezone.
If you placed the server in the SG, it will use the SG timezone.
```go
// Create a custom config.
cronx.New(cronx.Config{
    Address:  ":8998",
    Location: func() *time.Location { // Change timezone to Jakarta.
        jakarta, err := time.LoadLocation("Asia/Jakarta")
        if err != nil {
            secondsEastOfUTC := int((7 * time.Hour).Seconds())
            jakarta = time.FixedZone("WIB", secondsEastOfUTC)
        }
        return jakarta
    }(),
})
```

### My job requires certain information like current wave number, how can I get this information?
This kind of information is stored inside metadata, which stored automatically inside `context`. 
```go
type subscription struct{}

func (subscription) Run(ctx context.Context) error {
	md, ok := cronx.GetJobMetadata(ctx)
	if !ok {
		return errors.New("cannot job metadata")
	}

	log.WithLevel(zerolog.InfoLevel).
		Str("job", "subscription").
		Interface("metadata", md).
		Msg("is running")
	return nil
}
```
