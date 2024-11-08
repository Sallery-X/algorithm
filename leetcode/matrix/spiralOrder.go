package matrix

// 螺旋矩阵 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
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
