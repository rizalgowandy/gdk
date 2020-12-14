package errorx

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorString_Error(t *testing.T) {
	type fields struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success",
			fields: fields{
				s: "message",
			},
			want: "message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorString{
				s: tt.fields.s,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr(t *testing.T) {
	tests := []struct {
		name string
		args string
		want error
	}{
		{
			name: "Success",
			args: "standard",
			want: errors.New("standard"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Str(tt.args)
			if got == nil {
				t.Error("result should not be nil")
				return
			}

			if !reflect.DeepEqual(tt.want.Error(), got.Error()) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   string
		want   error
	}{
		{
			name:   "Success",
			format: "message: %s",
			args:   "standard",
			want:   errors.New("message: standard"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Errorf(tt.format, tt.args)
			if got == nil {
				t.Error("result should not be nil")
				return
			}

			if !reflect.DeepEqual(tt.want.Error(), got.Error()) {
				msg := "\nwant = %#v" + "\ngot  = %#v"
				t.Errorf(msg, tt.want, got)
			}
		})
	}
}
