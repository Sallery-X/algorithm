package array

import (
	"fmt"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	nums := []int{2, 4, 1, 2, 7, 8, 4}
	fmt.Println(findPeakElement(nums))
}
