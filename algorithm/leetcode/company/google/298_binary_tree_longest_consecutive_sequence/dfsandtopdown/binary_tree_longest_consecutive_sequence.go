package dfsandtopdown

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Top down method.
// Time: O(n)
// Space: O(logn)
func longestConsecutive(root *TreeNode) int {
	res := 0
	helper(root, nil, 0, &res)
	return res
}

func helper(root, pred *TreeNode, cur int, res *int) {
	if root == nil {
		return
	}
	if pred != nil && root.Val == pred.Val+1 {
		cur++
	} else {
		cur = 1
	}
	if cur > *res {
		*res = cur
	}
	helper(root.Left, root, cur, res)
	helper(root.Right, root, cur, res)
}
