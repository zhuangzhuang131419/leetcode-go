package main

import "fmt"

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				dfs(grid, i, j, visited)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i int, j int, visited [][]bool) {
	m := len(grid)
	n := len(grid[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if visited[i][j] {
		return
	}

	visited[i][j] = true
	if grid[i][j] == '1' {
		dfs(grid, i+1, j, visited)
		dfs(grid, i, j+1, visited)
		dfs(grid, i-1, j, visited)
		dfs(grid, i, j-1, visited)
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	fmt.Printf("%v\n", numIslands(grid))
}
