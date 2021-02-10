package part1

import "testing"

func Test_closestThreeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name        string
		args        args
		wantClosest int
	}{
		{
			name:        "",
			args:        args{nums: []int{-1, 2, 1, 0}, target: 1},
			wantClosest: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotClosest := closestThreeSum(tt.args.nums, tt.args.target); gotClosest != tt.wantClosest {
				t.Errorf("closestThreeSum() = %v, want %v", gotClosest, tt.wantClosest)
			}
		})
	}
}
