package leetcode

import (
	"fmt"
	"testing"
)

func Test_findLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 3, 4, 1}

	fmt.Println(findLength(nums2, nums1))

}
