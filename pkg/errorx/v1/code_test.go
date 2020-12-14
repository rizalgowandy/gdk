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
			want:  Internal,
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: Internal,
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: Internal,
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    Unknown,
				Message: "Unknown server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Permission,
					Message: "Permission error.",
					Op:      "accountGateway.FindUserByID",
					Err:     nil,
				},
			},
			want: Permission,
		},
		{
			name: "3 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Gateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err: &Error{
						Code:    Unknown,
						Message: "Unknown error.",
						Op:      "io.Write",
						Err:     nil,
					},
				},
			},
			want: Internal,
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
