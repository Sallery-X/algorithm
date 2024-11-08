package leetcode

import (
	"container/list"
	"math"
)

func longestValidParentheses(s string) int {

	stack := list.New()
	maxL := 0
	stack.PushFront(-1)
	for i, v := range s {
		if v == '(' {
			stack.PushFront(i)
		} else {
			stack.Remove(stack.Front())
			if stack.Len() == 0 {
				stack.PushFront(i)
			} else {
				maxL = int(math.Max(float64(maxL), float64(i)-float64(stack.Front().Value.(int))))
			}

		}
	}
	return maxL

}
