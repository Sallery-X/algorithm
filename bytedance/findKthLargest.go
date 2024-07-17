package bytedance

import "fmt"

func findKthLargest(nums []int, k int) {

}

func find(nums []int, left, right, k int) int {
	if left == right {
		return nums[k]
	}

	i := left
	j := right
	index := nums[l]
	for i < j {
		for i++; nums[i]<
	}

}


func quickselect(nums []int, l, r, k int) int{
	if (l == r){
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for (i < j) {
		for i++;nums[i]<partition;i++{}
		for j--;nums[j]>partition;j--{}
		if (i < j) {
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
	if (k <= j){
		return quickselect(nums, l, j, k)
	}else{
		return quickselect(nums, j + 1, r, k)
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

		return
		//第一次排完 左边都比k小 右边比k大
		for _, v := range arr {
			fmt.Print(v)
			fmt.Print(" ")
		}
		quickSort(start, i-1, arr)
		quickSort(j+1, end, arr)
	}
}
