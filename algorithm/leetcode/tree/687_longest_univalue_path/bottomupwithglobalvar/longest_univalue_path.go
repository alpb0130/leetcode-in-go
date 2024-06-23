package bottomupwithglobalvar

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Basically the same idea with another solution but different in two place:
// 1. Have a global value to store the len of current longest path
// 2. Pass the root val down the callee and determine what value to return to caller.
//    The return value is the len of max path start with root
// Time: O(n)
// Space: O(h)
func longestUnivaluePath(root *TreeNode) int {
	var res int
	var longestUnivaluePathHelper func(*TreeNode, int) int
	longestUnivaluePathHelper = func(root *TreeNode, parentVal int) int {
		if root == nil {
			return 0
		}
		leftLen := longestUnivaluePathHelper(root.Left, root.Val)
		rightLen := longestUnivaluePathHelper(root.Right, root.Val)
		if res < leftLen+rightLen {
			res = leftLen + rightLen
		}
		if root.Val != parentVal {
			return 0
		}
		return 1 + MaxInt(leftLen, rightLen)
	}
	longestUnivaluePathHelper(root, 0)
	return res
}

func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
