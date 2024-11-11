package array

import (
	"fmt"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}
