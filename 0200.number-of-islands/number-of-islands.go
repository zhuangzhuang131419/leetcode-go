package leetcode

func numIslands(grid [][]byte) int {

	n := len(grid)

	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {

				// right
				if grid[i][j+1] == 1 {

				}

			}
		}
	}

	return 0

}
