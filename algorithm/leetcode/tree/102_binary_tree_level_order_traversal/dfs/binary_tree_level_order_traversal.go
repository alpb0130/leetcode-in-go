package dfs


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use preorder recursive traversal. When we do recursive all, we only need to append value to
// corresponding result list of current height. Remember that you need pass a pointer of slice
// because you will need to modify the underlying length of the matrix.
// Time: O(n)
func levelOrder(root *TreeNode) [][]int {
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
		*res = append(*res, []int{})
	}
	(*res)[h] = append((*res)[h], root.Val)
	traversalHelper(root.Left, res, h+1)
	traversalHelper(root.Right, res, h+1)
}
