package recursive

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The constrain is the same as other house robber problem but the house is organized as a tree.
// For each subtree, we need to know the max sum with root and the max sum without root and return
// this back to callee. For max sum without root, it can be either sum include it's children or not.
// Time: O(n)
// Space: O(h)
func rob(root *TreeNode) int {
	a, b := robHelper(root)
	return maxInt(a, b)
}

func robHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left1, left2 := robHelper(root.Left)
	right1, right2 := robHelper(root.Right)
	// Be careful
	return root.Val + left2 + right2, maxInt(left1, left2) + maxInt(right1, right2)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
