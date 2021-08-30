package logx

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/tags"
	"github.com/stretchr/testify/assert"
)

func TestErrMetadata(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Success with nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "Success with standard error",
			args: args{
				err: fmt.Errorf("abc"),
			},
			want: map[string]interface{}{},
		},
		{
			name: "Success with custom error",
			args: args{
				err: errorx.E("abc", errorx.Fields{
					"k": "v",
				}, errorx.CodeInternal, errorx.Op("abc"), errorx.MetricStatusExpectedErr, errorx.Message("qwerty")),
			},
			want: map[string]interface{}{
				tags.Detail: map[string]interface{}{
					tags.Code:         errorx.CodeInternal,
					tags.Fields:       errorx.Fields{"k": "v"},
					tags.Message:      errorx.Message("qwerty"),
					tags.MetricStatus: errorx.MetricStatusExpectedErr,
					tags.Ops: []errorx.Op{
						errorx.Op("logx.TestErrMetadata"),
					},
					tags.ErrorLine: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrMetadata(tt.args.err)

			// Inject got err origin to want because it's gonna be different on each machine.
			if err, ok := got[tags.Detail]; ok {
				if origin, ok := (err.(map[string]interface{}))[tags.ErrorLine]; ok {
					if err2, ok := tt.want[tags.Detail]; ok {
						if _, ok := (err2.(map[string]interface{}))[tags.ErrorLine]; ok {
							(err2.(map[string]interface{}))[tags.ErrorLine] = origin
						}
					}
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata(t *testing.T) {
	type args struct {
		detail interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Success without detail",
			args: args{},
			want: map[string]interface{}{},
		},
		{
			name: "Success",
			args: args{
				detail: map[string]interface{}{
					"k": "v",
				},
			},
			want: map[string]interface{}{
				tags.Detail: map[string]interface{}{
					"k": "v",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Metadata(tt.args.detail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
