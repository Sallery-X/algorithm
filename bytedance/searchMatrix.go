package bytedance

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
