package inplace

// In place
// Time: O(m*n)
// Space: O(1)
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	dir := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			for _, d := range dir {
				if isValid(m, n, i+d[0], j+d[1]) {
					sum += board[i+d[0]][j+d[1]] & 1
				}
			}
			if board[i][j]&1 == 1 && (sum == 2 || sum == 3) {
				board[i][j] |= 2
			}
			if board[i][j]&1 == 0 && (sum == 3) {
				board[i][j] |= 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func isValid(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
