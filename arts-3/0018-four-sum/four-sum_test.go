package foursum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"Example 2", args{[]int{2, 2, 2, 2, 2}, 8}, [][]int{{2, 2, 2, 2}}},
		{"1", args{[]int{-9, -9, 10, -20, 5, 5, 25, -10, 9}, 30}, [][]int{{-10, 5, 10, 25}, {-9, 5, 9, 25}}},
		{"2", args{
			[]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0,
		}, [][]int{
			{-5, -4, 4, 5}, {-5, -3, 3, 5}, {-5, -2, 2, 5}, {-5, -2, 3, 4}, {-5, -1, 1, 5}, {-5, -1, 2, 4}, {-5, 0, 0, 5}, {-5, 0, 1, 4}, {-5, 0, 2, 3}, {-4, -3, 2, 5}, {-4, -3, 3, 4}, {-4, -2, 1, 5}, {-4, -2, 2, 4}, {-4, -1, 0, 5}, {-4, -1, 1, 4}, {-4, -1, 2, 3}, {-4, 0, 0, 4}, {-4, 0, 1, 3}, {-3, -2, 0, 5}, {-3, -2, 1, 4}, {-3, -2, 2, 3}, {-3, -1, 0, 4}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
