package leetcode

// 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return nil
	}
	rowStart := 0
	rowEnd := len(matrix) - 1
	colStart := 0
	colEnd := len(matrix[0]) - 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			res = append(res, matrix[rowStart][i])
		}
		rowStart++
		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, matrix[i][rowEnd])
		}
		colEnd--
		for i := colEnd; i >= colStart; i-- {
			res = append(res, matrix[rowEnd][i])
		}
		rowEnd--
		for i := rowEnd; i >= rowStart; i-- {
			res = append(res, matrix[i][colStart])
		}
		colStart++
	}

	return res

}
