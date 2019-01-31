package array

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidRow(i, board) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if !isValidColumn(i, board) {
			return false
		}
	}
	for n := 0; n < 3; n++ {
		for m := 0; m < 3; m++ {
			if !isValidGrid(n*3, m*3, board) {
				return false
			}
		}
	}

	return true
}

func isValidRow(i int, board [][]byte) bool {
	seen := make(map[byte]bool)
	for j := 0; j < 9; j++ {
		_, ok := seen[board[i][j]]
		if ok && board[i][j] != '.' {
			return false
		} else {
			seen[board[i][j]] = true
		}
	}
	return true
}

func isValidColumn(j int, board [][]byte) bool {
	seen := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		_, ok := seen[board[i][j]]
		if ok && board[i][j] != '.' {
			return false
		} else {
			seen[board[i][j]] = true
		}
	}
	return true
}

func isValidGrid(n int, m int, board [][]byte) bool {
	seen := make(map[byte]bool)
	for i := n; i < n+3; i++ {
		for j := m; j < m+3; j++ {
			_, ok := seen[board[i][j]]
			if ok && board[i][j] != '.' {
				return false
			} else {
				seen[board[i][j]] = true
			}
		}
	}
	return true
}
