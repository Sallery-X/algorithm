package bytedance

import (
	"fmt"
	"testing"
)

func Test_invertTree(t *testing.T) {
	root := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(invertTree(root).Right)
}
