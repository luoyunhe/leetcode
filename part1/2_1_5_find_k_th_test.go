package part1

import "testing"

func Test_findKth(t *testing.T) {
	type args struct {
		a []int
		b []int
		k int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{

		{
			name: "test1",
			args: args{
				a: []int{1, 3, 5, 7, 9},
				b: []int{2, 4, 5, 8, 10},
				k: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKth(tt.args.a, tt.args.b, tt.args.k); got != tt.want {
				t.Errorf("findKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
