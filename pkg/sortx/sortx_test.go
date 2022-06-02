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
		want Sorts
	}{
		{
			name: "Success",
			args: args{
				qs: "id:desc,status:asc,created_at",
			},
			want: Sorts{
				{
					Key:      "id",
					Order:    OrderDescending,
					Original: "id:desc",
				},
				{
					Key:      "status",
					Order:    OrderAscending,
					Original: "status:asc",
				},
				{
					Key:      "created_at",
					Order:    OrderAscending,
					Original: "created_at",
				},
			},
		},
		{
			name: "Success with different case",
			args: args{
				qs: "ID:DESC,Status:Asc,created_at",
			},
			want: Sorts{
				{
					Key:      "id",
					Order:    OrderDescending,
					Original: "id:desc",
				},
				{
					Key:      "status",
					Order:    OrderAscending,
					Original: "status:asc",
				},
				{
					Key:      "created_at",
					Order:    OrderAscending,
					Original: "created_at",
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
