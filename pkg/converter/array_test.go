package converter

import (
	"reflect"
	"testing"
)

func TestToArrayOfInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Args that passed is not string, return nil",
			args: args{
				v: 1,
			},
			want: nil,
		},
		{
			name: "Args that passed is string invalid, return nil",
			args: args{
				v: "[1,2",
			},
			want: nil,
		},
		{
			name: "Args that passed is string valid, return nil",
			args: args{
				v: "[1,2]",
			},
			want: []int{1, 2},
		},
		{
			name: "Args that passed is slice string valid, return nil",
			args: args{
				v: []string{"1", "2"},
			},
			want: []int{1, 2},
		},
		{
			name: "Args that passed is byte array, return nil",
			args: args{
				v: [][]byte{
					[]byte("1234"),
					[]byte("5678"),
				},
			},
			want: []int{1234, 5678},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayOfInt(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayOfInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayOfString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Args that passed is not string, return nil",
			args: args{
				v: 1,
			},
			want: nil,
		},
		{
			name: "Args that passed is string invalid, return nil",
			args: args{
				v: "[1,2",
			},
			want: nil,
		},
		{
			name: "Args that passed is int array, return nil",
			args: args{
				v: "[\"1\",\"2\"]",
			},
			want: []string{"1", "2"},
		},
		{
			name: "Args that passed is string array, return nil",
			args: args{
				v: "[\"1\",\"2\"]",
			},
			want: []string{"1", "2"},
		},
		{
			name: "Args that passed is byte array, return nil",
			args: args{
				v: [][]byte{
					[]byte("foo"),
					[]byte("bar"),
				},
			},
			want: []string{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayOfString(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
