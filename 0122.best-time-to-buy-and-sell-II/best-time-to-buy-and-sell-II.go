package leetcode

func maxProfit(prices []int) int {
	// dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i])

	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[n] = make([]int, 2) // k
	}

	// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
	// dp[i][1] = max(dp[i - 1][1], -prices[i])
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 0; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][1]-prices[i])
	}

	return dp[n-1][0]

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
