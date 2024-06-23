package twostack

import (
	"strconv"
	"strings"
)

// Time: O(n)
// Space: O(n)
func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	numStack := &intStack{}
	opStack := &strStack{}
	opOrder := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	i, j := 0, 0
	for j < len(s) {
		if opOrder[s[j]] != 0 {
			num, _ := strconv.Atoi(s[i:j])
			i = j + 1
			numStack.Push(num)
			for opStack.Len() > 0 && opOrder[opStack.Peek()] >= opOrder[s[j]] {
				right := numStack.Pop()
				left := numStack.Pop()
				numStack.Push(calc(left, right, opStack.Pop()))
			}
			opStack.Push(s[j])
		}
		j++
	}
	if i != j {
		num, _ := strconv.Atoi(s[i:j])
		numStack.Push(num)
	}
	for opStack.Len() > 0 {
		right := numStack.Pop()
		left := numStack.Pop()
		numStack.Push(calc(left, right, opStack.Pop()))
	}
	return numStack.Pop()
}

func calc(a, b int, op byte) int {
	if op == '+' {
		return a + b
	} else if op == '-' {
		return a - b
	} else if op == '*' {
		return a * b
	} else {
		return a / b
	}
}

type strStack []byte

func (s *strStack) Len() int {
	return len(*s)
}
func (s *strStack) Pop() byte {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *strStack) Push(x byte) {
	*s = append(*s, x)
}
func (s *strStack) Peek() byte {
	return (*s)[len(*s)-1]
}

type intStack []int

func (s *intStack) Len() int {
	return len(*s)
}
func (s *intStack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *intStack) Push(x int) {
	*s = append(*s, x)
}
func (s *intStack) Peek() int {
	return (*s)[len(*s)-1]
}
