package logx

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestContextWithRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{},
		},
		{
			name: "Success with context default",
			args: args{
				ctx: context.Background(),
			},
		},
		{
			name: "Success with context metadata",
			args: args{
				ctx: metadata.NewIncomingContext(
					context.Background(),
					metadata.Pairs(
						RequestID, "request-id",
					),
				),
			},
		},
		{
			name: "Success with context and built in request id",
			args: args{
				ctx: ContextWithRequestID(context.Background()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextWithRequestID(tt.args.ctx); GetRequestID(got) == "" {
				t.Errorf("ContextWithRequestID() = %v", GetRequestID(got))
			}
		})
	}
}

func TestGetRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success empty",
			args: args{},
			want: false,
		},
		{
			name: "Success with id",
			args: args{
				ctx: NewContext(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRequestID(tt.args.ctx); (got != "") != tt.want {
				t.Errorf("GetRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success with nil",
			args: args{
				ctx: nil,
				id:  "request-id",
			},
		},
		{
			name: "Success with context",
			args: args{
				ctx: context.Background(),
				id:  "request-id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetRequestID(tt.args.ctx, tt.args.id); GetRequestID(got) == "" {
				t.Errorf("SetRequestID() = %v", GetRequestID(got))
			}
		})
	}
}

func TestSetRequestIDFromMetadata(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "Nil",
			args: args{},
			want: nil,
		},
		{
			name: "Without metadata",
			args: args{
				ctx: context.Background(),
			},
			want: context.Background(),
		},
		{
			name: "With metadata",
			args: args{
				ctx: metadata.NewIncomingContext(
					context.Background(),
					metadata.Pairs(RequestID, "abc"),
				),
			},
			want: func() context.Context {
				c := metadata.NewIncomingContext(
					context.Background(),
					metadata.Pairs(RequestID, "abc"),
				)
				c = SetRequestID(c, "abc")
				return c
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetRequestIDFromMetadata(tt.args.ctx); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("SetRequestIDFromMetadata() = \n\n%v, want \n\n%v", got, tt.want)
			}
		})
	}
}

func TestGenRequestID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Success",
			want: "should not be empty",
		},
	}

	enableLog := false
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenRequestID()
			if (got != "") != (tt.want != "") {
				t.Errorf("GenRequestID() = %v, want %v", got, tt.want)
			}
			if enableLog {
				t.Logf("got = \n%#v\n\n", got)
			}
		})
	}
}

func TestNewContext(t *testing.T) {
	type args struct {
		in []context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "New",
			args: args{},
		},
		{
			name: "Exists",
			args: args{
				in: []context.Context{
					context.Background(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContext(tt.args.in...); GetRequestID(got) == "" {
				t.Errorf("NewContext() = %v", GetRequestID(got))
			}
		})
	}
}
