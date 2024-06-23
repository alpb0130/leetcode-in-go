package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder traversal
// Time: O(n) - n is the number of tree node
// Space: O(h) - h is the height of the tree
func isValidBST(root *TreeNode) bool {
	s := &stack{}
	cur := root
	var pre *TreeNode
	for cur != nil || s.Len() != 0 {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		cur = s.Pop()
		if pre != nil && cur.Val <= pre.Val {
			return false
		}
		pre = cur
		cur = cur.Right
	}
	return true
}

type stack []*TreeNode
func (s *stack) Push(n *TreeNode) {
	*s = append(*s, n)
}
func (s *stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[0:len(*s)-1]
	return n
}
func (s *stack) Len() int {
	return len(*s)
}
