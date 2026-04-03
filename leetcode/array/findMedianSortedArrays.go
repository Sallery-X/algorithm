package array

import "math"

/*
这题本质是在做：

找一个位置，把 nums1 切一刀
再自动算出 nums2 该切在哪里
切完以后要满足：
左边总个数正确
左边最大值 <= 右边最小值

只要这个条件满足：

奇数长度，中位数就是左边最大值
偶数长度，中位数就是左边最大值和右边最小值的平均
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 保证 nums1 是更短的那个数组
	// 这样后面只在 nums1 上二分，边界更容易处理
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}

	// x, y 分别表示两个数组长度
	x, y := len(nums1), len(nums2)

	// 在 nums1 上做二分
	// l, r 表示 nums1 可选切分位置范围 [0, x]
	l, r := 0, x

	for l <= r {
		// midX 表示 nums1 的切分位置
		// 意思是 nums1 左边取多少个元素
		midX := (l + r) >> 1

		// midY 表示 nums2 的切分位置
		// 让左右两边元素个数保持平衡
		// (x+y+1)/2 表示左半边总共应该有多少个元素

		//先规定左边总共应该有 (x+y+1)/2 个元素
		//nums1 已经切了 midX 个到左边
		//那 nums2 只能补剩下的数量
		//所以就是：
		midY := (x+y+1)/2 - midX

		// maxX 表示 nums1 左半部分的最大值
		// 如果 midX == 0，说明 nums1 左边一个元素都没有
		// 这时左边最大值不存在，用负无穷代替
		maxX := math.MinInt32
		if midX != 0 {
			maxX = nums1[midX-1]
		}

		// maxY 表示 nums2 左半部分的最大值
		// 如果 midY == 0，说明 nums2 左边没有元素
		// 同样用负无穷代替
		maxY := math.MinInt32
		if midY != 0 {
			maxY = nums2[midY-1]
		}

		// minX 表示 nums1 右半部分的最小值
		// 如果 midX == x，说明 nums1 右边没有元素
		// 这时右边最小值不存在，用正无穷代替
		minX := math.MaxInt32
		if midX != x {
			minX = nums1[midX]
		}

		// minY 表示 nums2 右半部分的最小值
		// 如果 midY == y，说明 nums2 右边没有元素
		// 用正无穷代替
		minY := math.MaxInt32
		if midY != y {
			minY = nums2[midY]
		}

		// 判断当前切分是否合理
		// 条件含义：
		// nums1 左边最大值 <= nums2 右边最小值
		// nums2 左边最大值 <= nums1 右边最小值
		// 如果都满足，说明左半部分所有元素都 <= 右半部分所有元素
		if maxX <= minY && maxY <= minX {

			// 如果总长度是偶数
			// 中位数 = 左边最大值 和 右边最小值 的平均数
			if (x+y)%2 == 0 {
				return float64(max(maxY, maxX)+min(minX, minY)) / 2.0
			} else {
				// 如果总长度是奇数
				// 因为左边元素个数比右边多 1
				// 所以中位数就是左边最大值
				return float64(max(maxX, maxY))
			}

		} else if maxX > minY {
			// 说明 nums1 左边取多了
			// 因为 nums1 左边最大值已经比 nums2 右边最小值还大
			// 所以要把切分点往左移动
			r = midX - 1
		} else {
			// 否则说明 nums1 左边取少了
			// 要把切分点往右移动
			l = midX + 1
		}

	}

	// 正常来说不会走到这里
	// 如果走到这里，说明输入数组不满足升序条件，或者逻辑有问题
	panic("not sorted arrays")
}
