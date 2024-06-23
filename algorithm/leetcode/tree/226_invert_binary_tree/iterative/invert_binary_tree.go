package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder traverse
// Time: O(n)
// Space: O(h)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	s := &stack{root}
	for s.Len() > 0 {
		node := s.Pop()
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
	}
	return root

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
