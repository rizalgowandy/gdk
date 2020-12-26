package timex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJakartaLocation(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetJakartaLocation()
			assert.NotNil(t, got)
		})
	}
}
