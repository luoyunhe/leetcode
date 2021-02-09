package part1

import "testing"

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name       string
		args       args
		wantIndex1 int
		wantIndex2 int
	}{
		{
			name:       "",
			args:       args{[]int{2, 3, 6, 7, 9, 111, 222}, 12},
			wantIndex1: 2,
			wantIndex2: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex1, gotIndex2 := twoSum(tt.args.nums, tt.args.target)
			if gotIndex1 != tt.wantIndex1 {
				t.Errorf("twoSum() gotIndex1 = %v, want %v", gotIndex1, tt.wantIndex1)
			}
			if gotIndex2 != tt.wantIndex2 {
				t.Errorf("twoSum() gotIndex2 = %v, want %v", gotIndex2, tt.wantIndex2)
			}
		})
	}
}
