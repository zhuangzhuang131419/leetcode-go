package leetcode

func candy(ratings []int) int {

	n := len(ratings)
	candy := make([]int, n)

	for i := 0; i < n; i++ {
		candy[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] && candy[i] <= candy[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}

	for i := n - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] && candy[i] >= candy[i-1] {
			candy[i-1] = candy[i] + 1
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		result += candy[i]
	}

	return result

}
