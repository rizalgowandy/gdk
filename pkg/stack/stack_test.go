package stack

import (
	"reflect"
	"testing"
)

func TestTrim(t *testing.T) {
	type args struct {
		stack []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Empty input",
			args: args{
				stack: nil,
			},
			want: nil,
		},
		{
			name: "Success, remove unnecessary start",
			args: args{
				stack: []byte(
					"random\nsrc/runtime/panic.go\ngithub.com/rizalgowandy\ntrace1\ntrace2\n",
				),
			},
			want: []byte("github.com/rizalgowandy\ntrace1\ntrace2\n"),
		},
		{
			name: "Success, remove unnecessary start and end",
			args: args{
				stack: []byte(
					"random\nsrc/runtime/panic.go\ngithub.com/rizalgowandy\ntrace1\ntrace2\nsrc/github.com/rizalgowandy/gdk/pkg/stack/stack.go\nrandom",
				),
			},
			want: []byte(
				"github.com/rizalgowandy\ntrace1\ntrace2\nsrc/github.com/rizalgowandy/gdk/pkg/stack/stack.go\n",
			),
		},
		{
			name: "Success, everything unnecessary, return without trim",
			args: args{
				stack: []byte("random\nabc\nqwerty"),
			},
			want: []byte("random\nabc\nqwerty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.stack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trim() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestToArr(t *testing.T) {
	type args struct {
		stack []byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty input",
			args: args{
				stack: nil,
			},
			want: nil,
		},
		{
			name: "Success, return with trim",
			args: args{
				stack: []byte("home/rizalgowandy/go/src/github.com/rizalgowandy/gdk/pkg/stack.go 130\n"),
			},
			want: []string{
				"github.com/rizalgowandy/gdk/pkg/stack.go 130",
			},
		},
		{
			name: "Success, 1 line match, return with trim",
			args: args{
				stack: []byte(
					"random\n/home/rizalgowandy/go/src/github.com/rizalgowandy/gdk/pkg/stack.go 130\nqwerty",
				),
			},
			want: []string{
				"random",
				"github.com/rizalgowandy/gdk/pkg/stack.go 130",
				"qwerty",
			},
		},
		{
			name: "Success, return without trim",
			args: args{
				stack: []byte("random\nabc\nqwerty"),
			},
			want: []string{
				"random",
				"abc",
				"qwerty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArr(tt.args.stack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
