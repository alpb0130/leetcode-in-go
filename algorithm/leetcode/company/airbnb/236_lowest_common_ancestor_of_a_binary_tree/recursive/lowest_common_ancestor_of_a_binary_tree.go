package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive postorder traverse. Always return the root of subtree if its val is
// equal to either p or q. This will not lost info because if no result return from
// another subtree, this root will be the final result. Pretty concise solution. If
// result from both left subtree and right subtree is not nil, this means the current
// root is the result because p and q are from different subtree.
// Time: O(n)
// Space: O(h)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}

