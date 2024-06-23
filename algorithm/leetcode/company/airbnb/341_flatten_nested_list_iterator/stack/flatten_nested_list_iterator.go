package stack

type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	GetList() []NestedInteger
}

type NestedIterator struct {
	s *stack
}

func NewNestedIterator(list []NestedInteger) NestedIterator {
	s := &stack{}
	for i := len(list)-1; i >= 0; i-- {
		s.Push(list[i])
	}
	return NestedIterator{
		s: s,
	}
}

func (n NestedIterator) Next() int {
	n.HasNext()
	return n.s.Pop().GetInteger()
}

func (n NestedIterator) HasNext() bool {
	for n.s.Len() > 0 {
		if n.s.Peek().IsInteger() {
			break
		}
		list := n.s.Pop().GetList()
		for i := len(list); i >= 0; i-- {
			n.s.Push(list[i])
		}
	}
	return n.s.Len() > 0
}

type stack []NestedInteger

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Pop() NestedInteger {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(x NestedInteger) {
	*s = append(*s, x)
}
func (s *stack) Peek() NestedInteger {
	return (*s)[len(*s)-1]
}
