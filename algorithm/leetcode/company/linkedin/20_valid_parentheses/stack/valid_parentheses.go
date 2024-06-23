package stack

// Use stack. If see left, push right to simplify.
// Time: O(n)
// Space: O(n)
func isValid(s string) bool {
	ss := &stack{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			ss.Push(')')
		case '{':
			ss.Push('}')
		case '[':
			ss.Push(']')
		default:
			if ss.Len() == 0 || ss.Pop() != s[i] {
				return false
			}
		}
	}
	return ss.Len() == 0
}

type stack []byte

func (s *stack) Push(n byte) {
	*s = append(*s, n)
}
func (s *stack) Pop() byte {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
func (s *stack) Len() int {
	return len(*s)
}
