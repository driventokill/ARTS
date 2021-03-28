package _1605_find_valid_matrix_given_row_and_column_sums

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	rowCount, colCount := len(rowSum), len(colSum)
	M := make([][]int, rowCount)
	e := make([]int, rowCount*colCount)

	for i := range M {
		M[i] = e[i*colCount : (i+1)*colCount]
	}

	colRemain := make([]int, colCount)
	copy(colRemain, colSum)

	for i, row := range rowSum {
		for j, col := range colRemain {
			if row > col {
				M[i][j] = col
				colRemain[j] = 0
				row -= col
			} else {
				M[i][j] = row
				colRemain[j] = colRemain[j] - row
				row = 0
			}
		}
	}

	return M
}
