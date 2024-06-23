package iterative

import "algorithm/leetcode/tree/models"

// Iterative
// Time: O(n)
// Space: O(h) - h is the height of the tree
func preorderTraversal(root *models.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	s := &stack{root}
	res := []int{}
	for s.Len() > 0 {
		node := s.Pop()
		res = append(res, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return res

}

type stack []*models.TreeNode

func (s *stack) Push(n *models.TreeNode) {
	*s = append(*s, n)
}
func (s *stack) Pop() *models.TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return n
}
func (s *stack) Len() int {
	return len(*s)
}
