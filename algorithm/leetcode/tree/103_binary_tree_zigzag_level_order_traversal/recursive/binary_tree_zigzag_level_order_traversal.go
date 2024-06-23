package recursive

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
// Time: O(n)
// Space: O(2 to the power of h)
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	zigzagHelper(root, &res, 0)
	return res
}

func zigzagHelper(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}
	if level == len(*res) {
		(*res) = append(*res, []int{})
	}
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		(*res)[level] = append([]int{root.Val}, (*res)[level]...)
	}
	zigzagHelper(root.Left, res, level+1)
	zigzagHelper(root.Right, res, level+1)
}
