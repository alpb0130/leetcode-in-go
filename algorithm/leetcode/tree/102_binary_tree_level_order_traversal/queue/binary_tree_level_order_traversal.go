package queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use a queue to store nodes of next level. Inorder to know which nodes are in the same
// level, you need to push the node at one level at the same for loop and process them in
// the same for loop. The outer for loop is loop over each level.
// Time: O(n)
// Space: O(2^h) - h is the height of the tree
func levelOrder(root *TreeNode) [][]int {
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
		res = append(res, subRes)
	}
	return res
}

type queue []*TreeNode

func (q *queue) Offer(n *TreeNode) {
	*q = append(*q, n)
}
func (q *queue) Poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:len(*q)]
	return n
}
func (q *queue) Len() int {
	return len(*q)
}
