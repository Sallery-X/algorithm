package bytedance

import (
	"sort"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 {
		if nums[i+1] > nums[i] {
			break
		}
		i--
	}
	j := len(nums) - 1

	if i >= 0 {
		for j > i {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	sort.Ints(nums[i+1:])

}
