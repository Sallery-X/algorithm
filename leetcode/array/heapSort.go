package array

// 导入 fmt 包，用来打印结果

// heapSort 是堆排序主函数
// 传入一个整型切片 nums，然后把它原地排成升序
func heapSort(nums []int) {
	n := len(nums) // n 表示数组长度

	// 第一步：把整个数组构造成一个大顶堆
	// 为什么从 n/2 - 1 开始？
	// 因为从这个位置开始，往前都是非叶子节点
	// 叶子节点自己本身就是堆，不需要调整
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i) // 对以下标 i 为根的子树做堆调整
	}

	// 第二步：不断把堆顶元素（最大值）交换到数组最后
	// end 表示“当前堆的最后一个位置”
	// 每交换一次，最大的数就被放到了 end 位置
	for end := n - 1; end > 0; end-- {
		nums[0], nums[end] = nums[end], nums[0] // 把堆顶最大值放到最后

		// 交换后，前面的堆可能被破坏了
		// 所以要对根节点重新做一次堆调整
		// 注意这里长度是 end，不包括已经排好序的尾部
		heapify(nums, end, 0)
	}
}

// heapify 的作用：
// 在 nums 的前 n 个元素里，
// 以 i 这个下标作为根节点，调整成“大顶堆”
func heapify(nums []int, n int, i int) {
	largest := i // 假设当前根节点 i 是最大值

	left := 2*i + 1  // 左孩子下标
	right := 2*i + 2 // 右孩子下标

	// 如果左孩子存在，并且左孩子比当前最大值还大
	// 那就更新 largest 为左孩子
	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	// 如果右孩子存在，并且右孩子比当前最大值还大
	// 那就更新 largest 为右孩子
	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// 如果 largest 不是 i
	// 说明当前根节点不是最大的，堆结构被破坏了
	// 需要把根节点和更大的孩子交换
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]

		// 交换后，子节点下面那棵树可能又不满足大顶堆
		// 所以继续递归向下调整
		heapify(nums, n, largest)
	}
}
