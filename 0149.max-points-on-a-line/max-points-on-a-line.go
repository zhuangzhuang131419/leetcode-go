package leetcode

func maxPoints(points [][]int) int {
	hash := map[float64]int{}

	n := len(points)
	maxCount := 0
	for i := 0; i < n; i++ {
		sameY := 1
		same := 1
		for j := i + 1; j < n; j++ {
			if points[i][1] == points[j][1] {
				sameY++
				if points[i][0] == points[j][0] {
					same++
				}
			} else {
				dx := points[i][0] - points[j][0]
				dy := points[i][1] - points[j][0]
				hash[float64(dx)/float64(dy)]++
			}
		}

		println("%v", hash)

		if maxCount < sameY {
			maxCount = sameY
		}
		for _, v := range hash {
			if v+same > maxCount {
				maxCount = v + same
			}
		}
		hash = map[float64]int{}
	}
	return maxCount
}
