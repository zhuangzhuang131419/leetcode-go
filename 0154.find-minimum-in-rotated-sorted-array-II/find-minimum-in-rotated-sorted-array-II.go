package leetcode

import "fmt"

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		if left == right {
			return nums[left]
		}

		if left+1 == right {
			if nums[left] > nums[right] {
				return nums[right]
			} else {
				return nums[left]
			}
		}

		mid := (left + right) / 2

		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] == nums[right] {
			right = right - 1
		} else {
			// mid > right
			left = mid + 1
		}
		fmt.Printf("left = %v, right = %v\n", left, right)
	}

	return -1

}
