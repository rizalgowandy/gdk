package errorx

import (
	"errors"
	"reflect"
	"runtime/debug"
	"testing"
)

func TestE(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want error
	}{
		{
			name: "No args",
			args: nil,
			want: nil,
		},
		{
			name: "1 layer",
			args: []interface{}{
				"message",
				Conflict,
				Op("userService.CreateUser"),
			},
			want: &Error{
				Code:    Conflict,
				Message: "message",
				Op:      "userService.CreateUser",
				Err:     nil,
			},
		},
		{
			name: "2 layer with standard error",
			args: []interface{}{
				"message",
				Conflict,
				Op("userService.CreateUser"),
				errors.New("standard-error"),
			},
			want: &Error{
				Code:    Conflict,
				Message: "message",
				Op:      "userService.CreateUser",
				Err:     errors.New("standard-error"),
			},
		},
		{
			name: "2 layer",
			args: []interface{}{
				"message",
				Conflict,
				Op("userService.CreateUser"),
				&Error{
					Code:    Gateway,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
			want: &Error{
				Code:    Conflict,
				Message: "message",
				Op:      "userService.CreateUser",
				Err: &Error{
					Code:    Gateway,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
		},
		{
			name: "Invalid type",
			args: []interface{}{
				123,
			},
			want: errors.New("unknown type int, value 123 in error call"),
		},
		{
			name: "Same code",
			args: []interface{}{
				"message",
				Conflict,
				Op("userService.CreateUser"),
				&Error{
					Code:    Conflict,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
			want: &Error{
				Code:    Conflict,
				Message: "message",
				Op:      "userService.CreateUser",
				Err: &Error{
					Code:    Unknown,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
		},
		{
			name: "Missing code",
			args: []interface{}{
				"message",
				Op("userService.CreateUser"),
				&Error{
					Code:    Conflict,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
			want: &Error{
				Code:    Conflict,
				Message: "message",
				Op:      "userService.CreateUser",
				Err: &Error{
					Code:    Unknown,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
		},
		{
			name: "Missing code",
			args: []interface{}{
				Internal,
				Op("userService.CreateUser"),
				&Error{
					Code:    Conflict,
					Message: "gateway-message",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
			want: &Error{
				Code:    Internal,
				Message: "gateway-message",
				Op:      "userService.CreateUser",
				Err: &Error{
					Code:    Conflict,
					Message: "",
					Op:      "userGateway.FindUser",
					Err:     nil,
				},
			},
		},
	}

	enableLog := false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					if enableLog {
						t.Logf("\nrecover = %#v \nstack =\n%s", err, string(debug.Stack()))
					}
				}
			}()

			got := E(tt.args...)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	type args struct {
		err1 error
		err2 error
	}
	tests := []struct {
		name string
		args
		want bool
	}{
		{
			name: "err1 is nil",
			args: args{
				err1: nil,
				err2: nil,
			},
			want: false,
		},
		{
			name: "err2 is nil",
			args: args{
				err1: &Error{},
				err2: nil,
			},
			want: false,
		},
		{
			name: "Message unequal",
			args: args{
				err1: &Error{
					Code:    "",
					Message: "err1",
					Op:      "",
					Err:     nil,
				},
				err2: &Error{
					Code:    "",
					Message: "err2",
					Op:      "",
					Err:     nil,
				},
			},
			want: false,
		},
		{
			name: "Op unequal",
			args: args{
				err1: &Error{
					Code:    "",
					Message: "",
					Op:      "userService.FindUser",
					Err:     nil,
				},
				err2: &Error{
					Code:    "",
					Message: "",
					Op:      "userService.FindUserByID",
					Err:     nil,
				},
			},
			want: false,
		},
		{
			name: "Code unequal",
			args: args{
				err1: &Error{
					Code:    Conflict,
					Message: "",
					Op:      "",
					Err:     nil,
				},
				err2: &Error{
					Code:    Gateway,
					Message: "",
					Op:      "",
					Err:     nil,
				},
			},
			want: false,
		},
		{
			name: "2 layers",
			args: args{
				err1: &Error{
					Code:    "",
					Message: "",
					Op:      "",
					Err: &Error{
						Code:    Gateway,
						Message: "",
						Op:      "",
						Err:     nil,
					},
				},
				err2: &Error{
					Code:    Gateway,
					Message: "",
					Op:      "",
					Err:     nil,
				},
			},
			want: false,
		},
		{
			name: "2 layers with standard error",
			args: args{
				err1: &Error{
					Code:    "",
					Message: "",
					Op:      "",
					Err:     errors.New("abc"),
				},
				err2: &Error{
					Code:    Gateway,
					Message: "",
					Op:      "",
					Err:     nil,
				},
			},
			want: false,
		},
		{
			name: "Match",
			args: args{
				err1: &Error{
					Code:    "",
					Message: "",
					Op:      "",
					Err:     nil,
				},
				err2: &Error{
					Code:    Gateway,
					Message: "",
					Op:      "",
					Err:     nil,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Match(tt.args.err1, tt.args.err2)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestIs(t *testing.T) {
	type args struct {
		code Code
		err  error
	}
	tests := []struct {
		name string
		args
		want bool
	}{
		{
			name: "Error is nil",
			args: args{
				code: "",
				err:  nil,
			},
			want: false,
		},
		{
			name: "Error is standard error",
			args: args{
				code: "",
				err:  errors.New("standard-error"),
			},
			want: false,
		},
		{
			name: "1 layer",
			args: args{
				code: Gateway,
				err: &Error{
					Code:    Gateway,
					Message: "",
					Op:      "",
					Err:     nil,
				},
			},
			want: true,
		},
		{
			name: "2 layer",
			args: args{
				code: Gateway,
				err: &Error{
					Code:    Unknown,
					Message: "",
					Op:      "",
					Err: &Error{
						Code:    Gateway,
						Message: "",
						Op:      "",
						Err:     nil,
					},
				},
			},
			want: true,
		},
		{
			name: "2 layer Unknown",
			args: args{
				code: Gateway,
				err: &Error{
					Code:    Unknown,
					Message: "",
					Op:      "",
					Err: &Error{
						Code:    Unknown,
						Message: "",
						Op:      "",
						Err:     nil,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Is(tt.args.code, tt.args.err)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
