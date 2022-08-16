package leetcode

func updateMatrix(mat [][]int) [][]int {

	m := len(mat)
	n := len(mat[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = m + n
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}

			if j-1 >= 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}

			if i-1 >= 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}

		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}

			if j+1 < n {
				dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
			}

			if i+1 < m {
				dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
			}
		}
	}

	return dp

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
