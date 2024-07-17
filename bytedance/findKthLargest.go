package bytedance

import "fmt"

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

func quickSort(start, end int, arr []int) {
	i := start
	j := end
	k := arr[start]
	if i < j {
		for i < j {
			for arr[i] < k {
				i++
			}
			for arr[j] > k {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
		}

		//第一次排完 左边都比k小 右边比k大
		for _, v := range arr {
			fmt.Print(v)
			fmt.Print(" ")
		}
		quickSort(start, i-1, arr)
		quickSort(j+1, end, arr)
	}
}
