package cronx

import (
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
)

func TestCommandController_Start(t *testing.T) {
	type fields struct {
		Commander *cron.Cron
		Address   string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success without server",
			fields: fields{
				Commander: cron.New(),
				Address:   "",
			},
		},
		{
			name: "Success with server",
			fields: fields{
				Commander: cron.New(),
				Address:   ":8998",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandController{
				Commander:        tt.fields.Commander,
				Interceptor:      nil,
				Address:          tt.fields.Address,
				Location:         nil,
				CreatedTime:      time.Time{},
				Parser:           cron.Parser{},
				UnregisteredJobs: nil,
			}
			c.Start()
		})
	}
}

func TestNewCommandController(t *testing.T) {
	type args struct {
		config       Config
		interceptors Interceptor
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
			got := NewCommandController(tt.args.config, tt.args.interceptors)
			assert.NotNil(t, got)
		})
	}
}
