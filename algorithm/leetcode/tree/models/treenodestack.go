package models

type TreeNodeStack []*TreeNode

func (s *TreeNodeStack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s *TreeNodeStack) Pop() *TreeNode {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *TreeNodeStack) Len() int {
	return len(*s)
}
