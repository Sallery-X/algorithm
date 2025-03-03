package array

import "math"

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//子数组
//是数组中的一个连续部分。
//
//
//
//示例 1：
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func MaxSubArray(nums []int) int {
	pre := 0
	maxA := nums[0]
	for _, v := range nums {
		pre = int(math.Max(float64(pre+v), float64(v)))
		maxA = int(math.Max(float64(maxA), float64(pre)))
	}
	return maxA
}
