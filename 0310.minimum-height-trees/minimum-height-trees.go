package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	result := []int{}

	minHeight := n

	for i := 0; i < n; i++ {
		used := map[int]bool{}
		queue := []int{}
		queue = append(queue, i)
		height := 0
		for len(queue) > 0 {

			size := len(queue)
			for k := 0; k < size; k++ {
				node := queue[0]
				used[node] = true
				queue = queue[1:]

				for j := 0; j < len(edges); j++ {
					if edges[j][0] == node && !used[edges[j][1]] {
						queue = append(queue, edges[j][1])
					}

					if edges[j][1] == node && !used[edges[j][0]] {
						queue = append(queue, edges[j][0])
					}
				}
			}
			height++

		}

		if height < minHeight {
			result = []int{}
			minHeight = height
			result = append(result, i)
		} else if height == minHeight {
			result = append(result, i)
		}

	}

	return result

}
