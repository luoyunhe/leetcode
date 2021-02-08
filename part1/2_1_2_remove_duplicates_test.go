package part1

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 2, 3, 3, 3, 3, 4},
			},
			want: []int{1, 1, 2, 2, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
