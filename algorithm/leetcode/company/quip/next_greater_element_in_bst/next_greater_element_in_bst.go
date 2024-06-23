package next_greater_element_in_bst

type TreeNode struct {
	val    int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

// Give a treenode, find the next greater node
func NextGreaterNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.right != nil {
		right := node.right
		for right.left != nil {
			right = right.left
		}
		return right
	}

	parent := node.parent
	for parent != nil && parent.val < node.val {
		parent = parent.parent
	}

	return parent
}
