package leetcode

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{4, 5, 6, 0, 1, 2, 3}
	x := search(nums, 2)
	fmt.Println(x)
}
