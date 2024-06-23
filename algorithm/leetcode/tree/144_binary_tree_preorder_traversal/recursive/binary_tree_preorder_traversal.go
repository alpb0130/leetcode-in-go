package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
// Time: O(n)
// Space: O(h)
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := []int{root.Val}
	left := preorderTraversal(root.Left)
	results = append(results, left...)
	right := preorderTraversal(root.Right)
	results = append(results, right...)
	return results
}
