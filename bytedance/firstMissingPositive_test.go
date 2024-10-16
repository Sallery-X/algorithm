package bytedance

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	nums := []int{9, 1, 2, 3, 7, 8}
	fmt.Println(firstMissingPositive(nums))
}
