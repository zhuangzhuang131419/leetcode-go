package leetcode

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	// base case
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i+1 < len(s) {
			dp[i][i+1] = (s[i] == s[i+1])
		}

	}
	for k := 2; k < len(s); k++ {
		for i := 0; i+k < len(s); i++ {
			j := i + k
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}

		}
	}

	start := -1
	end := -1
	length := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] == true && j-i+1 > length {
				length = j - i + 1
				start = i
				end = j
			}

		}
	}

	if length == 0 {
		return ""
	}

	return s[start : end+1]
}

// dp[i][j] = dp[i + 1][j - 1] && s[i] = s[j]
// dp[i][i] = true
// dp[i][i + 1] = s[i] == s[j]
