package recursivewithglobalvar

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Postorder recursive. Use a global variable to store the current max path.
// For each node, we need to know the max sum of path start with left child
// and the max sum of path start with right child. Cap the two child sum to
// not less than because if they are negative number, we can just abandon them
// which is the same as set it to 0. Update global max with root.Val+left+right.
// Return the max between left lean path and right lean path
// Time: O(n)
// Space: O(h)
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var maxPathSumHelper func(*TreeNode) int
	maxPathSumHelper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := maxInt(maxPathSumHelper(root.Left), 0)
		rightMax := maxInt(maxPathSumHelper(root.Right), 0)
		curMax := root.Val + leftMax + rightMax
		if curMax > res {
			res = curMax
		}
		return maxInt(leftMax, rightMax) + root.Val
	}
	maxPathSumHelper(root)
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
