package try

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	MaxRetries = 1
	tests := []struct {
		name  string
		input func(attempt int) (retry bool, err error)
		want  error
	}{
		{
			name: "No error",
			input: func(attempt int) (retry bool, err error) {
				return attempt < 1, nil
			},
			want: nil,
		},
		{
			name: "Error max retries reached",
			input: func(attempt int) (retry bool, err error) {
				return attempt < 5, errors.New("result")
			},
			want: errMaxRetriesReached,
		},
		{
			name: "Error",
			input: func(attempt int) (retry bool, err error) {
				return attempt < 1, errors.New("result")
			},
			want: errors.New("result"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Do(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsMaxRetries(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  bool
	}{
		{
			name:  "Success",
			input: errMaxRetriesReached,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsMaxRetries(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
