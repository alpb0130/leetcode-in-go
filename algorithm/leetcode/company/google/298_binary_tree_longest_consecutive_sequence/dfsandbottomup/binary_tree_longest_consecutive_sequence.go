package dfsandbottomup

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Bottom up method. For each node, we return the length of consecutive sequence that
// start with the node and also update the res.
// Time: O(n)
// Space: O(logn)
func longestConsecutive(root *TreeNode) int {
	res := 0
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := 1
	left := helper(root.Left, res)
	if root.Left != nil && root.Val+1 == root.Left.Val {
		l = maxInt(l, left+1)
	}
	right := helper(root.Right, res)
	if root.Right != nil && root.Val+1 == root.Right.Val {
		l = maxInt(l, right+1)
	}
	if l > *res {
		*res = l
	}
	return l
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
