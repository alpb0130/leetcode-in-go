package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder traversal
// Time: O(n)
// Space: O(h)
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, val int) int {
	if root.Left == nil && root.Right == nil {
		return 10*val + root.Val
	}
	var left, right int
	if root.Left != nil {
		left = sumNumbersHelper(root.Left, 10*val+root.Val)
	}
	if root.Right != nil {
		right = sumNumbersHelper(root.Right, 10*val+root.Val)
	}
	return left + right
}
