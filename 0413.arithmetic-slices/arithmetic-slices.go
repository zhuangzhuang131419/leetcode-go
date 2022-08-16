package leetcode

func numberOfArithmeticSlices(nums []int) int {

	return helper(nums, 0, len(nums)-1)

}

func helper(nums []int, start int, end int) int {
	if end-start < 2 {
		return 0
	}

	result := 0
	diff := nums[start+1] - nums[start]
	for i := start + 1; i < end; i++ {
		if nums[i+1]-nums[i] == diff {
			result++
		} else {
			break
		}
	}

	return helper(nums, start+1, end) + result
}

// dp[start][end] = dp[start + 1][end] + nums[start]...
