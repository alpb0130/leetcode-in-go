package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// According to master theorem, the time complexity is logn
// T(n) = T(n/2)+f(n)
// f(n) = O(1)
// O(fn) = n^0(logn)^0
// O(T(n)) = logn
// Time: O(logn)
// Space: O(h)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
