package bytedance

import (
	"fmt"
	"testing"
)

func Test_mergeInterval(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := mergeInterval(nums)
	fmt.Println(res)
}
