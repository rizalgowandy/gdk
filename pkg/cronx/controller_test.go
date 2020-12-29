package cronx

import (
	"context"
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
)

func TestCommandController_Start(t *testing.T) {
	type fields struct {
		Commander    *cron.Cron
		WorkerPool   chan struct{}
		PanicRecover func(ctx context.Context, j *Job)
		Address      string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success without server",
			fields: fields{
				Commander:    cron.New(),
				WorkerPool:   nil,
				PanicRecover: nil,
				Address:      "",
			},
		},
		{
			name: "Success with server",
			fields: fields{
				Commander:    cron.New(),
				WorkerPool:   nil,
				PanicRecover: nil,
				Address:      ":8998",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandController{
				Commander:    tt.fields.Commander,
				WorkerPool:   tt.fields.WorkerPool,
				PanicRecover: tt.fields.PanicRecover,
				Address:      tt.fields.Address,
			}
			c.Start()
		})
	}
}

func TestNewCommandController(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCommandController(tt.args.config)
			assert.NotNil(t, got)
		})
	}
}
