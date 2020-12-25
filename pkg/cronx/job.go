package cronx

import (
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type JobItf interface {
	Run() error
}

type Job struct {
	Name    string     `json:"name"`
	Status  StatusCode `json:"status"`
	Latency string     `json:"latency"`
	Error   string     `json:"error"`

	inner   JobItf
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
	case statusError:
		j.Status = StatusCodeError
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
	atomic.StoreUint32(&j.status, statusRunning)
	j.UpdateStatus()

	// Update job status after running.
	defer j.UpdateStatus()

	// Run the job.
	if err := j.inner.Run(); err != nil {
		j.Error = err.Error()
		atomic.StoreUint32(&j.status, statusError)
	} else {
		atomic.StoreUint32(&j.status, statusIdle)
	}

	// Record time needed to execute the whole process.
	j.Latency = time.Since(start).String()
}

// NewJob creates a new job with default status and name.
func NewJob(job JobItf) *Job {
	name := reflect.TypeOf(job).Name()
	if name == "Func" {
		name = "(nameless)"
	}

	return &Job{
		Name:   name,
		Status: StatusCodeUp,
		inner:  job,
		status: statusUp,
	}
}
