package bytedance

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	x := []int{2, 8, 9}
	y := []int{1, 3}
	fmt.Println(findMedianSortedArrays(x, y))
}
