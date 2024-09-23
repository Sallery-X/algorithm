package bytedance

func lengthOfLIS(nums []int) int {

	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	}
	res := make([]int, 0)
	res = append(res, nums[0])

	for _, v := range nums[1:] {
		if v > res[len(res)-1] {
			res = append(res, v)
		} else {
			l := 0
			r := len(res) - 1
			pos := 0
			for l <= r {
				mid := (l + r) >> 1
				if res[mid] < v {
					pos = mid + 1
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			res[pos] = v
		}

	}
	return len(res)

}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}

	}
	return maxLen

}
