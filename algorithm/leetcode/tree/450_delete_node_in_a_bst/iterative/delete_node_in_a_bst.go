package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Add a dummy node for edge case. Use a pointer to note parent.
// Time: O(logn)
// Space: O(h)
func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{}
	dummy.Left = root
	// find node
	pred := dummy
	cur := root
	for cur != nil && cur.Val != key {
		pred = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		return dummy.Left
	}
	if cur.Left == nil {
		if pred.Left == cur {
			pred.Left = cur.Right
		} else {
			pred.Right = cur.Right
		}
	} else {
		left := cur.Left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = cur.Right
		if pred.Left == cur {
			pred.Left = cur.Left
		} else {
			pred.Right = cur.Left
		}
	}
	return dummy.Left
}
