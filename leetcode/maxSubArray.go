package leetcode

import "math"

func MaxSubArray(nums []int) int {
	pre := 0
	maxA := nums[0]
	for _, v := range nums {
		pre = int(math.Max(float64(pre+v), float64(v)))
		maxA = int(math.Max(float64(maxA), float64(pre)))
	}
	return maxA
}
