package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(logn)
// Space: O(h)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if root == nil {
			break
		}
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
