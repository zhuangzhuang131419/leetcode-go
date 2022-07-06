package leetcode

func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		if left == right {
			return nums[right]
		}

		mid := (left + right) / 2
		if nums[mid-1] == nums[mid] {

			if mid%2 == 0 {
				right = mid - 2
			} else {
				left = mid + 1
			}

		} else if nums[mid+1] == nums[mid] {
			if mid%2 == 0 {
				left = mid + 2
			} else {
				right = mid - 1
			}

		} else {
			return nums[mid]
		}
	}
	return -1
}
