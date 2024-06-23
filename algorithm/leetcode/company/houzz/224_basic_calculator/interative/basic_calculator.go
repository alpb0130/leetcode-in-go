package interative

import (
	"strconv"
	"strings"
)

// Iterative and use stack. Just check every char and if we see '(', push current sum and sign into
// stack. If we see a ')', set current sum to num and pop out sum and number.
// Two tricks:
// 1. Set sign to be 1 and -1 so that we can do the calculation quickly
// 2. Set sign to be 1 and -2 so that we can use one int stack
// Time: O(n) - length of string
// Space: O(k) - k is number of "()"
func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	numOpStack := &stack{}
	sum := 0
	num := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit, _ := strconv.Atoi(string(s[i]))
			num = num*10 + digit
		} else if s[i] == '+' || s[i] == '-' {
			sum += (sign * num)
			num = 0
			if s[i] == '+' {
				sign = 1
			} else {
				sign = -1
			}
		} else if s[i] == '(' {
			numOpStack.Push(sum)
			numOpStack.Push(sign)
			sum = 0
			sign = 1
		} else {
			sum += (sign * num)
			num = sum
			sign = numOpStack.Pop()
			sum = numOpStack.Pop()
		}
	}
	sum += (sign * num)
	return sum

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
