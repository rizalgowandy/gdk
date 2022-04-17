package sortx

import (
	"reflect"
	"testing"
)

func TestNewSorts(t *testing.T) {
	type args struct {
		qs string
	}
	tests := []struct {
		name string
		args args
		want []Sort
	}{
		{
			name: "Success",
			args: args{
				qs: "id:desc,status:asc,created_at",
			},
			want: []Sort{
				{
					Key:   "id",
					Order: OrderDescending,
				},
				{
					Key:   "status",
					Order: OrderAscending,
				},
				{
					Key:   "created_at",
					Order: OrderAscending,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSorts(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
