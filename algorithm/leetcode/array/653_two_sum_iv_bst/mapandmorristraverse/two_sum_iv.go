package mapandmorristraverse

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func findTarget(root *TreeNode, k int) bool {
	numMap := map[int]bool{}
	for root != nil {
		if root.Left != nil {
			pred := root.Left
			for pred.Right != nil && pred.Right != root {
				pred = pred.Right
			}
			if pred.Right == nil {
				pred.Right = root
				root = root.Left
				continue
			} else {
				pred.Right = nil
			}
		}
		if numMap[k-root.Val] {
			return true
		}
		numMap[root.Val] = true
		root = root.Right
	}
	return false
}
