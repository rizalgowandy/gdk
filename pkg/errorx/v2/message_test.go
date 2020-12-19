package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  string
	}{
		{
			name:  "No error",
			input: nil,
			want:  "",
		},
		{
			name:  "Standard error",
			input: errors.New("standard-error"),
			want:  DefaultMessage,
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: "Internal server error.",
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    CodeInternal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: "Internal server error.",
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    CodeUnknown,
				Message: "",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    CodeGateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err:     nil,
				},
			},
			want: "Gateway server error.",
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
			want: "Internal server error.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMessage(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
