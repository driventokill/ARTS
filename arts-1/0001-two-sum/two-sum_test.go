package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0",
			args: args{
				nums:   []int{2, 7, 11, 9},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "test1",
			args: args{
				nums:   []int{1, 2, 3},
				target: 3,
			},
			want: []int{0, 1},
		},
		{
			name: "test2",
			args: args{
				nums:   []int{1, 2, 3, 3, 2},
				target: 6,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
