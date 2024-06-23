package queue

type TreeNode struct {
	Val  int
	Next []*TreeNode
}

func NaryTreeLevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := Queue{root}
	res := [][]int{}
	for queue.Len() != 0 {
		size := queue.Len()
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := queue.Poll()
			tmp = append(tmp, node.Val)
			for _, next := range node.Next {
				if next != nil {
					queue.Offer(next)
				}
			}
		}
		res = append(res, tmp)
	}
	return res
}

type Queue []*TreeNode

func (q *Queue) Len() int {
	return len(*q)
}
func (q *Queue) Offer(n *TreeNode) {
	*q = append(*q, n)
}
func (q *Queue) Poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}
