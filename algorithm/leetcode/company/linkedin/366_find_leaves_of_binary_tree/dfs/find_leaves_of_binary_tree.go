package dfs

import (
	"math"

	"algorithm/leetcode/company/linkedin/models"
)

// DFS. For each level, use bottom up method to determine current height
// and then put the value into correct array.
// Time: O(n)
// Space O(h) - h is the height of the tree
func findLeaves(root *models.TreeNode) [][]int {
	res := [][]int{}
	findLeavesHelper(root, &res)
	return res
}

func findLeavesHelper(root *models.TreeNode, res *[][]int) int {
	if root == nil {
		return -1
	}
	left := findLeavesHelper(root.Left, res)
	right := findLeavesHelper(root.Right, res)
	level := maxInt(left, right) + 1
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	return level
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
