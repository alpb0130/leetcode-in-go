package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Find the node recursively in O(logn) time and the delete should only take O(1) time.
// Time: O(logn)
// Space: O(h)
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	left := root.Left
	pred := left
	for pred.Right != nil {
		pred = pred.Right
	}
	pred.Right = root.Right
	return left
}
