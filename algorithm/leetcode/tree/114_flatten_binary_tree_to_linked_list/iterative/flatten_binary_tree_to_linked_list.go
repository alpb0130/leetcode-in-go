package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder morris traversal
// Time: O(n)
// Space: O(1)
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			left := cur.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
