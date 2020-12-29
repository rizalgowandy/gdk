package cronx

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// Config defines the config for the command controller.
type Config struct {
	// Address determines the address will we serve the json and frontend status.
	// Empty string meaning we won't serve the current job status.
	// Default ":8998".
	Address string

	// Location describes the timezone current cron is running.
	// By default the timezone will be the same timezone as the server.
	Location *time.Location
}

var (
	defaultConfig = Config{
		Address:  ":8998",
		Location: time.Local,
	}

	commandController *CommandController
)

// Default creates a cron with default config.
func Default(interceptors ...Interceptor) {
	New(defaultConfig, interceptors...)
}

// New creates a cron with custom config.
func New(config Config, interceptors ...Interceptor) {
	// If there is invalid config use the default config instead.
	if config.Location == nil {
		config.Location = time.Local
	}

	// Create new command controller and start the underlying jobs.
	commandController = NewCommandController(config, interceptors...)
	commandController.Start()
}

// Func is a type to allow callers to wrap a raw func.
// Example:
//	cronx.Schedule("@every 5m", cronx.Func(myFunc))
type Func func(ctx context.Context) error

func (r Func) Run(ctx context.Context) error {
	return r(ctx)
}

// Schedule sets a job to run at specific time.
// Example:
//  @every 5m
//  0 */10 * * * * => every 10m
func Schedule(spec string, job JobItf) error {
	if commandController == nil || commandController.Commander == nil {
		return errors.New("cronx has not been initialized")
	}

	// Check if spec is correct.
	schedule, err := commandController.Parser.Parse(spec)
	if err != nil {
		downJob := NewJob(job)
		downJob.Status = StatusCodeDown
		downJob.Error = err.Error()
		commandController.UnregisteredJobs = append(
			commandController.UnregisteredJobs,
			downJob,
		)
		return err
	}

	commandController.Commander.Schedule(schedule, NewJob(job))
	return nil
}

// Schedules sets a job to run multiple times at specific time.
// Symbol */,-? should never be used as separator character.
// These symbols are reserved for cron specification.
//
// Example:
//	Spec		: "0 0 1 * * *#0 0 2 * * *#0 0 3 * * *
//	Separator	: "#"
//	This input schedules the job to run 3 times.
func Schedules(spec, separator string, job JobItf) error {
	if spec == "" {
		return errors.New("invalid specification")
	}
	if separator == "" {
		return errors.New("invalid separator")
	}
	schedules := strings.Split(spec, separator)
	for _, v := range schedules {
		if err := Schedule(v, job); err != nil {
			return err
		}
	}
	return nil
}

// Every executes the given job at a fixed interval.
// The interval provided is the time between the job ending and the job being run again.
// The time that the job takes to run is not included in the interval.
// Minimal time is 1 sec.
func Every(duration time.Duration, job JobItf) {
	if commandController == nil || commandController.Commander == nil {
		return
	}

	commandController.Commander.Schedule(cron.Every(duration), NewJob(job))
}

// Stop stops active jobs from running at the next scheduled time.
func Stop() {
	if commandController == nil || commandController.Commander == nil {
		return
	}

	commandController.Commander.Stop()
}

// GetEntries return all the current registered jobs.
func GetEntries() []cron.Entry {
	if commandController == nil || commandController.Commander == nil {
		return nil
	}

	return commandController.Commander.Entries()
}

// Remove removes a specific job from running.
// Get EntryID from the list job entries cronx.GetEntries().
// If job is in the middle of running, once the process is finished it will be removed.
func Remove(id cron.EntryID) {
	if commandController == nil || commandController.Commander == nil {
		return
	}

	commandController.Commander.Remove(id)
}
