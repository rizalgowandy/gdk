package cronx

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct {
	Name    string     `json:"name"`
	Status  StatusCode `json:"status"`
	Latency string     `json:"latency"`

	inner   cron.Job
	status  uint32
	running sync.Mutex
}

// UpdateStatus updates the current job status to the latest.
func (j *Job) UpdateStatus() StatusCode {
	switch atomic.LoadUint32(&j.status) {
	case statusRunning:
		j.Status = StatusCodeRunning
	case statusIdle:
		j.Status = StatusCodeIdle
	default:
		j.Status = StatusCodeUp
	}
	return j.Status
}

// Run executes the current job operation.
func (j *Job) Run() {
	start := time.Now()
	defer commandController.PanicRecover(j)

	j.running.Lock()
	defer j.running.Unlock()

	commandController.WorkerPool <- struct{}{}
	defer func() {
		<-commandController.WorkerPool
	}()

	atomic.StoreUint32(&j.status, 1)
	j.UpdateStatus()

	defer j.UpdateStatus()
	defer atomic.StoreUint32(&j.status, 2)

	j.inner.Run()

	end := time.Now()
	j.Latency = end.Sub(start).String()
}

// NewJob creates a new job with default status and name.
func NewJob(job cron.Job) *Job {
	return &Job{
		Name:   reflect.TypeOf(job).Name(),
		Status: StatusCodeUp,
		inner:  job,
		status: statusUp,
	}
}
