package selfrecursive

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

// The tree is special, each node must have both left/right child or neither and the parent
// value is the smaller one of its left and right child. The problem requires us to find
// the second minimum value in the tree. We all know the root is the the minimum value.
// So we only to find the min value which is different from root in left and right subtree.
// For a subtree, if the current root val is not equal to the subtree root value, we know that
// we have found the subtree's smallest value and we don't need to find deeply. It's a potential
// second min. If the current root val is equal to the subtree root value, we need to keep search
// this subtree until we find a value which is different from the current root or reach the leaves.
// If reach leaves, return -1. Finally compare the left min and right min to get correct value.
// Time: O(n)
// Space: O(h)
func findSecondMinimumValue(root *models.TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	left, right := root.Left.Val, root.Right.Val
	if left == root.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if right == root.Val {
		right = findSecondMinimumValue(root.Right)
	}
	if left != -1 && right != -1 {
		return minInt(left, right)
	} else {
		return maxInt(left, right)
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
