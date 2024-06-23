package mapanddfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func findTarget(root *TreeNode, k int) bool {
	return findHelper(root, k, map[int]bool{})
}

func findHelper(root *TreeNode, k int, numMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if numMap[k-root.Val] {
		return true
	}
	numMap[root.Val] = true
	return findHelper(root.Left, k, numMap) || findHelper(root.Right, k, numMap)
}
