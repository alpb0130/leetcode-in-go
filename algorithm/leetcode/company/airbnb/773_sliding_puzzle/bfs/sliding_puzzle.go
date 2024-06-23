package bfs

// Build the transform graph and dfs. Basically the connection between two board
// exist if we can change one board to another by sway 0 with neighbor. Take care of
// the corner case like if the init board is good. Use 1-D array to represent board.
// Another way is to use string but go is not good at string operations.
func slidingPuzzle(board [][]int) int {
	nextMap := map[int][]int{
		0: []int{1, 3},
		1: []int{0, 2, 4},
		2: []int{1, 5},
		3: []int{0, 4},
		4: []int{1, 3, 5},
		5: []int{2, 4},
	}
	isVisited := map[int]bool{}
	flat := []int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			flat = append(flat, board[i][j])
		}
	}
	q := &queue{flat}
	isVisited[getBoardHash(flat)] = true
	res := 0
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			b := q.Poll()
			hash := getBoardHash(b)
			if hash == 123450 {
				return res
			}
			zero := findZeroIndex(b)
			for _, next := range nextMap[zero] {
				newBoard := clone(b)
				newBoard[zero], newBoard[next] = newBoard[next], newBoard[zero]
				newHash := getBoardHash(newBoard)
				if !isVisited[newHash] {
					isVisited[newHash] = true
					q.Offer(newBoard)
				}
			}
		}
		res++
	}
	return -1
}

type queue [][]int

func (q *queue) Offer(x []int) {
	*q = append(*q, x)
}
func (q *queue) Poll() []int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *queue) Len() int {
	return len(*q)
}

func getBoardHash(board []int) int {
	idx := 0
	for i := 0; i < len(board); i++ {
		idx = idx*10 + board[i]
	}
	return idx
}

func clone(board []int) []int {
	copyBoard := make([]int, 6)
	copy(copyBoard, board)
	return copyBoard
}

func findZeroIndex(board []int) int {
	for i, num := range board {
		if num == 0 {
			return i
		}
	}
	return 0
}
