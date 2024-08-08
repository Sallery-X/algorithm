package bytedance

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/

func treeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				line := make([]int, 0)
				line = append(line, nums[i], nums[l], nums[r])
				result = append(result, line)
				r--
				l++
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}
	return result
}
