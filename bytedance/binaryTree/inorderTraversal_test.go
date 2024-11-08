package bytedance

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	x := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(inorderTraversal(x))
}
