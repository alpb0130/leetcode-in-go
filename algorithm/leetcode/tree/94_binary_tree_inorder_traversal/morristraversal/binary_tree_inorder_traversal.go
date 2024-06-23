package morristraversal

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


// Use threaded tree to traverse the tree
// For reference:
// https://en.wikipedia.org/wiki/Threaded_binary_tree
// https://stackoverflow.com/questions/5502916/explain-morris-inorder-tree-traversal-without-using-stacks-or-recursion
// Time: O(n)
// Space: O(1)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	results := []int{}
	for cur != nil {
		if cur.Left != nil {
			left := cur.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = cur
			temp := cur
			cur = cur.Left
			// Don't forget this step!!!
			temp.Left = nil
		} else {
			results = append(results, cur.Val)
			cur = cur.Right
		}
	}
	return results
}
