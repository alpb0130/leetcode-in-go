package iterative

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Two pass. The pass path get the frequency of mode. The second pass add all modes into res
// Time: O(n)
// Space: O(1)
func findMode(root *TreeNode) []int {
	var res []int
	max := 0
	cur := 0
	curVal := 0
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if root.Val != curVal {
			curVal = root.Val
			cur = 0
		}
		cur++
		if cur > max {
			max = cur
		} else if cur == max {
			if res != nil {
				res = append(res, root.Val)
			}
		}
		inorder(root.Right)
	}
	inorder(root)
	cur = 0
	res = []int{}
	inorder(root)
	return res
}
