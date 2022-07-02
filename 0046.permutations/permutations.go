package leetcode

// import "fmt"

var result [][]int

func permute(nums []int) [][]int {
	path := []int{}
	result = make([][]int, 0)
	backtrack(nums, path)
	return result

}

func backtrack(choices []int, path []int) {

	if len(path) == len(choices) {
		r := []int{}
		for i := 0; i < len(path); i++ {
			r = append(r, path[i])
		}

		result = append(result, r)
		return
	}

	n := len(choices)
	for i := 0; i < n; i++ {

		isContain := false
		for j := 0; j < len(path); j++ {
			if path[j] == choices[i] {
				isContain = true
			}
		}

		if isContain {
			continue
		}

		path = append(path, choices[i])
		backtrack(choices, path)

		path = path[:len(path)-1]
	}
}
