package leetcode

import "math"

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	nums := 0

	for i < n && s[i] >= '0' && s[i] <= '9' {
		temp := int(s[i] - '0')

		if nums > (math.MaxInt32-temp)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		nums = nums*10 + temp
		i++

	}

	return sign * nums

}
