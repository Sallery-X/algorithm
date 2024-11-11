package array

import "math"

//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的
//子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

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
