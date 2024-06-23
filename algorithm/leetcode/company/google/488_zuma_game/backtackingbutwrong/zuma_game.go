package backtackingbutwrong

import "math"

// DFS and backtracking. But this solution is actually wrong.
func findMinStep(board string, hand string) int {
	handMap := map[byte]int{}
	for i := 0; i < len(hand); i++ {
		handMap[hand[i]]++
	}
	numHand := len(hand)
	res := math.MaxInt32
	// dfs
	dfs(board, handMap, numHand, 0, &res)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dfs(board string, handMap map[byte]int, numHand, cur int, res *int) {
	if len(board) == 0 {
		*res = minInt(*res, cur)
		return
	}
	if cur == numHand {
		return
	}
	i := 0
	for i < len(board) {
		j := i
		// The pruning here is not correct. We shouldn't do the greedy pruning and always try to
		// insert a ball at the place where the ball has the same color. We should just add a
		// ball at any place
		for j < len(board) && board[i] == board[j] {
			j++
		}
		neededHand := 3 - (j - i)
		if handMap[board[i]] >= neededHand {
			handMap[board[i]] -= neededHand
			newBoard := update(board[:i] + board[j:])
			dfs(newBoard, handMap, numHand, cur+neededHand, res)
			handMap[board[i]] += neededHand
		}
		i = j
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// delete consecutive balls with the same color
func update(board string) string {
	if len(board) <= 2 {
		return board
	}
	i := 0
	for i < len(board) {
		j := i
		for j < len(board) && board[i] == board[j] {
			j++
		}
		if j-i >= 3 {
			board = board[:i] + board[j:]
			i = 0
		} else {
			i = j
		}
	}
	return board
}
