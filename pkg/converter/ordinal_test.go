package converter

import (
	"testing"
)

func TestOrdinal(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1st",
			args: args{
				x: 1,
			},
			want: "1st",
		},
		{
			name: "2nd",
			args: args{
				x: 2,
			},
			want: "2nd",
		},
		{
			name: "3rd",
			args: args{
				x: 3,
			},
			want: "3rd",
		},
		{
			name: "11th",
			args: args{
				x: 11,
			},
			want: "11th",
		},
		{
			name: "22nd",
			args: args{
				x: 22,
			},
			want: "22nd",
		},
		{
			name: "33rd",
			args: args{
				x: 33,
			},
			want: "33rd",
		},
		{
			name: "21st",
			args: args{
				x: 21,
			},
			want: "21st",
		},
		{
			name: "101st",
			args: args{
				x: 101,
			},
			want: "101st",
		},
		{
			name: "202nd",
			args: args{
				x: 202,
			},
			want: "202nd",
		},
		{
			name: "303rd",
			args: args{
				x: 303,
			},
			want: "303rd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ordinal(tt.args.x); got != tt.want {
				t.Errorf("Ordinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
