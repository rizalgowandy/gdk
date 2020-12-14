package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name  string
		input *Error
		want  string
	}{
		{
			name:  "No error",
			input: &Error{},
			want:  "no error",
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
			want: "userService.FindUserByID: <internal> Internal server error.",
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
			want: "userService.FindUserByID: <internal> Internal server error. => standard-error",
		},
		{
			name: "2 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Gateway,
					Message: "Gateway server error.",
					Op:      "accountGateway.FindUserByID",
					Err:     nil,
				},
			},
			want: "userService.FindUserByID: <internal> Internal server error.:\n\taccountGateway.FindUserByID: <gateway> Gateway server error.",
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
			want: "userService.FindUserByID: <internal> Internal server error.:\n\taccountGateway.FindUserByID: <gateway> Gateway server error.:\n\tio.Write: Unknown error.",
		},
		{
			name: "3 layer with Unknown",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Unknown,
					Message: "",
					Op:      "io.Write",
					Err: &Error{
						Code:    Gateway,
						Message: "Gateway server error.",
						Op:      "accountGateway.FindUserByID",
						Err:     nil,
					},
				},
			},
			want: "userService.FindUserByID: <internal> Internal server error.:\n\tio.Write:\n\taccountGateway.FindUserByID: <gateway> Gateway server error.",
		},
		{
			name: "3 layer with no Op",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Internal,
					Message: "",
					Op:      "",
					Err: &Error{
						Code:    Gateway,
						Message: "Gateway server error.",
						Op:      "accountGateway.FindUserByID",
						Err:     nil,
					},
				},
			},
			want: "userService.FindUserByID: <internal> Internal server error.:\n\t<internal>:\n\taccountGateway.FindUserByID: <gateway> Gateway server error.",
		},
		{
			name: "3 layer with no Op and Unknown",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err: &Error{
					Code:    Unknown,
					Message: "Random error.",
					Op:      "",
					Err: &Error{
						Code:    Gateway,
						Message: "Gateway server error.",
						Op:      "accountGateway.FindUserByID",
						Err:     nil,
					},
				},
			},
			want: "userService.FindUserByID: <internal> Internal server error.:\n\tRandom error.:\n\taccountGateway.FindUserByID: <gateway> Gateway server error.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Error()
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v\n"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
