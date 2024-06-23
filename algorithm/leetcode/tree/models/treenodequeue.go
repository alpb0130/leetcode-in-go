package models

type TreeNodeQueue []*TreeNode

func (s *TreeNodeQueue) Offer(n *TreeNode) {
	*s = append(*s, n)
}

func (s *TreeNodeQueue) Poll() *TreeNode {
	n := (*s)[0]
	*s = (*s)[1:]
	return n
}

func (s *TreeNodeQueue) Len() int {
	return len(*s)
}
