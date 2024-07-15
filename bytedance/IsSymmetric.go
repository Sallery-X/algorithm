package bytedance

import "fmt"

type Tree struct {
	left  *Tree
	right *Tree
	value int
}

func IsSymmetric(root *Tree) bool {
	fmt.Println(judge2(root, root))
	return judge(root, root)
}

func judge(root1, root2 *Tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.value == root2.value && judge(root1.left, root2.right) && judge(root1.right, root2.left)
}
func judge2(root1, root2 *Tree) bool {
	list := make([]*Tree, 0)
	list = append(list, root1)
	list = append(list, root2)

	for len(list) > 0 {
		left := list[0]
		right := list[1]
		list = list[2:]
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		if left.value != right.value {
			return false
		}

		list = append(list, left.left, right.right)
		list = append(list, left.right, right.left)
	}
	return true
}
