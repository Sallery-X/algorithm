package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum := 0, 0
	change := false
	minLength := math.MaxInt

	for right := 0; right < n; right++ {
		sum = sum + nums[right]

		for sum >= target {
			change = true
			minLength = int(math.Min(float64(minLength), float64(right-left+1)))
			sum = sum - nums[left]
			left++
		}
	}

	if !change {
		return 0
	}

	return minLength

}
