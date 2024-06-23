package divideandconquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Divide and conquer. The straightforward way is to traverse the who tree but we don't need to
// do that because it's a complete tree. What we can do is identify the subtree which is not a
// perfect tree. We have two cases:
// 1. If height of the RIGHT subtree is not h-1 (h is the height of parent tree), the right subtree
//    must be a perfect tree but with height h-2, then the left subtree is not a perfect tree. We
//    can call the function recursively on left node.
// 2. If height of the RIGHT subtree is h-1 (h is the height of parent tree), the left subtree
//    must be a perfect tree with height h-1, then the right subtree is probably not a perfect tree.
//    We can call the function recursively on right node.
// We can always get the # of node of a perfect tree in O(1) time if we know the tree height.
// Time: O(logn*logn)
// Space: O(h)
func countNodes(root *TreeNode) int {
	h := getTreeHeight(root)
	return countNodesHelper(root, h)
}

func countNodesHelper(root *TreeNode, h int) int {
	if root == nil {
		return 0
	}
	if getTreeHeight(root.Right) == h-1 {
		return countNodesHelper(root.Right, h-1) + 1<<uint32(h-1)
	} else {
		return countNodesHelper(root.Left, h-1) + 1<<uint32(h-2)
	}
}

func getTreeHeight(root *TreeNode) int {
	h := 0
	for root != nil {
		root = root.Left
		h++
	}
	return h
}
