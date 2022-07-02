package main

var results []string

func generateParenthesis(n int) []string {

	results = []string{}

	backtrack(n, n, "")

	return results

}

func backtrack(left int, right int, s string) {

	if left == 0 && right == 0 {
		if isValidParenthesis(s) {
			results = append(results, s)
			return
		}
	}

	backtrack(left-1, right, s+"(")
	backtrack(left, right-1, s+")")

}

func isValidParenthesis(s string) bool {

	if s[0] != '(' {
		return false
	}

	stack := []byte{}
	stack = append(stack, s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
