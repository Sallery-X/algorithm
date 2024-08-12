package bytedance

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	key := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if v, ok := key[target-nums[i]]; ok && i != v {
			return []int{v, i}
		}

	}
	return nil
}
