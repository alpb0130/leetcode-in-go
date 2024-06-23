package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(logn)
func sumOfLeftLeaves(root *TreeNode) int {
	return sumHelper(root, false)
}

func sumHelper(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}
	return sumHelper(root.Left, true) + sumHelper(root.Right, false)
}
