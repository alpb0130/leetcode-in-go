package solution

// Time: O(1)
// Space: O(1)
func validTicTacToe(board []string) bool {
	xCount, oCount := 0, 0
	for _, b := range board {
		for i := 0; i < len(b); i++ {
			if b[i] == 'X' {
				xCount++
			}
			if b[i] == 'O' {
				oCount++
			}
		}
	}
	if xCount != oCount && xCount-oCount != 1 {
		return false
	}
	if checkWin('X', board) && xCount-oCount != 1 {
		return false
	}
	if checkWin('O', board) && xCount-oCount != 0 {
		return false
	}
	return true
}

func checkWin(target byte, board []string) bool {
	for i := 0; i < len(board); i++ {
		if board[0][i] == target && board[1][i] == target && board[2][i] == target {
			return true
		}
		if board[i][0] == target && board[i][1] == target && board[i][2] == target {
			return true
		}
	}
	if board[0][0] == target && board[1][1] == target && board[2][2] == target {
		return true
	}
	if board[0][2] == target && board[1][1] == target && board[2][0] == target {
		return true
	}
	return false
}
