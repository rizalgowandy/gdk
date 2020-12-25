package cronx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJob_Run(t *testing.T) {
	type fields struct {
		Name    string
		Status  StatusCode
		Latency string
		inner   JobItf
		status  uint32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success with run resulting error",
			fields: fields{
				Name:   "Func",
				Status: StatusCodeIdle,
				inner:  Func(func() error { return errors.New("error") }),
			},
		},
		{
			name: "Success",
			fields: fields{
				Name:   "Func",
				Status: StatusCodeIdle,
				inner:  Func(func() error { return nil }),
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
		inner   JobItf
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
				status: statusUp,
			},
			want: StatusCodeUp,
		},
		{
			name: "StatusCodeRunning",
			fields: fields{
				status: statusRunning,
			},
			want: StatusCodeRunning,
		},
		{
			name: "StatusCodeIdle",
			fields: fields{
				status: statusIdle,
			},
			want: StatusCodeIdle,
		},
		{
			name: "StatusCodeDown",
			fields: fields{
				status: statusDown,
			},
			want: StatusCodeDown,
		},
		{
			name: "StatusCodeError",
			fields: fields{
				status: statusError,
			},
			want: StatusCodeError,
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
		job JobItf
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				job: Func(func() error { return nil }),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewJob(tt.args.job))
		})
	}
}
