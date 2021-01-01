package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetCode(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  Code
	}{
		{
			name:  "No error",
			input: nil,
			want:  "",
		},
		{
			name:  "Standard error",
			input: errors.New("standard-error"),
			want:  CodeInternal,
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Err:     nil,
			},
			want: CodeInternal,
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Err:     errors.New("standard-error"),
			},
			want: CodeInternal,
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    CodeUnknown,
				Message: "Unknown server error.",
				Err: &Error{
					Code:    CodePermission,
					Message: "Permission error.",
					Err:     nil,
				},
			},
			want: CodePermission,
		},
		{
			name: "3 layer",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Err: &Error{
					Code:    CodeGateway,
					Message: "Gateway server error.",
					Err: &Error{
						Code:    CodeUnknown,
						Message: "Unknown error.",
						Err:     nil,
					},
				},
			},
			want: CodeInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCode(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
