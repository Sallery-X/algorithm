package leetcode

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, 0)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[topIndex] = i - topIndex
		}
		stack = append(stack, i)
	}

	return res
}
