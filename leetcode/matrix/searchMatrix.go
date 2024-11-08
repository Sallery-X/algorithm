package matrix

/*
给你一个满足下述两条属性的 m x n 整数矩阵：

每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

*/

func searchMatrix(matrix [][]int, target int) bool {

	temp := make([]int, 0)

	for _, v := range matrix {
		temp = append(temp, v...)

	}

	left := 0
	right := len(temp) - 1

	for left <= right {
		mid := (left + right) >> 1

		if temp[mid] == target {
			return true
		} else if temp[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return false

}
