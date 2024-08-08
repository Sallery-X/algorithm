package bytedance

import (
	"fmt"
	"testing"
)

func Test_treeSum(t *testing.T) {
	x := []int{-1, 0, 1, 2, -1, -4}
	r := treeSum(x)
	fmt.Println(r)
}
