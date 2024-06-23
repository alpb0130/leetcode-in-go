package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Level order traversal
// Time: O(n)
// Space: O(h)
func rightSideView(root *TreeNode) []int {
	res := []int{}
	rightSideViewHelper(root, &res, 0)
	return res
}

func rightSideViewHelper(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if level == len(*res) {
		*res = append(*res, root.Val)
	}
	rightSideViewHelper(root.Right, res, level+1)
	rightSideViewHelper(root.Left, res, level+1)
}
