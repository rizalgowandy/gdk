package cronx

import (
	"time"

	"github.com/robfig/cron/v3"
)

// StatusCode describes current job status.
type StatusCode string

const (
	// StatusCodeUp describes that current job has just been created.
	StatusCodeUp StatusCode = "UP"
	// StatusCodeIdle describes that current job is waiting for next execution time.
	StatusCodeIdle StatusCode = "IDLE"
	// StatusCodeRunning describes that current job is currently running.
	StatusCodeRunning StatusCode = "RUNNING"
	// StatusCodeDown describes that current job has failed to be registered.
	StatusCodeDown StatusCode = "DOWN"

	statusDown    uint32 = 0
	statusUp      uint32 = 1
	statusIdle    uint32 = 2
	statusRunning uint32 = 3
)

// StatusData defines current job status.
type StatusData struct {
	// ID is unique per job.
	ID cron.EntryID `json:"id,omitempty"`
	// Job defines current job.
	Job *Job `json:"job,omitempty"`
	// Next defines the next schedule to execute current job.
	Next time.Time `json:"next_run,omitempty"`
	// Prev defines the last run of the current job.
	Prev time.Time `json:"prev_run,omitempty"`
}

// GetStatusData returns all jobs status.
func GetStatusData() []StatusData {
	if commandController == nil || commandController.Commander == nil {
		return nil
	}

	entries := commandController.Commander.Entries()
	totalEntries := len(entries)

	downs := commandController.UnregisteredJobs
	totalDowns := len(downs)

	totalJobs := totalEntries + totalDowns
	listStatus := make([]StatusData, totalJobs)

	// Register down jobs.
	for k, v := range downs {
		listStatus[k].Job = v
	}

	// Register other jobs.
	for k, v := range entries {
		idx := totalDowns + k
		listStatus[idx].ID = v.ID
		listStatus[idx].Job = v.Job.(*Job)
		listStatus[idx].Next = v.Next
		listStatus[idx].Prev = v.Prev
	}

	return listStatus
}

// GetStatusJSON returns all jobs status as map[string]interface.
func GetStatusJSON() map[string]interface{} {
	return map[string]interface{}{
		"data": GetStatusData(),
	}
}
