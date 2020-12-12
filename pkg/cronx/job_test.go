package cronx

import (
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
)

func TestJob_Run(t *testing.T) {
	type fields struct {
		Name    string
		Status  StatusCode
		Latency string
		inner   cron.Job
		status  uint32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success",
			fields: fields{
				Name:   "Func",
				Status: StatusCodeIdle,
				inner:  Func(func() {}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Default()
			j := &Job{
				Name:    tt.fields.Name,
				Status:  tt.fields.Status,
				Latency: tt.fields.Latency,
				inner:   tt.fields.inner,
				status:  tt.fields.status,
			}
			j.Run()
		})
	}
}

func TestJob_UpdateStatus(t *testing.T) {
	type fields struct {
		Name    string
		Status  StatusCode
		Latency string
		inner   cron.Job
		status  uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   StatusCode
	}{
		{
			name: "StatusCodeUp",
			fields: fields{
				status: 0,
			},
			want: StatusCodeUp,
		},
		{
			name: "StatusCodeRunning",
			fields: fields{
				status: 1,
			},
			want: StatusCodeRunning,
		},
		{
			name: "StatusCodeIdle",
			fields: fields{
				status: 2,
			},
			want: StatusCodeIdle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				Name:    tt.fields.Name,
				Status:  tt.fields.Status,
				Latency: tt.fields.Latency,
				inner:   tt.fields.inner,
				status:  tt.fields.status,
			}
			if got := j.UpdateStatus(); got != tt.want {
				t.Errorf("UpdateStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJob(t *testing.T) {
	type args struct {
		job cron.Job
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				job: Func(func() {}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewJob(tt.args.job))
		})
	}
}
