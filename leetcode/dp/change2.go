package dp

/*
假设 coins = [1, 2, 5]，amount = 5。

初始化 dp 数组为 [1, 0, 0, 0, 0, 0]。

使用硬币 1 更新 dp 数组：

对于 i = 1 到 5，dp[i] = dp[i] + dp[i-1]。
更新后 dp 为 [1, 1, 1, 1, 1, 1]。
使用硬币 2 更新 dp 数组：

对于 i = 2 到 5，dp[i] = dp[i] + dp[i-2]。
更新后 dp 为 [1, 1, 2, 2, 3, 3]。
使用硬币 5 更新 dp 数组：

对于 i = 5，dp[i] = dp[i] + dp[i-5]。
更新后 dp 为 [1, 1, 2, 2, 3, 4]。
最终，dp[5] 的值为 4，表示有 4 种方法可以凑成金额 5。

*/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = dp[i] + dp[i-coin]
		}
	}

	return dp[amount]

}
