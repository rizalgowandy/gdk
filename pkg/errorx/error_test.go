package errorx

import (
	"errors"
	"testing"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name  string
		input *Error
	}{
		{
			name:  "No error",
			input: &Error{},
		},
		{
			name: "1 layer",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     nil,
			},
		},
		{
			name: "2 layer with standard error",
			input: &Error{
				Code:    Internal,
				Message: "Internal server error.",
				Op:      "userService.FindUserByID",
				Err:     errors.New("standard-error"),
			},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("\n\n" + tt.input.Error() + "\n")
		})
	}
}
