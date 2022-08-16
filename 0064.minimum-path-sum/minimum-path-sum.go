package leetcode

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	// dp[m-2][n-1] dp[m-1][n-2]
	// dp[m-3][n-1] dp[m-2][n-2] dp[m-1][n-3]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]

			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[i][j] = min(dp[i - 1][j], dp[i][j - 1])+grid[i][j]
