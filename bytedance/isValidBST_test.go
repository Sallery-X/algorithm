package bytedance

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	root := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(isValidBST(root))
}
