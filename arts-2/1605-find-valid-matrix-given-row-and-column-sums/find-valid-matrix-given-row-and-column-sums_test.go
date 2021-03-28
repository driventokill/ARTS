package _1605_find_valid_matrix_given_row_and_column_sums

import (
	"testing"
)

func Test_restoreMatrix(t *testing.T) {
	type args struct {
		rowSum []int
		colSum []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				[]int{5, 7, 10},
				[]int{8, 6, 8},
			},
		},
		{
			"2",
			args{
				[]int{14, 9},
				[]int{6, 9, 8},
			},
		},
		{
			"3",
			args{
				[]int{1, 0},
				[]int{1},
			},
		},
		{
			"4",
			args{
				[]int{0},
				[]int{0},
			},
		},
		{
			"5",
			args{
				[]int{},
				[]int{},
			},
		},
		{
			"6",
			args{
				[]int{100},
				[]int{0, 89, 1, 10},
			},
		},
		{
			"7",
			args{
				[]int{0, 80, 1, 19},
				[]int{100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreMatrix(tt.args.rowSum, tt.args.colSum); !sumEqual(tt.args.rowSum, tt.args.colSum, got) {
				t.Errorf("restoreMatrix() = %v, args %v", got, tt.args)
			}
		})
	}
}

func sumEqual(rowSum []int, colSum []int, m [][]int) bool {
	for i, r := range rowSum {
		if r != sumRow(m, i) {
			return false
		}
	}

	for j, c := range colSum {
		if c != sumCol(m, j) {
			return false
		}
	}
	return true
}

func sumRow(m [][]int, row int) (sum int) {
	sum = 0
	for _, r := range m[row] {
		sum += r
	}
	return
}

func sumCol(m [][]int, col int) (sum int) {
	sum = 0
	for _, r := range m {
		sum += r[col]
	}
	return
}
