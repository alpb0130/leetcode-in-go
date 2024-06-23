package iterativeandstack

import (
	"strconv"
	"unicode"
)

// Iterative. Use two stacks. One is for current total string. The other is for # of the substring.
// If we see a number, construct number
// If we see a letter, contribute the letter to current string
// If we see a '[', we know we should process the substring but before that we want to note down
// the current string and its number (push into stack). And then reset current str to empty and
// number to 0.
// If we see a ']', we know we have finish the substring (str). We can pop out the number and last
// total string. And then construct the new string using last_total_string+num*str. Finally we
// update the str to be the new string and number to be 0.
// Return str if we hit the string end.
// Time: O(n)
// Space: O(k) - k is depth of bracket
func decodeString(s string) string {
	if len(s) <= 1 {
		return s
	}
	intStack := &IntStack{}
	strStack := &StringStack{}
	num := 0
	str := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			digit, _ := strconv.Atoi(string(s[i]))
			num = num*10 + digit
		} else if unicode.IsLetter(rune(s[i])) {
			str += string(s[i])
		} else if s[i] == '[' {
			strStack.Push(str)
			intStack.Push(num)
			str = ""
			num = 0
		} else {
			num = intStack.Pop()
			sumStr := strStack.Pop()
			for i := 0; i < num; i++ {
				sumStr += str
			}
			str = sumStr
			num = 0
		}
	}
	return str
}

type IntStack []int

func (i *IntStack) Len() int {
	return len(*i)
}
func (i *IntStack) Push(x int) {
	*i = append(*i, x)
}
func (i *IntStack) Pop() int {
	x := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return x
}

type StringStack []string

func (s *StringStack) Len() int {
	return len(*s)
}
func (s *StringStack) Push(x string) {
	*s = append(*s, x)
}
func (s *StringStack) Pop() string {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
