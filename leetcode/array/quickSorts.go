package array

import "fmt"

func sort1(nums []int) {
	quickSort1(nums, 0, len(nums)-1)
}

func quickSort1(nums []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	base := nums[start]

	for i < j {
		//从右向左搞
		for i < j && nums[j] <= base {
			j--
		}
		for i < j && nums[i] >= base {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	fmt.Println(nums)

	nums[i], nums[start] = nums[start], nums[i]
	fmt.Println(nums)

	quickSort1(nums, start, i-1)
	quickSort1(nums, i+1, end)

}
