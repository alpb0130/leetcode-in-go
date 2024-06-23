package recursive

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Not use global value. Each call return the max in subtree and max between
// left lean path and right lean path
// Time: O(n)
// Space: O(h)
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, res := maxPathSumHelper(root)
	return res
}

func maxPathSumHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt32
	}
	left, leftMax := maxPathSumHelper(root.Left)
	right, rightMax := maxPathSumHelper(root.Right)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	max := maxInt(root.Val+left+right, maxInt(leftMax, rightMax))
	return maxInt(left, right) + root.Val, max

}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
