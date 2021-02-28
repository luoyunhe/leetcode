package part1

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name        string
		args        args
		wantNewNums []int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 5, 5, 6, 7, 9},
				target: 5,
			},
			wantNewNums: []int{1, 2, 3, 6, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNums := removeElement(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotNewNums, tt.wantNewNums) {
				t.Errorf("removeElement() = %v, want %v", gotNewNums, tt.wantNewNums)
			}
		})
	}
}
