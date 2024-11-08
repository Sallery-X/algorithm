package leetcode

func majorityElement(nums []int) int {
	res := make(map[int]int)
	majority := len(nums) / 2

	for _, v := range nums {
		res[v]++
		if res[v] > majority {
			return v
		}

	}
	return -1

}
