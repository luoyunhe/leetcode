package part1

import "testing"

func Test_getLongestConsecutiveSequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 3, 1, 7, 5, 9, 20, 8, 4, 6},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestConsecutiveSequence(tt.args.nums); got != tt.want {
				t.Errorf("getLongestConsecutiveSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
