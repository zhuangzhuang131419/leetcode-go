package leetcode

func sortColors(nums []int) {
	i := 0
	j := 0

	for m := 0; m < len(nums); m++ {

		if nums[m] == 0 {
			i++
			nums[i], nums[m] = nums[m], nums[i]
			if i < j {
				nums[j], nums[m] = nums[m], nums[j]
			}
		} else if nums[m] == 1 {
			j++
			nums[j], nums[m] = nums[m], nums[j]
		}
	}

}
