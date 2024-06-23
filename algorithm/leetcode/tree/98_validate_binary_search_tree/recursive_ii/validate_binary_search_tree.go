package recursive_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder traversal. Top-down method. min and max in input indicate current
// range of the tree with "root"
// Time: O(n) - n is the number of tree node
// Space: O(h) - h is the height of the tree
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root *TreeNode, min, max *int) bool {
	return root == nil ||
		((min == nil || root.Val > *min) &&
			(max == nil || root.Val < *max) &&
			isValid(root.Left, min, &root.Val) &&
			isValid(root.Right, &root.Val, max))
}
