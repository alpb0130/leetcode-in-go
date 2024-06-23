package stack

import "math"

// Use a stack. The idea is every time we see a ')', we check stack and see whether we
// have '(', if we do, we know we have valid string. But after we pop '(', it's two
// different conditions for stack is empty and non-empty.
// If '(', push current idx into stack.
// If ')'
//     If stack is empty, current substring is invalid and we record the index as our
//     substring boundary.
//     If stack is not empty. Pop peek '(' index.
//         If stack is empty after pop, we update max using curIdx-preIdx, because we know
//         the current valid substring is bounded by preIdx (like "()(())")
//         If stack is not empty after pop, we update max using curIdx-peekIdx because we
//         know the current valid substring is bounded by the peek '(' in stack (like "(()")
// Time: O(n)
// Space: O(n)
func longestValidParentheses(s string) int {
	st := &stack{}
	res := 0
	preIndex := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(i)
		} else {
			if st.Len() == 0 {
				preIndex = i
			} else {
				st.Pop()
				if st.Len() == 0 {
					res = maxInt(res, i-preIndex)
				} else {
					res = maxInt(res, i-st.Peek())
				}
			}
		}
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
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
