package leetcode

func maxArea(height []int) int {

	i := 0
	j := len(height) - 1

	lMax := height[i]
	rMax := height[j]
	ans := (j - i) * min(lMax, rMax)

	for i < j {
		if lMax < rMax {
			i++
			if height[i] > lMax {
				lMax = height[i]
			}
		} else {
			j--
			if height[j] > rMax {
				rMax = height[j]
			}
		}

		tempAns := (j - i) * min(lMax, rMax)
		if tempAns > ans {
			ans = tempAns
		}

	}

	return ans

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
