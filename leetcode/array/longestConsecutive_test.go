package array

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("The length of the longest consecutive sequence is:", longestConsecutive(nums))
}
