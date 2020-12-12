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
	Error   string     `json:"error"`

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
	case statusDown:
		j.Status = StatusCodeDown
	default:
		j.Status = StatusCodeUp
	}
	return j.Status
}

// Run executes the current job operation.
func (j *Job) Run() {
	start := time.Now()
	defer commandController.PanicRecover(j)

	// Lock current process.
	j.running.Lock()
	defer j.running.Unlock()

	// Wait for worker to be available.
	commandController.WorkerPool <- struct{}{}
	defer func() {
		<-commandController.WorkerPool
	}()

	// Update job status as running.
	atomic.StoreUint32(&j.status, 1)
	j.UpdateStatus()

	// Update job status after running.
	defer j.UpdateStatus()
	defer atomic.StoreUint32(&j.status, 2)

	// Run the job.
	j.inner.Run()

	// Record time needed to execute the whole process.
	j.Latency = time.Since(start).String()
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
