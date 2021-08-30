package build

import (
	"reflect"
	"testing"

	"github.com/rizalgowandy/gdk/pkg/tags"
)

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			name: "Success with empty",
			want: nil,
		},
		{
			name: "Success",
			want: map[string]string{
				tags.Hash: "H",
				tags.Time: "T",
				tags.User: "U",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			onceInfo.Reset()

			if tt.want != nil {
				md = &Metadata{
					Hash: tt.want[tags.Hash],
					User: tt.want[tags.User],
					Time: tt.want[tags.Time],
				}
			} else {
				md = nil
			}

			if got := Info(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		hash string
		user string
		time string
	}
	tests := []struct {
		name string
		args args
		want *Metadata
	}{
		{
			name: "Empty",
			args: args{},
			want: nil,
		},
		{
			name: "Success",
			args: args{
				hash: "A",
				user: "B",
				time: "C",
			},
			want: &Metadata{
				Hash: "A",
				User: "B",
				Time: "C",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.hash, tt.args.user, tt.args.time); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
