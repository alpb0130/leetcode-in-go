package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Similar to level order traversal but put the deep node in front list
// The destination index should be len(res)-h-1
// Time: O(n)
// Space: O(n)
func levelOrderBottom(root *TreeNode) [][]int {
	res := &[][]int{}
	if root == nil {
		return *res
	}
	// Preorder traversal to construct the results, need to pass height
	traversalHelper(root, res, 0)
	return *res
}

// pointer of slice
func traversalHelper(root *TreeNode, res *[][]int, h int) {
	if root == nil {
		return
	}
	if h >= len(*res) {
		*res = append([][]int{[]int{}}, (*res)...)
	}
	(*res)[len(*res)-h-1] = append((*res)[len(*res)-h-1], root.Val)
	traversalHelper(root.Left, res, h+1)
	traversalHelper(root.Right, res, h+1)
}
