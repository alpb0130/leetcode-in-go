package recursive


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(h)
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	left := postorderTraversal(root.Left)
	res = append(res, left...)
	right := postorderTraversal(root.Right)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}
