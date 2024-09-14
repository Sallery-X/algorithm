package bytedance

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	p := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(p, in))
}
