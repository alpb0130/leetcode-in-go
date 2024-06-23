package stack

import (
	"strconv"
	"strings"
)

// Understand the problem first!!! The problem requires to count the exclusive execution time of
// every function. Exclusive means the time does't include that of other nested functions. The
// execution time means time when the function is running. If the function is running twice
// sequentially, we count the time twice.
// Use a stack to record current running function (because there are nested). We solve this problem
// by handle every intervals. For each interval, either there is a new start or end, we need to
// decide which function is running during the interval. So we use prevTime to always note down
// most recent timestamp. Check which function is running by check the stack top and add the time
// to the function's map (slice). Be careful about the end time boundary because it means the end
// of a time.
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
		if strs[1] == "start" {
			if s.Len() > 0 {
				res[s.Peek()] += timestamp - prevTime
			}
			s.Push(id)
			prevTime = timestamp
		} else {
			res[s.Pop()] += timestamp - prevTime + 1
			prevTime = timestamp + 1
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
