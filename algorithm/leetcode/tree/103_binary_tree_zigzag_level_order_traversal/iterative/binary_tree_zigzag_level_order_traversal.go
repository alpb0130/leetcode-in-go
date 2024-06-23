package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree level order traversal but add value to the front or end of the list according
// to the level number
// Time: O(n)
// Space: O(2 to the power of h)
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := &queue{root}
	leftToRight := true
	for q.Len() != 0 {
		n := q.Len()
		level := []int{}
		for i := 0; i < n; i++ {
			node := q.Poll()
			if leftToRight {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				q.Offer(node.Left)
			}
			if node.Right != nil {
				q.Offer(node.Right)
			}
		}
		res = append(res, level)
		leftToRight = !leftToRight
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
