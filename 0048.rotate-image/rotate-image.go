package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	//
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}

	// 789
	// 456
	// 123

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[n-i-1][n-j-1] = matrix[n-i-1][n-j-1], matrix[i][j]
		}
	}
}
