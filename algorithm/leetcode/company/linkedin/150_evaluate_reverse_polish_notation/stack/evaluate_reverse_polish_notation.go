package stack

import "strconv"

// Time: O(n)
// Space: O(n) - worst case should still be n
func evalRPN(tokens []string) int {
	s := &stack{}
	for _, token := range tokens {
		if token == "+" {
			s.Push(s.Pop() + s.Pop())
		} else if token == "-" {
			a := s.Pop()
			b := s.Pop()
			s.Push(b - a)
		} else if token == "*" {
			s.Push(s.Pop() * s.Pop())
		} else if token == "/" {
			a := s.Pop()
			b := s.Pop()
			s.Push(b / a)
		} else {
			n, _ := strconv.Atoi(token)
			s.Push(n)
		}
	}
	return s.Pop()
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
