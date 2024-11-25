package leetcode

func res(nums []int) int {
	res := 1
	left := 0
	right := len(nums) - 1
	cur := nums[left] * nums[left]

	for left < right {
		leftV := nums[left] * nums[left]
		rightV := nums[right] * nums[right]

		if leftV > rightV {
			if cur != leftV {
				res++
				cur = nums[left] * nums[left]
			}
			left++
		} else {
			if cur != right {
				res++
				cur = nums[right] * nums[right]
			}
			right--
		}

	}

	return res

}
