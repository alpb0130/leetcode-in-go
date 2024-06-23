package recursive

import (
	"strconv"
	"strings"
)

// Every time we see a "(", do a recursive call.
// Time: O(n)
// Space: O(k) - k is number of "()"
func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	// calc
	i := 0
	return calc(s, &i)
}

func calc(s string, i *int) int {
	if *i == len(s) {
		return 0
	}
	st := &stack{}
	var sign byte
	sign = '+'
	num := 0
	for (*i) < len(s) {
		// fmt.Println("index:", *i, "char:", string(s[*i]))
		if s[*i] >= '0' && s[*i] <= '9' {
			digit, _ := strconv.Atoi(string(s[*i]))
			num = num*10 + digit
		} else if s[*i] == '(' {
			*i++
			num = calc(s, i)
		}
		if *i >= len(s)-1 || s[*i] == ')' || s[*i] == '-' || s[*i] == '+' || s[*i] == '*' || s[*i] == '/' {
			// fmt.Println("sign:", string(sign), "num:", num, "stack", *st)
			if sign == '-' {
				st.Push(-num)
			}
			if sign == '+' {
				st.Push(num)
			}
			if sign == '*' {
				st.Push(st.Pop() * num)
			}
			if sign == '/' {
				st.Push(st.Pop() / num)
			}
			if *i < len(s) && s[*i] == ')' {
				(*i)++
				break
			} else if (*i) < len(s) {
				sign = s[*i]
				num = 0
			}
		}
		(*i)++
	}
	sum := 0
	for st.Len() > 0 {
		sum += st.Pop()
	}
	return sum
}

type stack []int

func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}
func (s *stack) Len() int {
	return len(*s)
}
