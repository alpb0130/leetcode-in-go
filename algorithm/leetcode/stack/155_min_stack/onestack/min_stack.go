package onestack

import "math"

// Use only one stack. Before we push the current value, we need to push the min value if
// current value is smaller than cur min value because we need to record it otherwise we
// will lose it. When pop, if current top is equal to cur min, we pop it and update the current
// min to be the second value in the stack (need to pop the second value because it the old min).
// min always store the current min.
// All operations are O(1) time.
// Space: O(n)
type MinStack struct {
	s   *stack
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{&stack{}, math.MaxInt32}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.s.Push(this.min)
		this.min = x
	}
	this.s.Push(x)
}

func (this *MinStack) Pop() {
	if this.s.Pop() == this.min {
		this.min = this.s.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.s.Peek()
}

func (this *MinStack) GetMin() int {
	return this.min
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
