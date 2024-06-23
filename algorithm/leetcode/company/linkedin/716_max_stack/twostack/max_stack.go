package twostack

// A better method is two use tree map and double linked list (https://leetcode.com/problems/max-stack/solution/)
// Use two stack, one for exact value and another for snapshot max. All operations must
// be applied to both stacks.
// Example: If the real stack is [2, 1, 5, 3, 9], the snapshot max stack will be like
//          [2, 2, 5, 5, 9].
// Reference: https://leetcode.com/problems/max-stack/solution/
// Space: O(n)
type MaxStack struct {
	s        *stack
	maxStack *stack
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{&stack{}, &stack{}}
}

// For max stack, compare current value with current max. If current value is larger than current
// max, push the current value as max, else push the current max as max. Be careful of the maxStack
// is empty.
// Time: O(1)
// Space: O(1)
func (this *MaxStack) Push(x int) {
	this.s.Push(x)
	peek := x
	if this.maxStack.Len() > 0 && this.maxStack.Peek() > peek {
		peek = this.maxStack.Peek()
	}
	this.maxStack.Push(peek)
}

// Time: O(1)
// Space: O(1)
func (this *MaxStack) Pop() int {
	this.maxStack.Pop()
	return this.s.Pop()
}

// Time: O(1)
// Space: O(1)
func (this *MaxStack) Top() int {
	return this.s.Peek()
}

// Time: O(1)
// Space: O(1)
func (this *MaxStack) PeekMax() int {
	return this.maxStack.Peek()
}

// Find the max and pop it from the exact value stack. And then pop other value back.
// Time: O(1)
// Space: O(n)
func (this *MaxStack) PopMax() int {
	max := this.maxStack.Peek()
	tmp := &stack{}
	for this.s.Len() > 0 && this.s.Peek() != max {
		this.maxStack.Pop()
		tmp.Push(this.s.Pop())
	}
	x := this.s.Pop()
	this.maxStack.Pop()
	for tmp.Len() > 0 {
		this.Push(tmp.Pop())
	}
	return x
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
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
