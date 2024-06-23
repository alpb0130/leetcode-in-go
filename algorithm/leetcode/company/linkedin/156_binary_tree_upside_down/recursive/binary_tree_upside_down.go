package recursive

import "algorithm/leetcode/company/linkedin/models"

// The tree in this problems is all right node is either leaf node with sibling (left node)
// or empty. So it would be a tree very left-skew, like:
//				1               1
// 			   / \             /
//            2  (3)   =>     2 - (3)
//           / \             /
//          4  (5)          4 - (5)
//             ...
// We can solve this problem recursively in bottom-up style. For every function call,
// we return the new root which should be the lowest-leftmost node and do the rotation.
// Time: O(n) - the tree is a left-skewed tree so it size it equal to it's height
// Space: O(n)
func upsideDownBinaryTree(root *models.TreeNode) *models.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	newRoot := upsideDownBinaryTree(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil
	return newRoot
}
