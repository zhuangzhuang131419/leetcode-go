package leetcode

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {

		if n1, ok := hashTable[target - nums[i]]; ok {
			return []int{n1, i}
		}
		hashTable[nums[i]] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
