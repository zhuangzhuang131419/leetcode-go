package leetcode

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 0; i < len(coins); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
