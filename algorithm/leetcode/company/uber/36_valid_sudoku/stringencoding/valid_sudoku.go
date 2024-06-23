package stringencoding

import "strconv"

// The idea is to check whether every row, column and box have duplicate numbers.
// We can use separate map for every row, column and box. But we can also use only one.
// Just make sure the same number in different row/col/box has different encoding string.
// Time: O(1)
// Space: O(1)
func isValidSudoku(board [][]byte) bool {
	strMap := map[string]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row, col, box := encoding(board, i, j)
				if strMap[row] || strMap[col] || strMap[box] {
					return false
				}
				strMap[row] = true
				strMap[col] = true
				strMap[box] = true
			}
		}
	}
	return true
}

func encoding(board [][]byte, i, j int) (string, string, string) {
	row := strconv.Itoa(i) + "r" + string(board[i][j])
	col := strconv.Itoa(j) + "c" + string(board[i][j])
	box := strconv.Itoa(i/3) + "b" + string(board[i][j]) + "b" + strconv.Itoa(j/3)
	return row, col, box
}
