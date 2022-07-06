package leetcode

func findCircleNum(isConnected [][]int) int {

	n := len(isConnected)
	visited := make([]bool, n)

	count := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, visited, isConnected, n)
			count++
		}

	}

	return count

}

func dfs(target int, visited []bool, isConnected [][]int, size int) {
	visited[target] = true

	for i := 0; i < size; i++ {
		if i == target {
			continue
		}
		if isConnected[i][target] == 1 && !visited[i] {
			dfs(i, visited, isConnected, size)
		}
	}

	for i := 0; i < size; i++ {
		if i == target {
			continue
		}
		if isConnected[target][i] == 1 && !visited[i] {
			dfs(i, visited, isConnected, size)
		}
	}
}
