package leetcode

// f(n)= f(n-1)+f(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
