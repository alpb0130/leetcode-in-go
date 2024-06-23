package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(h)
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	return (root.Left == nil && root.Right == nil && sum == 0) ||
		hasPathSum(root.Left, sum) ||
		hasPathSum(root.Right, sum)
}
