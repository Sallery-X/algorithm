package bytedance

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{4, 5, 2, 6, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
