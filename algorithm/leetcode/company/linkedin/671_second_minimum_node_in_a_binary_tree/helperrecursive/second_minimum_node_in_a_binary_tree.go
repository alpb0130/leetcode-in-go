package helperrecursive

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

// Time: O(n)
// Space: O(h)
func findSecondMinimumValue(root *models.TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	left := secondMinimumHelper(root.Left, root.Val)
	right := secondMinimumHelper(root.Right, root.Val)
	if left != -1 && right != -1 {
		return minInt(left, right)
	} else {
		return maxInt(left, right)
	}
}

func secondMinimumHelper(root *models.TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val != min {
		return root.Val
	}
	left := secondMinimumHelper(root.Left, min)
	right := secondMinimumHelper(root.Right, min)
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
