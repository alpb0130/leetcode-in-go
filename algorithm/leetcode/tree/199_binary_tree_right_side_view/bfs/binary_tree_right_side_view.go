package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Level order traversal
// Time: O(n)
// Space: O(h)
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := &queue{root}
	for q.Len() > 0 {
		num := q.Len()
		for i := 0; i < num; i++ {
			node := q.Poll()
			if i == 0 {
				res = append(res, node.Val)
			}
			if node.Right != nil {
				q.Offer(node.Right)
			}
			if node.Left != nil {
				q.Offer(node.Left)
			}
		}
	}
	return res
}

type queue []*TreeNode

func (q *queue) Poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}
func (q *queue) Offer(n *TreeNode) {
	*q = append(*q, n)
}
func (q *queue) Len() int {
	return len(*q)
}
