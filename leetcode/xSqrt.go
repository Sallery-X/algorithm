package leetcode

import "fmt"

func mySqrt(x int) int {

	l := 0
	r := x

	ans := 0
	for l <= r {
		mid := (l + r) >> 1
		if mid*mid > x {
			//ans = mid
			r = mid - 1
			fmt.Println(r)
		} else {
			ans = mid
			l = mid + 1
			fmt.Println(l)
		}
	}
	return ans

}
