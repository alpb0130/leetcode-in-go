package dfs

import (
	"algorithm/leetcode/company/linkedin/models"
	"math"
	"strconv"
)

// DFS. Get the height of the tree and they pre-define the result array. Basically this problem
// require use to print out tree if it's a perfect binary tree. We pre-define the array and just
// fill number into it. The index can be easy identified recursively if we know the boundary of
// current subtree. The root must be in the middle.
// Time: O(h*2^h)
// Space: o(h) - h is the height of the tree
func printTree(root *models.TreeNode) [][]string {
	h := treeHeight(root)
	l := int(math.Pow(2, float64(h))) - 1
	res := make([][]string, h)
	for i := 0; i < h; i++ {
		res[i] = make([]string, l)
	}
	fill(root, res, 0, 0, l-1)
	return res
}

func fill(root *models.TreeNode, res [][]string, l, start, end int) {
	if root == nil {
		return
	}
	mid := (start + end) / 2
	res[l][mid] = strconv.Itoa(root.Val)
	fill(root.Left, res, l+1, start, mid-1)
	fill(root.Right, res, l+1, mid+1, end)
}

func treeHeight(root *models.TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeHeight(root.Left)
	right := treeHeight(root.Right)
	return maxInt(left, right) + 1
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
