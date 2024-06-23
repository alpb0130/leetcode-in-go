package recursive

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursively
// Time: O(n)
// Space: O(1)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := []int{}
	left := inorderTraversal(root.Left)
	results = append(left, root.Val)
	right := inorderTraversal(root.Right)
	results = append(results, right...)
	return results
}
