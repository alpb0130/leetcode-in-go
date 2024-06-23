package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	s   *stack
	ptr *TreeNode
}

// Reference: https://leetcode.com/problems/binary-search-tree-iterator/discuss/52647/Nice-Comparison-(and-short-Solution)
// Time: O(1)
// Space: O(h)
func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{&stack{}, root}
}

// Time: O(1) - amortized!!!
// Space: O(h)
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for this.ptr != nil {
		this.s.Push(this.ptr)
		this.ptr = this.ptr.Left
	}
	node := this.s.Pop()
	this.ptr = node.Right
	return node.Val
}

// Time: O(1)
// Space: O(h)
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.s.Len() > 0 || this.ptr != nil
}

type stack []*TreeNode

func (s *stack) Push(t *TreeNode) {
	*s = append(*s, t)
}
func (s *stack) Pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return t
}
func (s *stack) Len() int {
	return len(*s)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
