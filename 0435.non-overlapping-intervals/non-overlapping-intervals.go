package leetcode

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	result := 0
	i := 1
	fmt.Printf("intervals = %v", intervals)
	preIntervals := intervals[0]
	for i < len(intervals) {
		if preIntervals[1] <= intervals[i][0] {
			preIntervals = intervals[i]
		} else if preIntervals[1] >= intervals[i][1] {
			preIntervals = intervals[i]
			result++
		} else if preIntervals[1] < intervals[i][1] {
			result++
		}
		fmt.Printf("pre = %v", preIntervals)
		i++
	}

	return result
}
