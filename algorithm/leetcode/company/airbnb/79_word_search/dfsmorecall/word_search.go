package dfsmorecall

// DFS but put all checks (index check, visit check, true check in the callee). This would lead
// to more calls but much cleaner.
// Time: worst case O(m*n*3^(m*n)) - [["a", "a"]["a", "a"]] but want to find "aaab". The actual
//       time should be less than that because it impossible for each position we always have 4
//       choice.
// Space: O(m*n) - m*n is the size of the board
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	m := len(board)
	n := len(board[0])
	diffs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(board, word, map[int]bool{}, diffs, i, j, m, n) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, isVisited map[int]bool, diffs [][]int, i, j, m, n int) bool {
	if len(word) == 0 {
		return true
	}
	if isVisited[i*n+j] || !validIndex(i, j, m, n) || board[i][j] != word[0] {
		return false
	}
	isVisited[i*n+j] = true
	for _, diff := range diffs {
		ii := i + diff[0]
		jj := j + diff[1]
		if searchWord(board, word[1:], isVisited, diffs, ii, jj, m, n) {
			return true
		}
	}
	isVisited[i*n+j] = false
	return false
}

func validIndex(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
