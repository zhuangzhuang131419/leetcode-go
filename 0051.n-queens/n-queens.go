package leetcode

var results [][]string

func solveNQueens(n int) [][]string {
	results = [][]string{}
	board := make([]string, n)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}

	for i := 0; i < n; i++ {
		board[i] = s
	}
	return results
}

func backtrack(n int, board []string, chessNum int) {

	if chessNum == n {
		result := []string{}
		for i := 0; i < n; i++ {
			result = append(result, board[i])
		}

		results = append(results, result)
		return
	}

	for i := 0; i < n; i++ {
		if checkChessValid(board, chessNum, i) {
			s := []rune(board[chessNum])
			s[i] = 'Q'
			board[chessNum] = string(s)
			backtrack(n, board, chessNum+1)
			s[i] = '.'
			board[chessNum] = string(s)
		}
	}
}

func checkChessValid(board []string, row int, col int) bool {
	n := len(board)

	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	i := row - 1
	j := col + 1

	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	i = row - 1
	j = col - 1

	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	return true
}
