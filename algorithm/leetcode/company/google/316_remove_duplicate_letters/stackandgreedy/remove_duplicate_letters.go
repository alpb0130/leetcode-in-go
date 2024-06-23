package stackandgreedy

import "strings"

// We use greedy method. Given an example "bcacb", we know the result should be "acb". But why?
// Let's start with "b". Let say if we put "b" in the final result and then "c". When we see "a",
// we find "a" is smaller than "b" and "c" and we probably want to try put "a" first. But we need
// to verify whether there are "b" and "c" after "a". This can be done using the letter count.
// Every time we process a letter, we reduce the counter by one. When we process "a", if we found
// "b" and "c" both have counter >= 1, we know there are "b", "c". So we can pop out previous "b"
// and "c" and start with "a".
// What if "c" is not in the substring after "a"? In this case we want keep both "b" and "c"
// because "b" is smaller than "c" and we know this is a smaller results even there is one more "b".
// Time: O(n)
// Space: O(n)
func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}
	charCnt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		charCnt[s[i]]++
	}
	st := &stack{}
	isVisited := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if isVisited[s[i]] {
			charCnt[s[i]]--
			continue
		}
		for st.Len() > 0 && st.Peek() > s[i] && charCnt[st.Peek()] > 0 {
			isVisited[st.Pop()] = false
		}
		st.Push(s[i])
		charCnt[s[i]]--
		isVisited[s[i]] = true
	}
	var res strings.Builder
	for _, b := range *st {
		res.WriteString(string(b))
	}
	return res.String()
}

type stack []byte

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() byte {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(x byte) {
	*s = append(*s, x)
}
func (s *stack) Peek() byte {
	return (*s)[len(*s)-1]
}
