package recursive

import (
	"strconv"
	"strings"
)

// Recursive. Every time we see a left "(", start a new func and if we see a ")" or hit the end
// return the current function. Don't forget to process the last number.
// Time: O(n) - length of string
// Space: O(k) - k is number of "()"
func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}
	index := 0
	return calc(s, &index)
}

func calc(s string, i *int) int {
	if *i == len(s) {
		return 0
	}
	sum := 0
	num := 0
	sign := byte('+')
	for (*i) < len(s) {
		if s[*i] >= '0' && s[*i] <= '9' {
			digit, _ := strconv.Atoi(string(s[*i]))
			num = num*10 + digit
		} else if s[*i] == '+' || s[*i] == '-' {
			if sign == '+' {
				sum += num
			} else {
				sum -= num
			}
			num = 0
			sign = s[*i]
		} else if s[*i] == '(' {
			(*i)++
			num = calc(s, i)
		} else if s[*i] == ')' {
			if sign == '+' {
				sum += num
			} else {
				sum -= num
			}
			return sum
		}
		(*i)++
	}
	if sign == '+' {
		sum += num
	} else {
		sum -= num
	}
	return sum
}
