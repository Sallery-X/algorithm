package array

func sort1(nums []int) {

	quickSort1(nums, 0, len(nums)-1)

}

func quickSort1(nums []int, start, end int) {
	if start >= end {
		return
	}

	i := start
	j := end

	base := nums[start]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}

		for nums[i] <= base && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[start] = nums[start], nums[i]
	quickSort1(nums, start, i-1)
	quickSort1(nums, j+1, end)

}
