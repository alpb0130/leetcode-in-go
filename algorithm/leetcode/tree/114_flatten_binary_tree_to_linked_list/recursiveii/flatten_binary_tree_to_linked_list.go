package recursiveii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Not a straightforward solution. Use recursive postorder traversal.
// Use a global variable "successor" to store the most recent process node in post order.
// Then we will all know the next process node need to set it's right to the "successor",
// set the left to nil and update the successor.
// Time: O(n)
// Space: O(h)
func flatten(root *TreeNode) {
	var successor *TreeNode
	var flattenHelper func(*TreeNode)
	flattenHelper = func(root *TreeNode) {
		if root == nil {
			return
		}
		flattenHelper(root.Right)
		flattenHelper(root.Left)
		root.Right = successor
		root.Left = nil
		successor = root
	}
	flattenHelper(root)
}
