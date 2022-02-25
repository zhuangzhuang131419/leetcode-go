package leetcode

// func dailyTemperatures(temperatures []int) []int {
// 	answers := make([]int, len(temperatures))

// 	answers[len(answers) - 1] = 0

// 	for i := len(temperatures) - 2; i >= 0; i-- {

// 		if temperatures[i] < temperatures[i + 1] {
// 			answers[i] = 1
// 		} else {

// 			j := i + 1
// 			for answers[j] != 0 {
// 				if temperatures[i] < temperatures[j + answers[j]] {
// 					answers[i] = answers[j] + j - i
// 					break
// 				}

// 				j += answers[j]
// 			}
// 		}
// 	}
// 	return answers
// }

func dailyTemperatures(temperatures []int) []int {
	answers := make([]int, len(temperatures))

	answers[len(answers)-1] = 0

	stack := []int{}

	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answers[preIndex] = i - preIndex
		}
		stack = append(stack, temperatures[i])

	}
	return answers
}
