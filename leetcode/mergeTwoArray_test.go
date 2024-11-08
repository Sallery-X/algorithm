package leetcode

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	n1 := []int{3, 6, 7, 8, 0, 0}
	n2 := []int{1, 10}
	merge(n1, 4, n2, 2)
	fmt.Println(n1)
}
