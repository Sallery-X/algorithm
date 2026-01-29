package array

func moveZero(nums []int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}

	return nums

}
