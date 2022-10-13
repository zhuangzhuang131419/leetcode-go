package leetcode

func maxChunksToSorted(arr []int) int {

	result := 0

	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if max == i {
			result++
		}
	}

	return result
}
