package onestack

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(s string) float64 {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0.0
	}
	st := &stack{}
	num := 0.0
	sign := byte('+')
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit, _ := strconv.Atoi(string(s[i]))
			num = num*10 + float64(digit)
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if sign == '+' {
				st.Push(float64(num))
			} else if sign == '-' {
				st.Push(float64(-num))
			} else if sign == '*' {
				st.Push(num * st.Pop())
			} else {
				st.Push(st.Pop() / num)
			}
			sign = s[i]
			num = 0.0
		}
	}
	fmt.Println(*st)
	res := 0.0
	for st.Len() != 0 {
		value := st.Pop()
		res += value
	}

	return res
}

type stack []float64

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() float64 {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(x float64) {
	*s = append(*s, x)
}
