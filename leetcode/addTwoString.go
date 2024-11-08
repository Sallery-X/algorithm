package leetcode

import (
	"strconv"
)

// 123 45
func AddTwoString(num1, num2 string) string {
	len1 := len(num1) - 1
	len2 := len(num2) - 1
	res := ""
	carry := 0
	for len1 >= 0 || len2 >= 0 {
		x, y := 0, 0
		if len1 >= 0 {
			x, _ = strconv.Atoi(string(num1[len1]))
		}
		if len2 >= 0 {
			y, _ = strconv.Atoi(string(num2[len2]))
		}
		sum := x + y + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res

		len1--
		len2--
	}

	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}
