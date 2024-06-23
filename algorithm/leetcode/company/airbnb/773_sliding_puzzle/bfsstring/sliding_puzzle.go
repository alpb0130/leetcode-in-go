package bfsstring

import (
	"strconv"
	"strings"
)

// Same as another way but use string to represent the board so that we don't need to generate hash
// Time: O(1)
// Space: O(1)
func slidingPuzzle(board [][]int) int {
	var str strings.Builder
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			str.WriteString(strconv.Itoa(board[i][j]))
		}
	}
	if str.String() == "123450" {
		return 0
	}
	moveMap := map[int][]int{
		0: []int{1, 3},
		1: []int{0, 2, 4},
		2: []int{1, 5},
		3: []int{0, 4},
		4: []int{1, 3, 5},
		5: []int{2, 4},
	}
	isVisited := map[string]bool{str.String(): true}
	step := 0
	q := &queue{str.String()}
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			b := q.Poll()
			zeroIdx := strings.Index(b, "0")
			for _, next := range moveMap[zeroIdx] {
				newBoard := generateNewBoardString(b, zeroIdx, next)
				if newBoard == "123450" {
					return step + 1
				}
				if !isVisited[newBoard] {
					isVisited[newBoard] = true
					q.Offer(newBoard)
				}
			}
		}
		step++
	}
	return -1
}

type queue []string

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Poll() string {
	s := (*q)[0]
	*q = (*q)[1:]
	return s
}
func (q *queue) Offer(s string) {
	*q = append(*q, s)
}
func generateNewBoardString(board string, i, j int) string {
	bytes := []byte(board)
	bytes[i], bytes[j] = bytes[j], bytes[i]
	return string(bytes)
}
