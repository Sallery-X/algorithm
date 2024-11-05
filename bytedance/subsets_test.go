package bytedance

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}
