package stackandlinear

import "math"

// The problem requires us to verify whether a preorder array is from a
// valid BST. Our thought is if we can simulate the inorder traverse and
// make sure current number is larger than prev number, then it should be
// valid. But how do we get the prev num? Looking at the array, we know the
// first element should be root and smaller elements after root should be
// in left subtree. Larger elements after roo should be in right subtree.
// We use a stack to keep push num into it till we find a number larger
// than the stack peek(prev). At this point, we know we are on the right
// subtree and we can update min which should be the last number pushed out.
// Keep doing this process. If we find a num which is smaller than current min,
// we know it's invalid.
// Time: O(n)
// Space: O(logn)
func verifyPreorder(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	s := &stack{}
	min := math.MinInt32
	for _, val := range preorder {
		for s.Len() > 0 && s.Peek() < val {
			min = s.Pop()
		}
		if val <= min {
			return false
		}
		s.Push(val)
	}
	return true
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) Push(i int) {
	*s = append(*s, i)
}
func (s *stack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}
