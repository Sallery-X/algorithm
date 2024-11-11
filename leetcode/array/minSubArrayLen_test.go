package array

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	nums := []int{3, 3}
	fmt.Println(minSubArrayLen(5, nums))
}
