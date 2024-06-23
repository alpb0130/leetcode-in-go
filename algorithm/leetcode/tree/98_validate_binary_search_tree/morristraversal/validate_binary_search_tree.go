package morristraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use morris traversal
// Time: O(n) - n is the number of tree node
// Space: O(1)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Left != nil {
			right := cur.Left
			for right.Right != nil {
				right = right.Right
			}
			right.Right = cur
			left := cur.Left
			cur.Left = nil
			cur = left
		} else {
			if pre != nil && cur.Val <= pre.Val {
				return false
			}
			pre = cur
			cur = cur.Right
		}
	}
	return true
}
