package dfs

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

// Use DFS and the feature of BST. If current root is > target, no need to go to right tree because
// all right children will be much larger. If current root is < target, no need to go to left tree
// because all left children will be much larger.
// Time: O(h) - h is the height of the tree
// Space: O(h)
func closestValue(root *models.TreeNode, target float64) int {
	var left, right int
	if float64(root.Val) == target {
		return root.Val
	}
	res := root.Val
	diff := math.Abs(float64(res) - target)
	if root.Left != nil && float64(root.Val) > target {
		left = closestValue(root.Left, target)
		if diff > math.Abs(float64(left)-target) {
			res = left
		}
	}
	if root.Right != nil && float64(root.Val) < target {
		right = closestValue(root.Right, target)
		if diff > math.Abs(float64(right)-target) {
			res = right
		}
	}
	return res
}
