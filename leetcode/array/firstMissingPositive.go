package array

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

//示例 1：
//
//输入：nums = [1,2,0]
//输出：3
//解释：范围 [1,2] 中的数字都在数组中。

func firstMissingPositive(nums []int) int {
	n := len(nums)

	//把数放到应该放的位置上
	for i := 0; i < n; {
		if nums[i] <= n && nums[i] > 0 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}

	//找到第一个位置就是缺失的最小的
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}

	}
	//没找到
	return n + 1

}
