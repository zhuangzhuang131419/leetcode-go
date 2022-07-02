package leetcode

func trap(height []int) int {
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))

	lMax[0] = height[0]
	rMax[len(rMax)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {

		if lMax[i-1] > height[i] {
			lMax[i] = lMax[i-1]
		} else {
			lMax[i] = height[i]
		}
	}

	for i := len(height) - 2; i >= 0; i-- {

		if rMax[i+1] > height[i] {
			rMax[i] = rMax[i+1]
		} else {
			rMax[i] = height[i]
		}
	}

	result := 0
	for i := 0; i < len(height); i++ {
		h := 0
		if lMax[i] < rMax[i] {
			h = lMax[i]
		} else {
			h = rMax[i]
		}

		result += h - height[i]
	}

	return result
}
