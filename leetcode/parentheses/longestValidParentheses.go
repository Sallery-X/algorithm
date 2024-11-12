package parentheses

// / 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串
// 的长度。
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	maxL := 0
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
