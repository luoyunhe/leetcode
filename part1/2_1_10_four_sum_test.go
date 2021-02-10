package part1

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	var tests = []struct {
		name           string
		args           args
		wantResultList [][]int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			wantResultList: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResultList := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotResultList, tt.wantResultList) {
				t.Errorf("fourSum() = %v, want %v", gotResultList, tt.wantResultList)
			}
		})
	}
}
