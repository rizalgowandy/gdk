package cronx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusData(t *testing.T) {
	tests := []struct {
		name string
		mock func()
	}{
		{
			name: "Success without any job",
			mock: func() {
				Start()
			},
		},
		{
			name: "Success",
			mock: func() {
				Start()
				_ = Schedule("@every 5m", Func(func() {}))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			assert.NotNil(t, GetStatusData())
		})
	}
}

func TestGetStatusJSON(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, GetStatusJSON())
		})
	}
}
