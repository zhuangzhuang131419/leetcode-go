package leetcode

func maxArea(height []int) int {
	result := 0
	for left := 0; left < len(height); left++ {
		for right := left + 1; right < len(height); right++ {
			tmp := computeArea(left, right, height)
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

func computeArea(left int, right int, height []int) int {
	if height[right] > height[left] {
		return height[left] * (right - left)
	} else {
		return height[right] * (right - left)
	}
}
