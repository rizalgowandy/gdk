package cronx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusData(t *testing.T) {
	tests := []struct {
		name string
		mock func()
		want bool
	}{
		{
			name: "Uninitialized",
			mock: func() {
				commandController = nil
			},
		},
		{
			name: "Success without any job",
			mock: func() {
				Default()
			},
			want: true,
		},
		{
			name: "Success",
			mock: func() {
				Default()
				_ = Schedule("@every 5m", Func(func(ctx context.Context) error { return nil }))
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got := GetStatusData()
			if tt.want {
				assert.NotNil(t, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestGetStatusJSON(t *testing.T) {
	tests := []struct {
		name string
		mock func()
		want bool
	}{
		{
			name: "Uninitialized",
			mock: func() {
				commandController = nil
			},
		},
		{
			name: "Success without any job",
			mock: func() {
				Default()
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got := GetStatusJSON()
			if tt.want {
				assert.NotNil(t, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}
