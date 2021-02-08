package part1

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 2, 2, 3, 3, 4, 5, 5, 5, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates1() = %v, want %v", got, tt.want)
			}
		})
	}
}
