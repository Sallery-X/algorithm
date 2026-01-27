package leetcode

// "有序数组平方后不同值的个数" 问题（可以看作是 LeetCode 977 "有序数组的平方" 的变种）。
//输入假设：nums 是一个升序排列的整数数组（可能包含负数）
//例如：[-4, -2, 0, 1, 3]

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
			if cur != rightV {
				res++
				cur = nums[right] * nums[right]
			}
			right--
		}

	}

	return res

}
