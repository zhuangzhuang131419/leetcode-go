package main

import "fmt"

var results []string

func generateParenthesis(n int) []string {

	results = []string{}

	return results

}

func backtrack(left int, right int, s string) {

	if left == 0 && right == 0 {
		if isValidParenthesis(s) {
			results = append(results, s)
		}
	}

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

func main() {
	a := []string{}
	b := "abc"
	a = append(a, b)
	b = "ab"
	fmt.Print(a)
}
