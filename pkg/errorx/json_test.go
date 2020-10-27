package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  []byte
	}{
		{
			name:  "No error",
			input: nil,
			want:  nil,
		},
		{
			name:  "Standard error",
			input: errors.New("standard-error"),
			want:  nil,
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: []byte(
				`[{"code":"internal","message":"Internal server error.","op":"userService.FindUserByID"}]`,
			),
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: []byte(
				`[{"code":"internal","message":"Internal server error.","op":"userService.FindUserByID"},{"code":"standard","message":"standard-error"}]`,
			),
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
			want: []byte(
				`[{"message":"Unknown server error.","op":"userService.FindUserByID"},{"code":"permission","message":"Permission error.","op":"accountGateway.FindUserByID"}]`,
			),
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
			want: []byte(
				`[{"code":"internal","message":"Internal server error.","op":"userService.FindUserByID"},{"code":"gateway","message":"Gateway server error.","op":"accountGateway.FindUserByID"},{"message":"Unknown error.","op":"io.Write"}]`,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetJSON(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, string(tt.want), string(got))
			}
			t.Log(string(got))
		})
	}
}
