func singleNonDuplicate(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		if low == high {
			return nums[low]
		}

		mid := (low + high) / 2
		if nums[mid-1] == nums[mid] {

			if mid%2 == 0 {
				right = mid - 2
			} else {
				left = mid + 1
			}

		} else if nums[mid+1] == nums[mid] {
			if mid%2 == 0 {
				right = mid - 2
			} else {
				left = mid + 1
			}

		} else {
			return nums[mid]
		}
	}

}