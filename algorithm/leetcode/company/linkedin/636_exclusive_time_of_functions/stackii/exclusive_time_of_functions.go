package stackii

import (
	"strconv"
	"strings"
)

// The same as another method but consolidate some logic
// Time: O(n) - n is the num of logs
// Space: O(h) - h is the depth of function recursive call
func exclusiveTime(n int, logs []string) []int {
	s := &stack{}
	res := make([]int, n)
	prevTime := 0
	for _, log := range logs {
		strs := strings.Split(log, ":")
		id, _ := strconv.Atoi(strs[0])
		timestamp, _ := strconv.Atoi(strs[2])
		if s.Len() > 0 {
			res[s.Peek()] += timestamp - prevTime
		}
		prevTime = timestamp
		if strs[1] == "start" {
			s.Push(id)
		} else {
			res[s.Pop()]++
			prevTime++
		}
	}
	return res
}

type stack []int

func (s *stack) Push(n int) {
	*s = append(*s, n)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}
