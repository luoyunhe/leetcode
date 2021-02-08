package part1

import "testing"

func Test_searchRotatedSortedArray(t *testing.T) {
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
				nums: []int{4, 5, 6, 7, 8, 0, 1, 2, 3},
				target: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRotatedSortedArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchRotatedSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
