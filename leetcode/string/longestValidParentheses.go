package string

// / 最长有效括号
func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	maxL := 0
	stack = append(stack, -1)
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxL = max(maxL, i-stack[len(stack)-1])
			}
		}

	}
	return maxL

}
