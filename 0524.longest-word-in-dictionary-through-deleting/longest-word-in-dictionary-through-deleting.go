package leetcode

import "fmt"

func findLongestWord(s string, dictionary []string) string {
	result := ""
	for _, word := range dictionary {

		i := 0
		j := 0
		fmt.Printf("word=%v\n", word)

		for j < len(word) {
			if i < len(s) {
				if s[i] == word[j] {
					i++
					j++
				} else {
					i++
				}
			} else {
				break
			}
			fmt.Printf("i=%v, j=%v\n", i, j)
		}

		if j == len(word) {
			if len(word) > len(result) || len(word) == len(result) && result > word {
				result = word
			}
		}

	}
	return result
}
