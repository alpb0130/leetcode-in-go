package recursivewithglobalvar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(h)
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var sumNumbersHelper func(*TreeNode, int)
	sumNumbersHelper = func(root *TreeNode, val int) {
		if root.Left == nil && root.Right == nil {
			res += (10*val + root.Val)
		}
		if root.Left != nil {
			sumNumbersHelper(root.Left, 10*val+root.Val)
		}
		if root.Right != nil {
			sumNumbersHelper(root.Right, 10*val+root.Val)
		}
	}
	sumNumbersHelper(root, 0)
	return res
}
