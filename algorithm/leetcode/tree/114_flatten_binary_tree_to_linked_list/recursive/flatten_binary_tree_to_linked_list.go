package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive post traversal but need to return the last node so that
// we can connect the tree
// Time: O(n)
// Space: O(h)
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenHelper(root)
}

func flattenHelper(root *TreeNode) *TreeNode {
	end := root
	var leftEnd, rightEnd *TreeNode
	if root.Left != nil {
		leftEnd = flattenHelper(root.Left)
		end = leftEnd
	}
	if root.Right != nil {
		rightEnd = flattenHelper(root.Right)
		end = rightEnd
	}
	if leftEnd != nil {
		leftEnd.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	return end
}
