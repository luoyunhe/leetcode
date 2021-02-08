package part1

import "testing"

func Test_searchRotatedSortedArray2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 1, 2, 2, 3, 4, -1, 0, 1},
				target: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRotatedSortedArray2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchRotatedSortedArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
