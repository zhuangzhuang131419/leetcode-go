package leetcode

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var queue = [][]int{}

func shortestBridge(grid [][]int) int {
	sizeX := len(grid)
	sizeY := len(grid[0])
	visited := make([][]bool, sizeX)
	for i := 0; i < sizeX; i++ {
		visited[i] = make([]bool, sizeY)
	}

	isDone := false
	queue = [][]int{}

	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			if grid[i][j] == 1 {
				dfs(i, j, visited, sizeX, sizeY, grid)
				isDone = true
				break
			}
		}

		if isDone {
			break
		}
	}

	// fmt.Printf("grid=%v\n", grid)
	// fmt.Printf("queue=%v\n", queue)

	result := 0
	for len(queue) != 0 {
		result++
		size := len(queue)
		for i := 0; i < size; i++ {
			for _, dir := range dirs {
				nx := queue[0][0] + dir.x
				ny := queue[0][1] + dir.y

				if nx >= 0 && nx < sizeX && ny >= 0 && ny < sizeY && grid[nx][ny] != 2 {
					if grid[nx][ny] == 1 {
						return result - 1
					}

					grid[nx][ny] = 2

					queue = append(queue, []int{nx, ny})
				}
			}
			queue = queue[1:]
		}
	}
	return result
}

func dfs(x int, y int, visited [][]bool, sizeX int, sizeY int, grid [][]int) {
	if visited[x][y] {
		return
	}
	grid[x][y] = 2
	queue = append(queue, []int{x, y})
	// fmt.Printf("queue=%v\n", queue)

	visited[x][y] = true

	for _, dir := range dirs {
		nx := x + dir.x
		ny := y + dir.y

		if nx >= 0 && nx < sizeX && ny >= 0 && ny < sizeY && grid[nx][ny] == 1 {
			dfs(nx, ny, visited, sizeX, sizeY, grid)
		}

	}

}
