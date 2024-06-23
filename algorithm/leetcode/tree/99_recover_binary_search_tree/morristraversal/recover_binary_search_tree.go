package morristraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris traversal without breaking the tree
// Time: O(n)
// Space: O(h)
func recoverTree(root *TreeNode) {
	var pre, first, second *TreeNode
	cur := root
	for cur != nil {
		if cur.Left != nil {
			predecessor := cur.Left
			for predecessor.Right != nil && predecessor.Right != cur {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = cur
				cur = cur.Left
				continue
			} else {
				predecessor.Right = nil
			}
		}
		if pre != nil && pre.Val >= cur.Val {
			if first == nil {
				first = pre
			}
			second = cur
		}
		pre = cur
		cur = cur.Right
	}
	first.Val, second.Val = second.Val, first.Val
}
