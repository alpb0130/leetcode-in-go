package onestack

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	numStack := &stack{}
	num := 0
	sign := byte('+')
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit, _ := strconv.Atoi(string(s[i]))
			num = num*10 + digit
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if sign == '-' {
				numStack.Push(-num)
			}
			if sign == '+' {
				numStack.Push(num)
			}
			if sign == '*' {
				numStack.Push(numStack.Pop() * num)
			}
			if sign == '/' {
				numStack.Push(numStack.Pop() / num)
			}
			num = 0
			sign = s[i]
		}
	}
	res := 0
	for numStack.Len() != 0 {
		res += numStack.Pop()
	}
	return res
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