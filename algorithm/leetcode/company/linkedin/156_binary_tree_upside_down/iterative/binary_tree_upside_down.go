package iterative

import "algorithm/leetcode/company/linkedin/models"

// For each node, set its left to the parent's right and its right to the parent.
// Use two pointers to record that. Always process the init root because you need
// to set the left and right of init root to nil so that we can traverse the tee.
// Time: O(n)
// Space: O(1)
func upsideDownBinaryTree(root *models.TreeNode) *models.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	var prev, right *models.TreeNode
	for root != nil {
		next := root.Left
		root.Left = right
		right = root.Right
		root.Right = prev

		prev = root
		root = next
	}
	return prev
}
