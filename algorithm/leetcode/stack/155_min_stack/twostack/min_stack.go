package twostack

// Use two stacks. One stack stores the raw value and another stack stores snapshot min.
// All operations are O(1) time.
// Space: O(n)
type MinStack struct {
	rawStack *stack
	minStack *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{&stack{}, &stack{}}
}

func (this *MinStack) Push(x int) {
	this.rawStack.Push(x)
	min := x
	if this.minStack.Len() > 0 && this.minStack.Peek() < x {
		min = this.minStack.Peek()
	}
	this.minStack.Push(min)
}

func (this *MinStack) Pop() {
	this.rawStack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	return this.rawStack.Peek()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Peek()
}

type stack []int

func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
