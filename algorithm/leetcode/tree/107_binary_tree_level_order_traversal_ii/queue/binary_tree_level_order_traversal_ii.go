package queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Similar to level order traversal but put the later list in the head
// Time: O(n)
// Space: O(n)
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := &queue{root}
	for q.Len() != 0 {
		num := q.Len()
		subRes := []int{}
		for i := 0; i < num; i++ {
			node := q.Poll()
			subRes = append(subRes, node.Val)
			if node.Left != nil {
				q.Offer(node.Left)
			}
			if node.Right != nil {
				q.Offer(node.Right)
			}
		}
		res = append([][]int{subRes}, res...)
	}
	return res
}

type queue []*TreeNode

func (q *queue) Poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:len(*q)]
	return n
}
func (q *queue) Offer(n *TreeNode) {
	*q = append(*q, n)
}
func (q *queue) Len() int {
	return len(*q)
}
