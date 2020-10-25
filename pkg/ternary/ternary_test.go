package ternary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	type args struct {
		condition bool
		yes       string
		no        string
	}
	tests := []struct {
		name     string
		input    args
		expected string
	}{
		{
			name: "Condition true, return true",
			input: args{
				condition: true,
				yes:       "true",
				no:        "false",
			},
			expected: "true",
		},
		{
			name: "Condition false, return false",
			input: args{
				condition: false,
				yes:       "true",
				no:        "false",
			},
			expected: "false",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Str(test.input.condition, test.input.yes, test.input.no)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		condition bool
		yes       int64
		no        int64
	}
	tests := []struct {
		name     string
		input    args
		expected int64
	}{
		{
			name: "Condition true, return 1",
			input: args{
				condition: true,
				yes:       int64(1),
				no:        int64(0),
			},
			expected: int64(1),
		},
		{
			name: "Condition false, return 0",
			input: args{
				condition: false,
				yes:       int64(1),
				no:        int64(0),
			},
			expected: int64(0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Int64(test.input.condition, test.input.yes, test.input.no)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		condition bool
		yes       int
		no        int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{
			name: "Condition true, return 1",
			input: args{
				condition: true,
				yes:       1,
				no:        0,
			},
			expected: 1,
		},
		{
			name: "Condition false, return 0",
			input: args{
				condition: false,
				yes:       1,
				no:        0,
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Int(test.input.condition, test.input.yes, test.input.no)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		condition bool
		yes       float64
		no        float64
	}
	tests := []struct {
		name     string
		input    args
		expected float64
	}{
		{
			name: "Condition true, return 1",
			input: args{
				condition: true,
				yes:       float64(1),
				no:        float64(0),
			},
			expected: float64(1),
		},
		{
			name: "Condition false, return 0",
			input: args{
				condition: false,
				yes:       float64(1),
				no:        float64(0),
			},
			expected: float64(0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Float64(test.input.condition, test.input.yes, test.input.no)
			assert.Equal(t, test.expected, actual)
		})
	}
}
