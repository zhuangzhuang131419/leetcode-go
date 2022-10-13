package leetcode

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	i := m - 1
	j := 0

	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			i--
			continue
		}

		if matrix[i][j] < target {
			j++
			continue
		}
	}

	return false

}
