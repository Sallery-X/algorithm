package leetcode

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for v := range dp {
		dp[v] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]

}
