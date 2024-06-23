package iterative

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iteratively
// Time: O(n)
// Space: O(h) - h is the height of the tree
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := []int{}
	s := &stack{}
	cur := root
	for cur != nil || s.Len() != 0 {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		cur = s.Pop()
		results = append(results, cur.Val)
		cur = cur.Right
	}
	return results
}

type stack []*TreeNode
func (s *stack) Push(t *TreeNode) {
	*s = append(*s, t)
}
func (s *stack) Pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[0:len(*s)-1]
	return t
}
func (s *stack) Len() int {
	return len(*s)
}
