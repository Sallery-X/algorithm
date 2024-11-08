package leetcode

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
