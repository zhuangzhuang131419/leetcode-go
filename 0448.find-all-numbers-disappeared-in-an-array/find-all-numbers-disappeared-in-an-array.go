package leetcode

func findDisappearedNumbers(nums []int) []int {

	n := len(nums)

	for i := 0; i < n; i++ {

		for nums[i] != i + 1 && nums[i] != nums[nums[i] - 1]  {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}

		// i = 0 [4,3,2,7,8,2,3,1]
		// i = 0 [7,3,2,4,8,2,3,1]
		// [3,3,2,4,8,2,7,1]
		// [2,3,3,4,8,2,7,1]
		// [3,2,3,4,8,2,7,1]
		// [3,2,3,4,1,2,7,8]
		// [1,2,3,4,3,2,7,8]
	}

	result := []int{}
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			result = append(result, i + 1)
		}
	}

	return result
}