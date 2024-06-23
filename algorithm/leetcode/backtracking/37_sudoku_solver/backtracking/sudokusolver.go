package backtracking

// DFS in recursive style. If the current solution is not valid anymore, backtracking.
// Time: ???
// Space: O(n*n) - n is the length of the board
func solveSudoku(board [][]byte) {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		columns[i] = map[byte]bool{}
		squares[i] = map[byte]bool{}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			rows[i][board[i][j]] = true
			columns[j][board[i][j]] = true
			squares[(i/3)*3+j/3][board[i][j]] = true
		}
	}
	// recursive call
	solveSudokuHelper(board, 0, rows, columns, squares)
}

func solveSudokuHelper(board [][]byte, idx int, rows, columns, squares []map[byte]bool) bool {
	if idx == 81 {
		return true
	}
	i := idx / 9
	j := idx % 9
	if board[i][j] != '.' {
		return solveSudokuHelper(board, idx+1, rows, columns, squares)
	}
	for n := '1'; n <= '9'; n++ {
		b := byte(n)
		if !rows[i][b] && !columns[j][b] && !squares[(i/3)*3+j/3][b] {
			board[i][j] = b
			rows[i][b], columns[j][b], squares[(i/3)*3+j/3][b] = true, true, true
			if solveSudokuHelper(board, idx+1, rows, columns, squares) {
				return true
			}
			delete(rows[i], b)
			delete(columns[j], b)
			delete(squares[(i/3)*3+j/3], b)
		}
	}
	board[i][j] = '.'
	return false
}
