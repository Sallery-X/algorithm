package bytedance

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	//nums := []int{5, 6, 7, 8, 2, 4, 9}
	//fmt.Println(findKthLargest(nums, 3))

	nums := []int{5, 6, 7, 8, 2, 4, 9}
	fmt.Println(quickSort(0, len(nums)-1, nums, 3))
	fmt.Println(nums)

	nums = []int{0, 11, 5, 6, 7, 8, 2, 4, 9, 10}
	quiks(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
