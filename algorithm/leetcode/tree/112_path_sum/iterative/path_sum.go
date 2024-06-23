package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(h)
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	s := stack{root}
	for s.Len() > 0 {
		node := s.Pop()
		if node.Left == nil && node.Right == nil && node.Val == sum {
			return true
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			s.Push(node.Right)
		}
		if node.Left != nil {
			node.Left.Val += node.Val
			s.Push(node.Left)
		}
	}
	return false
}

type stack []*TreeNode

func (s *stack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return n
}
func (s *stack) Push(n *TreeNode) {
	*s = append(*s, n)
}
func (s *stack) Len() int {
	return len(*s)
}
