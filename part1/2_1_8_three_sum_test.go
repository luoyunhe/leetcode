package part1

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "",
			args: args{
				nums:   []int{0, 1, 2, -1, -4, -1},
				target: 0,
			},
			wantResult: [][]int{
				{
					-1, -1, 2,
				},
				{
					-1, 0, 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := threeSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("threeSum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
