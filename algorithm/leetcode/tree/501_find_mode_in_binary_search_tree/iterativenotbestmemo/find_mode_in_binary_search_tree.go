package iterativenotbestmemo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n) worst case
func findMode(root *TreeNode) []int {
	res := []int{}
	curMax := 0
	cur := 0
	var prev *TreeNode
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
		if prev != nil && root.Val == prev.Val {
			cur++
		} else {
			if cur > curMax && prev != nil {
				res = []int{prev.Val}
				curMax = cur
			} else if cur == curMax && prev != nil {
				res = append(res, prev.Val)
			}
			cur = 1
		}
		prev = root
		root = root.Right
	}
	if cur > curMax && prev != nil {
		res = []int{prev.Val}
		curMax = cur
	} else if cur == curMax && prev != nil {
		res = append(res, prev.Val)
	}
	return res
}
