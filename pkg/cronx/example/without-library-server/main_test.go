package main

import (
	"context"
	"testing"

	"github.com/peractio/gdk/pkg/cronx"
)

func TestAlwaysError_Run(t *testing.T) {
	type args struct {
		in context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Error",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AlwaysError{}
			if err := a.Run(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEveryJob_Run(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := EveryJob{}
			if err := ev.Run(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscription_Run(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				in0: cronx.SetJobMetadata(context.Background(), cronx.JobMetadata{
					EntryID:    1,
					Wave:       2,
					TotalWave:  3,
					IsLastWave: true,
				}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := Subscription{}
			if err := su.Run(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendEmail_Run(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := SendEmail{}
			if err := e.Run(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPayBill_Run(t *testing.T) {
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PayBill{}
			if err := p.Run(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegisterJobs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cronx.Default()
			RegisterJobs()
		})
	}
}
