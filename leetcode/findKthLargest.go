package leetcode

func findKthLargest(nums []int, k int) int {
	return Partition(nums, k)
}

func Partition(nums []int, k int) int { //快排
	i, j := 0, len(nums)-1
	pivot := nums[0]
	for i < j {
		for i < j && nums[j] < pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		for i < j && nums[i] > pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	if i+1 == k {
		return pivot //   位子恰好就是K-1返回结果
	} else if i+1 < k {
		return Partition(nums[i+1:], k-i-1) //以pivot为断点，在后半部分数组中进行查找
	} else {
		return Partition(nums[:i], k) //在前半部分数组中进行查找
	}
}

func quickSort(start, end int, nums []int, k int) int {
	//if start >= end {
	//	return start
	//}

	i := start
	j := end
	base := nums[start]

	for i < j {
		// 从右向左找第一个大于 k 的元素
		for i < j && nums[j] <= base {
			j--
		}
		// 从左向右找第一个小于 k 的元素
		for i < j && nums[i] >= base {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i], nums[start] = nums[start], nums[i]

	if i == k-1 {
		return nums[i]
	} else if k-1 > i {
		return quickSort(j+1, end, nums, k)
	} else {
		return quickSort(start, i-1, nums, k)
	}

}

func quiks(nums []int, start, end int) {
	if start >= end {
		return
	}
	i := start
	j := end
	base := nums[start]
	for i < j {
		for i < j && nums[j] >= base {
			j--
		}
		for i < j && nums[i] <= base {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	quiks(nums, start, i-1)
	quiks(nums, j+1, end)
}
