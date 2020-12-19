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
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: CodeInternal,
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: CodeInternal,
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    CodeUnknown,
				Message: "Unknown server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    CodePermission,
					Message: "Permission error.",
					Op:      "accountGateway.FindUserByID",
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
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    CodeGateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err: &Error{
						Code:    CodeUnknown,
						Message: "Unknown error.",
						Op:      "io.Write",
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
