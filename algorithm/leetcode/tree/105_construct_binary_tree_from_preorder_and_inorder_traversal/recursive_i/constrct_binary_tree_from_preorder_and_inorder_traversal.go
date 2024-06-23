package recursive_i

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n^2) - if the tree is left skew, we will get n square time because we need get the inorder
// index for each node in linear time
// Space: O(h)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for i < len(inorder) && inorder[i] != preorder[0] {
		i++
	}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
