package recursive_i

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder traversal. Bottom-up method. Each call return the max and min value of the tree
// Time: O(n) - n is the number of tree node
// Space: O(h) - h is the height of the tree
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isValid, _, _ := isValidBSTAndMaxMin(root)
	return isValid
}

func isValidBSTAndMaxMin(root *TreeNode) (bool, int, int) {
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	} else if root.Left == nil && root.Right != nil {
		isRightValid, rightMax, rightMin := isValidBSTAndMaxMin(root.Right)
		return isRightValid && root.Val < rightMin, rightMax, root.Val
	} else if root.Left != nil && root.Right == nil {
		isLeftValid, leftMax, leftMin := isValidBSTAndMaxMin(root.Left)
		return isLeftValid && root.Val > leftMax, root.Val, leftMin
	} else {
		isRightValid, rightMax, rightMin := isValidBSTAndMaxMin(root.Right)
		isLeftValid, leftMax, leftMin := isValidBSTAndMaxMin(root.Left)
		return isRightValid && isLeftValid && root.Val > leftMax && root.Val < rightMin, rightMax, leftMin
	}
}
