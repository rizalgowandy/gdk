package netx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPv4(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	enableLog := false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIPv4()
			assert.NotEmpty(t, got)

			if enableLog {
				t.Logf("got = \n%#v\n\n", got)
			}
		})
	}
}

func TestGetIPv16(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	enableLog := false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIPv16()
			assert.NotEmpty(t, got)

			if enableLog {
				t.Logf("got = \n%#v\n\n", got)
			}
		})
	}
}
