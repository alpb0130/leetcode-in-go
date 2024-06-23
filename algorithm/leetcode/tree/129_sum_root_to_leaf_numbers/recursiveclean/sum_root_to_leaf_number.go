package recursiveclean

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder traversal
// Time: O(n)
// Space: O(h)
func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 10*val + root.Val
	}
	return sumNumbersHelper(root.Left, 10*val+root.Val) + sumNumbersHelper(root.Right, 10*val+root.Val)
}
