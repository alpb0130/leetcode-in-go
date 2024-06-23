package bottomup

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// This problem can be solved recursively. For each node, we need to:
// 1. Check whether the root has the same value as it's child and recompute the
//    local path len.
// 2. Compare the local path len with existing longest path from left and right subtree
//    and update the current max
// 3. Return the len of max path start with root up to the caller
// Time: O(n)
// Space: O(h)
func longestUnivaluePath(root *TreeNode) int {
	_, res := longestUnivaluePathHelper(root)
	return res
}

func longestUnivaluePathHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftRootLen, leftMaxLen := longestUnivaluePathHelper(root.Left)
	rightRootLen, rightMaxLen := longestUnivaluePathHelper(root.Right)

	leftLen, rightLen := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		leftLen = leftRootLen + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rightLen = rightRootLen + 1
	}

	return maxInt(leftLen, rightLen), maxInt(maxInt(leftMaxLen, rightMaxLen), leftLen+rightLen)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
