package array

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	key := make(map[int]int, 0)
	for i, v := range nums {
		key[v] = i
	}
	for i, v := range nums {
		if index, ok := key[target-v]; ok && index != i {
			return []int{i, index}
		}
	}
	return []int{-1, -1}
}
