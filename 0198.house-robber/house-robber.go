package leetcode

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。


输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/

func rob(nums []int) int {
	value := map[int]int{}
	return robHelp(nums, 0, value)
}

func robHelp(nums []int, start int, value map[int]int) int {

	if v, exist := value[start]; exist {
		return v
	}

	if start+1 == len(nums)-1 {
		if nums[start] > nums[start+1] {
			return nums[start]
		} else {
			return nums[start+1]
		}
	}

	if start == len(nums)-1 {
		return nums[start]
	}

	a1 := nums[start] + robHelp(nums, start+2, value)
	a2 := robHelp(nums, start+1, value)

	if a2 > a1 {
		value[start] = a2
		return a2
	} else {
		value[start] = a1
		return a1
	}
}
