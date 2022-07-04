package leetcode

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	intervals := points[0]
	n := len(points)
	result := 0
	fmt.Printf("points = %v\n", points)
	for i := 1; i < n; i++ {
		if points[i][0] == intervals[1] {
			intervals[1] = points[i][0]
		} else if points[i][0] >= intervals[1] {
			intervals = points[i]
			result++
		} else if points[i][1] < intervals[1] {
			intervals = points[i]
		} else if points[i][1] >= intervals[1] {
			intervals[0] = points[i][0]
		}
		fmt.Printf("inter = %v\n", intervals)
	}
	return result + 1
}
