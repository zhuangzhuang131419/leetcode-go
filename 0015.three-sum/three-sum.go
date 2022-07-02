package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	return threeSumTarget(nums, 0)

}

func threeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	for i := 0; i < len(nums); i++ {
		for _, v := range twoSumTarget(nums, i+1, target-nums[i]) {
			v = append(v, nums[i])
			result = append(result, v)
		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result

}

func twoSumTarget(nums []int, start, target int) [][]int {
	i := start
	j := len(nums) - 1
	result := [][]int{}
	for i < j {
		if nums[i]+nums[j] == target {
			result = append(result, []int{nums[i], nums[j]})
		}

		if nums[i]+nums[j] < target {
			i++
			for i < len(nums)-1 && nums[i+1] == nums[i] {
				i++
			}
		} else {
			j--
			for j > 1 && nums[j-1] == nums[j] {
				j--
			}
		}
	}
	return result
}
