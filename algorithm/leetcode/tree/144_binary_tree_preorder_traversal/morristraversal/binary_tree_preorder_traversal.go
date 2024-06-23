package morristraversal

import "algorithm/leetcode/tree/models"

// Time: O(n)
// Space: O(1)
func PreorderTraversal(root *models.TreeNode) []int {
	cur := root
	results := []int{}
	for cur != nil {
		results = append(results, cur.Val)
		if cur.Left != nil {
			predecessor := cur.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = cur.Right
			cur.Right = cur.Left
			cur = cur.Right
		} else {
			cur = cur.Right
		}
	}
	return results
}
