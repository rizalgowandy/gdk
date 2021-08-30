package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	cases := []struct {
		err1        error
		err2        error
		expectMatch bool
	}{
		{
			err1:        E(errors.New("error")),
			err2:        nil,
			expectMatch: false,
		},
		{
			err1:        E(errors.New("error")),
			err2:        errors.New("error"),
			expectMatch: true,
		},
		{
			err1:        E(errors.New("error")),
			err2:        errors.New("something is different"),
			expectMatch: false,
		},
	}

	for _, val := range cases {
		if match := Match(val.err1, val.err2); match != val.expectMatch {
			t.Errorf("Match() want = %vv, got %v", val.expectMatch, match)
		}
	}
}

func TestE(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
	}{
		{
			name: "No args",
			args: nil,
		},
		{
			name: "1 layer",
			args: []interface{}{
				Message("message"),
				CodeConflict,
				MetricStatusSuccess,
				Fields{
					"K": "V",
				},
				Op("userService.CreateUser"),
			},
		},
		{
			name: "2 layer with standard error",
			args: []interface{}{
				errors.New("standard-error"),
				Message("message"),
				CodeConflict,
				Op("userService.CreateUser"),
			},
		},
		{
			name: "2 layer",
			args: []interface{}{
				&Error{
					Code:     CodeGateway,
					Message:  "gateway-message",
					OpTraces: []Op{"userGateway.FindUser"},
					Err:      errors.New("standard-error"),
				},
				Message("message"),
				CodeConflict,
				Op("userService.CreateUser"),
			},
		},
		{
			name: "3 layer with fields merged",
			args: []interface{}{
				E(
					E(
						New("timeout"),
						Fields{
							"user_id":   1,
							"random_id": 2,
							"actor_id":  3,
						},
						Op("userGateway.FindUser"),
					),
					Fields{
						"actor_id": 777,
					},
					Message("handler-message"),
					Op("userHandler.FindUser"),
				),
			},
		},
		{
			name: "Invalid type",
			args: []interface{}{
				123,
			},
		},
		{
			name: "Same code",
			args: []interface{}{
				&Error{
					Code:     CodeConflict,
					Message:  "gateway-message",
					OpTraces: []Op{"userGateway.FindUser"},
					Err:      errors.New("standard-error"),
				},
				Message("message"),
				CodeConflict,
				Op("userService.CreateUser"),
			},
		},
		{
			name: "Missing code",
			args: []interface{}{
				&Error{
					Code:    CodeConflict,
					Message: "gateway-message",
					Err:     errors.New("standard-error"),
				},
				Message("message"),
				Op("userService.CreateUser"),
			},
		},
		{
			name: "New from string",
			args: []interface{}{
				"this is an error",
				Message("message"),
				Op("userService.CreateUser"),
			},
		},
	}

	enableLog := false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := E(tt.args...)
			if enableLog {
				t.Logf("got = \n%#v\n\n", got)
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
				code: CodeGateway,
				err: &Error{
					Code:    CodeGateway,
					Message: "",
					Err:     nil,
				},
			},
			want: true,
		},
		{
			name: "2 layer",
			args: args{
				code: CodeGateway,
				err: &Error{
					Code:    CodeUnknown,
					Message: "",
					Err: &Error{
						Code:    CodeGateway,
						Message: "",
						Err:     nil,
					},
				},
			},
			want: true,
		},
		{
			name: "2 layer Unknown",
			args: args{
				code: CodeGateway,
				err: &Error{
					Code:    CodeUnknown,
					Message: "",
					Err: &Error{
						Code:    CodeUnknown,
						Message: "",
						Err:     nil,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Is(tt.args.err, tt.args.code)
			if !reflect.DeepEqual(tt.want, got) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
